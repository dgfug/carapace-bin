package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var join_phase_preflightCmd = &cobra.Command{
	Use:   "preflight",
	Short: "Run join pre-flight checks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_preflightCmd).Standalone()
	join_phase_preflightCmd.Flags().String("apiserver-advertise-address", "", "If the node should host a new control plane instance, the IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	join_phase_preflightCmd.Flags().Int32("apiserver-bind-port", 6443, "If the node should host a new control plane instance, the port for the API Server to bind to.")
	join_phase_preflightCmd.Flags().String("certificate-key", "", "Use this key to decrypt the certificate secrets uploaded by init.")
	join_phase_preflightCmd.Flags().String("config", "", "Path to kubeadm config file.")
	join_phase_preflightCmd.Flags().Bool("control-plane", false, "Create a new control plane instance on this node")
	join_phase_preflightCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	join_phase_preflightCmd.Flags().String("discovery-file", "", "For file-based discovery, a file or URL from which to load cluster information.")
	join_phase_preflightCmd.Flags().String("discovery-token", "", "For token-based discovery, the token used to validate cluster information fetched from the API server.")
	join_phase_preflightCmd.Flags().StringSlice("discovery-token-ca-cert-hash", []string{}, "For token-based discovery, validate that the root CA public key matches this hash (format: \"<type>:<value>\").")
	join_phase_preflightCmd.Flags().Bool("discovery-token-unsafe-skip-ca-verification", false, "For token-based discovery, allow joining without --discovery-token-ca-cert-hash pinning.")
	join_phase_preflightCmd.Flags().StringSlice("ignore-preflight-errors", []string{}, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	join_phase_preflightCmd.Flags().String("node-name", "", "Specify the node name.")
	join_phase_preflightCmd.Flags().String("tls-bootstrap-token", "", "Specify the token used to temporarily authenticate with the Kubernetes Control Plane while joining the node.")
	join_phase_preflightCmd.Flags().String("token", "", "Use this token for both discovery-token and tls-bootstrap-token when those values are not provided.")
	join_phaseCmd.AddCommand(join_phase_preflightCmd)

	carapace.Gen(join_phase_preflightCmd).FlagCompletion(carapace.ActionMap{
		"config":         carapace.ActionFiles(),
		"discovery-file": carapace.ActionFiles(),
		"ignore-preflight-errors": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return action.ActionChecks().Invoke(c).Filter(c.Parts).ToA().NoSpace()
			})
		}),
	})
}
