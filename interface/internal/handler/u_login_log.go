package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var LoginLog = loginLog{}

type loginLog struct {
}

func (h *loginLog) List(ctx context.Context, d *apiv1.LoginLogReq) (*res.PageRes, error) {
	total, result := service.LoginLog.List(ctx, d.Page, d.PageSize, &d.LoginLog)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *loginLog) Add(ctx context.Context, r *apiv1.LoginLogReq) (*res.DataRes, error) {
	if err := service.LoginLog.Add(ctx, &r.LoginLog); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *loginLog) Del(ctx context.Context, d *apiv1.LoginLogReq) (*res.DataRes, error) {
	if err := service.LoginLog.Del(ctx, d.LoginLog.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *loginLog) GetById(ctx context.Context, d *apiv1.LoginLogReq) (*res.DataRes, error) {
	data, err := service.LoginLog.GetById(ctx, d.LoginLog.Id)
	if err != nil {
		return nil, err
	}
	return res.OkData(data)
}

func (h *loginLog) Put(ctx context.Context, p *apiv1.LoginLogReq) (*res.DataRes, error) {
	if err := service.LoginLog.Put(ctx, p.LoginLog); err != nil {
		return nil, err
	}
	return res.Ok()
}
