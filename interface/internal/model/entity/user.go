// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint64      `json:"id"        description:""`
	Uname     string      `json:"uname"     description:""`
	Pwd       string      `json:"pwd"       description:""`
	Nickname  string      `json:"nickname"  description:""`
	Icon      string      `json:"icon"      description:""`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}