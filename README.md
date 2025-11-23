# Provider Harbor

`provider-harbor` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) and exposes XRM-conformant managed resources 
for the [Harbor](https://goharbor.io/) API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/globallogicuki/provider-harbor):
```
up ctp provider install globallogicuki/provider-harbor:v0.2.2
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-harbor
spec:
  package: xpkg.upbound.io/globallogicuki/provider-harbor:v0.2.3
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/globallogicuki/provider-harbor).

## Provider Config
Note, that the ProviderConfig uses basic auth and requires a local user. A robot user can be used, 
but has restrictions around what can be created (e.g. a Robot cannot create other Robots).

## Local Development & Testing

The provider includes automated setup for local development and testing using a Kind cluster with Harbor.

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [helm](https://helm.sh/docs/intro/install/)
- [Go](https://golang.org/doc/install) 1.21+

### Quick Start

Set up a complete local development environment with one command:

```bash
make dev-setup
```

This will:
- ✅ Create a Kind cluster named `harbor`
- ✅ Install Harbor with Helm (accessible at http://localhost:8080)
- ✅ Install Crossplane
- ✅ Apply all provider CRDs
- ✅ Create a ProviderConfig with local Harbor credentials

**Default credentials:**
- URL: http://localhost:8080
- Username: `admin`
- Password: `Harbor12345`

### Running the Provider

Start the provider in the background:

```bash
make dev-run
```

The provider will run in the background with logs saved to `.work/dev/provider.log`.

**View logs:**
```bash
make dev-logs
```

**Stop the provider:**
```bash
make dev-stop
```

### Testing with Examples

Apply examples from the examples or examples-generated directories:

```bash
kubectl apply -f examples-generated/project/project.yaml
```

Check the resources:

```bash
kubectl get projects
kubectl describe project main
```

### Useful Development Commands

```bash
make dev-status           # Check status of all components
make dev-restart-forward  # Restart port-forwarding if it stops
make dev-cleanup          # Tear down everything (cluster, provider, configs)
make dev-help             # Show all available commands
```

### Manual Steps (if needed)

If you prefer manual control:

```bash
# Ensure correct kubectl context
kubectl config use-context kind-harbor

# Run provider in foreground (for debugging)
make run

# Or run in background manually
make run > /tmp/provider.log 2>&1 &
```

### Code Generation

When updating provider configuration or Terraform provider version:

```bash
make generate
```

### Building & Testing

```bash
# Build binary
make build

# Run tests
make test

# Build, push, and install (for CI/CD)
make all
```

### Troubleshooting

**Port-forward not working?**
```bash
make dev-restart-forward
```

**Provider can't connect to API server?**
```bash
# Ensure correct context
kubectl config use-context kind-harbor

# Restart provider
make dev-stop && make dev-run
```

**Harbor not accessible?**
```bash
# Check Harbor pods
kubectl get pods -n harbor

# Check port-forward is running
ps aux | grep port-forward
```

**Need to start fresh?**
```bash
make dev-cleanup
make dev-setup
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/globallogicuki/provider-harbor/issues).
