package xredis

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestNewLock(t *testing.T) {
	lock := New("test")
	err := lock.Lock()
	if err != nil {
		t.Fatal(err.Error())
	}
	lock.Unlock()
}
func TestRedis(t *testing.T) {
	do, err := g.Redis().Do(context.TODO(), "set", "a", "b")
	if err != nil {
		panic(err)
	}
	g.Log().Info(nil, do)
}
