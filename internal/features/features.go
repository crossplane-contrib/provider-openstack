/*
 Copyright 2022 Upbound Inc
*/

package features

import (
	xpfeature "github.com/crossplane/crossplane-runtime/v2/pkg/feature"
)

// Feature flags.
const (
	// EnableBetaManagementPolicies enables beta support for
	// Management Policies. See the below design for more details.
	// https://github.com/crossplane/crossplane/pull/3531
	EnableBetaManagementPolicies xpfeature.Flag = xpfeature.EnableBetaManagementPolicies

	// EnableCloudChangelogs enables alpha support for
	// Cloud Changelogs. See the below design for more details.
	// https://github.com/crossplane/crossplane/pull/5822/changes
	EnableAlphaChangeLogs xpfeature.Flag = xpfeature.EnableAlphaChangeLogs
)
