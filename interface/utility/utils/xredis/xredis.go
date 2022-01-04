package xredis

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"time"
)

var (
	defaultTime = 10
)

type Lock struct {
	Key      string
	lockTime int
}

var (
	errFast    = errors.New("主人,您操作太快了")
	errUnknown = errors.New("抱歉，出现未知错误了")
	errUnlock  = errors.New("抱歉，解锁失败了")
)

func New(key string) *Lock {
	return &Lock{key, defaultTime}
}

func (l Lock) Lock() error {
	ok, err := g.Redis().Do(context.TODO(), "set", l.Key, time.Now(), "nx", "ex", l.lockTime)
	if err != nil {
		return errFast
	}
	if ok.String() != "OK" {
		return errFast
	}
	return nil
}
func (l Lock) Unlock() {
	ok, err := g.Redis().Do(context.TODO(), "del", l.Key)
	if err != nil {
		glog.Error(nil, errUnknown)
		return
	}
	if !(ok.Int() == 1 || ok.Int() == 0) {
		glog.Error(nil, errUnlock)
		return
	}
}
func UserLock(key uint64) (*Lock, error) {
	lock := New(fmt.Sprint("user_operation_lock_", key))
	if err := lock.Lock(); err != nil {
		glog.Warningf(nil, err.Error())
		return nil, err
	}
	return lock, nil
}
