package xdotool

import (
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/number"
)

// ActionDesktops completes desktops
func ActionDesktops() carapace.Action {
	return carapace.ActionExecCommand("xdotool", "get_num_desktops")(func(output []byte) carapace.Action {
		num, err := strconv.Atoi(strings.Split(string(output), "\n")[0])
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return number.ActionRange(number.RangeOpts{Format: "%d", Start: 0, End: num - 1})
	})
}
