package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"interface/internal/service/internal/dao"
)

var Book = NewBook()

type book struct{}

func NewBook() *book {
	a := book{}
	return &a
}

func (s *book) ListNewTop100(ctx context.Context) (gdb.Result, error) {
	return dao.Book.Ctx(ctx).Limit(100).OrderDesc("id").All()

}

func (s *book) ListAllBookIds(ctx context.Context) (gdb.Result, error) {
	return dao.Book.Ctx(ctx).Fields("id,name").All("status !=2")
}
