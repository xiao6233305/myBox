package encrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_md5Sum(t *testing.T) {
	ast := assert.New(t)
	q := []struct {
		s string
		r string
	}{
		{
			s: `abc123`,
			r: `e99a18c428cb38d5f260853678922e03`,
		},
	}
	for _, v := range q {
		ast.Equal(v.r, md5Sum(v.s), "输入:%v", v)
	}
}
