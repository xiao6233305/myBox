// 可以让我们把一些系统的密码加密保存在本地
// 同时让没有查看密码的人也看不到我们的密码
package main

import (
	"github.com/xiao6233305/mybox/src/cmd"
	"github.com/xiao6233305/mybox/src/common"
	_ "os"
)

func init() {
	//设置日志级别
	common.SetLogLevel(MODE)
}

func main() {
	//需要让程序一直跑  所以需要启动一个goruntine 空跑
	cmd.Execute()
}

/**
var (
	cmd                          string
	userName, passwd, secrectkey string
	sysName                      string
	num                          string
)

func getFlagVars() {
	flag.StringVar(&cmd, `cmd`, ``, `please chose cmd`)
	flag.StringVar(&userName, `u`, ``, `please input user name`)
	flag.StringVar(&passwd, `p`, ``, `please input password`)
	flag.StringVar(&secrectkey, `k`, ``, `please input secrectkey`)
	flag.StringVar(&sysName, `s`, ``, `please input systemName`)
	flag.StringVar(&num, `No`, ``, `please input No`)
	flag.Parse()
	common.Debug(`commond args is:`, os.Args)
}

func main() {
	getFlagVars()

	switch cmd {
	case `install`:
		//初始化  需要填写用户名  密码  以及对于的安全key  以及数据的保存目录  然后生成配置文件
		operating.Install(userName, passwd, secrectkey)
	case `login`:
		// 进行登录操作
		login.Login(passwd)
	case `version`, `v`:
		//打印程序版本
		fmt.Println(`version is:`, VERSION)
	case `logout`:
	//进行退出操作
	default:
		if !login.CheckLogin() {
			panic("please login")
		}
		switch cmd {
		case `list`:
			//列出所有的账号列表
			operating.ListAccountList()
		case `add`:
			//添加需要保存的密码  放入一个文件夹里面
			//先只支持保存账号  密码  以及系统名字
			operating.Add(userName, passwd, sysName)
		case `del`:
		//删除某个账号

		case `backend`:
		//备份
		case `query`:
			//查看某个账号
			operating.Query(userName, sysName, num)
		default:
			// 不支持的命令
			fmt.Println(`can not support this cmd,`, cmd)
		}
	}
}
*/
