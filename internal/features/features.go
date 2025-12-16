/*
 Copyright 2022 Upbound Inc
*/

package features

import (
	xpfeature "github.com/crossplane/crossplane-runtime/v2/pkg/feature"
)

// Feature flags.
const (
	// EnableAlphaExternalSecretStores enables alpha support for
	// External Secret Stores. See the below design for more details.
	// https://github.com/crossplane/crossplane/blob/390ddd/design/design-doc-external-secret-stores.md
	EnableAlphaExternalSecretStores xpfeature.Flag = "EnableAlphaExternalSecretStores"

	// EnableBetaManagementPolicies enables beta support for
	// Management Policies. See the below design for more details.
	// https://github.com/crossplane/crossplane/pull/3531
	EnableBetaManagementPolicies xpfeature.Flag = xpfeature.EnableBetaManagementPolicies
)
