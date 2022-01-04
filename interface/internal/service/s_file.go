package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"interface/apiv1"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
	"interface/utility/utils/res"
	"interface/utility/utils/xstr"
	"path"
	"time"
)

var File = NewFile()

type file struct{}

func NewFile() *file {
	a := file{}
	return &a
}

func (s *file) List(page, size int, p *entity.File) (int, gdb.List) {
	db := g.DB().Model(dao.File.Table())
	if p.Id != 0 {
		db = db.Where("id", p.Id)
	}
	if p.Name != "" {
		db = db.WhereLike("name", xstr.Like(p.Name))
	}
	if p.Group != "" {
		db = db.WhereLike("group", xstr.Like(p.Group))
	}
	count, _ := db.Count()
	all, _ := db.Limit(size).Offset((page - 1) * size).Order("id desc").All()
	if all.IsEmpty() {
		return count, gdb.List{}
	}
	return count, all.List()
}

func (s *file) Add(ctx context.Context, data *entity.File) error {
	if _, err := dao.File.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *file) Del(ctx context.Context, id uint64) error {
	value, _ := dao.File.Ctx(ctx).Value("name", "id", id)
	path, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return err
	}
	p := path.String() + "/" + value.String()
	if gfile.Exists(p) {
		_ = gfile.Remove(p)
	}
	_, err = dao.File.Ctx(ctx).Delete("id", id)
	return err
}

func (s *file) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.File.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *file) Put(ctx context.Context, data entity.File) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.File.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}

func (s *file) Upload(ctx context.Context, p *apiv1.FileReq) error {
	r := g.RequestFromCtx(ctx)
	files := r.GetUploadFiles("files")
	for _, file := range files {
		fileName := fmt.Sprint(grand.S(6), path.Ext(file.Filename))
		file.Filename = fileName
	}
	datePre := time.Now().Format("2006/01")
	group := r.Get("group").String()
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return err
	}
	mixPath := fmt.Sprintf("%s/%s/%s/", rootFilePath, group, datePre)
	_, err = files.Save(mixPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		dbName := fmt.Sprintf("%s/%s/%s", group, datePre, file.Filename)
		_, err := dao.File.Ctx(ctx).Insert(entity.File{
			Name:   dbName,
			Group:  group,
			Status: 1,
		})
		if err != nil {
			res.ErrMsg(err.Error(), r)
		}
	}
	return nil
}
