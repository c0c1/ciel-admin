package res

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	codeErr = -1
)

type DataRes struct {
	NormalData interface{} `json:"normal_data"`
}

func Ok() (*DataRes, error) {
	return nil, nil
}
func OkData(data interface{}) (*DataRes, error) {
	return &DataRes{data}, nil
}

func ErrMsg(msg string, c *ghttp.Request) {
	Json(c, codeErr, msg)
	c.Exit()
}

func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(g.Map{
		"code": code,
		"msg":  message,
		"data": responseData,
	})
}
