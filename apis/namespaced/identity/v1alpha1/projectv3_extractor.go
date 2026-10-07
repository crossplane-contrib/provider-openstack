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

// ExtractDomainID returns the external name of a ProjectV3 only when it
// represents a domain (isDomain: true). Referencing a regular project
// yields an empty value, causing reference resolution to fail.
func ExtractDomainID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		p, ok := any(mg).(*ProjectV3)
		if !ok {
			return ""
		}
		isDomain := (p.Spec.ForProvider.IsDomain != nil && *p.Spec.ForProvider.IsDomain) ||
			(p.Status.AtProvider.IsDomain != nil && *p.Status.AtProvider.IsDomain)
		if !isDomain {
			return ""
		}
		return reference.ExternalName()(mg)
	}
}
