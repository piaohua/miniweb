package models

import (
	"time"
	"strconv"
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
	Ctime  time.Time `bson:"ctime" json:"ctime"`   //创建时间
}

//GateInfo gate info
type GateInfo struct {
	Gateid string `bson:"gateid" json:"gateid"` //unique
	Type   int32  `bson:"type" json:"type"`     //type
	Star   int32  `bson:"star" json:"star"`     //星数量
	Num    int32  `bson:"num" json:"num"`       //完成次数num
}

//GeteKey unique key
func GateKey(Type, Gateid int32) string {
    return strconv.Itoa(int(Type)) + strconv.Itoa(int(Gateid))
}
