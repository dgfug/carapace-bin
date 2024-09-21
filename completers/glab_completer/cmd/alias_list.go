package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alias_listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "List the available aliases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_listCmd).Standalone()

	aliasCmd.AddCommand(alias_listCmd)
}
