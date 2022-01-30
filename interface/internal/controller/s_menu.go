package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var Menu = menu{table: "s_menu"}

type menu struct {
	table string
}

func (c *menu) Add(r *ghttp.Request) {
	var d entity.Menu
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *menu) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *menu) Put(r *ghttp.Request) {
	var d entity.Menu
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *menu) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *menu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	total, data := service.List(gctx.New(), service.SearchConf{
		Table: c.table,
		Page:  page, Size: size,
		Conditions: []*service.Condition{
			{Field: "id", Value: r.GetQuery("id")},
			{Field: "pid", Value: r.GetQuery("pid")},
			{Field: "type", Value: r.GetQuery("type")},
			{Field: "status", Value: r.GetQuery("status")},
			{Field: "name", Value: r.GetQuery("name"), Like: true},
		},
		OrderBy: "sort desc,id desc",
	})
	res.OkPage(data, total, page, size, r)
}
