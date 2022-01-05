package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var User = user{}

type user struct {
}

func (h *user) List(ctx context.Context, d *apiv1.UserReq) (*res.PageRes, error) {
	total, result := service.User.List(d.Page, d.PageSize, &d.User)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *user) Add(ctx context.Context, r *apiv1.UserReq) (*res.DataRes, error) {
	if err := service.User.Add(ctx, &r.User); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *user) Del(ctx context.Context, d *apiv1.UserReq) (*res.DataRes, error) {
	if err := service.User.Del(ctx, d.User.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *user) GetById(ctx context.Context, d *apiv1.UserReq) (*res.DataRes, error) {
	data, err := service.User.GetById(ctx, d.User.Id)
	if err != nil {
		return nil, err
	}
	return res.OkData(data)
}

func (h *user) Put(ctx context.Context, p *apiv1.UserReq) (*res.DataRes, error) {
	if err := service.User.Put(ctx, p.User); err != nil {
		return nil, err
	}
	return res.Ok()
}
