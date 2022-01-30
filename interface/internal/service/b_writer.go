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

var Writer = NewWriter()

type writer struct{}

func NewWriter() *writer {
	a := writer{}
	return &a
}
func (s *writer) List(ctx context.Context, page, size int, p *entity.Writer) (int, gdb.List) {
	db := g.DB().Model(dao.Writer.Table())
	if p.Id != 0 {

		db = db.Where("id", p.Id)
	}
	if p.Name != "" {
		db = db.WhereLike("name", xstr.Like(p.Name))
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
func (s *writer) Add(ctx context.Context, data *entity.Writer) error {
	timeId, err := TimeHistory.NewWriterBirthday(ctx, data.Name, data.Ex1)
	if err != nil {
		return err
	}
	data.Tid = uint64(timeId)
	if _, err := dao.Writer.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *writer) Del(ctx context.Context, id uint64) error {
	_, err := dao.Writer.Ctx(ctx).Delete("id", id)
	return err
}

func (s *writer) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.Writer.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *writer) Put(ctx context.Context, data entity.Writer) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.Writer.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}

func (s *writer) All(ctx context.Context) (gdb.Result, error) {
	return dao.Writer.Ctx(ctx).All()
}
