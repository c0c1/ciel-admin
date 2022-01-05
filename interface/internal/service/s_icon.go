package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
)

var Icon = NewIcon()

type icon struct{}

func NewIcon() *icon {
	a := icon{}
	return &a
}

func (s *icon) List(page, size int, p *entity.Icon) (int, gdb.List) {
	db := g.DB().Model(dao.Icon.Table())
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

func (s *icon) Add(ctx context.Context, data *entity.Icon) error {
	data.Status = 1
	if _, err := dao.Icon.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *icon) Del(ctx context.Context, id uint64) error {
	_, err := dao.Icon.Ctx(ctx).Delete("id", id)
	return err
}

func (s *icon) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.Icon.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *icon) Put(ctx context.Context, data entity.Icon) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.Icon.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
