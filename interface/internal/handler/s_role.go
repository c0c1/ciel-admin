package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var Role = role{}

type role struct {
}

func (h *role) List(ctx context.Context, d *apiv1.RoleReq) (*res.PageRes, error) {
	total, result := service.Role.List(d.Page, d.PageSize, &d.Role)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *role) Add(ctx context.Context, r *apiv1.RoleReq) (*res.DataRes, error) {
	if err := service.Role.Add(ctx, &r.Role); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *role) Del(ctx context.Context, d *apiv1.RoleReq) (*res.DataRes, error) {
	if err := service.Role.Del(ctx, d.Role.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *role) GetById(ctx context.Context, d *apiv1.RoleReq) (*res.DataRes, error) {
	data, _ := service.Role.GetById(ctx, d.Role.Id)
	return res.OkData(data)
}

func (h *role) Put(ctx context.Context, p *apiv1.RoleReq) (*res.DataRes, error) {
	if err := service.Role.Put(ctx, p.Role); err != nil {
		return nil, err
	}
	return res.Ok()
}
