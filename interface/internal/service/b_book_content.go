package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
)

var BookContent = NewBookContent()

type bookContent struct{}

func NewBookContent() *bookContent {
	a := bookContent{}
	return &a
}

func (s *bookContent) List(ctx context.Context, page, size int, p *entity.BookContent) (int, gdb.List) {
	db := g.DB().Model(dao.BookContent.Table())
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

func (s *bookContent) Add(ctx context.Context, data *entity.BookContent) error {
	if _, err := dao.BookContent.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *bookContent) Del(ctx context.Context, id uint64) error {
	_, err := dao.BookContent.Ctx(ctx).Delete("id", id)
	return err
}

func (s *bookContent) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.BookContent.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *bookContent) Put(ctx context.Context, data entity.BookContent) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.BookContent.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}

func (s *bookContent) GetByChapterId(ctx context.Context, bookId int64, chapterId int64) (gdb.Record, error) {
	if chapterId != 0 {
		return g.DB().Model(dao.BookChapter.Table()+" t1").LeftJoin(dao.BookContent.Table()+" t2 on t1.id = t2.chapter_id").
			LeftJoin(dao.Book.Table()+" t3 on t1.book_id = t3.id").
			Fields("t1.name,t2.content,t2.summary,t3.icon,t3.name bookName,t3.theme").
			One("t2.id", chapterId)
	}
	return g.DB().Model(dao.BookChapter.Table()+" t1").LeftJoin(dao.BookContent.Table()+" t2 on t1.id = t2.chapter_id").
		LeftJoin(dao.Book.Table()+" t3 on t1.book_id = t3.id").
		OrderAsc("t2.id").
		Fields("t1.name,t2.content,t2.summary,t3.icon,t3.name bookName,t3.theme").
		One("t1.book_id", bookId)
}
