package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var Api = api{}

type api struct {
}

func (h api) List(ctx context.Context, d *apiv1.ApiReq) (*res.PageRes, error) {
	total, result := service.Api.List(d.Page, d.PageSize, &d.Api)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h api) Add(ctx context.Context, r *apiv1.ApiReq) (*res.DataRes, error) {
	if err := service.Api.Add(ctx, &r.Api); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h api) Del(ctx context.Context, d *apiv1.ApiReq) (*res.DataRes, error) {
	if err := service.Api.Del(ctx, d.Api.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h api) GetById(ctx context.Context, d *apiv1.ApiReq) (*res.DataRes, error) {
	data, _ := service.Api.GetById(ctx, d.Api.Id)
	return res.OkData(data)
}

func (h api) Put(ctx context.Context, p *apiv1.ApiReq) (*res.DataRes, error) {
	if err := service.Api.Put(ctx, p.Api); err != nil {
		return nil, err
	}
	return res.Ok()
}
