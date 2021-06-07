package cmd

//登陆命令

import (
	"fmt"

	"github.com/spf13/cobra"
)

var password string

func init() {
	rootCmd.AddCommand(loginCmd)

	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "input the secrect password")

}

var loginCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mybox",
	Long:  `All software has versions. This is mybox's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mybox Static Site Generator v0.9 -- HEAD")
	},
}
