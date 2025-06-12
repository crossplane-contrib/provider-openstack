# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]

### Changed

- Update terraform-provider-openstack to `v3.2.0`
- Update to golang `v1.24.4`
- Update golang dependencies

## [0.7.1] - 2025-06-18

### Fixed

- Fix wrong reference from Neutron Subnets to Networks (#137)
- Use correct sdk for terraform-provider-openstack `v3.0.0` (Fix in 0.7.0 was not enough)

## [0.7.0] - 2025-05-22

### Changed

- Update golang dependencies
- Update to golang `v1.23.9`

### Fixed

- Use correct sdk for terraform-provider-openstack `v3.0.0`

## [0.6.0] - 2025-04-17

This release upgrades the internally used terraform-provider-openstack from v1 to v3. This includes removals of many fields that were previously deprecated. If you didn't use any deprecated fields, this upgrade should break non of your workloads.

View the [upstream changelogs](https://github.com/terraform-provider-openstack/terraform-provider-openstack/blob/main/CHANGELOG.md#300--25-september-2024-) for detailed changes between `v1.54.1` and `v3.0.0`.

### Changed

- Update terraform-provider-openstack to `v3.0.0`

### Removed

- Removed support for deprecated `openstack_compute_floatingip_associate_v2` resource. Use `openstack_networking_floatingip_associate_v2` instead.
- Removed support for deprecated `openstack_compute_floatingip_v2` resource. Use `openstack_networking_floatingip_v2` instead.
- Removed support for deprecated `openstack_compute_secgroup_v2` resource. Use `openstack_networking_secgroup_v2` instead.

## [0.5.0] - 2025-03-24

### Migration from 0.4.0

- In compliance with the Community Extension Projects Policies, we moved the provider to the new crossplane package registry. The provider will be available at `xpkg.crossplane.io/crossplane-contrib/provider-openstack`. The previously published packages in the Upbound Marketplace will no longer be updated but kept there for backwards compatbility.

### Added

- All remaining config parameters that are supported via [terraform-provider-openstack](https://registry.terraform.io/providers/terraform-provider-openstack)
- Made it easier to configure authentication by not requiring so many fields.
- Migrated to Upjets new No-Fork/Provider v2 SDK Architecture. The provider now does not ship with Terraform CLI anymore.
- Added many cross resource references (autmatically generated from Terraform provider schema)
- Add native metrics for managed resources (see [Upjet 1.3.0](https://github.com/crossplane/upjet/releases/tag/v1.3.0) for more details)

### Changed

- Update terraform-provider-openstack to `v1.54.1`
- Update Terraform CLI to `v1.5.7`
- Update provider to golang v1.23
- Update upjet to v1.4.0
- Update crossplane runtime to 1.16
- Update various other go dependencies
- Modernize CI Pipeline
- Rebase dockerfle to alpine 3.12

### Removed

- Removed Terraform CLI specific metrics, as there is no Terraform CLI involved anymore

## [0.4.0] - 2024-06-11

### Added

- Support setting a custom CA certificate for the OpenStack API
- Add cross references for security groups on InstanceV2 resources (Thanks @etesami)
- Option to specify "insecure" option (Thanks @FCosta999)

## [0.3.0] - 2023-12-04

See Migration Guide

### Added

- Support for Beta Management Policies (Granular Management Policies available since Crossplane 1.14) and initProvider
- Added all reminaing resources provided by terraform-provider-openstack

### Changed

- **Breaking:** Change api group of all managed resources to `openstack.crossplane.io` as this provider has no official affiliations with Upbound.
- Update to upjet `v1.0.0` and its dependencies
- Updated to golang `v1.20`
- Update Terraform CLI to `v1.5.5`

### Removed

- Support for Alpha Management Policies (e.g. `ObservedOnly`)

## [0.2.0] - 2023-10-27

### Added

- Add more supported values for provider configuration, including authentication options.

### Changed

- Update terraform-provider-openstack to v1.53.0
- Update upjet and corresponding dependencies

## [0.1.x]

See [Commitlog](https://github.com/crossplane-contrib/provider-openstack/commits/release-v0.1)

# Migration

## 0.2.x to 0.3.x

- Change all apiversions from `*.openstack.upbound.io` to `*.openstack.crossplane.io`. See https://docs.crossplane.io/knowledge-base/guides/import-existing-resources/.
