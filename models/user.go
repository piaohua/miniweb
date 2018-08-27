/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-25 21:26:30
 * Filename      : user.go
 * Description   : 玩家数据
 * *******************************************************/

package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//User 用户数据
type User struct {
	ID          string    `bson:"_id", json:"id"`
	NickName    string    `bson:"nick_name", json:"nick_name"`
	AvatarUrl   string    `bson:"avatar_url", json:"avatar_url"`
	Gender      int32     `bson:"gender", json:"gender"`
	Session     string    `bson:"session", json:"session"`
	SessionTime time.Time `bson:"session_time", json:"session_time"`
	OpenId      string    `bson:"openid", json:"openid"`
	SessionKey  string    `bson:"session_key", json:"session_key"`
	UnionId     string    `bson:"unionid", json:"unionid"`
	City        string    `bson:"city", json:"city"`
	Province    string    `bson:"province", json:"province"`
	Country     string    `bson:"country", json:"country"`
}

//Save 保存
func (u *User) Save() bool {
	return Upsert(Users, bson.M{"_id": u.ID}, u)
}

//HasOpenid openid是否存在
func HasOpenid(openid string) bool {
	return Has(Users, bson.M{"openid": openid})
}

//GetByOpenid  通过Openid获取
func (u *User) GetByOpenid(openid string) {
	GetByQ(Users, bson.M{"openid": openid}, u)
}

//HasSession session是否存在
func HasSession(session string) bool {
	return Has(Users, bson.M{"session": session})
}

//HasID id是否存在
func HasID(id string) bool {
	return Has(Users, bson.M{"_id": id})
}

//GetBySession  通过session获取
func (u *User) GetBySession(session string) {
	GetByQ(Users, bson.M{"session": session}, u)
}

//UpdateSessionKey 更新session_key
func (u *User) UpdateSessionKey() bool {
	return Update(Users, bson.M{"_id": u.ID},
		bson.M{"$set": bson.M{"session_key": u.SessionKey}})
}
