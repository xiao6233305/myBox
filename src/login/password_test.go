package login

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
		s: `123456`,
		r: `vXLMGbvKxSucxRUO8I+AEImAOxwmlhQK6C8ZFheZaK4=`,
	},
}

func Test_CalPassword(t *testing.T) {
	ast := assert.New(t)
	for _, v := range q {
		r := CalPassword(v.s, key)
		ast.Equal(v.r, r, "输入:%v", v)
	}
}
