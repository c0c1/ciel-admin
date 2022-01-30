package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"interface/internal/service/internal/dao"
)

var BookChapter = NewBookChapter()

type bookChapter struct{}

func NewBookChapter() *bookChapter {
	a := bookChapter{}
	return &a
}
func (s *bookChapter) ListChapter(ctx context.Context, id string) (gdb.Result, error) {
	all, err := dao.BookChapter.Ctx(ctx).Fields("id,book_id,name,hidden_words,type").Order("sort desc,id").All("book_id", id)
	if err != nil {
		return nil, err
	}
	return all, nil
}
