package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var RoleMenu = roleMenu{}

type roleMenu struct {
}

func (h *roleMenu) List(ctx context.Context, d *apiv1.RoleMenuReq) (*res.PageRes, error) {
	total, result := service.RoleMenu.List(d.Page, d.PageSize, &d.RoleMenu)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *roleMenu) Add(ctx context.Context, r *apiv1.RoleMenuAddReq) (*res.DataRes, error) {
	if err := service.RoleMenu.Add(ctx, r.Rid, r.Mid); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *roleMenu) Del(ctx context.Context, d *apiv1.RoleMenuReq) (*res.DataRes, error) {
	if err := service.RoleMenu.Del(ctx, d.RoleMenu.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *roleMenu) GetById(ctx context.Context, d *apiv1.RoleMenuReq) (*res.DataRes, error) {
	data, err := service.RoleMenu.GetById(ctx, d.RoleMenu.Id)
	if err != nil {
		return nil, err
	}
	return res.OkData(data)
}

func (h *roleMenu) Put(ctx context.Context, p *apiv1.RoleMenuReq) (*res.DataRes, error) {
	if err := service.RoleMenu.Put(ctx, p.RoleMenu); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *roleMenu) NoMenus(ctx context.Context, p *apiv1.RoleMenuReq) (*res.DataRes, error) {
	data := service.RoleMenu.NoMenus(ctx, p.Rid)
	return res.OkData(data)
}
