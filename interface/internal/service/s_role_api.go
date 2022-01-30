package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"interface/internal/model/entity"
	"interface/internal/service/internal/dao"
)

var RoleApi = NewRoleApi()

type roleApi struct{}

func NewRoleApi() *roleApi {
	a := roleApi{}
	return &a
}

func (s *roleApi) List(page, size int, rid, aid int) (int, gdb.List) {
	db := g.DB().Model(dao.RoleApi.Table() + " t1").
		LeftJoin(dao.Role.Table() + " t2 on t2.id = t1.rid").
		LeftJoin(dao.Api.Table() + " t3 on t3.id = t1.aid")
	if rid != 0 {
		db = db.Where("t1.rid", rid)
	}
	if aid != 0 {
		db = db.Where("t1.aid", aid)
	}
	count, _ := db.Count()
	db.Fields("t1.id id,t2.name role_name")
	db.Fields("t3.url,t3.method,t3.group,t3.desc")
	all, _ := db.Limit(size).Offset((page - 1) * size).Order("t1.id desc,t3.group").All()
	if all.IsEmpty() {
		return count, gdb.List{}
	}
	return count, all.List()
}

func (s *roleApi) Add(ctx context.Context, rid int, aids []int) error {
	var d []*entity.RoleApi
	for _, aid := range aids {
		d = append(d, &entity.RoleApi{
			Rid: uint64(rid),
			Aid: uint64(aid),
		})
	}
	if _, err := dao.RoleApi.Ctx(ctx).Batch(len(aids)).Replace(d); err != nil {
		return err
	}
	return nil
}

func (s *roleApi) NoApis(ctx context.Context, rid uint64) interface{} {
	array, _ := dao.RoleApi.Ctx(ctx).Array("aid", "rid", rid)
	if len(array) == 0 {
		all, _ := dao.Api.Ctx(ctx).All()
		if all.IsEmpty() {
			return make([]int, 0)
		}
		return all
	}
	all, _ := dao.Api.Ctx(ctx).WhereNotIn("id", array).All()
	return all
}
