package models

import (
	"strconv"
	"time"

	"miniweb/libs"
	"miniweb/pb"

	"github.com/astaxie/beego"
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

//GateCount 关卡统计
type GateCount struct {
	Type   int32 `bson:"type" json:"type"`     //关卡类型
	Number int32 `bson:"number" json:"number"` //关卡数量
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

//Upsert 更新数据库
func (t *Gate) Upsert() bool {
	return Upsert(Gates, bson.M{"type": t.Type, "gateid": t.Gateid}, t)
}

//Delete 删除数据
func (t *Gate) Delete() bool {
	return Delete(Gates, bson.M{"type": t.Type, "gateid": t.Gateid})
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
	if user.Gate == nil {
		user.Gate = make(map[string]GateInfo)
	}
	//单人
	key := GateKey(int32(pb.GATE_TYPE1), 1)
	if _, ok := user.Gate[key]; !ok {
		AddGate(user, int32(pb.GATE_TYPE1), 1, 0)
	}
	//副本
	key = GateKey(int32(pb.GATE_TYPE2), 1)
	if _, ok := user.Gate[key]; !ok {
		AddGate(user, int32(pb.GATE_TYPE2), 1, 0)
	}
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

//LoadGateList load gate info by gate.json
func LoadGateList() []Gate {
	filePath := "static/gate.json"
	list := make([]Gate, 0)
	err := libs.Load(filePath, &list)
	if err != nil {
		beego.Error("load gate err ", err)
	}
	return list
}

//InitGateList init gate to cache
func InitGateList() {
	list := GetGateList()
	//test use
	if !RunMode() {
		if len(list) == 0 {
			//SetGateList()
			list = LoadGateList()
		}
	}
	Cache.Put("gate", list, 0)
	for k, v := range list {
		Cache.Put(GateUniqueKey(v.Type, v.Gateid), &list[k], 0)
	}
	//count init
	InitGateCount(list)
}

//GetGateCount get gate numbers by type
func GetGateCount() (list []GateCount) {
	if v := Cache.Get("gatecount"); v != nil {
		if val, ok := v.([]GateCount); ok {
			list = val
		}
	}
	return
}

//InitGateCount gate numbers
func InitGateCount(list []Gate) {
	m := make(map[int32]int32)
	for _, v := range list {
		m[v.Type]++
	}
	l := make([]GateCount, 0)
	for k, v := range m {
		g := GateCount{
			Type:   k,
			Number: v,
		}
		l = append(l, g)
	}
	Cache.Put("gatecount", l, 0)
}

//UpsertGate upsert gate
func UpsertGate(gate Gate) bool {
	if len(gate.ID) == 0 {
		beego.Error("gate id err: ", gate)
		return false
	}
	key := GateUniqueKey(gate.Type, gate.Gateid)
	list := GetGates()
	if gate.Del != 0 {
		if !gate.Delete() {
			beego.Error("gate delete err: ", gate)
			return false
		}
		Cache.Delete(key)
		for k, v := range list {
			if v.Type == gate.Type && v.Gateid == gate.Gateid {
				list = append(list[:k], list[k+1:]...)
				break
			}
		}
		Cache.Put("gate", list, 0)
		return true
	}
	if !gate.Upsert() {
		beego.Error("gate upsert err: ", gate)
		return false
	}
	Cache.Put(key, &gate, 0)
	for k, v := range list {
		if v.Type == gate.Type && v.Gateid == gate.Gateid {
			list[k] = gate
			break
		}
	}
	Cache.Put("gate", list, 0)
	//count init
	InitGateCount(list)
	return true
}

//GetGates from cache
func GetGates() (l []Gate) {
	if v := Cache.Get("gate"); v != nil {
		if val, ok := v.([]Gate); ok {
			l = val
		}
	}
	return
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
