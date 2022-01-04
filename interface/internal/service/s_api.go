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

var Api = NewApi()

type api struct{}

func NewApi() *api {
	a := api{}
	return &a
}

func (s *api) List(page, size int, p *entity.Api) (int, gdb.List) {
	db := g.DB().Model(dao.Api.Table())
	if p.Id != 0 {
		db = db.Where("id", p.Id)
	}
	if p.Status != 0 {
		db = db.Where("status", p.Status)
	}
	if p.Url != "" {
		db = db.WhereLike("url", xstr.Like(p.Url))
	}
	if p.Method != "" {
		db = db.Where("method", p.Method)
	}
	if p.Group != "" {
		db = db.WhereLike("group", xstr.Like(p.Group))
	}
	if p.Desc != "" {
		db = db.WhereLike("desc", xstr.Like(p.Desc))
	}
	count, _ := db.Count()
	all, _ := db.Limit(size).Offset((page - 1) * size).Order("id desc").All()
	if all.IsEmpty() {
		return count, gdb.List{}
	}
	return count, all.List()
}

func (s *api) Add(ctx context.Context, data *entity.Api) error {
	if err := dao.Api.Insert(ctx, data); err != nil {
		return err
	}
	return nil
}

func (s *api) Del(ctx context.Context, id uint64) error {
	_, err := dao.Api.Ctx(ctx).Delete("id", id)
	return err
}

func (s *api) GetById(ctx context.Context, id uint64) (*entity.Api, error) {
	one, err := dao.Api.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	var res entity.Api
	if err = one.Struct(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *api) Put(ctx context.Context, data entity.Api) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.Api.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
