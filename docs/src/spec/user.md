# User

User defined [Specs] are automatically loaded by `carapace _carapace` from [`${UserConfigDir}/carapace/specs`](https://pkg.go.dev/os#UserConfigDir).

> Added files initially need a new shell to be started for it to be registered with `carapace _carapace`.
> Afterwards any change to it has an immediate effect.
>
> It is **mandatory** that the file name matches the name defined in the spec (e.g. `myspec.yaml` for `name: myspec`).

## Override

[Specs] override an internal completer with the same name. E.g. if the internal `kubectl` completer does not work as expected it can be [bridged](./bridge.md) instead:

```yaml
name: kubectl
description: kubectl controls the Kubernetes cluster manager
completion:
  positionalany: ["$carapace.bridge.Cobra([kubectl])"]
```

## JSON Schema

A [JSON Schema](http://json-schema.org/) can be used by adding the following header to the [Specs]:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
```

[Specs]:../spec.md
