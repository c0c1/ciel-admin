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

var User = NewUser()

type user struct{}

func NewUser() *user {
	a := user{}
	return &a
}

func (s *user) List(page, size int, p *entity.User) (int, gdb.List) {
	db := g.DB().Model(dao.User.Table())
	if p.Id != 0 {
		db = db.Where("id", p.Id)
	}
	if p.Uname != "" {
		db = db.WhereLike("uname", xstr.Like(p.Uname))
	}
	if p.Status != 0 {
		db = db.Where("status", p.Status)
	}
	count, _ := db.Count()
	all, _ := db.Limit(size).Offset((page - 1) * size).Order("id desc").All()
	if all.IsEmpty() {
		return count, gdb.List{}
	}
	return count, all.List()
}

func (s *user) Add(ctx context.Context, data *entity.User) error {
	if _, err := dao.User.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *user) Del(ctx context.Context, id uint64) error {
	_, err := dao.User.Ctx(ctx).Delete("id", id)
	return err
}

func (s *user) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.User.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *user) Put(ctx context.Context, data entity.User) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.User.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
