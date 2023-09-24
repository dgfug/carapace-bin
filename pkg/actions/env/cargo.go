package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
	"github.com/rsteube/carapace-bin/pkg/conditions"
	"github.com/rsteube/carapace/pkg/style"
)

func init() {
	_bool := carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
	knownVariables["cargo"] = variables{
		Condition: conditions.ConditionPath("cargo"),
		Variables: map[string]string{
			"CARGO_BIN_NAME":                         "The name of the binary that is currently being compiled",
			"CARGO_BUILD_DEP_INFO_BASEDIR":           "Dep-info relative directory, see build.dep-info-basedir",
			"CARGO_BUILD_INCREMENTAL":                "Incremental compilation, see build.incremental",
			"CARGO_BUILD_JOBS":                       "Number of parallel jobs, see build.jobs",
			"CARGO_BUILD_RUSTC":                      "The rustc executable, see build.rustc",
			"CARGO_BUILD_RUSTC_WORKSPACE_WRAPPER":    "The rustc wrapper for workspace members only, see build.rustc-workspace-wrapper",
			"CARGO_BUILD_RUSTC_WRAPPER":              "The rustc wrapper, see build.rustc-wrapper",
			"CARGO_BUILD_RUSTDOCFLAGS":               "Extra rustdoc flags, see build.rustdocflags",
			"CARGO_BUILD_RUSTDOC":                    "The rustdoc executable, see build.rustdoc",
			"CARGO_BUILD_RUSTFLAGS":                  "Extra rustc flags, see build.rustflags",
			"CARGO_BUILD_TARGET_DIR":                 "The default output directory, see build.target-dir",
			"CARGO_BUILD_TARGET":                     "The default target platform, see build.target",
			"CARGO_CACHE_RUSTC_INFO":                 "If this is set to 0 then Cargo will not try to cache compiler version information",
			"CARGO_CARGO_NEW_VCS":                    "The default source control system with cargo new, see cargo-new.vcs",
			"CARGO_CFG_TARGET_ARCH":                  "The CPU target architecture",
			"CARGO_CFG_TARGET_ENDIAN":                "The CPU target endianness",
			"CARGO_CFG_TARGET_ENV":                   "The target environment ABI",
			"CARGO_CFG_TARGET_FAMILY":                "The target family",
			"CARGO_CFG_TARGET_FEATURE":               "List of CPU target features enabled",
			"CARGO_CFG_TARGET_OS":                    "The target operating system",
			"CARGO_CFG_TARGET_POINTER_WIDTH":         "The CPU pointer width",
			"CARGO_CFG_TARGET_VENDOR":                "The target vendor",
			"CARGO_CFG_UNIX":                         "Set on unix-like platforms",
			"CARGO_CFG_WINDOWS":                      "Set on windows-like platforms",
			"CARGO_CRATE_NAME":                       "The name of the crate that is currently being compiled",
			"CARGO_ENCODED_RUSTDOCFLAGS":             "A list of custom flags separated by 0x1f (ASCII Unit Separator) to pass to all rustdoc invocations that Cargo performs",
			"CARGO_ENCODED_RUSTFLAGS":                "A list of custom flags separated by 0x1f (ASCII Unit Separator) to pass to all compiler invocations that Cargo performs",
			"CARGO_FUTURE_INCOMPAT_REPORT_FREQUENCY": "How often we should generate a future incompat report notification, see future-incompat-report.frequency",
			"CARGO_HOME":                             "Cargo maintains a local cache of the registry index and of git checkouts of crates",
			"CARGO_HTTP_CAINFO":                      "The TLS certificate Certificate Authority file, see http.cainfo",
			"CARGO_HTTP_CHECK_REVOKE":                "Disables TLS certificate revocation checks, see http.check-revoke",
			"CARGO_HTTP_DEBUG":                       "Enables HTTP debugging, see http.debug",
			"CARGO_HTTP_LOW_SPEED_LIMIT":             "The HTTP low-speed limit, see http.low-speed-limit",
			"CARGO_HTTP_MULTIPLEXING":                "Whether HTTP/2 multiplexing is used, see http.multiplexing",
			"CARGO_HTTP_PROXY":                       "Enables HTTP proxy, see http.proxy",
			"CARGO_HTTP_SSL_VERSION":                 "The TLS version to use, see http.ssl-version",
			"CARGO_HTTP_TIMEOUT":                     "The HTTP timeout, see http.timeout",
			"CARGO_HTTP_USER_AGENT":                  "The HTTP user-agent header, see http.user-agent",
			"CARGO_INCREMENTAL":                      "If this is set to 1 then Cargo will force incremental compilation to be enabled for the current compilation, and when set to 0 it will force disabling it",
			"CARGO_INSTALL_ROOT":                     "The default directory for cargo install, see install.root",
			"CARGO_LOG":                              "Cargo uses the env_logger crate to display debug log messages",
			"CARGO_MAKEFLAGS":                        "Contains parameters needed for Cargo’s jobserver implementation to parallelize subprocesses",
			"CARGO_MANIFEST_DIR":                     "The directory containing the manifest of your package",
			"CARGO_MANIFEST_LINKS":                   "the manifest links value",
			"CARGO_NET_GIT_FETCH_WITH_CLI":           "Enables the use of the git executable to fetch, see net.git-fetch-with-cli",
			"CARGO_NET_OFFLINE":                      "Offline mode, see net.offline",
			"CARGO_NET_RETRY":                        "Number of times to retry network errors, see net.retry",
			"CARGO_PKG_AUTHORS":                      "Colon separated list of authors from the manifest of your package",
			"CARGO_PKG_DESCRIPTION":                  "The description from the manifest of your package",
			"CARGO_PKG_HOMEPAGE":                     "The home page from the manifest of your package",
			"CARGO_PKG_LICENSE_FILE":                 "The license file from the manifest of your package",
			"CARGO_PKG_LICENSE":                      "The license from the manifest of your package",
			"CARGO_PKG_NAME":                         "The name of your package",
			"CARGO_PKG_README":                       "Path to the README file of your package",
			"CARGO_PKG_REPOSITORY":                   "The repository from the manifest of your package",
			"CARGO_PKG_RUST_VERSION":                 "The Rust version from the manifest of your package",
			"CARGO_PKG_VERSION_MAJOR":                "The major version of your package",
			"CARGO_PKG_VERSION_MINOR":                "The minor version of your package",
			"CARGO_PKG_VERSION_PATCH":                "The patch version of your package",
			"CARGO_PKG_VERSION_PRE":                  "The pre-release version of your package",
			"CARGO_PKG_VERSION":                      "The full version of your package",
			"CARGO_PRIMARY_PACKAGE":                  "This environment variable will be set if the package being built is primary",
			"CARGO_REGISTRY_DEFAULT":                 "Default registry for the --registry flag, see registry.default",
			"CARGO_REGISTRY_TOKEN":                   "Authentication token for crates.io, see registry.token",
			"CARGO_TARGET_DIR":                       "Location of where to place all generated artifacts, relative to the current working directory",
			"CARGO_TARGET_TMPDIR":                    "Only set when building integration test or benchmark code",
			"CARGO_TERM_COLOR":                       "The default color mode, see term.color",
			"CARGO_TERM_PROGRESS_WHEN":               "The default progress bar showing mode, see term.progress.when",
			"CARGO_TERM_PROGRESS_WIDTH":              "The default progress bar width, see term.progress.width",
			"CARGO_TERM_QUIET":                       "Quiet mode, see term.quiet",
			"CARGO_TERM_VERBOSE":                     "The default terminal verbosity, see term.verbose",
		},
		VariableCompletion: map[string]carapace.Action{
			"CARGO_BUILD_DEP_INFO_BASEDIR": carapace.ActionDirectories(),
			"CARGO_LOG":                    carapace.ActionValues("debug", "info", "warn", "error", "trace").StyleF(style.ForLogLevel),
			"CARGO_HOME":                   carapace.ActionDirectories(),
			"CARGO_TARGET_DIR":             carapace.ActionDirectories(),
			"CARGO_INCREMENTAL": carapace.ActionStyledValuesDescribed(
				"0", "force disabled,", style.Red,
				"1", "force enabled", style.Green,
			),
			"CARGO_CARGO_NEW_VCS":   carapace.ActionValues("git", "hg", "pijul", "fossil", "none"),
			"CARGO_HTTP_USER_AGENT": http.ActionUserAgents(),
			"CARGO_TERM_COLOR":      carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
			"CARGO_TERM_QUIET":      _bool,
			"CARGO_TERM_VERBOSE":    _bool,
			// TODO more completions
		},
	}

}
