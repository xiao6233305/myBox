package login

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var q1 = []struct {
	s string
	r bool
}{
	{
		s: `123456`,
		r: false,
	},
	{
		s: `yajun900220`,
		r: true,
	},
}

func Test_CheckLogin(t *testing.T) {
	ast := assert.New(t)
	for _, v := range q {
		r := CalPassword(v.s, key)
		ast.Equal(v.r, r, "输入:%v", v)
	}
}

func Test_login(t *testing.T) {
	ast := assert.New(t)
	for _, v := range q1 {
		Login(v.s)
		r := CheckLogin()
		ast.Equal(v.r, r, "输入:%v", v)
	}
}
