# Crossplane Provider OpenStack

<div align="center">

[![Upbound Marketplace](https://img.shields.io/badge/provider--openstack-xxx?label=upbound%20marketplace&color=blue)](https://marketplace.upbound.io/providers/crossplane-contrib/provider-openstack)

</div>

https://marketplace.upbound.io/providers/crossplane-contrib/provider-openstack/

`provider-openstack` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
OpenStack API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/crossplane-contrib/provider-openstack):

```
up ctp provider install crossplane-contrib/provider-openstack:vX.X.X
```

Alternatively, you can use declarative installation:

```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-openstack
spec:
  package: xpkg.upbound.io/crossplane-contrib/provider-openstack:vX.Y.Z
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-openstack).

## Configuration

```yaml
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

The secret key must contain a json dictionary that provides the authentication data

```json
{
  "auth_url": "https://auth.openstack.example/",
  "application_credential_id": "123456789",
  "application_credential_secret": "secret-key"
}
```

Check [Terraform OpenStack provider docs](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs#configuration-reference) to see available configuration settings. Currently not all options of the upstream provider are supported. Check [client code](https://github.com/crossplane-contrib/provider-openstack/blob/main/internal/clients/openstack.go#L66) to see if your option is supported. If something is missing, please open a new issue.

## Developing

Install the required submodules to build and run:

```console
make submodules
```

Apply the Current CRDs and a providerConfig:

```console
kubectl apply -f package/crds
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

Run against a Kubernetes cluster: (make sure to apply CRDs and providerConfig)

```console
make run
```

Run a testbuild with linting:

```console
make reviewable
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-openstack/issues).
