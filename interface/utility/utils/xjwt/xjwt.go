package xjwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"interface/internal/consts"
	"net/http"
	"time"
)

var (
	TokenExpireDuration = time.Second * 10
	MySecret            = []byte("夏天夏天悄悄过去")
	tokenKey            = "token"
	authorizationKey    = "Authorization"
)

func init() {
	get, err := g.Cfg().Get(nil, "server.jwtExpireDuration")
	if err != nil {
		panic(err)
	}
	duration := get.Duration()
	if duration == 0 {
		panic("jwt.ExpireDuration is none")
	}
	TokenExpireDuration = duration
}

type MyClaims struct {
	Uname string `json:"username"`
	Uid   uint64 `json:"uid"`
	Rid   uint64
	jwt.StandardClaims
}

func UserInfo(r *ghttp.Request) (*MyClaims, error) {
	return parseToken(getToken(r))
}

func getToken(c *ghttp.Request) string {
	token := c.GetHeader(tokenKey)
	if token == "" {
		token = c.GetHeader(authorizationKey)
	}
	if token == "" {
		token = c.Get(tokenKey).String()
		if token == "" {
			token = c.Get(authorizationKey).String()
		}
	}
	if token == "" {
		c.Response.WriteStatus(http.StatusForbidden, consts.ErrAuth.Error())
		c.Exit()
	}
	return token
}

// parseToken 解析JWT
func parseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		//glog.Error(err.Error())
		return nil, consts.ErrAuth
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	glog.Error(nil, err.Error())
	return nil, consts.ErrAuth
}

func GenToken(username string, uid uint64, rid uint64) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		uid,
		rid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}
