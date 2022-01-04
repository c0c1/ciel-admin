package middleware

import "github.com/gogf/gf/v2/net/ghttp"

func CORS(c *ghttp.Request) {
	c.Response.CORSDefault()
	c.Middleware.Next()
}
