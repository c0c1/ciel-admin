# ciel_admin

锲而舍之朽木不折，锲而不舍金石可镂

登录页面 <br>
![img](document/login.png)

导航 <br>
![img](document/menu.png)

详情页面  <br>
![img](document/menu_list.png)

修改页面 <br>
![img](document/update.png)

## 说明

ciel-admin 是一个简洁的后台管理系统，提供了简单的权限管理功能，和本地文件上传。

前端页面风格参考 https://www.random.org/

使用到的技术如下

- 前端 react,next.js,axios,bootstrap,swr,react-toastify,node-sass
- 后端 goframe 2.0

真的追求简洁，一个表的CURD 前端+后端代码不到100行。例如菜单页面

前端 menu.js

```text
import DefaultPage from "../../components/defaultPage";

const Menu = () => <DefaultPage
    fields={[
        {field: 'id', editHidden: 1},                                              // filed 字段名 ,editHidden:编辑时是否隐藏（id一般不允许编辑）
        {field: 'pid', search: 1, type: 'number'},                                 // type: 默认为 input 可不填，可选 number|select(当为select时需要指定items)  
        {field: 'name', search: 1},                                                // search: 是否搜索
        {field: 'sort', type: 'number', step: '0.1', detailsDesc: "请给一个排序"}, // detailsDesc  input框的简要说明
        {
            field: 'type', title: '类型', search: 1, type: "select",   
            items: [
                {
                    label: 'group',                                                // 可选时的标签
                    value: 1,                                                      // 对应数据库的值
                    bgGround: "#80deea"                                            // 背景颜色，不给则使用系统自动匹配   
                },
                {
                    label: 'point',
                    value: 2,
                    bgGround: "#e0e0e0"
                }
            ]
        },
        {field: 'path'},
        {field: 'status', type: 'select', items: [{label: 'yes', value: 1}, {label: 'no', value: 2}]},
        {field: 'created_at', editHidden: 1},
    ]} name={'Menu'} urlPrefix={'menu'}/>
export default Menu
```

只需要将数据库中使用到自动添加到js中即可。

后端 s_menu.go

```text
package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var Menu = menu{table: "s_menu"}

type menu struct {
	table string
}

func (c *menu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	total, data := service.List(gctx.New(), service.SearchConf{
		Table: c.table,
		Page:  page, Size: size,
		Conditions: []*service.Condition{
			{Field: "id", Value: r.GetQuery("id")},
			{Field: "pid", Value: r.GetQuery("pid")},
			{Field: "type", Value: r.GetQuery("type")},
			{Field: "status", Value: r.GetQuery("status")},
			{Field: "name", Value: r.GetQuery("name"), Like: true},
		},
		OrderBy: "sort desc,id desc",
	})
	res.OkPage(data, total, page, size, r)
}

func (c *menu) Add(r *ghttp.Request) {
	var d entity.Menu
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *menu) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *menu) Put(r *ghttp.Request) {
	var d entity.Menu
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *menu) GetById(r *ghttp.Request) {
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}
```

## 主要的两个文件

当前简洁的背后，主要在于将相同的操作都分装在了两个文件中。

### 前端 defaultPage.js

```js
import useSWR from "swr";
import {useEffect, useState} from "react";
import {toast} from "react-toastify";
import Meta from "./meta";
import {Nav} from "./nav";
import {TableFooter} from "./table";
import {defaultPageParam} from "../data/const";
import request from "../utils/request";
import {DetailsLi2} from "./detailsLi";
import {SearchLi} from "./searchLi";
import {TableTd} from "./tableTd";

// defaultPage
const DefaultPage = ({
                         name,      //must 页面标题名
                         fields,    //must 字段信息
                         urlPrefix, //must 请求前缀
                         noEdit,    // option  没有编辑按钮     1则没有 0 或不不传则有编辑按钮
                         noNav,     // option  没有导航条？     同上
                         noSearch,  // option  没有寿搜索按钮？ 同上
                         query,     // option  默认查询参数  用法见 Role角色页面
                     }) => {
    const fetcher = async ({url, params}) => request({url: url, method: 'get', params}),                     // 获取列表请求
        update = async (data) => request({url: `/${urlPrefix}/update`, method: 'put', data}),                // 修改请求
        add = async (data) => request({url: `/${urlPrefix}/add`, method: 'post', data}),                     // 添加请求
        del = async (id) => request({url: `/${urlPrefix}/del?id=${id}`, method: 'delete'})                   // 删除请求
    const [params, setParams] = useState({...defaultPageParam, ...query}),                                   // 查询参数
        [paramsTemp, setParamsTemp] = useState(params),                                                      // 查询临时参数
        [details, setDetails] = useState({showType: "main"}),                                                // 修改和添加参数
        {data, mutate} = useSWR({url: `/${urlPrefix}/list`, params: params}, fetcher),                       // swr 查询（用法请看 swr官网）
        {currPage, list, totalPage, totalCount} = data ? data.data : ""                                      // 返回数据
    useEffect(() => setParams({...params, ...query}), [query])                                               // 到用到 query时将查询参数添加到 params中
    const onSearch = () => setParams(paramsTemp),                                                            // 搜索操作
        onChange = (e, name) => setParamsTemp({...paramsTemp, [name]: e.target.value}),                      // 查询临时参数变动时
        onDetailsChange = (e, name) => setDetails({...details, [name]: e.target.value}),                     // 修改和添加参数变动时
        onKeyDown = (e) => (e.key === 'Enter') ? params.page = 1 && onSearch() : "",                         // 搜索键盘事件enter事件 
        onDelete = async (id) => (confirm("Are you sure?")) ? await del(id) && toast('ok') && mutate() : "", // 删除操作
        onSubmit = async () => {                                                                             // 添加或修改操作
            let {code, msg} = details.showType === "update" ? await update(details) : await add(details)     // 判断修改还是添加
            if (code === 0) {
                toast(msg)
                mutate()
                setDetails({showType: "main"})
            }
        }
    let tbody = ''
    let search = <ul>// 搜索部分
        {fields.filter(item => item.search).map(item =>
            <SearchLi key={item.field}
                      field={item.field} label={item.title}
                      type={item.type} items={item.items}
                      onChange={onChange} paramsTemp={paramsTemp} onKeyDown={onKeyDown}
            />
        )}
    </ul>
    if (list) {                                                                                              // 当数据存在则填充tbody
        tbody = list.map(item => <tr key={item.id}>
            {fields.filter(item => !item.hidden).map(i => (<TableTd key={i.field + 's'} value={item[i.field]} type={i.type} items={i.items}/>))}
            <td>
                <input hidden={noEdit} type={'submit'} value={'Edit'} onClick={() => setDetails({...item, showType: "update"})}/>
                <input type={'submit'} value={'Del'} onClick={() => onDelete(item.id)}/>
            </td>
        </tr>)
    }
    return <>
        <Meta/>
        {noNav ? '' : <Nav/>}
        <header className={' hr'}>
            <h1>{name}</h1>
            <section>
                {search}
                <ul>
                    <input hidden={noSearch} type="submit" value={'Search'} onClick={() => onSearch()}/>
                    <input type="submit" value={`Add ${name}`} onClick={() => setDetails({showType: "add", rid: ''})}/>
                </ul>
            </section>
        </header>
        <main className={' hr'} hidden={!(details.showType === "main")}>// 主数据区域
            <table>
                <thead>
                    <tr>
                        {[...fields.filter(item => !item.hidden), {field: "Operation"}].map(item =>
                            <th key={item.field}>{item.title ? item.title : item.field} </th>)}
                    </tr>
                </thead>
                <tbody>{tbody}</tbody>
                <TableFooter params={params} setParams={setParams} current={currPage} tCount={totalCount} tPage={totalPage}/>
            </table>
        </main>
        <section className={'details'} hidden={!(details.showType === "add" || details.showType === "update")}>// 修改和添加区域默认隐藏
            <h2>Step1 Set info</h2>
            <section>
                <ul> {fields.filter(item => !item.editHidden).map(item =>
                    <DetailsLi2
                        key={item.field} field={item.field} type={item.type} details={details} onDetailsChange={onDetailsChange}
                        label={item.title} placeholder={item.detailsDesc}
                        items={item.items} step={item.step} height={item.detailsHeight}
                    />)}
                </ul>
            </section>
            <h2>Step2 Go!</h2>
            <p>Be patient!Operation may be some time...</p>
            <input type={'button'} value={'Go Back'} onClick={() => setDetails({showType: "main"})}/>
            <input type={'button'} value={'Submit'} onClick={onSubmit}/>
        </section>
    </>
}
export default DefaultPage
```

### 后台

#### 路由注册
添加新的数据表时才此文件进行注册
```go
package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"interface/internal/controller"
	"interface/utility/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.BindMiddlewareDefault(ghttp.MiddlewareHandlerResponse, middleware.CORS)
			s.Group("/menu", func(g *ghttp.RouterGroup) {
				g.Middleware(middleware.Auth)
				g.GET("/list", controller.Menu.List)
				g.GET("/getById", controller.Menu.GetById)
				g.Middleware(middleware.LockAction)
				g.POST("/add", controller.Menu.Add)
				g.PUT("/update", controller.Menu.Put)
				g.DELETE("/del", controller.Menu.Del)
			})
			s.Group("/api", func(g *ghttp.RouterGroup) {
				g.Middleware(middleware.Auth)
				g.GET("/list", controller.Api.List)
				g.GET("/getById", controller.Api.GetById)
				g.Middleware(middleware.LockAction)
				g.POST("/add", controller.Api.Add)
				g.PUT("/update", controller.Api.Put)
				g.DELETE("/del", controller.Api.Del)
			})
			// ...
			s.Run()
			return nil
		},
	}
)

```

#### controller
```go
package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/model/entity"
	"interface/internal/service"
	"interface/utility/res"
	"interface/utility/utils/xparam"
)

var Menu = menu{table: "s_menu"}

type menu struct {
	table string
}

func (c *menu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	total, data := service.List(gctx.New(), service.SearchConf{
		Table: c.table,
		Page:  page, Size: size,
		Conditions: []*service.Condition{
			{Field: "id", Value: r.GetQuery("id")},                              //如果有新的查询需求在此处添加即可
			{Field: "pid", Value: r.GetQuery("pid")},
			{Field: "type", Value: r.GetQuery("type")},
			{Field: "status", Value: r.GetQuery("status")},
			{Field: "name", Value: r.GetQuery("name"), Like: true},
		},
		OrderBy: "sort desc,id desc",                                            // 排序规则 默认 `desc id`
	})
	res.OkPage(data, total, page, size, r)
}
func (c *menu) Add(r *ghttp.Request) {
	var d entity.Menu
	_ = r.Parse(&d)
	if err := service.Add(gctx.New(), c.table, &d); err != nil {                 // 调用通用 service添加方法 (若有其他业务则创建对应的service文件即可)
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *menu) Del(r *ghttp.Request) {
	if err := service.Del(gctx.New(), c.table, xparam.ID(r)); err != nil {       // 同上
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *menu) Put(r *ghttp.Request) {
	var d entity.Menu
	_ = r.Parse(&d)
	if err := service.Update(gctx.New(), c.table, d.Id, &d); err != nil {        // 同上
		res.Error(err, r)
	}
	res.OK(r)
}

func (c *menu) GetById(r *ghttp.Request) {                                       // 同上
	data, _ := service.GetById(gctx.New(), c.table, xparam.ID(r))
	res.OKData(data, r)
}
```
#### 公共 service方法 
如果没有特殊需求，只需要写 `controller` 和`cmd` 路由注册文件即可，因为所有 `controller` 层都可以调用该通用`service` 方法
```go
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
	Field string                                                                            // 需要搜索的字段
	Value interface{}                                                                       // 具体值
	Like  bool                                                                              // false:where true: whereLike 
}
type SearchConf struct { 
	Table      string                                                                       // 表名
	Page, Size int                                                                          // 页数和大小
	Conditions []*Condition                                                                 // 查询条件
	OrderBy    string                                                                       // 排序方式
	FieldsEx   string                                                                       // 过滤的字段
}

func QueryCondition(filed string, r *ghttp.Request, like ...bool) *Condition {              // 新添查询条件方法
	condition := Condition{Field: filed, Value: r.GetQuery(filed)}
	if len(like) > 0 {
		condition.Like = true
	}
	return &condition
}

func List(ctx context.Context, s SearchConf) (int, gdb.List) {                              // 通用List查询
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
func GetById(ctx context.Context, table, id interface{}) (gdb.Record, error) {              // 通用根据ID查参数方法
	one, err := g.DB().Ctx(ctx).Model(table).One("id", id)
	if err != nil {
		glog.Error(ctx, err)
		return nil, err
	}
	return one, nil
}
func Add(ctx context.Context, table, data interface{}) error {                              // 通用添加方法
	_, err := g.DB().Ctx(ctx).Model(table).Insert(data)
	if err != nil {
		glog.Error(ctx, err)
		return err
	}
	return nil
}
func Del(ctx context.Context, table, id interface{}) (err error) {                         // 通用删除方法
	if _, err = g.DB().Ctx(ctx).Model(table).Delete("id", id); err != nil {
		glog.Error(ctx, err)
		return
	}
	return
}
func Update(ctx context.Context, table string, id, data interface{}) error {              // 通用修改方法
	_, err := g.DB().Update(ctx, table, data, "id", id)
	if err != nil {
		glog.Error(ctx, err)
		return err
	}
	return nil
}

```