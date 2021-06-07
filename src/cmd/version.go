package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mybox",
	Long:  `All software has versions. This is mybox's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mybox Static Site Generator v0.9 -- HEAD")
	},
}
