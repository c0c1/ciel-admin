package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var BookChapter = bookChapter{table: "b_book_chapter"}

type bookChapter struct {
	table string
}

func (c *bookChapter) Add(r *ghttp.Request) {
	var d entity.BookChapter
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *bookChapter) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *bookChapter) Put(r *ghttp.Request) {
	var d entity.BookChapter
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *bookChapter) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *bookChapter) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	total, data := service.List(gctx.New(), service.SearchConf{
		Table: c.table, Page: page, Size: size,
		Conditions: []*service.Condition{
			service.QueryCondition("id", r),
			service.QueryCondition("status", r),
		},
	})
	res.OkPage(data, total, page, size, r)
}

func (c *bookChapter) ListChapter(r *ghttp.Request) {
	data, err := service.BookChapter.ListChapter(gctx.New(), r.GetQuery("bookId").String())
	if err != nil {
		res.Error(err, r)
	}
	res.OKData(data, r)
}
