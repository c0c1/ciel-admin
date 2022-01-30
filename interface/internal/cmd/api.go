package cmd

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"interface/internal/controller"
)

func ApiRouters(s *ghttp.Server) {
	s.Group("/book", func(g *ghttp.RouterGroup) {
		g.GET("/category", controller.BookCategory.ListHomeData)
		g.GET("/listChapter", controller.BookChapter.ListChapter)
		g.GET("/details", controller.BookContent.Details)
		g.GET("/allBookIds", controller.Book.ListAllBookIds)
	})
}
