package models

import (
	"strconv"
	"time"

	"miniweb/pb"

	"github.com/globalsign/mgo/bson"
)

//Gate gate
type Gate struct {
	ID       string          `bson:"_id" json:"id"`              //unique ID
	Gateid   int32           `bson:"gateid" json:"gateid"`       //unique
	Type     int32           `bson:"type" json:"type"`           //type
	Star     int32           `bson:"star" json:"star"`           //星数量
	Data     []byte          `bson:"data" json:"data"`           //数据
	TempShop []string        `bson:"temp_shop" json:"temp_shop"` //temp shop ids
	Prize    []GatePrizeProp `bson:"prize" json:"prize"`         //prize
	Incr     bool            `bson:"incr" json:"incr"`           //有序递增
	Del      int             `bson:"del" json:"del"`             //是否移除
	Ctime    time.Time       `bson:"ctime" json:"ctime"`         //创建时间
}

//GatePrizeProp 奖励
type GatePrizeProp struct {
	Type   int32 `bson:"type" json:"type"`     //物品类型
	Number int32 `bson:"number" json:"number"` //物品数量
}

//GateInfo gate info
type GateInfo struct {
	Gateid int32 `bson:"gateid" json:"gateid"` //unique
	Type   int32 `bson:"type" json:"type"`     //type
	Star   int32 `bson:"star" json:"star"`     //星数
	Num    int32 `bson:"num" json:"num"`       //完成次数
}

//GetGateList get prop list
func GetGateList() []Gate {
	var list []Gate
	ListByQ(Gates, bson.M{"del": 0}, &list)
	return list
}

//Save 写入数据库
func (t *Gate) Save() bool {
	t.Ctime = bson.Now()
	return Insert(Gates, t)
}

//GateKey user gate unique key
func GateKey(Type, Gateid int32) string {
	return strconv.Itoa(int(Type)) + strconv.Itoa(int(Gateid))
}

//GateUniqueKey cache gate unique key
func GateUniqueKey(Type, Gateid int32) string {
	return "gate" + strconv.Itoa(int(Type)) + strconv.Itoa(int(Gateid))
}

//GateInit gate init
func GateInit(user *User) {
	if user.Gate != nil {
		return
	}
	user.Gate = make(map[string]GateInfo)
	AddGate(user, int32(pb.GATE_TYPE1), 1, 0)
}

//AddGate add new gate
func AddGate(user *User, Type, id, star int32) {
	key := GateKey(Type, id)
	if val, ok := user.Gate[key]; ok {
		val.Num++
		if val.Star < star {
			val.Star = star
		}
		user.Gate[key] = val
		return
	}
	user.Gate[key] = GateInfo{
		Gateid: id,
		Type:   Type,
		Star:   star,
	}
}

//AddNewGate add new gate
func AddNewGate(user *User, Type, id int32) {
	key := GateKey(Type, id)
	if _, ok := user.Gate[key]; !ok {
		user.Gate[key] = GateInfo{
			Gateid: id,
			Type:   Type,
		}
	}
}

//GetGate get gate by type and id
func GetGate(Type, Gateid int32) (gate *Gate) {
	if v := Cache.Get(GateUniqueKey(Type, Gateid)); v != nil {
		if val, ok := v.(*Gate); ok {
			gate = val
		}
	}
	return
}
