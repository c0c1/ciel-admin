package apiv1

import "interface/internal/model/entity"

type UserReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.User
}

type LoginLogReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.LoginLog
}

type UserDetailsReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.UserDetails
}
