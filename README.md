# Provider Harbor

`provider-harbor` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) and exposes XRM-conformant managed resources 
for the [Harbor](https://goharbor.io/) API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/globallogicuki/provider-harbor):
```
up ctp provider install globallogicuki/provider-harbor:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-harbor
spec:
  package: globallogicuki/provider-harbor:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/globallogicuki/provider-harbor).

## Provider Config
Note, that the ProviderConfig uses basic auth and requires a local user. A robot user can be used, 
but has restrictions around what can be created (e.g. a Robot cannot create other Robots).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/globallogicuki/provider-harbor/issues).
