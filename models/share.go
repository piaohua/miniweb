package models

import (
	"time"

	"miniweb/libs"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
)

//Share 每日分享配置
type Share struct {
	ID     string      `bson:"_id" json:"id"`        //ID
	Number int32       `bson:"number" json:"number"` //数量
	Info   string      `bson:"info" json:"info"`     //描述信息
	Prize  []PrizeProp `bson:"prize" json:"prize"`   //prize
	Del    int         `bson:"del" json:"del"`       //是否移除
	Ctime  time.Time   `bson:"ctime" json:"ctime"`   //创建时间
}

//ShareInfo share info
type ShareInfo struct {
	ID     string `bson:"id,omitempty" json:"id,omitempty"`         //id
	Status int32  `bson:"status,omitempty" json:"status,omitempty"` //领取奖励
}

//GetShareList 每日分享配置列表
func GetShareList() []Share {
	var list []Share
	ListByQ(Shares, bson.M{"del": 0}, &list)
	return list
}

//Save 写入数据库
func (t *Share) Save() bool {
	t.Ctime = bson.Now()
	return Insert(Shares, t)
}

//Upsert 更新数据库
func (t *Share) Upsert() bool {
	return Upsert(Shares, bson.M{"_id": t.ID}, t)
}

//Delete 删除数据
func (t *Share) Delete() bool {
	return Delete(Shares, bson.M{"_id": t.ID})
}

//LoadShareList load share info by share.json
func LoadShareList() []Share {
	filePath := "static/share.json"
	list := make([]Share, 0)
	err := libs.Load(filePath, &list)
	if err != nil {
		beego.Error("load share err ", err)
	}
	return list
}

//InitShareList init share to cache
func InitShareList() {
	list := GetShareList()
	//test use
	if !RunMode() {
		if len(list) == 0 {
			list = LoadShareList()
		}
	}
	Cache.Put("share", list, 0)
	for k, v := range list {
		Cache.Put(ShareKey(v.ID), &list[k], 0)
	}
}

//UpsertShare upsert share
func UpsertShare(share Share) bool {
	if len(share.ID) == 0 {
		beego.Error("share id err: ", share)
		return false
	}
	key := ShareKey(share.ID)
	list := GetShares()
	if share.Del != 0 {
		if !share.Delete() {
			beego.Error("share delete err: ", share)
			return false
		}
		Cache.Delete(key)
		for k, v := range list {
			if v.ID == share.ID {
				list = append(list[:k], list[k+1:]...)
				break
			}
		}
		Cache.Put("share", list, 0)
		return true
	}
	if !share.Upsert() {
		beego.Error("share upsert err: ", share)
		return false
	}
	Cache.Put(key, &share, 0)
	var has bool
	for k, v := range list {
		if v.ID == share.ID {
			list[k] = share
			has = true
			break
		}
	}
	if !has {
		list = append(list, share)
	}
	Cache.Put("share", list, 0)
	return true
}

//ShareKey cache share unique key
func ShareKey(id string) string {
	return "share" + id
}

//GetShares from cache
func GetShares() (l []Share) {
	if v := Cache.Get("share"); v != nil {
		if val, ok := v.([]Share); ok {
			l = val
		}
	}
	return
}

//GetShare get share from cache by id
func GetShare(id string) (share *Share) {
	if v := Cache.Get(ShareKey(id)); v != nil {
		if val, ok := v.(*Share); ok {
			share = val
		}
	}
	return
}
