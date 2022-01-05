package cmd

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"interface/internal/handler"
	"interface/utility/utils/middleware"
)

func SysRouters(s *ghttp.Server) {
	s.Group("/menu", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.Menu.List)
		g.GET("/getById", handler.Menu.GetById)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.Menu.Add)
		g.PUT("/update", handler.Menu.Put)
		g.DELETE("/del", handler.Menu.Del)
	})

	s.Group("/api", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.Api.List)
		g.GET("/getById", handler.Api.GetById)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.Api.Add)
		g.PUT("/update", handler.Api.Put)
		g.DELETE("/del", handler.Api.Del)
	})

	s.Group("/role", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.Role.List)
		g.GET("/getById", handler.Role.GetById)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.Role.Add)
		g.PUT("/update", handler.Role.Put)
		g.DELETE("/del", handler.Role.Del)
	})

	s.Group("/roleMenu", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.RoleMenu.List)
		g.GET("/getById", handler.RoleMenu.GetById)
		g.GET("/noMenus", handler.RoleMenu.NoMenus)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.RoleMenu.Add)
		g.PUT("/update", handler.RoleMenu.Put)
		g.DELETE("/del", handler.RoleMenu.Del)
	})
	s.Group("/roleApi", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.RoleApi.List)
		g.GET("/getById", handler.RoleApi.GetById)
		g.GET("/noApis", handler.RoleApi.NoApis)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.RoleApi.Add)
		g.PUT("/update", handler.RoleApi.Put)
		g.DELETE("/del", handler.RoleApi.Del)
	})

	s.Group("/admin", func(g *ghttp.RouterGroup) {
		g.POST("/login", handler.Admin.Login)
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.Admin.List)
		g.GET("/getById", handler.Admin.GetById)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.Admin.Add)
		g.PUT("/update", handler.Admin.Put)
		g.DELETE("/del", handler.Admin.Del)
		g.PUT("/pwd", handler.Admin.UpdatePwd)
	})
	s.Group("/dict", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.Dict.List)
		g.GET("/getById", handler.Dict.GetById)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.Dict.Add)
		g.PUT("/update", handler.Dict.Put)
		g.DELETE("/del", handler.Dict.Del)
	})

	s.Group("/file", func(g *ghttp.RouterGroup) {
		g.POST("/upload", handler.File.Upload)
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.File.List)
		g.GET("/getById", handler.File.GetById)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.File.Add)
		g.PUT("/update", handler.File.Put)
		g.DELETE("/del", handler.File.Del)
	})

	s.Group("/icon", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.Icon.List)
		g.GET("/getById", handler.Icon.GetById)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.Icon.Add)
		g.PUT("/update", handler.Icon.Put)
		g.DELETE("/del", handler.Icon.Del)
	})
	s.Group("/user", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.User.List)
		g.GET("/getById", handler.User.GetById)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.User.Add)
		g.PUT("/update", handler.User.Put)
		g.DELETE("/del", handler.User.Del)
	})
}
