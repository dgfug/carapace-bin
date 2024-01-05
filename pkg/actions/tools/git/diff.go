package git

import (
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionDiffAlgorithms completes diff algorithms
//
//	myers (The basic greedy diff algorithm)
//	minimal (Spend extra time to make sure the smallest possible diff is produced)
func ActionDiffAlgorithms() carapace.Action {
	return carapace.ActionValuesDescribed(
		"myers", "The basic greedy diff algorithm",
		"minimal", "Spend extra time to make sure the smallest possible diff is produced",
		"patience", "Use patience diff algorithm when generating patches",
		"histogram", "This algorithm extends the patience algorithm to support low-occurrence common elements",
	)
}

// ActionColorMovedModes completes color moved modes
//
//	no (Moved lines are not highlighted)
//	default (default mode)
func ActionColorMovedModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"no", "Moved lines are not highlighted",
		"default", "default mode",
		"plain", "plain mode",
		"blocks", "greedily detects blocks",
		"zebra", "Blocks of moved text are detected as in blocks mode",
		"dimmed-zebra", "Similar to zebra, but additional dimming of uninteresting parts of moved code",
	)
}

// ActionColorMovedWsModes completed color moved whitespace modes
//
//	no (Do not ignore whitespace when performing move detection)
//	ignore-space-at-eol (Ignore changes in whitespace at EOL)
func ActionColorMovedWsModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"no", "Do not ignore whitespace when performing move detection",
		"ignore-space-at-eol", "Ignore changes in whitespace at EOL",
		"ignore-space-change", "Ignore changes in amount of whitespace.",
		"ignore-all-space", "Ignore whitespace when comparing lines.",
		"allow-indentation-change", "Initially ignore any whitespace in the move detection",
	)
}

// ActionWordDiffModes completes word diff modes
//
//	no (Do not ignore whitespace when performing move detection)
//	ignore-space-at-eol (Ignore changes in whitespace at EOL)
func ActionWordDiffModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"color", "Highlight changed words using only colors",
		"plain", "Show words as [-removed-] and {+added+}",
		"porcelain", "Use a special line-based format intended for script consumption",
		"none", "Disable word diff again",
	)
}

// ActionWsErrorHighlightModes completes whitespace error highlight modes
//
//	context (context lines)
//	old (old lines)
func ActionWsErrorHighlightModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"context", "context lines",
		"old", "old lines",
		"new", "new linex",
		"none", "reset previous values",
		"default", "reset to new",
		"all", "shorthand for old,new,context",
	)
}

// ActionDiffTools completes diff tools
//
//	meld
//	nvimdiff
func ActionDiffTools() carapace.Action {
	return carapace.ActionValues(
		"araxis",
		"bc",
		"codecompare",
		"deltawalker",
		"diffmerge",
		"diffuse",
		"ecmerge",
		"emerge",
		"examdiff",
		"guiffy",
		"gvimdiff",
		"kdiff3",
		"kompare",
		"meld",
		"nvimdiff",
		"opendiff",
		"p4merge",
		"smerge",
		"tkdiff",
		"vimdiff",
		"winmerge",
		"xxdiff",
	)
}

// ActionRefDiffs completes changes beetween refs
// Accepts up to two refs
// 0: compare current workspace to HEAD
// 1: compare current workspace to given ref
// 2: compare first ref to second ref
func ActionRefDiffs(refs ...string) carapace.Action {
	return actionRefDiffs(false, refs...)
}

// ActionCachedDiffs completes changes between stage and given ref
func ActionCachedDiffs(ref string) carapace.Action {
	return actionRefDiffs(true, ref)
}

func actionRefDiffs(cached bool, refs ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"diff", "--name-status"}
		if cached {
			args = append(args, "--cached")
			if len(refs) != 1 {
				return carapace.ActionMessage("only up to two refs allowed [ActionCachedDiffs]")
			}
		}

		switch len(refs) {
		case 0:
			args = append(args, "HEAD")
		case 1:
			args = append(args, refs[0])
		case 2:
			args = append(args, refs[0], refs[1])
		default:
			return carapace.ActionMessage("only up to two refs allowed [ActionRefDiffs]")
		}

		return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			root, err := rootDir(c)
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, line := range lines {
				splitted := strings.Split(line, "\t")
				if len(splitted) < 2 {
					continue
				}

				relativePath, err := filepath.Rel(c.Dir, root+"/"+splitted[1])
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}

				switch {
				case strings.HasPrefix(relativePath, "../"):
					vals = append(vals, relativePath, splitted[0])
				case strings.HasPrefix(c.Value, "."):
					vals = append(vals, "./"+relativePath, splitted[0])
				default:
					vals = append(vals, relativePath, splitted[0])
				}

			}
			return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPathExt)
		})
	}).Tag("changed files")
}
