package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"interface/internal/consts"
	"interface/utility/utils/xjwt"
	"net/http"
)

var (
	Uid = "userInfoKey"
)

func Auth(c *ghttp.Request) {
	userInfo, err := xjwt.UserInfo(c)
	if err != nil {
		c.Response.WriteStatus(http.StatusForbidden, consts.ErrAuth.Error())
		c.Exit()
	}
	c.SetParam(Uid, userInfo.Uid)
	c.Middleware.Next()
}
