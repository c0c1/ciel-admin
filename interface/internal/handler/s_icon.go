package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var Icon = icon{}

type icon struct {
}

func (h *icon) List(ctx context.Context, d *apiv1.IconReq) (*res.PageRes, error) {
	total, result := service.Icon.List(d.Page, d.PageSize, &d.Icon)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *icon) Add(ctx context.Context, r *apiv1.IconReq) (*res.DataRes, error) {
	if err := service.Icon.Add(ctx, &r.Icon); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *icon) Del(ctx context.Context, d *apiv1.IconReq) (*res.DataRes, error) {
	if err := service.Icon.Del(ctx, d.Icon.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *icon) GetById(ctx context.Context, d *apiv1.IconReq) (*res.DataRes, error) {
	data, err := service.Icon.GetById(ctx, d.Icon.Id)
	if err != nil {
		return nil, err
	}
	return res.OkData(data)
}

func (h *icon) Put(ctx context.Context, p *apiv1.IconReq) (*res.DataRes, error) {
	if err := service.Icon.Put(ctx, p.Icon); err != nil {
		return nil, err
	}
	return res.Ok()
}
