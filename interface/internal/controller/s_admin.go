package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
	"interface/utility/utils/xpwd"
	"interface/utility/utils/xuser"
)

var Admin = admin{table: "s_admin"}

type admin struct {
	table string
}

func (c *admin) Add(r *ghttp.Request) {
	var d entity.Admin
	_ = r.Parse(&d)
	d.Pwd = xpwd.GenPwd(d.Pwd)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *admin) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *admin) Put(r *ghttp.Request) {
	var d entity.Admin
	_ = r.Parse(&d)
	if err := service.Admin.Put(gctx.New(), &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *admin) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}

func (c *admin) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	s := service.SearchConf{
		Table: c.table,
		Page:  page, Size: size, Conditions: []*service.Condition{
			{Field: "id", Value: r.GetQuery("id")},
			{Field: "status", Value: r.GetQuery("status")},
			{Field: "rid", Value: r.GetQuery("rid")},
			{Field: "pwd"},
		},
		FieldsEx: "pwd",
	}
	total, data := service.List(gctx.New(), s)
	res.OkPage(data, total, page, size, r)
}

func (c admin) Login(r *ghttp.Request) {
	var d struct {
		Uname string
		Pwd   string
	}
	_ = r.Parse(&d)
	data, err := service.Admin.Login(gctx.New(), d.Uname, d.Pwd)
	if err != nil {
		res.Error(err, r)
	}
	res.OKData(data, r)
}

func (c admin) UpdatePwd(r *ghttp.Request) {
	var d struct {
		Pwd string
	}
	_ = r.Parse(&d)
	if err := service.Admin.UpdatePwd(gctx.New(), d.Pwd, xuser.Uid(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}
