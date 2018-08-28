package models

import (
	"strconv"
	"time"
)

//Prop prop info
type Prop struct {
	ID     string    `bson:"_id" json:"id"`        //unique ID
	Propid int32     `bson:"propid" json:"propid"` //unique
	Name   string    `bson:"name" json:"name"`     //name
	Type   int32     `bson:"type" json:"type"`     //type unique
	Star   int32     `bson:"star" json:"star"`     //星数量
	Ctime  time.Time `bson:"ctime" json:"ctime"`   //创建时间
}

//PropInfo prop info
type PropInfo struct {
	Propid int32 `bson:"propid" json:"propid"` //unique
	Type   int32 `bson:"type" json:"type"`     //type
	Star   int32 `bson:"star" json:"star"`     //星数量
	Num    int32 `bson:"num" json:"num"`       //num
}

//PropKey unique key
func PropKey(Type, Propid int32) string {
	return strconv.Itoa(int(Type)) + strconv.Itoa(int(Propid))
}
