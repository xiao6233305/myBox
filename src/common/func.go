package common

import (
	"math/rand"
	"myBox/src/myError"
	"time"
	"github.com/astaxie/beego/config"
)

// 产生一个随机字符串  用于加解密
func RandEncryptStr(l int) string  {
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


//获取配置信息  这里可以缓存一下  免得每次都去读取配置文件
func GetConfigByName(s string) string  {
	return GetStringFromFile(CONFFILE,s)
}

// 从某个文件里面读取某个记录
func GetStringFromFile(fileName,key string) string  {
	iniconf, err := config.NewConfig("ini", fileName)
	myError.ErrorOut(err)
	return iniconf.String(key)
}

// 从加密的文件里面读取加密的key
func GetEncryptPassword(fileName string) string  {
	return GetStringFromFile(fileName,PASSWORDKEYNAME)
}

// 获取配置到加密key  这里不解密
func GetEncryptKey() (s1,s2 string)  {
	return GetConfigByName(ENCRYPTCONFIGNAME),GetConfigByName(SECRCTKEYCONFIGNAME)
}