package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var Dict = dict{table: "s_dict"}

type dict struct {
	table string
}

func (c *dict) Add(r *ghttp.Request) {
	var d entity.Dict
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *dict) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *dict) Put(r *ghttp.Request) {
	var d entity.Dict
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *dict) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *dict) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	total, data := service.List(gctx.New(), service.SearchConf{
		Table: c.table,
		Page:  page, Size: size,
		Conditions: []*service.Condition{
			service.QueryCondition("id", r),
			service.QueryCondition("status", r),
			service.QueryCondition("k", r, true),
			service.QueryCondition("v", r, true),
			service.QueryCondition("group", r, true),
		},
	})
	res.OkPage(data, total, page, size, r)
}
