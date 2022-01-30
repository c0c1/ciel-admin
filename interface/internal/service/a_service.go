package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"interface/utility/utils/xstr"
)

type Condition struct {
	Field string
	Value interface{}
	Like  bool // false:where true: whereLike
}
type SearchConf struct {
	Table      string
	Page, Size int
	Conditions []*Condition
	OrderBy    string
	FieldsEx   string
}

func List(ctx context.Context, s SearchConf) (int, gdb.List) {
	db := g.DB().Model(s.Table).Ctx(ctx)
	if len(s.Conditions) > 0 {
		for _, item := range s.Conditions {
			if g.IsEmpty(item.Value) {
				continue
			}
			if item.Like {
				db = db.WhereLike(item.Field, xstr.Like(gconv.String(item.Value)))
			} else {
				db = db.Where(item.Field, item.Value)
			}

		}
	}
	count, _ := db.Count()
	var o = "id desc"
	if s.OrderBy != "" {
		o = s.OrderBy
	}
	if s.FieldsEx != "" {
		db.FieldsEx(s.FieldsEx)
	}
	all, _ := db.Limit(s.Size).Offset((s.Page - 1) * s.Size).Order(o).All()
	if all.IsEmpty() {
		return 0, gdb.List{}
	}
	return count, all.List()
}
func QueryCondition(filed string, r *ghttp.Request, like ...bool) *Condition {
	condition := Condition{Field: filed, Value: r.GetQuery(filed)}
	if len(like) > 0 {
		condition.Like = true
	}
	return &condition
}
func Add(ctx context.Context, table, data interface{}) error {
	_, err := g.DB().Ctx(ctx).Model(table).Insert(data)
	if err != nil {
		glog.Error(ctx, err)
		return err
	}
	return nil
}
func Del(ctx context.Context, table, id interface{}) (err error) {
	if _, err = g.DB().Ctx(ctx).Model(table).Delete("id", id); err != nil {
		glog.Error(ctx, err)
		return
	}
	return
}
func Update(ctx context.Context, table string, id, data interface{}) error {
	_, err := g.DB().Update(ctx, table, data, "id", id)
	if err != nil {
		glog.Error(ctx, err)
		return err
	}
	return nil
}
func GetById(ctx context.Context, table, id interface{}) (gdb.Record, error) {
	one, err := g.DB().Ctx(ctx).Model(table).One("id", id)
	if err != nil {
		glog.Error(ctx, err)
		return nil, err
	}
	return one, nil
}
