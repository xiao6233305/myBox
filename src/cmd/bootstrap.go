package cmd

// 执行安装命令
import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"myBox/src/common"
	"myBox/src/file"
	"myBox/src/myError"
)

const LOCKFILE  = `./install.lock`

var (
	user string
	path string
	secrectkey string
)

func init() {

	rootCmd.PersistentFlags().StringVar(&user, "user", "", "must input user")

	rootCmd.AddCommand(bootstrapCmd)
}
var bootstrapCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mybox",
	Long:  `All software has versions. This is mybox's`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap(user,password,secrectkey)
		fmt.Println("mybox Static Site Generator v0.9 -- HEAD")
	},
}

// 要设置很多东西
// 最好是根prompt结合起来使用
func bootstrap(userName,passwd,secrectkey string)  {
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


