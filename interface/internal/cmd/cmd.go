package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"interface/internal/controller"
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
			s.Group("/menu", func(g *ghttp.RouterGroup) {
				g.Middleware(middleware.Auth)
				g.GET("/list", controller.Menu.List)
				g.GET("/getById", controller.Menu.GetById)
				g.Middleware(middleware.LockAction)
				g.POST("/add", controller.Menu.Add)
				g.PUT("/update", controller.Menu.Put)
				g.DELETE("/del", controller.Menu.Del)
			})
			s.Group("/api", func(g *ghttp.RouterGroup) {
				g.Middleware(middleware.Auth)
				g.GET("/list", controller.Api.List)
				g.GET("/getById", controller.Api.GetById)
				g.Middleware(middleware.LockAction)
				g.POST("/add", controller.Api.Add)
				g.PUT("/update", controller.Api.Put)
				g.DELETE("/del", controller.Api.Del)
			})
			s.Group("/role", func(g *ghttp.RouterGroup) {
				g.Middleware(middleware.Auth)
				g.GET("/all", controller.Role.All)
				g.GET("/list", controller.Role.List)
				g.GET("/getById", controller.Role.GetById)
				g.Middleware(middleware.LockAction)
				g.POST("/add", controller.Role.Add)
				g.PUT("/update", controller.Role.Put)
				g.DELETE("/del", controller.Role.Del)
			})
			s.Group("/roleMenu", func(g *ghttp.RouterGroup) {
				g.Middleware(middleware.Auth)
				g.GET("/list", controller.RoleMenu.List)
				g.GET("/getById", controller.RoleMenu.GetById)
				g.GET("/noMenus", controller.RoleMenu.NoMenus)
				g.Middleware(middleware.LockAction)
				g.POST("/add", controller.RoleMenu.Add)
				g.PUT("/update", controller.RoleMenu.Put)
				g.DELETE("/del", controller.RoleMenu.Del)
			})
			s.Group("/roleApi", func(g *ghttp.RouterGroup) {
				g.Middleware(middleware.Auth)
				g.GET("/list", controller.RoleApi.List)
				g.GET("/getById", controller.RoleApi.GetById)
				g.GET("/noApis", controller.RoleApi.NoApis)
				g.Middleware(middleware.LockAction)
				g.POST("/add", controller.RoleApi.Add)
				g.PUT("/update", controller.RoleApi.Put)
				g.DELETE("/del", controller.RoleApi.Del)
			})
			s.Group("/admin", func(g *ghttp.RouterGroup) {
				g.POST("/login", controller.Admin.Login)
				g.Middleware(middleware.Auth)
				g.GET("/list", controller.Admin.List)
				g.GET("/getById", controller.Admin.GetById)
				g.Middleware(middleware.LockAction)
				g.POST("/add", controller.Admin.Add)
				g.PUT("/update", controller.Admin.Put)
				g.DELETE("/del", controller.Admin.Del)
				g.PUT("/pwd", controller.Admin.UpdatePwd)
			})
			s.Group("/dict", func(g *ghttp.RouterGroup) {
				g.Middleware(middleware.Auth)
				g.GET("/list", controller.Dict.List)
				g.GET("/getById", controller.Dict.GetById)
				g.Middleware(middleware.LockAction)
				g.POST("/add", controller.Dict.Add)
				g.PUT("/update", controller.Dict.Put)
				g.DELETE("/del", controller.Dict.Del)
			})
			s.Group("/file", func(g *ghttp.RouterGroup) {
				g.Middleware(middleware.Auth)
				g.POST("/upload", controller.File.Upload)
				g.GET("/list", controller.File.List)
				g.GET("/getById", controller.File.GetById)
				g.Middleware(middleware.LockAction)
				g.POST("/add", controller.File.Add)
				g.PUT("/update", controller.File.Put)
				g.DELETE("/del", controller.File.Del)
			})
			s.Run()
			return nil
		},
	}
)
