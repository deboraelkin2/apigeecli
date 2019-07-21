package getsf

import (
	"../../shared"
	"github.com/spf13/cobra"
	"net/url"
	"path"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a shared flow by name",
	Long:  "Gets a shared flow by name, including a list of its revisions.",
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := url.Parse(shared.BaseURL)
		u.Path = path.Join(u.Path, shared.RootArgs.Org, "sharedflows", name)
		_ = shared.HttpClient(u.String())
	},
}

var name string

func init() {
	Cmd.Flags().StringVarP(&name, "name", "n",
		"", "Shared flow name")

	_ = Cmd.MarkFlagRequired("name")

}
