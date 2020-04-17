package operating

import (
	"fmt"
	"myBox/src/common"
	"myBox/src/encrypt"
	"myBox/src/file"
)

// 查询某个账号的密码
func Query(userName,sysName,num string)  {
	list := getAccountList()
	Account := AccountStruct{}
	exits := false
	for _,v := range list{
		if v.No == num || v.Account == userName || v.SysName == sysName {
			exits = true
			Account = v
			break
		}
	}
	if exits{
		switch  {
		//先判断文件是否存在  出错的情况需要删除掉出错的记录
		case !file.CheckFileExists(Account.FileName):
			common.Error(fmt.Sprintf("system %s file has delete",Account.SysName))
		case Account.Md5 != file.GetFileMd5(Account.FileName):
			// 在判断文件是否有被修改过
			common.Error(fmt.Sprintf("system %s file has damage",Account.SysName))
		default:
			// 先读取文件  然后把密码读出来
			encryptPasswd := common.GetEncryptPassword(Account.FileName)
			common.Debug(Account.FileName,`===`,encryptPasswd)
			passwd := encrypt.DecryptPasswd(encryptPasswd)
			fmt.Println(`password is:`,passwd)
		}
	}else{
		fmt.Println("can not find accountInfo")
	}
}
