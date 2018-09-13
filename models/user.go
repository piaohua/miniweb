/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-30 01:09:26
 * Filename      : user.go
 * Description   : 玩家数据
 * *******************************************************/

package models

import (
	"time"

	"miniweb/pb"

	"github.com/globalsign/mgo/bson"
)

//User 用户数据
type User struct {
	ID          string    `bson:"_id", json:"id"`
	NickName    string    `bson:"nick_name", json:"nick_name"`
	AvatarUrl   string    `bson:"avatar_url", json:"avatar_url"`
	Gender      int32     `bson:"gender", json:"gender"`
	Session     string    `bson:"session", json:"session"`
	SessionTime time.Time `bson:"session_time,omitempty", json:"session_time,omitempty"`
	OpenId      string    `bson:"openid", json:"openid"`
	SessionKey  string    `bson:"session_key", json:"session_key"`
	UnionId     string    `bson:"unionid", json:"unionid"`
	City        string    `bson:"city", json:"city"`
	Province    string    `bson:"province", json:"province"`
	Country     string    `bson:"country", json:"country"`
	//
	RegistIP  string    `bson:"regist_ip" json:"regist_ip"`   // 注册账户时的IP地址
	LoginIP   string    `bson:"login_ip" json:"login_ip"`     // 登录账户时的IP地址
	Ctime     time.Time `bson:"ctime" json:"ctime"`           // 注册时间
	LoginTime time.Time `bson:"login_time" json:"login_time"` // 最后登录时间
	//
	LoginTimes uint32 `bson:"login_times" json:"login_times"` //连续登录次数
	LoginPrize uint32 `bson:"login_prize" json:"login_prize"` //连续登录奖励
	LoginCount uint32 `bson:"login_count" json:"login_count"` //累计登录次数
	//
	Diamond    int64 `bson:"diamond" json:"diamond"`                             // 钻石
	Coin       int64 `bson:"coin" json:"coin"`                                   // 金币
	Energy     int64 `bson:"energy" json:"energy"`                               // 精力
	EnergyTime int64 `bson:"energy_time,omitempty" json:"energy_time,omitempty"` // 精力恢复时间
	//
	Gate map[string]GateInfo `bson:"gate,omitempty" json:"gate,omitempty"` // 关卡
	Prop map[string]PropInfo `bson:"prop,omitempty" json:"prop,omitempty"` // 道具
	//
	TempProp map[string]TempPropInfo `bson:"temp_prop,omitempty" json:"temp_prop,omitempty"` // 道具
	//
	ShareNum    int32                 `bson:"share_num,omitempty" json:"share_num,omitempty"`       //当天分享次数
	ShareTime   time.Time             `bson:"share_time,omitempty" json:"share_time,omitempty"`     //当天分享时间
	ShareInfo   map[string]ShareInfo  `bson:"share_info,omitempty" json:"share_info,omitempty"`     // share info
	Invite      string                `bson:"invite,omitempty" json:"invite,omitempty"`             //邀请userid
	InviteNum   int32                 `bson:"invite_num,omitempty" json:"invite_num,omitempty"`     //当天邀请总数
	InviteTime  time.Time             `bson:"invite_time,omitempty" json:"invite_time,omitempty"`   //当天邀请时间
	InviteCount int32                 `bson:"invite_count,omitempty" json:"invite_count,omitempty"` //累计邀请总数
	InviteInfo  map[string]InviteInfo `bson:"invite_info,omitempty" json:"invite_info,omitempty"`   // invite info
}

//Get 加载
func (u *User) Get() {
	Get(Users, u.ID, u)
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

//UpdateCurrency 更新Currency
func UpdateCurrency(arg *pb.ChangeCurrency) bool {
	return Update(Users, bson.M{"_id": arg.Userid},
		bson.M{"$inc": bson.M{"coin": arg.Coin, "diamond": arg.Diamond}})
}

//GetBySession  通过session获取
func (u *User) GetBySession(session string) {
	GetByQ(Users, bson.M{"session": session}, u)
}

//GetOpenid  通过session获取openid
func GetOpenid(session string) string {
	u := new(User)
	GetByQWithFields(Users, bson.M{"session": session}, []string{"openid"}, &u)
	return u.OpenId
}

//GetSessionKey  通过session获取session_key
func GetSessionKey(session string) (sessionKey string) {
	u := bson.M{}
	GetByQWithFields(Users, bson.M{"session": session}, []string{"session_key"}, &u)
	if val, ok := u["_id"]; ok {
		if s, ok := val.(string); ok {
			if len(s) == 0 {
				return
			}
		}
	}
	if val, ok := u["session_key"]; ok {
		if s, ok := val.(string); ok {
			sessionKey = s
		}
	}
	return
}

//UpdateSessionKey 更新session_key
func (u *User) UpdateSessionKey() bool {
	return Update(Users, bson.M{"_id": u.ID},
		bson.M{"$set": bson.M{"session_key": u.SessionKey}})
}

//AddEnergy add energy
func (u *User) AddEnergy(num int64) {
	u.Energy += num
	if u.Energy < 0 {
		u.Energy = 0
	}
	if u.Energy > 30 {
		u.Energy = 30
	}
}

//AddCoin add coin
func (u *User) AddCoin(num int64) {
	u.Coin += num
	if u.Coin < 0 {
		u.Coin = 0
	}
}

//AddDiamond add diamond
func (u *User) AddDiamond(num int64) {
	u.Diamond += num
	if u.Diamond < 0 {
		u.Diamond = 0
	}
}

//AddCoinMsg add coin msg
func AddCoinMsg(user *User, num int64) (msg *pb.SPushProp) {
	user.AddCoin(num)
	msg = &pb.SPushProp{
		//Type: pb.LOG_TYPE0,
		Num: num,
		PropInfo: &pb.PropData{
			Type: pb.PROP_TYPE2,
			Num:  user.Coin,
		},
	}
	return
}

//AddDiamondMsg add diamond msg
func AddDiamondMsg(user *User, num int64) (msg *pb.SPushProp) {
	user.AddDiamond(num)
	msg = &pb.SPushProp{
		//Type: pb.LOG_TYPE0,
		Num: num,
		PropInfo: &pb.PropData{
			Type: pb.PROP_TYPE1,
			Num:  user.Diamond,
		},
	}
	return
}
