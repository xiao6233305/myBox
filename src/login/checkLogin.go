package login

import (
	"myBox/src/common"
	"myBox/src/encrypt"
	"myBox/src/file"
	"strconv"
	"strings"
	"time"
)

/**
	登录只保留15分钟
	通过判断文件名是否存在以及信息是否被篡改来判断
 */
//判断是否完成登录
func CheckLogin() (b bool) {
	//先判断是否存在登录都文件
	if file.CheckFileExists(common.LOGINLOCKFILE){
		data,_ := file.ReadAllData(common.LOGINLOCKFILE)
		nowTimeStamp := time.Now().Unix()
		dataStr := strings.TrimSpace(string(data))
		t := encrypt.DecryptPasswd(dataStr)
		oldTimestamp,_ := strconv.Atoi(t)
		if int(nowTimeStamp)-oldTimestamp<=common.MAXLOGINTIME{
			b = true
		}
	}
	return
}


func Login(passwd string)  {
	encryptkey,t := common.GetEncryptKey()
	encryptkey,_ = encrypt.AesDecrypt(encryptkey,t)
	common.Debug("encryptkey is:",encryptkey,`---password is:`,passwd)
	passwd = CalPassword(passwd,encryptkey)
	oldPasswd := common.GetConfigByName(`password`)
	common.Debug("old password is:",oldPasswd,`--cal password is:`,passwd)
	if oldPasswd == passwd{
		// 把这些信息写入文件
		sSlice := make([]string,1)
		timestamp := time.Now().Unix()
		encryptData,_ := encrypt.AesEncrypt(strconv.Itoa(int(timestamp)),encryptkey)
		sSlice[0] = encryptData
		file.RewriteFile(common.LOGINLOCKFILE,sSlice)
	}else{
		common.Error("password is error")
	}
}
