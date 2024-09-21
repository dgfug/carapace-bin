package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
	"github.com/carapace-sh/carapace/pkg/style"
)

func init() {
	knownVariables["docker"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("docker"),
			Variables: map[string]string{
				"DOCKER_API_VERSION":          "Override the negotiated API version to use for debugging",
				"DOCKER_CERT_PATH":            "Location of your authentication keys",
				"DOCKER_CONFIG":               "The location of your client configuration files",
				"DOCKER_CONTENT_TRUST_SERVER": "The URL of the Notary server to use",
				"DOCKER_CONTENT_TRUST":        "When set Docker uses notary to sign and verify images",
				"DOCKER_CONTEXT":              "Name of the docker context to use",
				"DOCKER_DEFAULT_PLATFORM":     "Default platform for commands that take the --platform flag",
				"DOCKER_HIDE_LEGACY_COMMANDS": "When set, Docker hides \"legacy\" top-level commands",
				"DOCKER_HOST":                 "Daemon socket to connect to",
				"DOCKER_TLS_VERIFY":           "When set Docker uses TLS and verifies the remote",
				"BUILDKIT_PROGRESS":           "Set type of progress output",
			},
			VariableCompletion: map[string]carapace.Action{
				"DOCKER_API_VERSION":      carapace.ActionValues(),
				"DOCKER_CERT_PATH":        carapace.ActionDirectories(),
				"DOCKER_CONFIG":           carapace.ActionFiles(),
				"DOCKER_CONTENT_TRUST":    carapace.ActionValues("1"),
				"DOCKER_CONTEXT":          docker.ActionContexts(),
				"DOCKER_DEFAULT_PLATFORM": carapace.ActionValues(), // TODO
				"DOCKER_HIDE_LEGACY_COMMANDS": carapace.ActionStyledValuesDescribed(
					"0", "show", style.Carapace.KeywordNegative,
					"1", "hide", style.Carapace.KeywordPositive,
				),
				"DOCKER_TLS_VERIFY": carapace.ActionValues("1"),
				"BUILDKIT_PROGRESS": carapace.ActionValues("auto", "plain", "tty"),
			},
		}
	}
}
