package encrypt

import (
	"crypto/md5"
	"fmt"
	"myBox/src/common"
)

func md5Sum(data string) string  {
	tmp := md5.Sum([]byte(data))
	s := fmt.Sprintf("%x", tmp)
	return s
}

// 计算加盐以后的md5值
func Md5WithSalt(data string) string  {
	encryptKey,secretkey := common.GetEncryptKey()
	key,_ := AesDecrypt(encryptKey,secretkey)
	return md5Sum(data+key)
}
