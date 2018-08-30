package models

import (
	"strconv"
	"time"

	"miniweb/pb"

	"github.com/globalsign/mgo/bson"
)

//LoginPrize 连续登录奖励配置
type LoginPrize struct {
	ID    string           `bson:"_id" json:"id"`      //unique ID
	Day   uint32           `bson:"day" json:"day"`     //unique
	Prize []LoginPrizeProp `bson:"prize" json:"prize"` //prize
	Del   int              `bson:"del" json:"del"`     //是否移除
	Ctime time.Time        `bson:"ctime" json:"ctime"` //创建时间
}

type LoginPrizeProp struct {
	Type   int32 `bson:"type" json:"type"`     //物品类型
	Number int32 `bson:"number" json:"number"` //物品数量
}

//GetLoginPrizeList 获取连续登录奖励配置
func GetLoginPrizeList() []LoginPrize {
	var list []LoginPrize
	ListByQ(LoginPrizes, bson.M{"del": 0}, &list)
	return list
}

//Save 写入数据库
func (t *LoginPrize) Save() bool {
	t.Ctime = bson.Now()
	return Insert(LoginPrizes, t)
}

//SetLoginPrizeList 添加登录奖励
func SetLoginPrizeList() {
	prize2 := NewLoginPrizeProp(int32(pb.PROP_TYPE11), 1)
	prize3 := NewLoginPrizeProp(int32(pb.PROP_TYPE4), 1)
	prize4 := NewLoginPrizeProp(int32(pb.PROP_TYPE6), 1)
	var i uint32
	for i = 0; i < 7; i++ {
		prize := []LoginPrizeProp{}
		var num int32 = 10000 + int32(i)*5000
		prize1 := NewLoginPrizeProp(int32(pb.PROP_TYPE2), num)
		NewLoginPrize(i, append(prize, prize1, prize2, prize3, prize4))
	}
}

//NewLoginPrizeProp 添加奖励
func NewLoginPrizeProp(Type, num int32) LoginPrizeProp {
	return LoginPrizeProp{
		Type:   Type,
		Number: num,
	}
}

//NewLoginPrize 添加天数
func NewLoginPrize(day uint32, prize []LoginPrizeProp) {
	t := LoginPrize{
		ID:    bson.NewObjectId().String(),
		Day:   day,
		Prize: prize,
	}
	t.Save()
}

//InitLoginPrizeList init login prize to cache
func InitLoginPrizeList() {
	list := GetLoginPrizeList()
	Cache.Put("prize", list, 0)
	for k, v := range list {
		Cache.Put(PrizeKey(v.Day), &list[k], 0)
	}
}

//PrizeKey unique key
func PrizeKey(day uint32) string {
	return "prize" + strconv.Itoa(int(day))
}

//GetLoginPrizes from cache
func GetLoginPrizes() (l []LoginPrize) {
	if v := Cache.Get("prize"); v != nil {
		if val, ok := v.([]LoginPrize); ok {
			l = val
		}
	}
	return
}

//GetLoginPrize get login prize from cache by id
func GetLoginPrize(day uint32) (prize *LoginPrize) {
	if v := Cache.Get(PrizeKey(day)); v != nil {
		if val, ok := v.(*LoginPrize); ok {
			prize = val
		}
	}
	return
}
