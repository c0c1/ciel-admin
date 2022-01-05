package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"interface/internal/consts"
	"interface/utility/utils/res"
	"interface/utility/utils/xjwt"
	"net/http"
	"strings"
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
	if userInfo.Rid != 0 {
		if !checkPath(int64(userInfo.Rid), c.RequestURI, c.Method) {
			res.ErrMsg(consts.ErrAuthNotEnough.Error(), c)
		}
	}
	c.SetParam(Uid, userInfo.Uid)
	c.Middleware.Next()
}
func checkPath(rid int64, uri string, method string) bool {
	if strings.Contains(uri, "?") {
		uri = strings.Split(uri, "?")[0]
	}
	if uri == "/" {
		return true
	}
	count, _ := g.DB().Model("s_role t1").
		LeftJoin("s_role_api t2 on t1.id = t2.rid").
		LeftJoin("s_api t3 on t2.aid = t3.id").
		Where("t3.url = ? and t3.method = ? and t1.id = ?  ", uri, method, rid).
		Count()
	if count == 1 {
		return false
	}
	return true
}
