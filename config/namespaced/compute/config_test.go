package compute

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/resource/json"

	"github.com/crossplane-contrib/provider-openstack/apis/namespaced/compute/v1alpha1"
)

func TestInstanceV2LateInitializeIgnoresPlacementFields(t *testing.T) {
	observed := map[string]any{
		"availability_zone":   "row-d",
		"hypervisor_hostname": "some-sandbox-kvm-2.ostk.dc",
		"flavor_id":           "0fbaa4f6-169a-4991-be42-cede3c966309",
		"flavor_name":         "vm.m",
		"region":              "RegionOne",
		"power_state":         "active",
	}
	raw, err := json.TFParser.Marshal(observed)
	if err != nil {
		t.Fatalf("failed to marshal observed state: %v", err)
	}

	var tr resource.Terraformed = &v1alpha1.InstanceV2{}
	changed, err := tr.LateInitialize(raw)
	if err != nil {
		t.Fatalf("LateInitialize failed: %v", err)
	}
	if !changed {
		t.Fatal("expected LateInitialize to report changes for non-ignored fields")
	}

	inst := tr.(*v1alpha1.InstanceV2)
	if inst.Spec.ForProvider.AvailabilityZone != nil {
		t.Errorf("availabilityZone should not be late-initialized, got %q", *inst.Spec.ForProvider.AvailabilityZone)
	}
	if inst.Spec.ForProvider.HypervisorHostname != nil {
		t.Errorf("hypervisorHostname should not be late-initialized, got %q", *inst.Spec.ForProvider.HypervisorHostname)
	}
	if inst.Spec.ForProvider.FlavorID != nil {
		t.Errorf("flavorId should not be late-initialized, got %q", *inst.Spec.ForProvider.FlavorID)
	}
	if inst.Spec.ForProvider.FlavorName != nil {
		t.Errorf("flavorName should not be late-initialized, got %q", *inst.Spec.ForProvider.FlavorName)
	}
	if inst.Spec.ForProvider.Region == nil || *inst.Spec.ForProvider.Region != "RegionOne" {
		t.Errorf("region should still be late-initialized, got %v", inst.Spec.ForProvider.Region)
	}
	if inst.Spec.ForProvider.PowerState == nil || *inst.Spec.ForProvider.PowerState != "active" {
		t.Errorf("powerState should still be late-initialized, got %v", inst.Spec.ForProvider.PowerState)
	}
}
