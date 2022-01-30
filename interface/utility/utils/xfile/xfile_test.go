package xfile

import (
	"testing"
)

func TestGen(t *testing.T) {

	Gen(TemplateConfig{
		RootPath:      "/home/howl/learn/freekey",
		PathCmd:       "/interface/internal/cmd/admin.go",
		PathApi:       "/interface/apiv1/book_api.go",
		PathDataFetch: "/admin/data/book.js", // 如果文件不存在请先创建
		PathIndex:     "/admin/pages/book",
		group:         "book",

		NameZh:     "作者",
		FileName:   "b_writer",
		EntityName: "Writer",
		P1:         "name",
		P2:         "summary",
		P3:         "words",
		P1Name:     "名称",
		P2Name:     "概括",
		P3Name:     "语录",
	})
}
