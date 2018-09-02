package models

import (
	"strconv"
	"time"

	"miniweb/libs"
	"miniweb/pb"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
)

//Prop prop info
type Prop struct {
	ID    string    `bson:"_id" json:"id"`      //unique ID
	Name  string    `bson:"name" json:"name"`   //name
	Type  int32     `bson:"type" json:"type"`   //type unique
	Attr  int32     `bson:"attr" json:"attr"`   //属性
	Del   int       `bson:"del" json:"del"`     //是否移除
	Ctime time.Time `bson:"ctime" json:"ctime"` //创建时间
}

//PropInfo prop info
type PropInfo struct {
	Type int32 `bson:"type" json:"type"` //type
	Attr int32 `bson:"attr" json:"attr"` //属性
	Num  int32 `bson:"num" json:"num"`   //num
}

//TempPropInfo temp prop info
type TempPropInfo struct {
	GType  int32 `bson:"gtype" json:"gtype"`   //gate type
	Gateid int32 `bson:"gateid" json:"gateid"` //gate id
	Type   int32 `bson:"type" json:"type"`     //prop type
	Attr   int32 `bson:"attr" json:"attr"`     //属性
	Num    int32 `bson:"num" json:"num"`       //num
}

//GetPropList get prop list
func GetPropList() []Prop {
	var list []Prop
	ListByQ(Props, bson.M{"del": 0}, &list)
	return list
}

//Save 写入数据库
func (t *Prop) Save() bool {
	t.Ctime = bson.Now()
	return Insert(Props, t)
}

//Upsert 更新数据库
func (t *Prop) Upsert() bool {
	return Upsert(Props, bson.M{"type": t.Type}, t)
}

//Delete 删除数据
func (t *Prop) Delete() bool {
	return Delete(Props, bson.M{"type": t.Type})
}

//AddPropMsg add prop msg, TODO 日志
func AddPropMsg(user *User, key string, num int64,
	ptype pb.PropType) (msg *pb.SPushProp) {
	msg = addPropMsg(user, key, num, ptype, false)
	return
}

//AddTempPropMsg add temp prop msg, TODO 日志
func AddTempPropMsg(user *User, key string, num int64,
	ptype pb.PropType) (msg *pb.SPushProp) {
	msg = addPropMsg(user, key, num, ptype, true)
	return
}

//addPropMsg add prop msg, TODO 日志
func addPropMsg(user *User, key string, num int64,
	ptype pb.PropType, temp bool) (msg *pb.SPushProp) {
	switch ptype {
	case pb.PROP_TYPE0:
		return
	case pb.PROP_TYPE1:
		msg = AddDiamondMsg(user, num)
		return
	case pb.PROP_TYPE2:
		msg = AddCoinMsg(user, num)
		return
	case pb.PROP_TYPE3:
		msg = AddEnergyMsg(user, num)
		return
	}
	msg = &pb.SPushProp{
		//Type: pb.LOG_TYPE0,
		Ptype: ptype,
		Num:   num,
	}
	if num < 0 { //use temp, TODO gate
		if val, ok := user.TempProp[key]; ok {
			val.Num += int32(num)
			msg.Total = int64(val.Num)
			if val.Num <= 0 {
				msg.Total = 0
				delete(user.TempProp, key)
			} else {
				user.TempProp[key] = val
			}
			return
		}
	}
	if temp {
		if val, ok := user.TempProp[key]; ok {
			val.Num += int32(num)
			msg.Total = int64(val.Num)
			if val.Num <= 0 {
				msg.Total = 0
				delete(user.TempProp, key)
			} else {
				user.TempProp[key] = val
			}
		} else if num > 0 {
			msg.Total = int64(num)
			user.TempProp[key] = TempPropInfo{
				Type: int32(ptype),
				Num:  int32(num),
			}
		}
		return
	}
	if val, ok := user.Prop[key]; ok {
		val.Num += int32(num)
		msg.Total = int64(val.Num)
		if val.Num <= 0 {
			msg.Total = 0
			delete(user.Prop, key)
		} else {
			user.Prop[key] = val
		}
	} else if num > 0 {
		msg.Total = int64(num)
		user.Prop[key] = PropInfo{
			Type: int32(ptype),
			Num:  int32(num),
		}
	}
	return
}

//PropKey user prop unique key
func PropKey(Type int32) string {
	return strconv.Itoa(int(Type))
}

//PropUniqueKey cache prop unique key
func PropUniqueKey(Type int32) string {
	return "prop" + strconv.Itoa(int(Type))
}

//PropInit prop init
func PropInit(user *User) {
	if user.Prop == nil {
		user.Prop = make(map[string]PropInfo)
	}
	if user.TempProp == nil {
		user.TempProp = make(map[string]TempPropInfo)
	}
}

//LoadPropList load prop info by prop.json
func LoadPropList() []Prop {
	filePath := "static/prop.json"
	list := make([]Prop, 0)
	err := libs.Load(filePath, &list)
	if err != nil {
		beego.Error("load prop err ", err)
	}
	return list
}

//InitPropList init prop to cache
func InitPropList() {
	list := GetPropList()
	//test use
	if !RunMode() {
		if len(list) == 0 {
			//SetPropList()
			list = LoadPropList()
		}
	}
	Cache.Put("prop", list, 0)
	for k, v := range list {
		Cache.Put(PropUniqueKey(v.Type), &list[k], 0)
	}
}

//UpsertProp upsert prop
func UpsertProp(prop Prop) bool {
	key := PropUniqueKey(prop.Type)
	list := GetProps()
	if prop.Del != 0 {
		if !prop.Delete() {
			beego.Error("prop delete err: ", prop)
			return false
		}
		Cache.Delete(key)
		for k, v := range list {
			if v.Type == prop.Type {
				list = append(list[:k], list[k+1:]...)
				break
			}
		}
		Cache.Put("prop", list, 0)
		return true
	}
	if !prop.Upsert() {
		beego.Error("prop upsert err: ", prop)
		return false
	}
	Cache.Put(key, &prop, 0)
	for k, v := range list {
		if v.Type == prop.Type {
			list[k] = prop
			break
		}
	}
	Cache.Put("prop", list, 0)
	return true
}

//GetProps from cache
func GetProps() (l []Prop) {
	if v := Cache.Get("prop"); v != nil {
		if val, ok := v.([]Prop); ok {
			l = val
		}
	}
	return
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
