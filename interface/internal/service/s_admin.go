package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"interface/internal/consts"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
	"interface/utility/utils/middleware"
	"interface/utility/utils/xjwt"
	"interface/utility/utils/xpwd"
	"interface/utility/utils/xstr"
)

var Admin = NewAdmin()

type admin struct{}

func NewAdmin() *admin {
	a := admin{}
	return &a
}

func (s *admin) List(page, size int, p *entity.Admin) (int, gdb.List) {
	db := g.DB().Model(dao.Admin.Table() + " t1").LeftJoin(dao.Role.Table() + " t2 on t1.rid = t2.id")
	if p.Id != 0 {
		db = db.Where("t1.id", p.Id)
	}
	if p.Uname != "" {
		db = db.WhereLike("t1.uname", xstr.Like(p.Uname))
	}
	if p.Desc != "" {
		db = db.WhereLike("t1.desc", xstr.Like(p.Desc))
	}
	if p.Rid != 0 {
		db = db.Where("t1.rid", p.Rid)
	}
	if p.Status != 0 {
		db = db.Where("t1.status", p.Status)
	}
	count, _ := db.Count()
	db.Fields("t1.*,t2.name role_name")
	all, _ := db.Limit(size).Offset((page - 1) * size).Order("t1.id desc").All()
	if all.IsEmpty() {
		return count, gdb.List{}
	}
	return count, all.List()
}

func (s *admin) Add(ctx context.Context, data *entity.Admin) error {
	if count, err := dao.Admin.Ctx(ctx).Count("uname", data.Uname); err != nil {
		return err
	} else if count != 0 {
		return consts.ErrUnameAlreadyExist
	}
	if _, err := dao.Admin.Ctx(ctx).Insert(data); err != nil {
		return err
	}
	return nil
}

func (s *admin) Del(ctx context.Context, id uint64) error {
	_, err := dao.Admin.Ctx(ctx).Delete("id", id)
	return err
}

func (s *admin) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.Admin.Ctx(ctx).FieldsEx("pwd").One("id", id)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (s *admin) Put(ctx context.Context, data entity.Admin) error {
	m := gconv.Map(data)
	delete(m, "uname")
	delete(m, "pwd")
	if data.Pwd != "" {
		m["pwd"] = xpwd.GenPwd(data.Pwd)
	}
	if data.Ex1 == "true" {
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
	u, err := dao.Admin.FindByUname(ctx, uname)
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
	value := ctx.Value(middleware.Uid)
	g.Log().Info(ctx, value)
	_, err := dao.Admin.Ctx(ctx).Data("pwd", xpwd.GenPwd(pwd)).Where("id", uid).Update()
	if err != nil {
		return err
	}
	return nil
}
