package main

import (
	"fmt"
	"math/rand"
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

func install(userName,passwd,secrectkey string) (b bool,err error)  {

	if exitLockFile(){
		 err = fmt.Errorf(`has install this programe`)
		 return
	}
	dataPath := `./data`


	// 默认安装路径 以及后续的账号密码存在路径
	//flag.StringVar(&dataPath,`-path`,`./data`,`please input filepath for storage data`)
	//flag.Parse()

	if len(secrectkey) < 16{
		err = fmt.Errorf("secrectkey length must gt 16")
		myError.ErrorOut(err)
	}

	// 先判断目录是否存在  以及是否有写入权限
	b,err = createDir(dataPath)
	if !b {
		err = fmt.Errorf(`create dataPath fail,errinfo is:%s`,err)
		return
	}
	// 判断配置文件目录是否存在 不存在就创建一个  并且把配置写进去
	confPath := dataPath+PATHSEPARATOR+`conf`
	b,err = createDir(confPath)
	if !b {
		err = fmt.Errorf(`create conf dir fail,errinfo is:%s`,err)
		return
	}
	accountPath := dataPath+PATHSEPARATOR+`storage`
	b,err = createDir(accountPath)
	if !b {
		err = fmt.Errorf(`create storagePath fail,errinfo is:%s`,err)
		return
	}
	// 创建配置文件
	confFile := confPath+PATHSEPARATOR+`app.ini`
	sSlice := make([]string,1)
	sSlice[0] = `dataPath = "`+dataPath+`"`
	sSlice  = append(sSlice , `user = "`+userName+`"`)
	sSlice  = append(sSlice , `confPath = "`+confPath+`"`)
	sSlice  = append(sSlice , `confName = "`+confFile+`"`)
	sSlice  = append(sSlice , `accountPath = "`+accountPath+`"`)
	sSlice  = append(sSlice , `secrectkey = "`+secrectkey+`"`)
	// 密码肯定不能明文保存  需要用加密算法计算出来
	passwd = login.CalPassword(passwd,secrectkey)
	sSlice  = append(sSlice , `password = "`+passwd+`"`)
	encryptKey,_ := encrypt.AesEncrypt(randEncryptStr(32),secrectkey)
	sSlice  = append(sSlice , `encryptKey = "`+encryptKey+`"`)



	file.WriteFile(confFile,sSlice)
	// 生成lock文件
	file.WriteFile(LOCKFILE,nil)
	return
}

// 产生一个随机字符串  用于加解密
func randEncryptStr(l int) string  {
	s := ``
	rand.Seed(time.Now().UnixNano())

	for l>0{
		l1 := rand.Intn(9)
		var s1 byte
		n := rand.Intn(25)
		switch  {
		case l1%2==0:
			//产生的是偶数 就生成一个大写字母
			s1 = byte(n+65)
		default:
			//否则产生的是小写字母
			s1 = byte(n+97)
		}
		s = s+string(s1)
		l--
	}
	return s
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