package xfile

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestDel(t *testing.T) {
	f, err := os.OpenFile("/home/howl/learn/ciel-admin/interface/internal/cmd/admin.go", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	delim := "\n"
	var offset int64
	r := bufio.NewReader(f)
	for {
		offset--
		_, err := f.Seek(offset, os.SEEK_END)
		if err != nil {
			log.Println(err)
			break
		}
		r.Reset(f)
		b, err := r.Peek(1)
		if err != nil {
			log.Println(err)
			break
		}
		if string(b) == delim {
			break
		}
	}
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	err = f.Truncate(info.Size() + offset)
	if err != nil {
		log.Fatal(err)
	}
}

func TestGen(t *testing.T) {

	Gen(TemplateConfig{
		RootPath:      "/home/howl/learn/ciel-admin",
		FileName:      "u_user_details",
		EntityName:    "UserDetails",
		PathService:   "/interface/internal/service",
		PathHandler:   "/interface/internal/handler",
		PathCmd:       "/interface/internal/cmd/admin.go",
		PathApi:       "/interface/apiv1/user_api.go",
		PathDataFetch: "/admin/data/user.js", // 如果文件不存在请先创建
		PathIndex:     "/admin/pages/user",
		group:         "user",
		NameZh:        "用户详情",
		P1:            "uid",
		P2:            "real_name",
		P3:            "desc",
		P1Name:        "用户ID",
		P2Name:        "姓名",
		P3Name:        "备注",
	})
}
