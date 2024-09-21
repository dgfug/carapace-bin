package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var branch_listCmd = &cobra.Command{
	Use:     "list [OPTIONS] [NAMES]...",
	Short:   "List branches and their targets",
	Aliases: []string{"l"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_listCmd).Standalone()

	branch_listCmd.Flags().BoolP("all-remotes", "a", false, "Show all tracking and non-tracking remote branches including the ones whose targets are synchronized with the local branches")
	branch_listCmd.Flags().BoolP("conflicted", "c", false, "Show only conflicted branches")
	branch_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branch_listCmd.Flags().StringSliceP("revisions", "r", []string{}, "Show branches whose local targets are in the given revisions")
	branch_listCmd.Flags().StringP("template", "T", "", "Render each branch using the given template")

	branch_listCmd.MarkFlagsMutuallyExclusive("all-remotes", "conflicted")

	branchCmd.AddCommand(branch_listCmd)

	carapace.Gen(branch_listCmd).FlagCompletion(carapace.ActionMap{
		"revisions": jj.ActionRevSets(jj.RevOption{}.Default()).UniqueList(","),
	})

	carapace.Gen(branch_listCmd).PositionalAnyCompletion(
		jj.ActionLocalBranches().FilterArgs(),
	)
}
