package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var RoleApi = roleApi{}

type roleApi struct {
}

func (h *roleApi) List(ctx context.Context, d *apiv1.RoleApiReq) (*res.PageRes, error) {
	total, result := service.RoleApi.List(d.Page, d.PageSize, &d.RoleApi)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *roleApi) Add(ctx context.Context, r *apiv1.RoleApiAddReq) (*res.DataRes, error) {
	if err := service.RoleApi.Add(ctx, r.Rid, r.Aid); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *roleApi) Del(ctx context.Context, d *apiv1.RoleApiReq) (*res.DataRes, error) {
	if err := service.RoleApi.Del(ctx, d.RoleApi.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *roleApi) GetById(ctx context.Context, d *apiv1.RoleApiReq) (*res.DataRes, error) {
	data, err := service.RoleApi.GetById(ctx, d.RoleApi.Id)
	if data != nil {
		return nil, err
	}
	return res.OkData(data)
}

func (h *roleApi) Put(ctx context.Context, p *apiv1.RoleApiReq) (*res.DataRes, error) {
	if err := service.RoleApi.Put(ctx, p.RoleApi); err != nil {
		return nil, err
	}
	return res.Ok()
}
func (h roleApi) NoApis(ctx context.Context, p *apiv1.RoleApiReq) (*res.DataRes, error) {
	apis := service.RoleApi.NoApis(ctx, p.Rid)
	return res.OkData(apis)
}
