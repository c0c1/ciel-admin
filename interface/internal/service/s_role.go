package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"interface/internal/model/bo"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
)

var Role = NewRole()

type role struct{}

func NewRole() *role {
	a := role{}
	return &a
}

func (s *role) List(page, size int, p *entity.Role) (int, gdb.List) {
	db := g.DB().Model(dao.Role.Table())
	if p.Id != 0 {
		db = db.Where("id", p.Id)
	}
	count, _ := db.Count()
	all, _ := db.Limit(size).Offset((page - 1) * size).Order("id desc").All()
	if all.IsEmpty() {
		return count, gdb.List{}
	}
	return count, all.List()
}

func (s *role) Add(ctx context.Context, data *entity.Role) error {
	if _, err := dao.Role.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *role) Del(ctx context.Context, id uint64) error {
	_, err := dao.Role.Ctx(ctx).Delete("id", id)
	return err
}

func (s *role) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.Role.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (s *role) Put(ctx context.Context, data entity.Role) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.Role.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
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
