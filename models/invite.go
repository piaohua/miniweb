package models

import (
	"time"

	"miniweb/libs"
	"miniweb/pb"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
)

//Invite 邀请好友设定
type Invite struct {
	ID     string      `bson:"_id" json:"id"`        //ID
	Type   int32       `bson:"type" json:"type"`     //0每日,1累计状态
	Number int32       `bson:"number" json:"number"` //数量
	Info   string      `bson:"info" json:"info"`     //描述信息
	Prize  []PrizeProp `bson:"prize" json:"prize"`   //prize
	Del    int         `bson:"del" json:"del"`       //是否移除
	Ctime  time.Time   `bson:"ctime" json:"ctime"`   //创建时间
}

//InviteInfo invite info
type InviteInfo struct {
	ID     string `bson:"id,omitempty" json:"id,omitempty"`         //id
	Status int32  `bson:"status,omitempty" json:"status,omitempty"` //领取奖励
}

//GetInviteList 邀请配置列表
func GetInviteList() []Invite {
	var list []Invite
	ListByQ(Invites, bson.M{"del": 0}, &list)
	return list
}

//Save 写入数据库
func (t *Invite) Save() bool {
	t.Ctime = bson.Now()
	return Insert(Invites, t)
}

//Upsert 更新数据库
func (t *Invite) Upsert() bool {
	return Upsert(Invites, bson.M{"_id": t.ID}, t)
}

//Delete 删除数据
func (t *Invite) Delete() bool {
	return Delete(Invites, bson.M{"_id": t.ID})
}

//LoadInviteList load invite info by invite.json
func LoadInviteList() []Invite {
	filePath := "static/invite.json"
	list := make([]Invite, 0)
	err := libs.Load(filePath, &list)
	if err != nil {
		beego.Error("load invite err ", err)
	}
	return list
}

//InitInviteList init invite to cache
func InitInviteList() {
	list := GetInviteList()
	//test use
	if !RunMode() {
		if len(list) == 0 {
			list = LoadInviteList()
		}
	}
	Cache.Put("invite", list, 0)
	for k, v := range list {
		Cache.Put(InviteKey(v.ID), &list[k], 0)
	}
}

//UpsertInvite upsert invite
func UpsertInvite(invite Invite) bool {
	if len(invite.ID) == 0 {
		beego.Error("invite id err: ", invite)
		return false
	}
	key := InviteKey(invite.ID)
	list := GetInvites()
	if invite.Del != 0 {
		if !invite.Delete() {
			beego.Error("invite delete err: ", invite)
			return false
		}
		Cache.Delete(key)
		for k, v := range list {
			if v.ID == invite.ID {
				list = append(list[:k], list[k+1:]...)
				break
			}
		}
		Cache.Put("invite", list, 0)
		return true
	}
	if !invite.Upsert() {
		beego.Error("invite upsert err: ", invite)
		return false
	}
	Cache.Put(key, &invite, 0)
	var has bool
	for k, v := range list {
		if v.ID == invite.ID {
			list[k] = invite
			has = true
			break
		}
	}
	if !has {
		list = append(list, invite)
	}
	Cache.Put("invite", list, 0)
	return true
}

//InviteKey cache invite unique key
func InviteKey(id string) string {
	return "invite" + id
}

//GetInvites from cache
func GetInvites() (l []Invite) {
	if v := Cache.Get("invite"); v != nil {
		if val, ok := v.([]Invite); ok {
			l = val
		}
	}
	return
}

//GetInvite get invite from cache by id
func GetInvite(id string) (invite *Invite) {
	if v := Cache.Get(InviteKey(id)); v != nil {
		if val, ok := v.(*Invite); ok {
			invite = val
		}
	}
	return
}

//SetInviteByID set invite
func SetInviteByID(id string) {
	user := new(User)
	user.ID = id
	user.Get()
	beego.Info("user ", user)
	InviteInit(user)
	SetInvite(id, user)
}

//SetInvite set invite
func SetInvite(id string, user *User) {
	user.InviteNum++
	user.InviteCount++
	user.InviteTime = time.Now().Local()
	if user.InviteInfo == nil {
		beego.Error("SetInvite failed ", id)
		user.Save()
		return
	}
	list := GetInvites()
	for _, v := range list {
		if val, ok := user.InviteInfo[v.ID]; ok {
			switch val.Status {
			case int32(pb.PrizeGot):
				continue
			}
		}
		user.InviteInfo[v.ID] = InviteInfo{
			ID:     v.ID,
			Status: int32(pb.PrizeDone),
		}
	}
	beego.Info("user ", user)
	user.Save()
}

//InviteInit invite init
func InviteInit(user *User) {
	if user.InviteInfo == nil {
		user.InviteInfo = make(map[string]InviteInfo)
		return
	}
	today := libs.TodayTime()
	if user.InviteTime.After(today) {
		return
	}
	//reset
	user.InviteNum = 0
	for k, v := range user.InviteInfo {
		prize := GetInvite(k)
		if prize == nil {
			beego.Error("inviteInit failed ", k)
			delete(user.InviteInfo, k)
			continue
		}
		switch prize.Type {
		case int32(pb.InviteToday):
			v.Status = int32(pb.PrizeNone) //reset
			user.InviteInfo[k] = v
		}
	}
}
