package networking

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/resource/json"

	"github.com/crossplane-contrib/provider-openstack/apis/cluster/networking/v1alpha1"
)

func TestPortV2LateInitializeIgnoresNovaManagedFields(t *testing.T) {
	observed := map[string]any{
		"device_id":    "d7905377-120f-4d30-b813-201623d7a122",
		"device_owner": "compute:az-dedicated",
		"binding": map[string]any{
			"host_id":   "some-sandbox-kvm-2",
			"profile":   "{}",
			"vnic_type": "normal",
		},
		"region": "RegionOne",
	}
	raw, err := json.TFParser.Marshal(observed)
	if err != nil {
		t.Fatalf("failed to marshal observed state: %v", err)
	}

	var tr resource.Terraformed = &v1alpha1.PortV2{}
	changed, err := tr.LateInitialize(raw)
	if err != nil {
		t.Fatalf("LateInitialize failed: %v", err)
	}
	if !changed {
		t.Fatal("expected LateInitialize to report changes for non-ignored fields")
	}

	port := tr.(*v1alpha1.PortV2)
	if port.Spec.ForProvider.DeviceID != nil {
		t.Errorf("deviceId should not be late-initialized, got %q", *port.Spec.ForProvider.DeviceID)
	}
	if port.Spec.ForProvider.DeviceOwner != nil {
		t.Errorf("deviceOwner should not be late-initialized, got %q", *port.Spec.ForProvider.DeviceOwner)
	}
	if port.Spec.ForProvider.Binding == nil {
		t.Fatal("binding should be late-initialized with non-ignored fields")
	}
	if port.Spec.ForProvider.Binding.HostID != nil {
		t.Errorf("binding.hostId should not be late-initialized, got %q", *port.Spec.ForProvider.Binding.HostID)
	}
	if port.Spec.ForProvider.Binding.Profile == nil || *port.Spec.ForProvider.Binding.Profile != "{}" {
		t.Errorf("binding.profile should still be late-initialized, got %v", port.Spec.ForProvider.Binding.Profile)
	}
	if port.Spec.ForProvider.Binding.VnicType == nil || *port.Spec.ForProvider.Binding.VnicType != "normal" {
		t.Errorf("binding.vnicType should still be late-initialized, got %v", port.Spec.ForProvider.Binding.VnicType)
	}
	if port.Spec.ForProvider.Region == nil || *port.Spec.ForProvider.Region != "RegionOne" {
		t.Errorf("region should still be late-initialized, got %v", port.Spec.ForProvider.Region)
	}
}
