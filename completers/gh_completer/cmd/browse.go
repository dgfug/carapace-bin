package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action/git"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:   "browse [<number> | <path>]",
	Short: "Open the repository in the browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	browseCmd.Flags().StringP("branch", "b", "", "Select another branch by passing in the branch name")
	browseCmd.Flags().BoolP("commit", "c", false, "Open the last commit")
	browseCmd.Flags().BoolP("no-browser", "n", false, "Print destination URL instead of opening the browser")
	browseCmd.Flags().BoolP("projects", "p", false, "Open repository projects")
	browseCmd.Flags().BoolP("settings", "s", false, "Open repository settings")
	browseCmd.Flags().BoolP("wiki", "w", false, "Open repository wiki")
	browseCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(browseCmd)

	carapace.Gen(browseCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(browseCmd), // TODO merge with tags
		"repo":   action.ActionRepoOverride(browseCmd),
	})

	carapace.Gen(browseCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if browseCmd.Flag("projects").Changed ||
				browseCmd.Flag("settings").Changed ||
				browseCmd.Flag("wiki").Changed {
				return carapace.ActionValues()
			}

			ref := browseCmd.Flag("branch").Value.String()
			if browseCmd.Flag("commit").Changed {
				commit, err := git.LastCommit()
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				ref = commit.Sha
			}

			path := filepath.Dir(c.CallbackValue)
			path = strings.TrimPrefix(path, "/")
			if !strings.HasPrefix(c.CallbackValue, "/") {
				wd, err := os.Getwd()
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				root, err := util.FindReverse(wd, ".git")
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				rel, err := filepath.Rel(filepath.Dir(root), wd)
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				abs := fmt.Sprintf("%v/%v", rel, c.CallbackValue)
				r := regexp.MustCompile(`[^\/]+\/\.\.\/`)
				for {
					if match := r.FindString(abs); match != "" {
						abs = strings.Replace(abs, match, "", 1)
					} else {
						break
					}
				}
				path = filepath.Dir(abs)
			}
			// TODO merge with issue/pr ids (show all with prefix search using callbackvalue)
			return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
				return action.ActionContents(browseCmd, path, ref)
			})
		}),
	)
}
