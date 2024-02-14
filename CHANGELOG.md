# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Support setting a custom CA certificate for the OpenStack API

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
