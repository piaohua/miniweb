package models

import (
	"strconv"
	"time"

	"miniweb/libs"
	"miniweb/pb"

	"github.com/astaxie/beego"
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

//LoginPrizeProp 登录奖励
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

//Upsert 更新数据库
func (t *LoginPrize) Upsert() bool {
	return Upsert(LoginPrizes, bson.M{"day": t.Day}, t)
}

//Delete 删除数据
func (t *LoginPrize) Delete() bool {
	return Delete(LoginPrizes, bson.M{"day": t.Day})
}

//SetLoginPrizeList 添加登录奖励
func SetLoginPrizeList() {
	prize2 := NewLoginPrizeProp(int32(pb.PROP_TYPE11), 1)
	prize3 := NewLoginPrizeProp(int32(pb.PROP_TYPE4), 1)
	prize4 := NewLoginPrizeProp(int32(pb.PROP_TYPE6), 1)
	var i uint32
	for i = 0; i < 7; i++ {
		prize := []LoginPrizeProp{}
		var num = 10000 + int32(i)*5000
		prize1 := NewLoginPrizeProp(int32(pb.PROP_TYPE2), num)
		NewLoginPrize(i, append(prize, prize1, prize2, prize3, prize4))
	}
}

//LoadLoginPrizeList load login prize info by prize.json
func LoadLoginPrizeList() []LoginPrize {
	filePath := "static/prize.json"
	list := make([]LoginPrize, 0)
	err := libs.Load(filePath, &list)
	if err != nil {
		beego.Error("load prize err ", err)
	}
	return list
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
	//test use
	if !RunMode() {
		if len(list) == 0 {
			//SetLoginPrizeList()
			list = LoadLoginPrizeList()
		}
	}
	Cache.Put("prize", list, 0)
	for k, v := range list {
		Cache.Put(PrizeKey(v.Day), &list[k], 0)
	}
}

//UpsertPrize upsert prize
func UpsertPrize(prize LoginPrize) bool {
	if len(prize.ID) == 0 {
		beego.Error("prize id err: ", prize)
		return false
	}
	key := PrizeKey(prize.Day)
	list := GetLoginPrizes()
	if prize.Del != 0 {
		if !prize.Delete() {
			beego.Error("prize delete err: ", prize)
			return false
		}
		Cache.Delete(key)
		for k, v := range list {
			if v.Day == prize.Day {
				list = append(list[:k], list[k+1:]...)
				break
			}
		}
		Cache.Put("prize", list, 0)
		return true
	}
	if !prize.Upsert() {
		beego.Error("prize upsert err: ", prize)
		return false
	}
	Cache.Put(key, &prize, 0)
	for k, v := range list {
		if v.Day == prize.Day {
			list[k] = prize
			break
		}
	}
	Cache.Put("prize", list, 0)
	return true
}

//PrizeKey cache prize unique key
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
