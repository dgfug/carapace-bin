package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
)

func init() {
	knownVariables["maven"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("mvn"),
			Variables: map[string]string{
				"MAVEN_OPTS": "parameters used to start up the JVM running Maven",
				"MAVEN_ARGS": "arguments passed to Maven before CLI arguments",
			},
			VariableCompletion: map[string]carapace.Action{
				"MAVEN_OPTS": bridge.ActionCarapaceBin("java").Split(),
				"MAVEN_ARGS": bridge.ActionCarapaceBin("mvn").Split(),
			},
		}
	}
}
