package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_deleteCmd = &cobra.Command{
	Use:   "delete <id> [flags]",
	Short: "Delete a CI/CD pipeline",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_deleteCmd).Standalone()

	ci_deleteCmd.Flags().Bool("dry-run", false, "simulate process, but do not delete anything")
	ci_deleteCmd.Flags().StringP("status", "s", "", "delete pipelines by status: {running|pending|success|failed|canceled|skipped|created|manual}")
	ciCmd.AddCommand(ci_deleteCmd)

	carapace.Gen(ci_deleteCmd).FlagCompletion(carapace.ActionMap{
		"status": carapace.ActionValues("running", "pending", "success", "failed", "canceled", "skipped", "created", "manual"),
	})

	carapace.Gen(ci_deleteCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPipelines(ci_deleteCmd, ci_deleteCmd.Flag("status").Value.String()).UniqueList(",")
		}),
	)
}
