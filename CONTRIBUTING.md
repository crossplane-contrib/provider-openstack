# Contributing

THis project uses [Makefile](Makefile) to manage common tasks. You can see all available commands by running:

```bash
make help
```

Ensure your environment provides the encessary tools to build and test the project.

- [Go](https://golang.org/dl/) in the version specified in [go.mod](go.mod)
- [Docker](https://www.docker.com/get-started) or a compatible container runtime

```bash

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

## Environment Speicific Notes

### Apple Silicon (ARM64)

- When using Rancher Desktop, use QEMU as the emulator to ensure kind works correctly (see https://github.com/rancher-sandbox/rancher-desktop/issues/9192)