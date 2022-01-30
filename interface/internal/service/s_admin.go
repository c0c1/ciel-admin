package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"interface/internal/consts"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
	"interface/utility/utils/xjwt"
	"interface/utility/utils/xpwd"
)

var Admin = NewAdmin()

type admin struct{}

func NewAdmin() *admin {
	a := admin{}
	return &a
}

func (s *admin) Put(ctx context.Context, data *entity.Admin) error {
	m := gconv.Map(data)
	delete(m, "uname")
	delete(m, "pwd")
	if data.Pwd != "" {
		m["pwd"] = xpwd.GenPwd(data.Pwd)
	}
	glog.Info(nil, data.Ex1)
	if data.Ex1 == "1" {
		count, err := dao.Admin.Ctx(ctx).Count("uname", data.Uname)
		if err != nil {
			return err
		}
		if count != 0 {
			return consts.ErrUnameAlreadyExist
		}
		m["uname"] = data.Uname
	}
	data.UpdatedAt = gtime.Now()
	if _, err := dao.Admin.Ctx(ctx).Update(m, "id", data.Id); err != nil {
		return err
	}
	return nil
}

func (s *admin) Login(ctx context.Context, uname, pwd string) (g.Map, error) {
	u, err := s.FindByUname(ctx, uname)
	if err != nil {
		return nil, consts.ErrLogin
	}
	if u.Status == consts.UserStatusClose {
		return nil, consts.ErrClose
	}
	if !xpwd.ComparePassword(u.Pwd, pwd) {
		return nil, consts.ErrLogin
	}
	menus, err := Role.Menus(u.Rid, -1)
	if err != nil {
		return nil, err
	}
	token, err := xjwt.GenToken(u.Uname, u.Id, u.Rid)
	if err != nil {
		return nil, err
	}
	return g.Map{
		"u":     u,
		"menus": menus,
		"token": token,
	}, nil
}

func (s *admin) UpdatePwd(ctx context.Context, pwd string, uid uint64) error {
	_, err := dao.Admin.Ctx(ctx).Data("pwd", xpwd.GenPwd(pwd)).Where("id", uid).Update()
	if err != nil {
		return err
	}
	return nil
}

func (s *admin) FindByUname(ctx context.Context, uname string) (*entity.Admin, error) {
	var d entity.Admin
	err := dao.Admin.Ctx(ctx).Scan(&d, "uname", uname)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
