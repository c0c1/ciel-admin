package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var Writer = writer{table: "b_writer"}

type writer struct {
	table string
}

func (c *writer) Add(r *ghttp.Request) {
	var d entity.Writer
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *writer) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *writer) Put(r *ghttp.Request) {
	var d entity.Writer
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *writer) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *writer) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	s := service.SearchConf{
		Table: c.table, Size: size, Page: page,
		Conditions: []*service.Condition{
			service.QueryCondition("id", r),
			service.QueryCondition("status", r),
		},
	}
	total, data := service.List(gctx.New(), s)
	res.OkPage(data, total, page, size, r)
}
