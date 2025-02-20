package cmd

//go:generate go run ../../generate/gen.go

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rsteube/carapace/pkg/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "carapace [flags] [COMPLETER] [bash|elvish|fish|oil|powershell|tcsh|xonsh|zsh]",
	Long: "multi-shell multi-command argument completer",
	Example: fmt.Sprintf(`  Single completer:
    bash:       source <(carapace chmod bash)
    elvish:     eval (carapace chmod elvish | slurp)
    fish:       carapace chmod fish | source
    nushell:    carapace chmod | save chmod.nu ; nu -c 'source chmod.nu'
    oil:        source <(carapace chmod oil)
    powershell: carapace chmod powershell | Out-String | Invoke-Expression
    tcsh:       eval `+"`"+`carapace _chmod tcsh`+"`"+`
    xonsh:      exec($(carapace chmod xonsh))
    zsh:        source <(carapace chmod zsh)

  All completers:
    bash:       source <(carapace _carapace bash)
    elvish:     eval (carapace _carapace elvish | slurp)
    fish:       carapace _carapace fish | source
    nushell:    carapace _carapace | save carapace.nu ; nu -c 'source carapace.nu'
    oil:        source <(carapace _carapace oil)
    powershell: carapace _carapace powershell | Out-String | Invoke-Expression
    tcsh:       eval `+"`"+`carapace _carapace tcsh`+"`"+`
    xonsh:      exec($(carapace _carapace xonsh))
    zsh:        source <(carapace _carapace zsh)

  Shell parameter is optional and if left out carapace will try to detect it by parent process name.
  Some completions are cached at [%v/carapace].
`, os.TempDir()),
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: completers,
	Run: func(cmd *cobra.Command, args []string) {
		// since flag parsing is disabled do this manually
		switch args[0] {
		case "-h":
			cmd.Help()
		case "--help":
			cmd.Help()
		case "-v":
			println(cmd.Version)
		case "--version":
			println(cmd.Version)
		case "--list":
			printCompleters()
		case "_carapace":
			shell := ps.DetermineShell()
			if len(args) > 1 {
				shell = args[1]
			}
			switch shell {
			case "bash":
				fmt.Println(bash_lazy(completers))
			case "elvish":
				fmt.Println(elvish_lazy(completers))
			case "fish":
				fmt.Println(fish_lazy(completers))
			case "nushell":
				fmt.Println(nushell_lazy(completers))
			case "oil":
				fmt.Println(oil_lazy(completers))
			case "powershell":
				fmt.Println(powershell_lazy(completers))
			case "tcsh":
				fmt.Println(tcsh_lazy(completers))
			case "xonsh":
				fmt.Println(xonsh_lazy(completers))
			case "zsh":
				fmt.Println(zsh_lazy(completers))
			default:
				fmt.Fprintln(os.Stderr, "could not determine shell")
			}
		default:
			invokeCompleter(args[0])
		}

	},
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
	DisableFlagParsing: true,
}

func printCompleters() {
	maxlen := 0
	for _, name := range completers {
		if len := len(name); len > maxlen {
			maxlen = len
		}
	}

	for _, name := range completers {
		fmt.Printf("%-"+strconv.Itoa(maxlen)+"v %v\n", name, descriptions[name])
	}
}

func invokeCompleter(completer string) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	os.Args[1] = "_carapace"
	executeCompleter(completer)

	w.Close()
	out := <-outC
	os.Stdout = old

	executable, err := os.Executable()
	if err != nil {
		panic(err.Error()) // TODO exit with error message
	}
	executableName := filepath.Base(executable)
	patched := strings.Replace(string(out), fmt.Sprintf("%v _carapace", executableName), fmt.Sprintf("%v %v", executableName, completer), -1)      // general callback
	patched = strings.Replace(patched, fmt.Sprintf("'%v', '_carapace'", executableName), fmt.Sprintf("'%v', '%v'", executableName, completer), -1) // xonsh callback
	fmt.Print(patched)

}

func Execute(version string) error {
	rootCmd.Version = version
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().Bool("list", false, "list completers")
}
