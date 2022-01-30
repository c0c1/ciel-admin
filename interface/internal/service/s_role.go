package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"interface/internal/model/bo"
	"interface/internal/service/internal/dao"
)

var Role = NewRole()

type role struct{}

func NewRole() *role {
	a := role{}
	return &a
}

func (s *role) Menus(rid uint64, pid int) ([]*bo.Menu, error) {
	var data []*bo.Menu
	// 查询一级菜单
	if err := g.DB().Model("s_role_menu t1").LeftJoin("s_menu t2 on t1.mid = t2.id").Fields("t2.*").Order("sort").Scan(&data, "t1.rid=? and t2.pid= ?", rid, pid); err != nil {
		return nil, err
	}
	for _, item := range data {
		if item.Type == 1 {
			children, err := s.Menus(rid, int(item.Id))
			if err != nil {
				return nil, err
			}
			item.Children = children
			if len(item.Children) == 0 {
				continue
			}
		}
	}
	return data, nil
}

func (s *role) All(ctx context.Context) (gdb.Result, error) {
	all, err := dao.Role.Ctx(ctx).Fields("id,name").All()
	if err != nil {
		return nil, err
	}
	return all, nil
}
