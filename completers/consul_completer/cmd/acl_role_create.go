package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var acl_role_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an ACL role",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_role_createCmd).Standalone()
	addClientFlags(acl_roleCmd)
	addServerFlags(acl_roleCmd)

	acl_role_createCmd.Flags().String("description", "", "A description of the role")
	acl_role_createCmd.Flags().String("format", "", "Output format.")
	acl_role_createCmd.Flags().Bool("meta", false, "Indicates that role metadata such as the content hash and raft indices should be shown for each entry")
	acl_role_createCmd.Flags().String("name", "", "The new role's name.")
	acl_role_createCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_role_createCmd.Flags().String("node-identity", "", "Name of a node identity to use for this role.")
	acl_role_createCmd.Flags().String("policy-id", "", "ID of a policy to use for this role.")
	acl_role_createCmd.Flags().String("policy-name", "", "Name of a policy to use for this role.")
	acl_role_createCmd.Flags().String("service-identity", "", "Name of a service identity to use for this role.")
	acl_roleCmd.AddCommand(acl_role_createCmd)

	carapace.Gen(acl_role_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
		"node-identity": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionNodes(acl_role_createCmd)
			case 1:
				return action.ActionDatacenters(acl_role_createCmd)
			default:
				return carapace.ActionValues()
			}
		}),
		"service-identity": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionServices(acl_role_createCmd)
			case 1:
				return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
					return action.ActionDatacenters(acl_role_createCmd).Invoke(c).Filter(c.Parts).ToA()
				})
			default:
				return carapace.ActionValues()
			}
		}),
	})
}
