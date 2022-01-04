package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
)

var RoleMenu = NewRoleMenu()

type roleMenu struct{}

func NewRoleMenu() *roleMenu {
	a := roleMenu{}
	return &a
}

func (s *roleMenu) List(page, size int, p *entity.RoleMenu) (int, gdb.List) {
	rid := p.Rid
	mid := p.Mid

	db := g.DB().Model(dao.RoleMenu.Table() + " t1").
		LeftJoin(dao.Role.Table() + " t2 on t1.rid = t2.id").
		LeftJoin(dao.Menu.Table() + " t3 on t1.mid = t3.id")
	if rid != 0 {
		db = db.Where("t1.rid", rid)
	}
	if mid != 0 {
		db = db.Where("t1.mid", mid)
	}
	count, _ := db.Count()
	db.Fields("t2.name r_name,t3.name m_name,t1.id id")
	all, _ := db.Limit(size).Offset((page - 1) * size).Order("t1.id desc").All()
	if all.IsEmpty() {
		return count, gdb.List{}
	}
	return count, all.List()
}

func (s *roleMenu) Add(ctx context.Context, rid int, mids []int) error {
	var d []*entity.RoleMenu
	for _, item := range mids {
		if item != 0 {
			d = append(d, &entity.RoleMenu{Rid: uint64(rid), Mid: uint64(item)})
		}
	}
	if _, err := dao.RoleMenu.Ctx(ctx).Batch(len(mids)).Replace(d); err != nil {
		return err
	}
	return nil
}

func (s *roleMenu) Del(ctx context.Context, id uint64) error {
	_, err := dao.RoleMenu.Ctx(ctx).Delete("id", id)
	return err
}

func (s *roleMenu) GetById(ctx context.Context, id uint64) (gdb.Record, error) {
	one, err := dao.RoleMenu.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	return one, err
}

func (s *roleMenu) Put(ctx context.Context, data entity.RoleMenu) error {
	if _, err := dao.RoleMenu.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}

func (s *roleMenu) NoMenus(ctx context.Context, rid uint64) interface{} {
	array, _ := dao.RoleMenu.Ctx(ctx).Array("mid", "rid", rid)
	if len(array) == 0 {
		all, _ := dao.Menu.Ctx(ctx).All()
		if all.IsEmpty() {
			return make([]int, 0)
		}
		return all
	}
	all, _ := dao.Menu.Ctx(ctx).WhereNotIn("id", array).Fields("id,name").All()
	return all
}
