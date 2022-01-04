package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var Menu = menu{}

type menu struct {
}

func (h *menu) List(ctx context.Context, d *apiv1.MenuReq) (*res.PageRes, error) {
	total, result := service.Menu.List(d.Page, d.PageSize, &d.Menu)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *menu) Add(ctx context.Context, r *apiv1.MenuReq) (*res.DataRes, error) {
	if err := service.Menu.Add(ctx, &r.Menu); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *menu) Del(ctx context.Context, d *apiv1.MenuReq) (*res.DataRes, error) {
	if err := service.Menu.Del(ctx, d.Menu.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *menu) GetById(ctx context.Context, d *apiv1.MenuReq) (*res.DataRes, error) {
	data, _ := service.Menu.GetById(ctx, d.Menu.Id)
	return res.OkData(data)
}

func (h *menu) Put(ctx context.Context, p *apiv1.MenuReq) (*res.DataRes, error) {
	if err := service.Menu.Put(ctx, p.Menu); err != nil {
		return nil, err
	}
	return res.Ok()
}
