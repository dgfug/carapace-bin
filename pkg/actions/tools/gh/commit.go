package gh

import "github.com/carapace-sh/carapace"

// ActionCommitFields completes commit fields
//
//	author
//	commit
func ActionCommitFields() carapace.Action {
	return carapace.ActionValues(
		"author",
		"commit",
		"committer",
		"sha",
		"id",
		"parents",
		"repository",
		"url",
	)
}
