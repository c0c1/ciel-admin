package xuser

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"interface/utility/middleware"
)

func Uid(r *ghttp.Request) uint64 {
	return r.Get(middleware.Uid).Uint64()
}
