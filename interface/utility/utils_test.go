package utility

import (
	"interface/utility/utils/xpwd"
	"testing"
)

func Test1(t *testing.T) {
	pwd := xpwd.GenPwd("1")
	println(pwd)
	println(xpwd.ComparePassword("$2a$10$Ny0c.Yslfm4F2kl.dh7Z4edZQC.DxD.CFP.qYRD5VSrZNXl7o2bmm", "1"))
}

func Test2(t *testing.T) {
}
