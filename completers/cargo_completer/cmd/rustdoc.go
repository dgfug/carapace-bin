package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rustdocCmd = &cobra.Command{
	Use:   "rustdoc",
	Short: "Build a package's documentation, using specified custom flags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rustdocCmd).Standalone()

	rustdocCmd.Flags().Bool("all-features", false, "Activate all available features")
	rustdocCmd.Flags().Bool("all-targets", false, "Build all targets")
	rustdocCmd.Flags().StringSlice("bench", []string{}, "Build only the specified bench target")
	rustdocCmd.Flags().Bool("benches", false, "Build all benches")
	rustdocCmd.Flags().StringSlice("bin", []string{}, "Build only the specified binary")
	rustdocCmd.Flags().Bool("bins", false, "Build all binaries")
	rustdocCmd.Flags().StringSlice("example", []string{}, "Build only the specified example")
	rustdocCmd.Flags().Bool("examples", false, "Build all examples")
	rustdocCmd.Flags().StringSliceP("features", "F", []string{}, "Space or comma separated list of features to activate")
	rustdocCmd.Flags().BoolP("help", "h", false, "Print help")
	rustdocCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	rustdocCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	rustdocCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error (unstable)")
	rustdocCmd.Flags().Bool("lib", false, "Build only this package's library")
	rustdocCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	rustdocCmd.Flags().StringSlice("message-format", []string{}, "Error format")
	rustdocCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	rustdocCmd.Flags().Bool("open", false, "Opens the docs in a browser after the operation")
	rustdocCmd.Flags().StringP("package", "p", "", "Package to document")
	rustdocCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	rustdocCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	rustdocCmd.Flags().BoolP("release", "r", false, "Build artifacts in release mode, with optimizations")
	rustdocCmd.Flags().StringSlice("target", []string{}, "Build for the target triple")
	rustdocCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	rustdocCmd.Flags().StringSlice("test", []string{}, "Build only the specified test target")
	rustdocCmd.Flags().Bool("tests", false, "Build all tests")
	rustdocCmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	rustdocCmd.Flags().Bool("unit-graph", false, "Output build graph in JSON (unstable)")
	rustdocCmd.Flag("timings").NoOptDefVal = " "
	rootCmd.AddCommand(rustdocCmd)

	// TODO flag completion
	carapace.Gen(rustdocCmd).FlagCompletion(carapace.ActionMap{
		"bench":          action.ActionTargets(rustdocCmd, action.TargetOpts{Bench: true}),
		"bin":            action.ActionTargets(rustdocCmd, action.TargetOpts{Bin: true}),
		"example":        action.ActionTargets(rustdocCmd, action.TargetOpts{Example: true}),
		"features":       action.ActionFeatures(rustdocCmd).UniqueList(","),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(rustdocCmd, false),
		"profile":        action.ActionProfiles(rustdocCmd),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(rustdocCmd, action.TargetOpts{Test: true}),
	})
}
