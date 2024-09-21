package nix

import (
	"fmt"
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionChannelPackages completes local channels and their packages
//
//	nixos.g++
//	nixos.gacutil
func ActionChannelPackages() carapace.Action {
	return carapace.ActionMultiParts(".", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionLocalChannels().Invoke(c).Suffix(".").ToA()

		case 1:
			return ActionPackages(c.Parts[0])

		default:
			return carapace.ActionValues()
		}
	})
}

// ActionPackages completes packages for given channel
//
//	g++
//	gacutil
func ActionPackages(channel string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if channel == "" {
			channel = "nixos"
		}

		path, err := c.Abs(fmt.Sprintf("~/.nix-defexpr/channels/%v/programs.sqlite", channel))
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		if _, err := os.Stat(path); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		if c.Value == "" {
			return carapace.ActionMessage("search needs at least 1 character")
		}

		query := fmt.Sprintf(`SELECT DISTINCT package FROM Programs WHERE package LIKE '%v%%'`, c.Value) // TODO filter by system (architecture)
		return carapace.ActionExecCommand("sqlite3", path, query)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
