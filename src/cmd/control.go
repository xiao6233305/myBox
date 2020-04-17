package cmd

// 具体的数据操作  我理解应该是登陆以后才可以
// 所以需要有一个channel和登陆模块进行交互 如果登陆成功了

import (
	"fmt"
	"github.com/c-bata/go-prompt"

	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(controlCmd)
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "login", Description: "pleas input you password for check login"},
		{Text: "list", Description: "list all you account"},
		{Text: "add", Description: "add  you account  must input username systemname password"},
		{Text: "delete", Description: "delete  you account by keywords"},
		{Text: "search", Description: "search you account by keywords"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

// 登陆  添加  修改 删除 都需要在这里完成  所以要用prompt

// 弄两个chann  一个用来把需要运行的函数写入chann  另一个负责把函数的结果写入chann  然后打印
var controlCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mybox",
	Long:  `All software has versions. This is mybox's`,
	Run: func(cmd *cobra.Command, args []string) {
		control()
	},
}

func control()  {
	t := prompt.Input(`mybox> `, completer)
	fmt.Println("You selected " + t)

}

func exector(s string)  {
	switch s {
	case `add`:
		fmt.Println("add password")
	default:
		fmt.Println("can not support")
	}
}

