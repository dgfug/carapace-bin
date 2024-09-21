package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List merge requests",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_listCmd).Standalone()

	mr_listCmd.Flags().BoolP("all", "A", false, "Get all merge requests")
	mr_listCmd.Flags().StringSliceP("assignee", "a", []string{}, "Get only merge requests assigned to users")
	mr_listCmd.Flags().String("author", "", "Filter merge request by Author <username>")
	mr_listCmd.Flags().BoolP("closed", "c", false, "Get only closed merge requests")
	mr_listCmd.Flags().BoolP("draft", "d", false, "Filter by draft merge requests")
	mr_listCmd.PersistentFlags().StringP("group", "g", "", "Select a group/subgroup. This option is ignored if a repo argument is set.")
	mr_listCmd.Flags().StringSliceP("label", "l", []string{}, "Filter merge request by label <name>")
	mr_listCmd.Flags().BoolP("merged", "M", false, "Get only merged merge requests")
	mr_listCmd.Flags().StringP("milestone", "m", "", "Filter merge request by milestone <id>")
	mr_listCmd.Flags().Bool("mine", false, "Get only merge requests assigned to me")
	mr_listCmd.Flags().StringSlice("not-label", []string{}, "Filter merge requests by not having label <name>")
	mr_listCmd.Flags().BoolP("opened", "o", false, "Get only open merge requests")
	mr_listCmd.Flags().StringP("page", "p", "", "Page number")
	mr_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page")
	mr_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	mr_listCmd.Flags().StringSliceP("reviewer", "r", []string{}, "Get only merge requests with users as reviewer")
	mr_listCmd.Flags().String("search", "", "Filter by <string> in title and description")
	mr_listCmd.Flags().StringP("source-branch", "s", "", "Filter by source branch <name>")
	mr_listCmd.Flags().StringP("target-branch", "t", "", "Filter by target branch <name>")
	mr_listCmd.Flag("mine").Hidden = true
	mr_listCmd.Flag("opened").Hidden = true
	mrCmd.AddCommand(mr_listCmd)

	carapace.Gen(mr_listCmd).FlagCompletion(carapace.ActionMap{
		"assignee":      action.ActionProjectMembers(mr_listCmd).UniqueList(","),
		"author":        action.ActionUsers(mr_listCmd),
		"group":         action.ActionGroups(mr_listCmd),
		"label":         action.ActionLabels(mr_listCmd).UniqueList(","),
		"milestone":     action.ActionMilestones(mr_listCmd),
		"not-label":     action.ActionLabels(mr_listCmd).UniqueList(","),
		"repo":          action.ActionRepo(mr_listCmd),
		"reviewer":      action.ActionProjectMembers(mr_listCmd).UniqueList(","),
		"source-branch": action.ActionBranches(mr_listCmd),
		"target-branch": action.ActionBranches(mr_listCmd),
	})
}
