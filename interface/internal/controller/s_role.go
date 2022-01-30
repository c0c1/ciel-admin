package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var Role = role{table: "s_role"}

type role struct {
	table string
}

func (c *role) Add(r *ghttp.Request) {
	var d entity.Role
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *role) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *role) Put(r *ghttp.Request) {
	var d entity.Role
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *role) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *role) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	total, data := service.List(gctx.New(), service.SearchConf{
		Table: c.table, Page: page, Size: size,
		Conditions: []*service.Condition{
			{Field: "id", Value: r.GetQuery("id")},
		},
	})
	res.OkPage(data, total, page, size, r)
}

func (c *role) All(r *ghttp.Request) {
	all, err := service.Role.All(gctx.New())
	if err != nil {
		res.Error(err, r)
	}
	res.OKData(all, r)
}
