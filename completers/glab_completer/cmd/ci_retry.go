package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ci_retryCmd = &cobra.Command{
	Use:   "retry <job-id>",
	Short: "Retry a CI/CD job",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_retryCmd).Standalone()

	ciCmd.AddCommand(ci_retryCmd)

	// TODO positional completion
}
