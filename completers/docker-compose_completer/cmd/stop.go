package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [OPTIONS] [SERVICE...]",
	Short: "Stop services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()

	stopCmd.Flags().StringP("timeout", "t", "", "Specify a shutdown timeout in seconds")
	rootCmd.AddCommand(stopCmd)

	carapace.Gen(stopCmd).PositionalAnyCompletion(
		action.ActionServices(stopCmd).FilterArgs(),
	)
}
