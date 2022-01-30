package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"interface/internal/consts"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
	"strings"
)

var TimeHistory = NewTimeHistory()

type timeHistory struct{}

func NewTimeHistory() *timeHistory {
	a := timeHistory{}
	return &a
}
func (s *timeHistory) AllByTypeAndTypeDetails(ctx context.Context, i int, i2 int) (gdb.Result, error) {
	//todo
	return nil, nil
}

func (s *timeHistory) NewWriterBirthday(ctx context.Context, name string, ex1 string) (int64, error) {
	if ex1 == "" {
		return 0, consts.ErrWriterEx1IsEmpty
	}
	split := strings.Split(ex1, "/")
	d := entity.TimeHistory{
		Year:        gconv.Int(split[0]),
		Month:       gconv.Int(split[1]),
		Day:         gconv.Int(split[2]),
		Type:        9,
		TypeDetails: 2,
		Status:      1,
		Name:        name,
	}
	d.Century = d.Year/100 + 1
	id, err := dao.TimeHistory.Ctx(ctx).InsertAndGetId(d)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *timeHistory) NewBookPublish(ctx context.Context, year int, month int, name string) (int64, error) {
	if year == 0 {
		return 0, consts.ErrAddBookWithoutYear
	}
	d := entity.TimeHistory{
		Year:        year,
		Type:        7,
		TypeDetails: 4,
		Status:      1,
		Name:        name,
	}
	if month != 0 {
		d.Month = month
	}
	d.Century = d.Year/100 + 1
	id, err := dao.TimeHistory.Ctx(ctx).InsertAndGetId(d)
	if err != nil {
		return 0, err
	}
	return id, nil
}
