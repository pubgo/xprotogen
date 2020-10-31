package cmds

import (
	"github.com/pubgo/xerror"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Run() {
	rootCmd.Use = "xprotogen"
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return xerror.Wrap(rootCmd.Help())
	}
	xerror.Exit(rootCmd.Execute())
}
