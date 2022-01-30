package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var Api = api{table: "s_api"}

type api struct {
	table string
}

func (c *api) Add(r *ghttp.Request) {
	var d entity.Api
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *api) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *api) Put(r *ghttp.Request) {
	var d entity.Api
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *api) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *api) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	s := service.SearchConf{
		Table: c.table,
		Page:  page, Size: size,
		Conditions: []*service.Condition{
			{Field: "id", Value: r.GetQuery("id")},
			{Field: "url", Value: r.GetQuery("url"), Like: true},
			{Field: "method", Value: r.GetQuery("method")},
			{Field: "group", Value: r.GetQuery("group")},
			{Field: "desc", Value: r.GetQuery("desc"), Like: true},
			{Field: "status", Value: r.GetQuery("status")},
		},
	}
	total, data := service.List(gctx.New(), s)
	res.OkPage(data, total, page, size, r)
}
