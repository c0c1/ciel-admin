package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var File = file{}

type file struct {
}

func (h *file) List(ctx context.Context, d *apiv1.FileReq) (*res.PageRes, error) {
	total, result := service.File.List(d.Page, d.PageSize, &d.File)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *file) Add(ctx context.Context, r *apiv1.FileReq) (*res.DataRes, error) {
	if err := service.File.Add(ctx, &r.File); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *file) Del(ctx context.Context, d *apiv1.FileReq) (*res.DataRes, error) {
	if err := service.File.Del(ctx, d.File.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *file) GetById(ctx context.Context, d *apiv1.FileReq) (*res.DataRes, error) {
	data, err := service.File.GetById(ctx, d.File.Id)
	if err != nil {
		return nil, err
	}
	return res.OkData(data)
}

func (h *file) Put(ctx context.Context, p *apiv1.FileReq) (*res.DataRes, error) {
	if err := service.File.Put(ctx, p.File); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *file) Upload(ctx context.Context, p *apiv1.FileReq) (*res.DataRes, error) {
	err := service.File.Upload(ctx, p)
	if err != nil {
		return nil, err
	}
	return res.Ok()
}
