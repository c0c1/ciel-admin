package middleware

import (
	"errors"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"interface/utility/res"
	"interface/utility/utils/xredis"
)

func LockAction(r *ghttp.Request) {
	uid := r.Get(Uid).Uint64()
	if uid == 0 {
		err := errors.New("uid is empty")
		glog.Error(nil, err)
		res.Error(err, r)
	}
	lock, err := xredis.UserLock(uid)
	if err != nil {
		res.Error(err, r)
	}
	r.Middleware.Next()
	lock.Unlock()
}
