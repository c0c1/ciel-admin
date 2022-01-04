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

var Dict = NewDict()

type dict struct{}

func NewDict() *dict {
	a := dict{}
	return &a
}

func (s *dict) List(page, size int, p *entity.Dict) (int, gdb.List) {
	db := g.DB().Model(dao.Dict.Table())
	if p.Id != 0 {
		db = db.Where("id", p.Id)
	}
	if p.K != "" {
		db = db.WhereLike("k", xstr.Like(p.K))
	}
	if p.V != "" {
		db = db.WhereLike("v", xstr.Like(p.V))
	}
	if p.Desc != "" {
		db = db.WhereLike("desc", xstr.Like(p.Desc))
	}
	if p.Group != "" {
		db = db.WhereLike("group", xstr.Like(p.Group))
	}
	if p.Type != 0 {
		db = db.Where("type", p.Type)
	}
	count, _ := db.Count()
	all, _ := db.Limit(size).Offset((page - 1) * size).Order("id desc").All()
	if all.IsEmpty() {
		return count, gdb.List{}
	}
	return count, all.List()
}

func (s *dict) Add(ctx context.Context, data *entity.Dict) error {
	if _, err := dao.Dict.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *dict) Del(ctx context.Context, id uint64) error {
	_, err := dao.Dict.Ctx(ctx).Delete("id", id)
	return err
}

func (s *dict) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.Dict.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *dict) Put(ctx context.Context, data entity.Dict) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.Dict.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
