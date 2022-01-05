package xfile

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

var handleTemplate = `package handler

import (
	"context"
	"interface/apiv1"
	"interface/internal/service"
	"interface/utility/utils/res"
)

var $Name$ = $name${}

type $name$ struct {
}

func (h *$name$) List(ctx context.Context, d *apiv1.$Name$Req) (*res.PageRes, error) {
	total, result := service.$Name$.List(ctx,d.Page, d.PageSize, &d.$Name$)
	return res.NewPageRes(result, total, d.Page, d.PageSize), nil
}

func (h *$name$) Add(ctx context.Context, r *apiv1.$Name$Req) (*res.DataRes, error) {
	if err := service.$Name$.Add(ctx, &r.$Name$); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *$name$) Del(ctx context.Context, d *apiv1.$Name$Req) (*res.DataRes, error) {
	if err := service.$Name$.Del(ctx, d.$Name$.Id); err != nil {
		return nil, err
	}
	return res.Ok()
}

func (h *$name$) GetById(ctx context.Context, d *apiv1.$Name$Req) (*res.DataRes, error) {
	data, err := service.$Name$.GetById(ctx, d.$Name$.Id)
	if err != nil {
		return nil, err
	}
	return res.OkData(data)
}

func (h *$name$) Put(ctx context.Context, p *apiv1.$Name$Req) (*res.DataRes, error) {
	if err := service.$Name$.Put(ctx, p.$Name$); err != nil {
		return nil, err
	}
	return res.Ok()
}
`

func genHandler(c TemplateConfig) {
	file, err := gfile.Create(fmt.Sprint(c.RootPath + "/interface/internal/handler/" + c.FileName + ".go"))
	if err != nil {
		panic(err)
	}
	template := handleTemplate
	template = strings.ReplaceAll(template, "$Name$", c.EntityName)
	template = strings.ReplaceAll(template, "$name$", gstr.CaseCamelLower(c.EntityName))
	if _, err = file.WriteString(template); err != nil {
		panic(err)
	}
	if err = file.Close(); err != nil {
		panic(err)
	}
}
