package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
	"interface/utility/utils/xstr"
)

var Menu = NewMenu()

type menu struct{}

func NewMenu() *menu {
	a := menu{}
	return &a
}

func (s *menu) List(page, size int, p *entity.Menu) (int, gdb.List) {
	db := g.DB().Model(dao.Menu.Table())
	if p.Id != 0 {
		db = db.Where("id", p.Id)
	}
	if p.Pid != 0 {
		db = db.Where("pid", p.Pid)
	}
	if p.Name != "" {
		db = db.WhereLike("name", xstr.Like(p.Name))
	}
	if p.Status != 0 {
		db = db.Where("status", p.Status)
	}
	count, _ := db.Count()
	all, _ := db.Limit(size).Offset((page - 1) * size).Order("sort desc,id desc").All()
	if all.IsEmpty() {
		return count, gdb.List{}
	}
	return count, all.List()
}

func (s *menu) Add(ctx context.Context, data *entity.Menu) error {
	if _, err := dao.Menu.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *menu) Del(ctx context.Context, id uint64) error {
	_, err := dao.Menu.Ctx(ctx).Delete("id", id)
	return err
}

func (s *menu) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.Menu.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (s *menu) Put(ctx context.Context, data entity.Menu) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.Menu.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
