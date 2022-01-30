package apiv1

import "github.com/gogf/gf/v2/util/gmeta"

type ListByIdRes struct {
	gmeta.Meta `orm:"table:b_book_category"`
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Books      []*BookInfo `orm:"with:category_id=id,where:status=3,order:sort desc" json:"books"`
}
type BookInfo struct {
	gmeta.Meta     `orm:"table:b_book"`
	CategoryId     int    `json:"-"`
	Icon           string `json:"icon"`
	Name           string `json:"name"`
	Summary        string `json:"summary"`
	FirstChapterId int    `json:"first_chapter_id"`
	Id             int    `json:"id"`
}
type ListChapterReq struct {
	BookId string `v:"required"`
}
