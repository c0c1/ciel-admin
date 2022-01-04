package bo

type Menu struct {
	Id       uint64  `orm:"id,primary" json:"id"`     //
	Pid      uint64  `orm:"pid"        json:"pid"`    //
	Divide   int     `orm:"divide"     json:"divide"` //
	Name     string  `orm:"name"  json:"name"`        //
	Path     string  `json:"path"`
	Icon     string  `orm:"icon"       json:"icon"`   //
	Type     int     `orm:"type"      json:"type"`    //
	Sort     float64 `orm:"sort"       json:"sort"`   //
	Status   int     `orm:"status"     json:"status"` //
	Children []*Menu
}
