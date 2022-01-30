package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var RoleMenu = roleMenu{table: "s_role_menu"}

type roleMenu struct {
	table string
}

func (c *roleMenu) Add(r *ghttp.Request) {
	var d struct {
		Rid int
		Mid []int
	}
	_ = r.Parse(&d)
	if err := service.RoleMenu.Add(gctx.New(), d.Rid, d.Mid); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *roleMenu) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *roleMenu) Put(r *ghttp.Request) {
	var d entity.RoleMenu
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *roleMenu) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *roleMenu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	total, data := service.RoleMenu.List(page, size, r.GetQuery("rid").Uint64(), 0)
	res.OkPage(data, total, page, size, r)
}

func (c *roleMenu) NoMenus(r *ghttp.Request) {
	data := service.RoleMenu.NoMenus(gctx.New(), r.GetQuery("rid").Uint64())
	res.OKData(data, r)
}
