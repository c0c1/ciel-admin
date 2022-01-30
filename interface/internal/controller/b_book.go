package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var Book = book{table: "b_book"}

type book struct {
	table string
}

func (c *book) Add(r *ghttp.Request) {
	var d entity.Book
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *book) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *book) Put(r *ghttp.Request) {
	var d entity.Book
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *book) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *book) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	s := service.SearchConf{
		Table: c.table,
		Page:  page, Size: size,
		Conditions: []*service.Condition{
			service.QueryCondition("id", r),
			service.QueryCondition("name", r, true),
			service.QueryCondition("status", r),
		},
	}
	total, data := service.List(gctx.New(), s)
	res.OkPage(data, total, page, size, r)
}

func (c *book) ListAllBookIds(r *ghttp.Request) {
	ids, err := service.Book.ListAllBookIds(gctx.New())
	if err != nil {
		res.Error(err, r)
	}
	res.OKData(ids, r)
}
