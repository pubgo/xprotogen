package cmds

import (
	"github.com/spf13/cobra"
)

const tplDir = "tpl"

func InitCmd() *cobra.Command {
	var initCmd = &cobra.Command{}
	return initCmd
}

// .xprotogen
func init() {
	rootCmd.AddCommand(InitCmd())
}
