package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var Dict = dict{}

type dict struct {
}

func (h *dict) List(ctx context.Context, d *apiv1.DictReq) (*res.PageRes, error) {
	total, result := service.Dict.List(d.Page, d.PageSize, &d.Dict)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *dict) Add(ctx context.Context, r *apiv1.DictReq) (*res.DataRes, error) {
	if err := service.Dict.Add(ctx, &r.Dict); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *dict) Del(ctx context.Context, d *apiv1.DictReq) (*res.DataRes, error) {
	if err := service.Dict.Del(ctx, d.Dict.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *dict) GetById(ctx context.Context, d *apiv1.DictReq) (*res.DataRes, error) {
	data, err := service.Dict.GetById(ctx, d.Dict.Id)
	if err != nil {
		return nil, err
	}
	return res.OkData(data)
}

func (h *dict) Put(ctx context.Context, p *apiv1.DictReq) (*res.DataRes, error) {
	if err := service.Dict.Put(ctx, p.Dict); err != nil {
		return nil, err
	}
	return res.Ok()
}
