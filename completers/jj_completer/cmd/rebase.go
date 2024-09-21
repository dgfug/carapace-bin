package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var rebaseCmd = &cobra.Command{
	Use:   "rebase [OPTIONS] --destination <DESTINATION>",
	Short: "Move revisions to different parent(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebaseCmd).Standalone()

	rebaseCmd.Flags().StringArray("after", []string{}, "Alias for `--insert-after`")
	rebaseCmd.Flags().StringArray("before", []string{}, "Alias for `--insert-before`")
	rebaseCmd.Flags().StringSliceP("branch", "b", []string{}, "Rebase the whole branch relative to destination's ancestors (can be repeated)")
	rebaseCmd.Flags().StringSliceP("destination", "d", []string{}, "The revision(s) to rebase onto (can be repeated to create a merge commit)")
	rebaseCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rebaseCmd.Flags().StringArrayP("insert-after", "A", []string{}, "Revision(s) to insert after")
	rebaseCmd.Flags().StringArrayP("insert-before", "B", []string{}, "Revision(s) to insert before")
	rebaseCmd.Flags().StringP("revisions", "r", "", "Rebase only this revision, rebasing descendants onto this revision's parent(s)")
	rebaseCmd.Flags().Bool("skip-emptied", false, "Abandon empty commits created by the rebase")
	rebaseCmd.Flags().StringSliceP("source", "s", []string{}, "Rebase specified revision(s) together their tree of descendants (can be repeated)")
	rebaseCmd.MarkFlagRequired("destination")
	rootCmd.AddCommand(rebaseCmd)

	carapace.Gen(rebaseCmd).FlagCompletion(carapace.ActionMap{
		"after":         jj.ActionRevs(jj.RevOption{}.Default()),
		"before":        jj.ActionRevs(jj.RevOption{}.Default()),
		"branch":        jj.ActionRevs(jj.RevOption{LocalBranches: true, RemoteBranches: true, Tags: true}),
		"destination":   jj.ActionRevs(jj.RevOption{}.Default()),
		"insert-after":  jj.ActionRevs(jj.RevOption{}.Default()),
		"insert-before": jj.ActionRevs(jj.RevOption{}.Default()),
		"revisions":     jj.ActionRevs(jj.RevOption{}.Default()),
		"source":        jj.ActionRevs(jj.RevOption{}.Default()),
	})
}
