package xredis

import (
	"testing"
)

type Listener func()

func TestNewSubscription(t *testing.T) {
	NewSubscription("admin").Listener(func(msg NoticeMsgBo) {
	})
	select {}
}
func TestPublish(t *testing.T) {
	NewSubscription("admin").Publish(NoticeMsgBo{Msg: "123", Type: 0})
}
