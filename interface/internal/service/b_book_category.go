package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"interface/apiv1"
)

var BookCategory = NewBookCategory()

type bookCategory struct{}

func NewBookCategory() *bookCategory {
	a := bookCategory{}
	return &a
}

func (s *bookCategory) ListHomeData() ([]*apiv1.ListByIdRes, error) {
	var d = make([]*apiv1.ListByIdRes, 0)
	err := g.DB().Model(d).With(apiv1.BookInfo{}).Scan(&d)
	return d, err
}
