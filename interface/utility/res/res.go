package res

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	okCode  = 0
	okMsg   = "ok"
	errCode = -1
)

type defaultResData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func json(r *ghttp.Request, code int, message string, data interface{}) {
	_ = r.Response.WriteJson(defaultResData{code, message, data})
}
func GetPage(r *ghttp.Request) (page, size int) {
	page = r.GetQuery("page").Int()
	size = r.GetQuery("size").Int()
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	return
}

func OK(r *ghttp.Request) {
	json(r, okCode, okMsg, nil)
	r.Exit()
}
func OKData(data interface{}, r *ghttp.Request) {
	json(r, okCode, okMsg, data)
	r.Exit()
}
func Error(err error, c *ghttp.Request) {
	json(c, errCode, err.Error(), nil)
	c.Exit()
}

func Ok2() (*DataRes, error) {
	return nil, nil
}

type DataRes struct {
	NormalData interface{} `json:"normal_data"`
}

func OkData2(data interface{}) (*DataRes, error) {
	return &DataRes{data}, nil
}
