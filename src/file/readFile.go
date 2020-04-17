package file

import (
	"fmt"
	"io/ioutil"
	"myBox/src/encrypt"
	"myBox/src/myError"
)


// 把整个文件内容读取出来
func ReadAllData(fileName string) (data []byte,err error)  {
	switch  {
	case !CheckFileExists(fileName):
		err = fmt.Errorf("file %s not exist",fileName)
	default:
		data,err = ioutil.ReadFile(fileName)
	}
	return
}


// 判断文件的md5值
func GetFileMd5(filename string) (s string)  {
	switch  {
	case CheckFileExists(filename):
		data,_ := ReadAllData(filename)
		s = encrypt.Md5WithSalt(string(data))
	default:
		myError.ErrorOut(fmt.Errorf("file not exists ,can  not check md5"))
	}
	return s
}