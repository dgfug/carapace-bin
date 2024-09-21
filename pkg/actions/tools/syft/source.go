package syft

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/carapace-sh/carapace/pkg/condition"
)

// ActionSources completes sources
//
//	./local/file
//	alpine:3.6
func ActionSources() carapace.Action {
	return carapace.Batch(
		carapace.ActionFiles(),
		docker.ActionRepositoryTags().Unless(condition.CompletingPath),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			splitted := strings.SplitN(c.Value, ":", 2)
			switch len(splitted) {
			case 1:
				return carapace.ActionValuesDescribed(
					"docker", "use images from the Docker daemon",
					"podman", "use images from the Podman daemon",
					"docker-archive", "use a tarball from disk for archives created from \"docker save\"",
					"oci-archive", "use a tarball from disk for OCI archives (from Skopeo or otherwise)",
					"oci-dir", "read directly from a path on disk for OCI layout directories (from Skopeo or otherwise)",
					"singularity", "read directly from a Singularity Image Format (SIF) container on disk",
					"dir", "read directly from a path on disk (any directory)",
					"file", "read directly from a path on disk (any single file)",
					"registry", "pull image directly from a registry (no container runtime required)",
				).Suffix(":").Tag("sources")
			case 2:
				prefix := splitted[0] + ":"
				c.Value = strings.TrimPrefix(c.Value, prefix)
				return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					// TODO podman, registry
					switch splitted[0] {
					case "docker":
						return docker.ActionRepositoryTags()

					case "docker-archive", "oci-archive", "singularity", "file":
						return carapace.ActionFiles()

					case "oci-dir", "dir":
						return carapace.ActionDirectories()

					default:
						return carapace.ActionValues()
					}
				}).Invoke(c).Prefix(prefix).ToA()

			default:
				return carapace.ActionValues()
			}
		}),
	).ToA()
}
