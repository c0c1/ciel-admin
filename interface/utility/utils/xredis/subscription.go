package xredis

import (
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type (
	Subscription struct {
		channel string
	}
	NoticeMsgBo struct {
		Type int    `json:"type"`
		Msg  string `json:"msg"`
	}
)

func (s *Subscription) Publish(m NoticeMsgBo) {
	g.Redis().Do(nil, "publish", s.channel, m)
}
func (s *Subscription) Listener(f func(msg NoticeMsgBo)) {
	conn, err := g.Redis().Conn(nil)
	defer conn.Close(nil)
	_, err = conn.Do(nil, "SUBSCRIBE", s.channel)
	if err != nil {
		panic(err)
	}
	for {
		replay, err := conn.Receive(nil)
		if err != nil {
			glog.Error(nil, err.Error())
		}
		bo := NoticeMsgBo{}
		json.Unmarshal([]byte(replay.Strings()[2]), &bo)
		f(bo)
	}
}

func NewSubscription(channel string) *Subscription {
	return &Subscription{channel: channel}
}
