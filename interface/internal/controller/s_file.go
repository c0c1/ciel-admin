package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var File = file{table: "s_file"}

type file struct {
	table string
}

func (c *file) Add(r *ghttp.Request) {
	var d entity.File
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *file) Put(r *ghttp.Request) {
	var d entity.File
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *file) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *file) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	total, data := service.List(gctx.New(), service.SearchConf{
		Table: c.table,
		Page:  page, Size: size,
		Conditions: []*service.Condition{
			service.QueryCondition("id", r),
			service.QueryCondition("status", r),
			service.QueryCondition("group", r),
		},
	})
	res.OkPage(data, total, page, size, r)
}

func (c *file) Del(r *ghttp.Request) {
	if err := service.File.Del(gctx.New(), xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *file) Upload(r *ghttp.Request) {
	err := service.File.Upload(gctx.New(), r)
	if err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}
