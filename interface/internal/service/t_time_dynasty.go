
package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
)

var TimeDynasty = NewTimeDynasty()

type timeDynasty struct{}

func NewTimeDynasty() *timeDynasty {
	a := timeDynasty{}
	return &a
}

func (s *timeDynasty) List(ctx context.Context,page, size int, p *entity.TimeDynasty) (int, gdb.List) {
	db := g.DB().Model(dao.TimeDynasty.Table())
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

func (s *timeDynasty) Add(ctx context.Context, data *entity.TimeDynasty) error {
	if _, err := dao.TimeDynasty.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *timeDynasty) Del(ctx context.Context, id uint64) error {
	_, err := dao.TimeDynasty.Ctx(ctx).Delete("id", id)
	return err
}

func (s *timeDynasty) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.TimeDynasty.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *timeDynasty) Put(ctx context.Context, data entity.TimeDynasty) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.TimeDynasty.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
