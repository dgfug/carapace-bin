// package fs contains filesystem related actions
package fs

import (
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
)

// ActionTarFileContents completes contents of given tar file
//   fileA
//   dirA/fileB
func ActionTarFileContents(file string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("tar", "--list", "--file", file)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}).Cache(24*time.Hour, cache.FileStats(file)).Invoke(c).ToMultiPartsA("/")
	})
}
