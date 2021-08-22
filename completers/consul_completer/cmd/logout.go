package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Destroy a Consul token created with login",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()
	addClientFlags(logoutCmd)
	addServerFlags(logoutCmd)

	rootCmd.AddCommand(logoutCmd)
}
