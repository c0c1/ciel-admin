package xfile

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

var serviceTemplate = `
package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
)

var $Name$ = New$Name$()

type $name$ struct{}

func New$Name$() *$name$ {
	a := $name${}
	return &a
}

func (s *$name$) List(ctx context.Context,page, size int, p *entity.$Name$) (int, gdb.List) {
	db := g.DB().Model(dao.$Name$.Table())
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

func (s *$name$) Add(ctx context.Context, data *entity.$Name$) error {
	if _, err := dao.$Name$.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *$name$) Del(ctx context.Context, id uint64) error {
	_, err := dao.$Name$.Ctx(ctx).Delete("id", id)
	return err
}

func (s *$name$) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.$Name$.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *$name$) Put(ctx context.Context, data entity.$Name$) error {
	data.UpdatedAt = gtime.Now()
	if _, err := dao.$Name$.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
`

func genService(c TemplateConfig) {
	create, err := gfile.Create(fmt.Sprint(c.RootPath + c.PathService + "/" + c.FileName + ".go"))
	if err != nil {
		panic(err)
	}
	serviceTemplate = strings.ReplaceAll(serviceTemplate, "$Name$", c.EntityName)
	serviceTemplate = strings.ReplaceAll(serviceTemplate, "$name$", gstr.CaseCamelLower(c.EntityName))
	if _, err = create.WriteString(serviceTemplate); err != nil {
		panic(err)
	}
	if err = create.Close(); err != nil {
		panic(err)
	}
}
