package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//LoginPrize 连续登录奖励配置
type LoginPrize struct {
	ID      string    `bson:"_id" json:"id"`          //unique ID
	Day     uint32    `bson:"day" json:"day"`         //unique
	Coin    int64     `bson:"coin" json:"coin"`       //金币奖励
	Diamond int64     `bson:"diamond" json:"diamond"` //钻石奖励
	Del     int       `bson:"del" json:"del"`       //是否移除
	Ctime   time.Time `bson:"ctime" json:"ctime"`     //创建时间
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
	var i uint32
	for i = 0; i < 28; i++ {
		diamond := i * 50 + 300 //基本300,每天增加50
		var coin uint32
		if (i + 1) % 7 == 0 {
			coin += (i / 7) * 100 + 88 //第七天增加88,每周增加100
		}
		t := LoginPrize{
			ID:      bson.NewObjectId().String(),
			Diamond: int64(diamond),
			Coin:    int64(coin),
			Day:     i,
		}
		//config.SetLogin(t)
		t.Save()
	}
}