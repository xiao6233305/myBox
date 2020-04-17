package operating

import (
	"fmt"
	"myBox/src/common"
	"myBox/src/encrypt"
	"myBox/src/file"
	"myBox/src/login"
	"myBox/src/myError"
	"os"
	"strconv"
	"time"
)

const LOCKFILE  = `./install.lock`
const PATHSEPARATOR  = string(os.PathSeparator)

func checkParame(userName,passwd,secrectkey string) (err error)  {
	switch  {
	case exitLockFile():
		err = fmt.Errorf(`has install this programe`)
	case len(userName) == 0:
		err = fmt.Errorf("userName is empty")
	case len(secrectkey) == 0:
		err = fmt.Errorf("secrectkey is empty")
	case len(secrectkey) != 16:
		err = fmt.Errorf("secrectkey length must eq 16")
	case len(passwd)==0:
		err = fmt.Errorf("password is empty")
	}
	return err
}


func Install(userName,passwd,secrectkey string)  {
	err := checkParame(userName,passwd,secrectkey)
	myError.ErrorOut(err)
	dataPath := common.DATAPATH

	//以及是否有写入权限
	common.Debug(`begin create dataPath`)
	b,err := createDir(dataPath)
	if !b {
		myError.ErrorOut(fmt.Errorf(`create dataPath fail,errinfo is:%s`,err))
	}
	// 判断配置文件目录是否存在 不存在就创建一个  并且把配置写进去
	confPath := dataPath+PATHSEPARATOR+`conf`
	b,err = createDir(confPath)
	if !b {
		myError.ErrorOut(fmt.Errorf(`create conf dir fail,errinfo is:%s`,err))
	}
	accountPath := dataPath+PATHSEPARATOR+`storage`
	b,err = createDir(accountPath)
	if !b {
		myError.ErrorOut(fmt.Errorf(`create storagePath fail,errinfo is:%s`,err))
	}
	// 创建配置文件
	confFile := common.CONFFILE
	sSlice := make([]string,1)
	sSlice[0] = `dataPath = "`+dataPath+`"`
	sSlice  = append(sSlice , `user = "`+userName+`"`)
	//sSlice  = append(sSlice , `confPath = "`+confPath+`"`)
	//sSlice  = append(sSlice , `confName = "`+confFile+`"`)
	sSlice  = append(sSlice , `accountPath = "`+accountPath+`"`)
	sSlice  = append(sSlice , `secrectkey = "`+secrectkey+`"`)

	encryptKey := common.RandEncryptStr(16)
	encryptedKey,_ := encrypt.AesEncrypt(encryptKey,secrectkey)
	sSlice  = append(sSlice , `encryptKey = "`+encryptedKey+`"`)

	// 密码肯定不能明文保存  需要用加密算法计算出来  这里不能用secrectkey
	passwd = login.CalPassword(passwd, encryptKey)
	sSlice  = append(sSlice , `password = "`+passwd+`"`)

	file.RewriteFile(confFile,sSlice)
	listFile := common.ACCOUNTLISTFILE
	file.RewriteFile(listFile,nil)

	// 生成lock文件
	file.WriteFile(LOCKFILE,nil)
}



func testTouchFile(path string) (bool) {
	name := time.Now().Unix()
	fileName := path+strconv.Itoa(int(name))
	file,err:=os.Create(fileName)
	if err!=nil{
		return false
	}
	_ = file.Close()
	defer func() {
		// 删除文件
		_ = os.Remove(fileName);
	}()
	return true
}


// 如果存在安装文件  说明已经安装了  不能重装
func exitLockFile() (bool)  {
	_, err := os.Stat(LOCKFILE)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return false
}

func createDir(path string) (b bool,err error)  {
	b = true
	if !IsDir(path) {
		err = os.Mkdir(path,0777)
		if err != nil{
			b = false
			return
		}
	}
	// 试着创建一个文件  看看是否能写入
	if !testTouchFile(path){
		err = fmt.Errorf(`can write data into %s`,path)
		return
	}
	return
}
// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	if len(path) == 0{
		return false
	}
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}