package models

import (
	"strconv"
	"time"

	"miniweb/pb"

	"github.com/globalsign/mgo/bson"
)

//Prop prop info
type Prop struct {
	ID    string    `bson:"_id" json:"id"`      //unique ID
	Name  string    `bson:"name" json:"name"`   //name
	Type  int32     `bson:"type" json:"type"`   //type unique
	Attr  int32     `bson:"attr" json:"attr"`   //属性
	Ctime time.Time `bson:"ctime" json:"ctime"` //创建时间
}

//PropInfo prop info
type PropInfo struct {
	Type int32 `bson:"type" json:"type"` //type
	Attr int32 `bson:"attr" json:"attr"` //属性
	Num  int32 `bson:"num" json:"num"`   //num
}

//GetPropList get prop list
func GetPropList() []Prop {
	var list []Prop
	ListByQ(Props, nil, &list)
	return list
}

//Save 写入数据库
func (t *Prop) Save() bool {
	t.Ctime = bson.Now()
	return Insert(Props, t)
}

//AddPropMsg add prop msg
func AddPropMsg(user *User, key string, n int32,
	ptype pb.PropType) (msg *pb.SPushProp) {
	if val, ok := user.Prop[key]; ok {
		val.Num += n
		if val.Num <= 0 {
			delete(user.Prop, key)
		}
	}
	msg = &pb.SPushProp{
		//Type: pb.LOG_TYPE0,
		Ptype: ptype,
		Num:   int64(n),
	}
	return
}

//PropKey unique key
func PropKey(Type int32) string {
	return strconv.Itoa(int(Type))
}

//PropUniqueKey unique key
func PropUniqueKey(Type int32) string {
	return "prop" + strconv.Itoa(int(Type))
}

//PropInit prop init
func PropInit(user *User) {
	if user.Prop != nil {
		return
	}
	user.Prop = make(map[string]PropInfo)
}

//GetProp get prop by type
func GetProp(Type int32) (prop *Prop) {
	if v := Cache.Get(PropUniqueKey(Type)); v != nil {
		if val, ok := v.(*Prop); ok {
			prop = val
		}
	}
	return
}

//GetPropName get prop name
func GetPropName(Type int32) (name string) {
	if v := GetProp(Type); v != nil {
		name = v.Name
	}
	return
}
