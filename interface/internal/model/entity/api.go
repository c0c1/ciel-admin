// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-01-30 15:53:54
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Api is the golang structure for table api.
type Api struct {
	Id        uint64      `json:"id"        description:""`
	Url       string      `json:"url"       description:""`
	Method    string      `json:"method"    description:""`
	Group     string      `json:"group"     description:""`
	Desc      string      `json:"desc"      description:""`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
