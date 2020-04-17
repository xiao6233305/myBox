package encrypt


// 验证aes算法是否正确
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var key = `lEKukKEDelDQIeQl`

var q = []struct {
	s string
	r string
}{
	{
		s: `1577692368`,
		r: `8IeAqhX9SilHm3Zuknnrbg==`,
	},
}

func Test_aesEncrypt(t *testing.T) {
	ast := assert.New(t)
	for _, v := range q {
		r, _ := AesEncrypt(v.s, key)
		ast.Equal(v.r, r, "输入:%v", v)
	}
}

func Test_aesDecrypt(t *testing.T) {
	ast := assert.New(t)
	for _, v := range q {
		r, _ := AesDecrypt(v.r, key)
		ast.Equal(v.s, r, "输入:%v", v)
	}
}


