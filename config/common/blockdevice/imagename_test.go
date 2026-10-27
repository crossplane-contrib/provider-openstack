package blockdevice

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
)

func testResource(create schema.CreateContextFunc, resolve imageResolver) *schema.Resource {
	return testResourceWithRead(create, nil, resolve, nil)
}

func testResourceWithRead(create schema.CreateContextFunc, read schema.ReadContextFunc, resolve imageResolver, name imageNamer) *schema.Resource {
	res := &schema.Resource{
		CreateContext: create,
		ReadContext:   read,
		Schema: map[string]*schema.Schema{
			"block_device": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_type": {Type: schema.TypeString, Required: true, ForceNew: true},
						"uuid":        {Type: schema.TypeString, Optional: true, ForceNew: true},
					},
				},
			},
		},
	}
	extendResource(res, resolve, name)

	return res
}

func fakeResolver(ids map[string]string) imageResolver {
	return func(_ context.Context, _ *schema.ResourceData, _ any, name string) (string, error) {
		id, ok := ids[name]
		if !ok {
			return "", errors.Errorf("image %q not found", name)
		}

		return id, nil
	}
}

func TestExtendSchema(t *testing.T) {
	res := testResource(nil, nil)
	bd := res.Schema["block_device"].Elem.(*schema.Resource)

	in, ok := bd.Schema["image_name"]
	if !ok {
		t.Fatal("image_name was not added to block_device schema")
	}
	if !in.Optional || !in.ForceNew || in.Type != schema.TypeString {
		t.Errorf("unexpected image_name schema: %+v", in)
	}
	if bd.Schema["uuid"].DiffSuppressFunc == nil {
		t.Error("uuid DiffSuppressFunc was not set")
	}
	if res.CreateContext != nil || res.ReadContext != nil {
		t.Error("nil CreateContext/ReadContext must stay nil")
	}
}

func TestCreateResolvesImageName(t *testing.T) {
	var gotUUID, gotName string
	create := func(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
		gotUUID = d.Get("block_device.0.uuid").(string)
		gotName = d.Get("block_device.0.image_name").(string)
		d.SetId("id")

		return nil
	}
	res := testResource(create, fakeResolver(map[string]string{"ubuntu": "1111"}))

	d := schema.TestResourceDataRaw(t, res.Schema, map[string]any{
		"block_device": []any{map[string]any{
			"source_type": "image",
			"image_name":  "ubuntu",
		}},
	})
	if diags := res.CreateContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}
	if gotUUID != "1111" {
		t.Errorf("uuid = %q, want 1111", gotUUID)
	}
	if gotName != "ubuntu" {
		t.Errorf("image_name = %q, want ubuntu", gotName)
	}
}

func TestCreateDoesNotTouchUUID(t *testing.T) {
	called := false
	create := func(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
		called = d.Get("block_device.0.uuid").(string) == "explicit"

		return nil
	}
	res := testResource(create, fakeResolver(nil))

	d := schema.TestResourceDataRaw(t, res.Schema, map[string]any{
		"block_device": []any{map[string]any{
			"source_type": "image",
			"uuid":        "explicit",
		}},
	})
	if diags := res.CreateContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}
	if !called {
		t.Error("explicit uuid was not preserved")
	}
}

func TestCreateValidation(t *testing.T) {
	cases := map[string]map[string]any{
		"WrongSource": {"source_type": "volume", "image_name": "ubuntu"},
		"NotFound":    {"source_type": "image", "image_name": "missing"},
	}
	for name, bd := range cases {
		t.Run(name, func(t *testing.T) {
			create := func(context.Context, *schema.ResourceData, any) diag.Diagnostics {
				t.Error("upstream create must not be called on validation error")

				return nil
			}
			res := testResource(create, fakeResolver(map[string]string{"ubuntu": "1111"}))
			d := schema.TestResourceDataRaw(t, res.Schema, map[string]any{"block_device": []any{bd}})
			if diags := res.CreateContext(context.Background(), d, nil); !diags.HasError() {
				t.Error("expected error")
			}
		})
	}
}

// TestCreateImageNameWinsOverLateInitUUID covers the recreation scenario where
// imageName is set in initProvider and uuid was late-initialized into forProvider
// from the observed TF state. On recreation the merged params contain both fields.
// image_name must take precedence: its resolved UUID replaces the stale UUID in
// Terraform ResourceData. Late initialization does not overwrite the existing
// forProvider UUID; suppressUUID ignores its diff while image_name is configured.
func TestCreateImageNameWinsOverLateInitUUID(t *testing.T) {
	var gotUUID string
	create := func(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
		gotUUID = d.Get("block_device.0.uuid").(string)
		d.SetId("new-id")

		return nil
	}
	res := testResource(create, fakeResolver(map[string]string{"ubuntu": "1111"}))

	// Simulate merged params: initProvider.imageName + forProvider.uuid (late-initialized).
	d := schema.TestResourceDataRaw(t, res.Schema, map[string]any{
		"block_device": []any{map[string]any{
			"source_type": "image",
			"image_name":  "ubuntu",
			"uuid":        "stale-late-initialized-uuid",
		}},
	})
	if diags := res.CreateContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("expected success when both set (late-init scenario): %v", diags)
	}
	// Resolver returns "1111" for "ubuntu"; the stale uuid must be replaced
	// inside the upstream create call.
	if gotUUID != "1111" {
		t.Errorf("uuid inside create = %q, want 1111 (resolved from image_name)", gotUUID)
	}
	// ResourceData must retain the resolved UUID as the created Terraform state.
	// This does not update the existing UUID in spec.forProvider. The subsequent
	// UUID diff is suppressed while image_name is configured (see TestDiff).
	if got := d.Get("block_device.0.uuid"); got != "1111" {
		t.Errorf("uuid in ResourceData after create = %v, want 1111 (resolved image ID)", got)
	}
}

// TestCreateImageNameWinsWhenImageReplaced covers the scenario where an image
// is replaced (same name, new UUID). forProvider.uuid holds the old UUID from a
// previous late-initialization. wrapCreate must use the freshly resolved UUID
// and retain it in Terraform ResourceData. The stored forProvider UUID remains
// unchanged; suppressUUID prevents drift while image_name is configured.
func TestCreateImageNameWinsWhenImageReplaced(t *testing.T) {
	const oldUUID = "old-uuid-1111"
	const newUUID = "new-uuid-2222"

	var gotUUID string
	create := func(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
		gotUUID = d.Get("block_device.0.uuid").(string)
		d.SetId("new-id")

		return nil
	}
	// Image "ubuntu" now resolves to the new UUID.
	res := testResource(create, fakeResolver(map[string]string{"ubuntu": newUUID}))

	// forProvider.uuid = oldUUID (from previous late-init), imageName still set.
	d := schema.TestResourceDataRaw(t, res.Schema, map[string]any{
		"block_device": []any{map[string]any{
			"source_type": "image",
			"image_name":  "ubuntu",
			"uuid":        oldUUID,
		}},
	})
	if diags := res.CreateContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}
	if gotUUID != newUUID {
		t.Errorf("uuid inside create = %q, want %q (freshly resolved)", gotUUID, newUUID)
	}
	// The created Terraform state must hold the new UUID. Late initialization
	// does not overwrite the old UUID in spec.forProvider; TestDiff covers the
	// suppression of that stale UUID while image_name remains configured.
	if got := d.Get("block_device.0.uuid"); got != newUUID {
		t.Errorf("uuid in ResourceData after create = %v, want %q (resolved image ID)", got, newUUID)
	}
}

func serverState(attrs map[string]string) *terraform.InstanceState {
	a := map[string]string{
		"id":                         "server-id",
		"block_device.#":             "1",
		"block_device.0.source_type": "image",
		"block_device.0.uuid":        "1111",
	}
	for k, v := range attrs {
		a[k] = v
	}

	return &terraform.InstanceState{ID: "server-id", Attributes: a}
}

// diffInputs builds the diff config and attaches its cty value to the state as
// RawConfig, the way upjet does, so diff suppress functions can read it.
func diffInputs(t *testing.T, res *schema.Resource, state *terraform.InstanceState, bd map[string]any) (*terraform.InstanceState, *terraform.ResourceConfig) {
	t.Helper()
	bd["source_type"] = "image"
	raw := map[string]any{"block_device": []any{bd}}
	v, err := schema.JSONMapToStateValue(raw, res.CoreConfigSchema())
	if err != nil {
		t.Fatal(err)
	}
	if state != nil {
		state.RawConfig = v
	}

	return state, terraform.NewResourceConfigRaw(raw)
}

func TestDiff(t *testing.T) {
	res := testResource(nil, nil)

	cases := map[string]struct {
		state      *terraform.InstanceState
		cfg        map[string]any
		wantDiff   bool
		wantNewKey bool
	}{
		"NewResource": {
			state:      nil,
			cfg:        map[string]any{"image_name": "ubuntu"},
			wantDiff:   true,
			wantNewKey: true,
		},
		"SameName": {
			state: serverState(map[string]string{"block_device.0.image_name": "ubuntu"}),
			cfg:   map[string]any{"image_name": "ubuntu"},
		},
		"ChangedName": {
			state:      serverState(map[string]string{"block_device.0.image_name": "ubuntu"}),
			cfg:        map[string]any{"image_name": "debian"},
			wantDiff:   true,
			wantNewKey: true,
		},
		"NoObservedName": {
			state: serverState(nil),
			cfg:   map[string]any{"image_name": "ubuntu"},
		},
		"UUIDOnlyConfigWithObservedName": {
			state: serverState(map[string]string{"block_device.0.image_name": "ubuntu"}),
			cfg:   map[string]any{"uuid": "1111"},
		},
		// The image was replaced under the same name and the server recreated
		// with the new ID, while late initialization left the old ID in spec.
		"StaleUUIDWithImageName": {
			state: serverState(map[string]string{
				"block_device.0.uuid":       "2222",
				"block_device.0.image_name": "ubuntu",
			}),
			cfg: map[string]any{"image_name": "ubuntu", "uuid": "1111"},
		},
		"ChangedNameWithStaleUUID": {
			state:      serverState(map[string]string{"block_device.0.image_name": "ubuntu"}),
			cfg:        map[string]any{"image_name": "debian", "uuid": "1111"},
			wantDiff:   true,
			wantNewKey: true,
		},
		"ChangedUUID": {
			state:    serverState(map[string]string{"block_device.0.image_name": "ubuntu"}),
			cfg:      map[string]any{"uuid": "2222"},
			wantDiff: true,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			state, cfg := diffInputs(t, res, tc.state, tc.cfg)
			diff, err := res.Diff(context.Background(), state, cfg, nil)
			if err != nil {
				t.Fatal(err)
			}
			gotDiff := diff != nil && !diff.Empty()
			if gotDiff != tc.wantDiff {
				t.Fatalf("diff = %s, want diff: %v", diff.GoString(), tc.wantDiff)
			}
			if tc.wantNewKey && !diff.RequiresNew() {
				t.Errorf("expected diff to require a new resource, got %v", diff)
			}
		})
	}
}

func TestReadPopulatesImageName(t *testing.T) {
	read := func(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
		// Emulates the upstream read which rebuilds block_device without image_name.
		_ = d.Set("block_device", []any{
			map[string]any{"source_type": "image", "uuid": "1111"},
			map[string]any{"source_type": "volume", "uuid": "vol"},
		})

		return nil
	}
	lookups := 0
	namer := func(_ context.Context, _ *schema.ResourceData, _ any, id string) (string, error) {
		lookups++
		if id != "1111" {
			return "", errors.New("unexpected id " + id)
		}

		return "ubuntu", nil
	}
	res := testResourceWithRead(nil, read, nil, namer)

	d := res.Data(serverState(nil))
	if diags := res.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}
	if got := d.Get("block_device.0.image_name"); got != "ubuntu" {
		t.Errorf("image_name = %v, want ubuntu", got)
	}
	if got := d.Get("block_device.1.image_name"); got != "" {
		t.Errorf("volume block device must not get image_name, got %v", got)
	}
	if lookups != 1 {
		t.Errorf("lookups = %d, want 1", lookups)
	}
}

func TestReadKeepsStateOnLookupError(t *testing.T) {
	read := func(context.Context, *schema.ResourceData, any) diag.Diagnostics { return nil }
	namer := func(context.Context, *schema.ResourceData, any, string) (string, error) {
		return "", errors.New("not found")
	}
	res := testResourceWithRead(nil, read, nil, namer)

	d := res.Data(serverState(map[string]string{"block_device.0.image_name": "ubuntu"}))
	if diags := res.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("lookup failure must not fail read: %v", diags)
	}
	if got := d.Get("block_device.0.image_name"); got != "ubuntu" {
		t.Errorf("image_name = %v, want previous value ubuntu", got)
	}
}
