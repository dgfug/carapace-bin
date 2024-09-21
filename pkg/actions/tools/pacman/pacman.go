// package pacman contains arch package manager related actions
package pacman

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionPackages completes installed packages
//
//	gnome-common
//	gstreamer
func ActionPackages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("pacman", "-Qs", "^"+c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<group>[^/]+)/(?P<name>[^ ]+) (?P<version>[^ ]+)(?P<context>.*)$`)

			vals := make([]string, 0)
			for i := 0; i < len(lines)-1; i += 2 {
				if matches := r.FindStringSubmatch(lines[i]); matches != nil {
					s := style.Default
					if strings.Contains(matches[4], "[installed]") {
						s = style.Blue
					}
					vals = append(vals, matches[2], strings.TrimSpace(lines[i+1]), s)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}

// ActionPackageGroups completes package group names
//
//	i3-manjaro
//	linux49-extramodules
func ActionPackageGroups() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("sh", "-c", "pacman -Qg | cut -d' ' -f1 | sort |  uniq")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
