# Provider OpenStack

`provider-openstack` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
OpenStack API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/dusky-mate/provider-openstack):
```
up ctp provider install dusky-mate/provider-openstack:v0.1.3
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-openstack
spec:
  package: dusky-mate/provider-openstack:v0.1.3
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/dusky-mate/provider-openstack).

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
open an [issue](https://github.com/dusky-mate/provider-openstack/issues).
