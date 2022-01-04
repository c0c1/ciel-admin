package apiv1

import "interface/internal/model/entity"

type AdminReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.Admin
}

type LoginReq struct {
	Uname string `v:"required"`
	Pwd   string `v:"required"`
}

type PwdReq struct {
	Pwd string `v:"required"`
}

type MenuReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.Menu
}

type RoleReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.Role
}

type ApiReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.Api
}

type RoleMenuReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.RoleMenu
}

type RoleMenuAddReq struct {
	Rid int
	Mid []int
}

type RoleApiReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.RoleApi
}

type RoleApiAddReq struct {
	Rid int
	Aid []int
}

type DictReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.Dict
}

type FileReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.File
}
