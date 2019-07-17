package delsf

import (
	"../../shared"
	"github.com/spf13/cobra"
	"net/url"
	"path"
)

var Cmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a shared flow",
	Long:  "Deletes a shared flow and all associated policies, resources, and revisions. The flow must be undeployed first.",
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := url.Parse(shared.BaseURL)
		u.Path = path.Join(u.Path, shared.RootArgs.Org, "sharedflows", name)
		shared.HttpClient(u.String(), "", "DELETE")
	},
}

var name string

func init() {
	Cmd.Flags().StringVarP(&name, "name", "n",
		"", "shared flow name")

	Cmd.MarkFlagRequired("name")

}