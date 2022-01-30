// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-01-27 15:11:39
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Book is the golang structure of table b_book for DAO operations like Where/Data.
type Book struct {
	g.Meta         `orm:"table:b_book, do:true"`
	Id             interface{} //
	Tid            interface{} //
	CategoryId     interface{} //
	FirstChapterId interface{} //
	TypeDetails    interface{} // - 小说             - 诗歌             - 散文             - 童话             - 传记
	WriterId       interface{} //
	Name           interface{} //
	Icon           interface{} //
	HiddenWords    interface{} //
	Summary        interface{} //
	Sort           interface{} //
	Scope          interface{} //
	Status         interface{} // 1 ok 2 off 3 home show
	Theme          interface{} // #96b97d,dark,#f6f4f0,dark  top:light top:dark nav:light nav:dark
	ClickNum       interface{} //
	PublishTime    interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	Ex1            interface{} //
}