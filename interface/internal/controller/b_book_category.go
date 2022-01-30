package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var BookCategory = bookCategory{table: "b_book_category"}

type bookCategory struct {
	table string
}

func (c *bookCategory) Add(r *ghttp.Request) {
	var d entity.BookCategory
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *bookCategory) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *bookCategory) Put(r *ghttp.Request) {
	var d entity.BookCategory
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *bookCategory) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *bookCategory) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	s := service.SearchConf{
		Table: c.table,
		Page:  page, Size: size,
		Conditions: []*service.Condition{
			service.QueryCondition("id", r),
			service.QueryCondition("status", r),
		},
	}
	total, data := service.List(gctx.New(), s)
	res.OkPage(data, total, page, size, r)
}
func (c *bookCategory) ListHomeData(r *ghttp.Request) {
	data, err := service.BookCategory.ListHomeData()
	if err != nil {
		res.Error(err, r)
	}
	res.OKData(data, r)
}
