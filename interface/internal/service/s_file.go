package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
	"interface/utility/res"
	"path"
	"time"
)

var File = NewFile()

type file struct{}

func NewFile() *file {
	a := file{}
	return &a
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

func (s *file) Upload(ctx context.Context, r *ghttp.Request) error {
	files := r.GetUploadFiles("files")
	for _, file := range files {
		fileName := fmt.Sprint(grand.S(6), path.Ext(file.Filename))
		file.Filename = fileName
	}
	datePre := time.Now().Format("2006/01")
	group := r.Get("group").String()
	if group == "" || group == "undefined" {
		group = "1"
	}
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
			Group:  gconv.Int(group),
			Status: 1,
		})
		if err != nil {
			res.Error(err, r)
		}
	}
	return nil
}
