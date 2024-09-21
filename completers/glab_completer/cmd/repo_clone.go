package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_cloneCmd = &cobra.Command{
	Use:   "clone <repo> [flags] [<dir>] [-- [<gitflags>...]]",
	Short: "Clone a GitLab repository/project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_cloneCmd).Standalone()

	repo_cloneCmd.Flags().BoolP("archived", "a", false, "Limit by archived status. Used with --group flag")
	repo_cloneCmd.Flags().StringP("group", "g", "", "Specify group to clone repositories from")
	repo_cloneCmd.Flags().BoolP("include-subgroups", "G", false, "Include projects in subgroups of this group. Default is true. Used with --group flag")
	repo_cloneCmd.Flags().BoolP("mine", "m", false, "Limit by projects in the group owned by the current authenticated user. Used with --group flag")
	repo_cloneCmd.Flags().String("page", "", "Page number")
	repo_cloneCmd.Flags().Bool("paginate", false, "Make additional HTTP requests to fetch all pages of projects before cloning. Respects --per-page")
	repo_cloneCmd.Flags().String("per-page", "", "Number of items to list per page")
	repo_cloneCmd.Flags().BoolP("preserve-namespace", "p", false, "Clone the repo in a subdirectory based on namespace")
	repo_cloneCmd.Flags().StringP("visibility", "v", "", "Limit by visibility {public, internal, or private}. Used with --group flag")
	repo_cloneCmd.Flags().BoolP("with-issues-enabled", "I", false, "Limit by projects with issues feature enabled. Default is false. Used with --group flag")
	repo_cloneCmd.Flags().BoolP("with-mr-enabled", "M", false, "Limit by projects with merge requests feature enabled. Default is false. Used with --group flag")
	repo_cloneCmd.Flags().BoolP("with-shared", "S", false, "Include projects shared to this group. Default is false. Used with --group flag")
	repoCmd.AddCommand(repo_cloneCmd)

	carapace.Gen(repo_cloneCmd).FlagCompletion(carapace.ActionMap{
		"group":      action.ActionGroups(repo_cloneCmd),
		"visibility": carapace.ActionValues("public", "internal", "private"),
	})

	carapace.Gen(repo_cloneCmd).PositionalCompletion(
		action.ActionRepo(repo_cloneCmd), // TODO not yet correct - also needs to consider flags being set
		carapace.ActionDirectories(),
	)
}
