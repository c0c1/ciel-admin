package middleware

import (
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gsession"
)

const (
	userSession = "userSession"
)

type (
	session  struct{}
	UserInfo struct {
		Id    int
		Uname string
		Rid   int64 // 角色id
		Icon  string
		Menu  []*Menu
	}
	Menu struct {
		Id       uint64  `orm:"id,primary" json:"id"`  //
		Pid      uint64  `orm:"pid"        json:"pid"` //
		Name     string  `orm:"name"  json:"name"`     //
		Path     string  `json:"path"`
		Icon     string  `orm:"icon"       json:"icon"`   //
		Group    int     `orm:"group"      json:"group"`  //
		Sort     float64 `orm:"sort"       json:"sort"`   //
		Status   int     `orm:"status"     json:"status"` //
		Children []*Menu
	}
)

func nSession() *session {
	return &session{}
}

var (
	Session = session{}
)

func (s *session) Set(session *gsession.Session, id int, uname string, rid int64, icon string, menus []*Menu) {
	info := UserInfo{id, uname, rid, icon, menus}
	session.Set(userSession, &info)
}
func (s *session) Get(session *gsession.Session) *UserInfo {
	var u UserInfo
	get, err := session.Get(userSession, &u)
	if err != nil {
		return nil
	}
	if err = get.Struct(&u); err != nil {
		return nil
	}
	return &u
}
func (s *session) IsAuth(session *gsession.Session) bool {
	get, err := session.Get(userSession)
	if err != nil {
		return false
	}
	return get.IsNil()
}
func (s *session) Logout(session *gsession.Session) {
	if err := session.Close(); err != nil {
		glog.Error(nil, err.Error())
	}
}
