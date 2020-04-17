package operating

import (
	"fmt"
	"myBox/src/common"
	"myBox/src/encrypt"
	"myBox/src/file"
	"myBox/src/myError"
	"strconv"
)


// 添加账号密码放入文件里面
func Add(user,passwd,sysName string) bool  {
	switch  {
	case len(user)==0:
		myError.ErrorOut(fmt.Errorf("users is empty"))
		return false
	case len(sysName)==0:
		myError.ErrorOut(fmt.Errorf("sysName is empty"))
		return false
	case len(passwd)==0:
		myError.ErrorOut(fmt.Errorf("password is empty"))
		return false
	default:
		// 先判断是否存在  存在就进行更新  不存在就添加
		sSlice := make([]string,3)
		sSlice[0] = `user = "`+user+`"`
		sSlice[1] = `sysName = "`+sysName+`"`
		sSlice[2] = common.PASSWORDKEYNAME+` = "`+encrypt.EncryptPassword(passwd)+`"`
		list := getAccountList()
		maxNo := ``
		//把密码写入文件里面 防止文件名有空格这类系统不支持对情况  所以base64
		fileName := ``
		existsKey := -1
		newAccount := new(AccountStruct)
		for k,v := range list{
			if v.SysName == sysName{
				existsKey = k
				newAccount = &v
				fileName = v.FileName
				// 这里不要跳出去  因为需要获取最大的NO
			}
			maxNo = v.No
		}
		if existsKey == -1{
			fileName = common.ACCOUNTPATH+common.RandEncryptStr(32)
			tmpNo,_ := strconv.Atoi(maxNo)
			tmpNo++
			newAccount.No = strconv.Itoa(tmpNo)
		}
		newAccount.FileName = fileName
		newAccount.SysName = sysName
		newAccount.Account = user
		file.RewriteFile(fileName,sSlice)
		md5Value := file.GetFileMd5(fileName)
		newAccount.Md5 = md5Value
		switch  {
		case len(list)==0:
			// 完全没有记录
			list = append(list,AccountStruct{})
			list[0] = *newAccount
		case existsKey==-1:
			//新添加的一条
			list = append(list,*newAccount)
		default:
			//修改已有的
			list[existsKey] = *newAccount
		}
		writeAccountList(list)
	}
	return true
}
