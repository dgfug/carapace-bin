package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:   "secret <command>",
	Short: "Manage GitHub secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	secretCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(secretCmd)

	carapace.Gen(secretCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepoOverride(secretCmd),
	})
}
