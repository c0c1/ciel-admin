package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"interface/utility/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.BindMiddlewareDefault(ghttp.MiddlewareHandlerResponse, middleware.CORS)
			s.Group("/", func(g *ghttp.RouterGroup) {
				g.ALL("/hello", func(r *ghttp.Request) {
					r.Response.WriteJson(ghttp.DefaultHandlerResponse{
						Code:    1,
						Message: "hello",
						Data:    nil,
					})
				})
			})
			SysRouters(s)
			ApiRouters(s)
			s.Run()
			return nil
		},
	}
)
