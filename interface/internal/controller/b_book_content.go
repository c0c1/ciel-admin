package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var BookContent = bookContent{table: "b_book_content"}

type bookContent struct {
	table string
}

func (c *bookContent) Add(r *ghttp.Request) {
	var d entity.BookContent
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *bookContent) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *bookContent) Put(r *ghttp.Request) {
	var d entity.BookContent
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *bookContent) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *bookContent) List(r *ghttp.Request) {
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

func (c *bookContent) Details(r *ghttp.Request) {
	var d struct {
		ChapterId int64
		BookId    int64
	}
	r.Parse(&d)
	data, err := service.BookContent.GetByChapterId(gctx.New(), d.BookId, d.ChapterId)
	if err != nil {
		res.Error(err, r)
	}
	res.OKData(data, r)
}
