package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestGetById(t *testing.T) {
	data, err := GetById(gctx.New(), "s_api", "2")
	if err != nil {
		panic(err)
	}
	g.Dump(data)
}
