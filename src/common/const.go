package common

import "os"

const (
	DATAPATH  = `/Users/xiaoyajun/work/go/src/myBox/src/data`
	CONFFILE  = DATAPATH+string(os.PathSeparator)+`conf/app.ini`
	ACCOUNTPATH = DATAPATH+string(os.PathSeparator)+`storage`+string(os.PathSeparator)
	ENCRYPTCONFIGNAME = `encryptKey`
	SECRCTKEYCONFIGNAME = `secrectkey`
	ACCOUNTLISTFILE = ACCOUNTPATH+`listfile`
	PASSWORDKEYNAME = `password`

	LOGINLOCKFILE = DATAPATH+string(os.PathSeparator)+`login.lock`

	MAXLOGINTIME = 900
)


