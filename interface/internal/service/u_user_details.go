
package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
)

var UserDetails = NewUserDetails()

type userDetails struct{}

func NewUserDetails() *userDetails {
	a := userDetails{}
	return &a
}

func (s *userDetails) List(ctx context.Context,page, size int, p *entity.UserDetails) (int, gdb.List) {
	db := g.DB().Model(dao.UserDetails.Table())
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

func (s *userDetails) Add(ctx context.Context, data *entity.UserDetails) error {
	if _, err := dao.UserDetails.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *userDetails) Del(ctx context.Context, id uint64) error {
	_, err := dao.UserDetails.Ctx(ctx).Delete("id", id)
	return err
}

func (s *userDetails) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.UserDetails.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *userDetails) Put(ctx context.Context, data entity.UserDetails) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.UserDetails.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
