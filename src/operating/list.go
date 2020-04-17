package operating

import (
	"encoding/json"
	"fmt"
	"myBox/src/common"
	"myBox/src/file"
	"myBox/src/myError"
)

func getAccountList() (list []AccountStruct) {
	data,err := file.ReadAllData(common.ACCOUNTLISTFILE)
	myError.ErrorOut(err)
	if len(data)>0{
		err = json.Unmarshal(data,&list)
		if err != nil{
			myError.ErrorOut(fmt.Errorf("unmarshal listfile json fail,errinfo is:",err))
		}
	}
	return
}

// 把信息都写进去
func writeAccountList(list []AccountStruct)  {
	jsonBytes,_ := json.Marshal(list)
	file.RewriteFile(common.ACCOUNTLISTFILE,[]string{string(jsonBytes)})
}


// 展示所有的账号
func ListAccountList() {
	list := getAccountList()
	fmt.Println("No  SYSTEMNAME ACCOUNT")
	for _,v := range list{
		fmt.Printf("%s  %s %s",v.No,v.SysName,v.Account)
		fmt.Println()
	}
}
