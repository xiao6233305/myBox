package file

import (
	"bufio"
	"fmt"
	"myBox/src/myError"
	"os"
)

// 反正出错就退出了  都不需要返回值了
func WriteFile(filename string ,sSlice []string) {
	var err error
	// 如果文件存在  就覆盖  否则创建文件
	if CheckFileExists(filename) { //如果文件存在
		err = fmt.Errorf("file %s has Exist",filename)
		myError.ErrorOut(err)
	}
	f, err := os.Create(filename) //创建文件
	myError.ErrorOut(err)
	defer f.Close()


	w := bufio.NewWriter(f)
	for _, v := range sSlice {
		lineStr := fmt.Sprintf("%s",  v)
		_,_ = fmt.Fprintln(w, lineStr)
	}
	err = w.Flush()
	myError.ErrorOut(err)
}
/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileExists(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 如果文件存在  就删了  重新写  如果不存在就写入
func RewriteFile(filename string ,sSlice []string)  {
	if CheckFileExists(filename) { //如果文件存在
		err := os.Remove(filename)
		myError.ErrorOut(err)
	}
	WriteFile(filename,sSlice)
}
