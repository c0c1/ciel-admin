package xparam

import "github.com/gogf/gf/v2/net/ghttp"

func ID(r *ghttp.Request) uint64 {
	return r.GetQuery("id").Uint64()
}
