package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
)

var LoginLog = NewLoginLog()

type loginLog struct{}

func NewLoginLog() *loginLog {
	a := loginLog{}
	return &a
}

func (s *loginLog) List(ctx context.Context, page, size int, p *entity.LoginLog) (int, gdb.List) {
	g.Log().Info(ctx, "123")
	db := g.DB().Model(dao.LoginLog.Table())
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

func (s *loginLog) Add(ctx context.Context, data *entity.LoginLog) error {
	if _, err := dao.LoginLog.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *loginLog) Del(ctx context.Context, id uint64) error {
	_, err := dao.LoginLog.Ctx(ctx).Delete("id", id)
	return err
}

func (s *loginLog) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.LoginLog.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *loginLog) Put(ctx context.Context, data entity.LoginLog) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.LoginLog.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
