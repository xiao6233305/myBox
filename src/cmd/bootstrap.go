package cmd

// 执行安装命令
import (
	"fmt"
	"github.com/xiao6233305/mybox/src/operating"

	"github.com/spf13/cobra"
)

const LOCKFILE = `./install.lock`

var (
	user       string
	path       string
	secrectkey string
)

func init() {

	rootCmd.PersistentFlags().StringVar(&user, "user", "", "must input user")

	rootCmd.AddCommand(bootstrapCmd)
}

var bootstrapCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mybox",
	Long:  `All software has versions. This is mybox's`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap(user, password, secrectkey)
		fmt.Println("mybox Static Site Generator v0.9 -- HEAD")
	},
}

// 要设置很多东西
// 最好是根prompt结合起来使用
func bootstrap(userName, passwd, secrectkey string) {
	operating.Install(userName,password,secrectkey)
}