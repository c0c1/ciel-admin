package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestMenus(t *testing.T) {
	menus, err := Role.Menus(8, 0)
	if err != nil {
		panic(err)
	}
	g.Log().Info(nil, menus)
}
