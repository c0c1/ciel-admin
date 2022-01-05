package apiv1

import "interface/internal/model/entity"

type UserReq struct {
	Page     int `d:"1"`
	PageSize int `d:"10"`
	entity.User
}
