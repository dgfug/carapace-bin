package cmd

import (
	"github.com/rsteube/carapace-bin/completers/halt_completer/cmd"
)

/**
Description for go:generate
	Use: "poweroff",
	Short: "poweroff the machine",
*/

func Execute() error {
	return cmd.ExecutePoweroff()
}
