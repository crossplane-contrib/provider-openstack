/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer
Copyright 2025 Yannick Schlosser
*/

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

// ExtractFloatingIPAddress returns the allocated floating IP address from
// status.atProvider.address of a FloatingipV2 managed resource.
func ExtractFloatingIPAddress() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		f, ok := any(mg).(*FloatingipV2)
		if !ok {
			return ""
		}
		if f.Status.AtProvider.Address == nil {
			return ""
		}
		return *f.Status.AtProvider.Address
	}
}
