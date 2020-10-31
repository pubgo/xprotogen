package cmds

import (
	"encoding/json"
	"fmt"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xprotogen/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(string(xerror.PanicBytes(json.MarshalIndent(map[string]interface{}{
				"build_time": version.BuildTime,
				"version":    version.Version,
				"go_version": version.GoVersion,
				"go_path":    version.GoPath,
				"go_root":    version.GoROOT,
				"commit_id":  version.CommitID,
				"project":    version.Project,
			}, "", "  "))))
		},
	})
}
