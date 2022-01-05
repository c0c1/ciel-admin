package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var UserDetails = userDetails{}

type userDetails struct {
}

func (h *userDetails) List(ctx context.Context, d *apiv1.UserDetailsReq) (*res.PageRes, error) {
	total, result := service.UserDetails.List(ctx,d.Page, d.PageSize, &d.UserDetails)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *userDetails) Add(ctx context.Context, r *apiv1.UserDetailsReq) (*res.DataRes, error) {
	if err := service.UserDetails.Add(ctx, &r.UserDetails); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *userDetails) Del(ctx context.Context, d *apiv1.UserDetailsReq) (*res.DataRes, error) {
	if err := service.UserDetails.Del(ctx, d.UserDetails.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *userDetails) GetById(ctx context.Context, d *apiv1.UserDetailsReq) (*res.DataRes, error) {
	data, err := service.UserDetails.GetById(ctx, d.UserDetails.Id)
	if err != nil {
		return nil, err
	}
	return res.OkData(data)
}

func (h *userDetails) Put(ctx context.Context, p *apiv1.UserDetailsReq) (*res.DataRes, error) {
	if err := service.UserDetails.Put(ctx, p.UserDetails); err != nil {
		return nil, err
	}
	return res.Ok()
}
