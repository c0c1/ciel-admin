package res

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"math"
)

type PageRes struct {
	Data       interface{} `json:"data,omitempty"`
	Other      interface{} `json:"other,omitempty"`
	TotalCount int64       `json:"totalCount"`
	PageSize   int64       `json:"pageSize,omitempty"`
	TotalPage  int64       `json:"totalPage,omitempty"`
	CurrPage   int64       `json:"currPage,omitempty"`
	List       interface{} `json:"list"`
}

func NewPageRes(data interface{}, total int, page, pageSize int) *PageRes {
	if pageSize == 0 {
		pageSize = 10
	}
	totalPage := math.Ceil(float64(total) / float64(pageSize)) //这里计算总页数时，要向上取整
	if totalPage <= 0 {
		totalPage = 1
	}
	if total == 0 {
		data = make([]interface{}, 0)
	}
	return &PageRes{
		TotalCount: int64(total),
		PageSize:   int64(pageSize),
		CurrPage:   int64(page),
		List:       data,
		TotalPage:  int64(totalPage),
	}
}
func OkPage(data interface{}, total int, page, pageSize int, r *ghttp.Request) {
	if pageSize == 0 {
		pageSize = 10
	}
	totalPage := math.Ceil(float64(total) / float64(pageSize)) //这里计算总页数时，要向上取整
	if totalPage <= 0 {
		totalPage = 1
	}
	if total == 0 {
		data = make([]interface{}, 0)
	}
	OKData(PageRes{TotalCount: int64(total), PageSize: int64(pageSize), CurrPage: int64(page), List: data, TotalPage: int64(totalPage)}, r)
}

func (v *PageRes) AddOther(data interface{}) interface{} {
	v.Other = data
	return v
}
