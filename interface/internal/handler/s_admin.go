package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
	"interface/utility/utils/xuser"
)

var Admin = admin{}

type admin struct {
}

func (h admin) List(ctx context.Context, d *apiv1.AdminReq) (*res.PageRes, error) {
	total, result := service.Admin.List(d.Page, d.PageSize, &d.Admin)
	pageRes := res.NewPageRes(result, total, d.Page, d.PageSize)
	all, err := service.Role.All(ctx)
	if err != nil {
		return nil, err
	}
	pageRes.AddOther(all)
	return pageRes, nil
}

func (h admin) Add(ctx context.Context, r *apiv1.AdminReq) (*res.DataRes, error) {
	if err := service.Admin.Add(ctx, &r.Admin); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h admin) Del(ctx context.Context, d *apiv1.AdminReq) (*res.DataRes, error) {
	if err := service.Admin.Del(ctx, d.Admin.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h admin) GetById(ctx context.Context, d *apiv1.AdminReq) (*res.DataRes, error) {
	data, _ := service.Admin.GetById(ctx, d.Admin.Id)
	return res.OkData(data)
}

func (h admin) Put(ctx context.Context, p *apiv1.AdminReq) (*res.DataRes, error) {
	if err := service.Admin.Put(ctx, p.Admin); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h admin) Login(ctx context.Context, p *apiv1.LoginReq) (*res.DataRes, error) {
	data, err := service.Admin.Login(ctx, p.Uname, p.Pwd)
	if err != nil {
		return nil, err
	}
	return res.OkData(data)
}

func (h admin) UpdatePwd(ctx context.Context, p *apiv1.PwdReq) (*res.DataRes, error) {
	if err := service.Admin.UpdatePwd(ctx, p.Pwd, xuser.Uid(ctx)); err != nil {
		return nil, err
	}
	return nil, nil
}
