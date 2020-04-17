package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"myBox/src/common"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key string) (string, error) {
	r,e := aesEncrypt([]byte(origData),[]byte(key))
	return base64.StdEncoding.EncodeToString(r), e
}


func aesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func aesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func AesDecrypt(crypted, key string) (string, error) {
	s,_ := base64.StdEncoding.DecodeString(crypted)
	r,e := aesDecrypt(s,[]byte(key))
	return string(r),e
}


// 对密码进行加密  不能让人通过代码就能解密出密钥  必须获取到key
func EncryptPassword(s string) string  {
	encryptKey,secretkey := common.GetEncryptKey()
	key,_ := AesDecrypt(encryptKey,secretkey)
	r,_ := AesEncrypt(s,key)
	return r
}

//对密码进行解密
func DecryptPasswd(s string) string  {
	encryptKey,secretkey := common.GetEncryptKey()
	key,_ := AesDecrypt(encryptKey,secretkey)
	r,_ := AesDecrypt(s,key)
	return r
}