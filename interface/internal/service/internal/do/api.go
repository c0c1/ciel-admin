// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-01-27 15:11:39
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Api is the golang structure of table s_api for DAO operations like Where/Data.
type Api struct {
	g.Meta    `orm:"table:s_api, do:true"`
	Id        interface{} //
	Url       interface{} //
	Method    interface{} //
	Group     interface{} //
	Desc      interface{} //
	Status    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}