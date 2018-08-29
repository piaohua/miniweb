package models

import (
	"strconv"
	"time"
)

//Gate gate
type Gate struct {
	ID     string    `bson:"_id" json:"id"`        //unique ID
	Gateid int32     `bson:"gateid" json:"gateid"` //unique
	Type   int32     `bson:"type" json:"type"`     //type
	Star   int32     `bson:"star" json:"star"`     //星数量
	Data   []byte    `bson:"data" json:"data"`     //数据
	Ptype  int32     `bson:"ptype" json:"ptype"`   //prop type
	Num    int32     `bson:"num" json:"num"`       //prop number
	Incr   bool      `bson:"incr" json:"incr"`     //有序递增
	Ctime  time.Time `bson:"ctime" json:"ctime"`   //创建时间
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
	ListByQ(Gates, nil, &list)
	return list
}

//Save 写入数据库
func (t *Gate) Save() bool {
	t.Ctime = bson.Now()
	return Insert(Gates, t)
}

//GeteKey unique key
func GateKey(Type, Gateid int32) string {
	return strconv.Itoa(int(Type)) + strconv.Itoa(int(Gateid))
}

//GateUniqueKey unique key
func GateUniqueKey(Type, Gateid int32) string {
	return "gate" + strconv.Itoa(int(Type)) + strconv.Itoa(int(Gateid))
}

//GateInit gate init
func GateInit(user *User) {
    if user.Gate != nil {
        return
    }
    user.Gate = make(map[string]GateInfo)
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
