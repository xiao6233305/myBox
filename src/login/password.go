package login

import (
	"crypto/sha256"
	"encoding/base64"
)

func CalPassword(passwd,key string) string  {
	str := passwd+key
	sum := sha256.Sum256([]byte(str))
	s := make([]byte,len(sum))
	for k,v := range sum{
		s[k] = v
	}
	return base64.StdEncoding.EncodeToString(s)
}
