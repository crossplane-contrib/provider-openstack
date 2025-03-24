# Crossplane Provider OpenStack

<div align="center">

[![GitHub release](https://img.shields.io/github/release/crossplane-contrib/provider-openstack/all.svg)](https://github.com/crossplane-contrib/provider-openstack/releases)

</div>

`provider-openstack` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
OpenStack API.

## Getting Started

### Installation

You can use declarative installation to install the provider:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-openstack
spec:
  package: xpkg.crossplane.io/crossplane-contrib/provider-openstack:vX.Y.Z
```

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-openstack).

### Configuration

```yaml
---
# Providerconfig that referers to the secret
apiVersion: openstack.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: provider-openstack-config
spec:
  credentials:
    source: Secret
    secretRef:
      key: config
      name: provider-openstack-config
      namespace: crossplane

---
# Secret that stores credentials and other configuration
apiVersion: v1
kind: Secret
metadata:
  name: provider-openstack-config
  namespace: crossplane
type: Opaque
data:
  config: <see below>
```

The secret key must contain a json dictionary that provides the authentication data.
You can create the secret via this command:

```bash
kubectl create secret generic provider-openstack-config --from-file=config=config.json --namespace crossplane
```

```json
// config.json
{
  "auth_url": "https://auth.openstack.example/",
  "application_credential_id": "123456789",
  "application_credential_secret": "secret-key"
}
```

Check [Terraform OpenStack provider docs](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs#configuration-reference) to see available configuration settings. Currently not all options of the upstream provider are supported. Check [client code](https://github.com/crossplane-contrib/provider-openstack/blob/main/internal/clients/openstack.go#L66) to see if your option is supported. If something is missing, please open a new issue.


### Deployment Customization

You can use a `DeploymentRuntimeConfig` to provide custom arguments or otherwise modify the provider deployment

Available command line arguments can be found [here](cmd/provider/main.go)

```yaml
---
# Create a DeploymentRuntimeConfig to customize the provider deployment
apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
  name: provider-openstack
spec:
  deploymentTemplate:
    spec:
      # Control replica count to temporary disable deployment. Do not scale more than 1 replica.
      replicas: 1
      selector: {}
      template:
        metadata:
          annotations:
            # Add annotations, e.g. to enable metrics scraping
            prometheus.io/path: /metrics
            prometheus.io/port: "8080"
            prometheus.io/scrape: "true"
        spec:
          containers:
          - args:
            # Add command line arguments, e.g. to enable management policies
            - --enable-management-policies
            name: package-runtime

---
# Add this to your provider resource to reference the DeploymentRuntimeConfig
spec:
  runtimeConfigRef:
    apiVersion: pkg.crossplane.io/v1beta1
    kind: DeploymentRuntimeConfig
    name: provider-openstack
```

## Developing

Install the required submodules to build and run:

```bash
make submodules
```

Apply the Current CRDs and a providerConfig:

```bash
kubectl apply -f package/crds
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

Run against a Kubernetes cluster: (make sure to apply CRDs and providerConfig)

```bash
make run
```

Run a testbuild with linting:

```bash
make reviewable
```

Build binary:

```bash
make build
```

## Release a new version (Maintainer)

- Update Changelog (Add new Version & Date)
- Create or merge to existing release branch (release-v(major).(minor))
- Run Release pipeline on release branch, using specific version as parameter

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-openstack/issues).
