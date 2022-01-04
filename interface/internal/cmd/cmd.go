package cmd

import (
	"context"
	"interface/utility/utils/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"interface/internal/handler"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.BindMiddlewareDefault(ghttp.MiddlewareHandlerResponse, middleware.CORS)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(handler.Hello)
			})
			SysRouters(s)
			s.Run()
			return nil
		},
	}
)
