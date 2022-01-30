package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var RoleApi = roleApi{table: "s_role_api"}

type roleApi struct {
	table string
}

func (c *roleApi) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *roleApi) Put(r *ghttp.Request) {
	var d entity.RoleApi
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *roleApi) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *roleApi) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	total, result := service.RoleApi.List(page, size, r.GetQuery("rid").Int(), 0)
	res.OkPage(result, total, page, size, r)
}

func (c *roleApi) Add(r *ghttp.Request) {
	var d struct {
		Rid int
		Aid []int
	}
	_ = r.Parse(&d)
	if err := service.RoleApi.Add(gctx.New(), d.Rid, d.Aid); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c roleApi) NoApis(r *ghttp.Request) {
	apis := service.RoleApi.NoApis(gctx.New(), r.GetQuery("rid").Uint64())
	res.OKData(apis, r)
}
