package vagrant

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionSnapshots completes snapshots
func ActionSnapshots(machine string) carapace.Action {
	return carapace.ActionExecCommand("vagrant", "snapshot", "list", machine)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[1 : len(lines)-1]...)
	})
}
