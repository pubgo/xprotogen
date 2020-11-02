package cmds

import (
	"github.com/pubgo/xerror"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

const tplDir = "tpl"
const cfg = ".xprotogen"

func InitCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "init",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			defer xerror.RespErr(&err)
			// 根目录创建一个配置文件
			// 创建目录
			// 创建Makefile
			// 写入配置文件
			//可以考虑prompt的方式

			// 创建配置文件
			xerror.Panic(ioutil.WriteFile(cfg, nil, 0644))

			// 创建目录
			xerror.Panic(os.MkdirAll(tplDir, 0755))

			// 创建main.go

			return nil
		},
	}

	return cmd
}

// .xprotogen
func init() {
	rootCmd.AddCommand(InitCmd())
}
