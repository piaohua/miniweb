package models

import (
	"errors"
	"time"

	"miniweb/libs"
	"miniweb/pb"

	"github.com/astaxie/beego"
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

//HasOpenid 是否存在
func HasOpenid(openid string) bool {
	return Has(Users, bson.M{"openid": openid})
}

//GetByOpenid  通过Openid获取
func (u *User) GetByOpenid(openid string) {
	GetByQ(Users, bson.M{"openid": openid}, u)
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

//GetSession 获取session
func GetSession(jscode string) (session string, err error) {
	wxs, err := WX.Jscode2Session(jscode)
	if err != nil {
		return
	}
	if len(wxs.OpenId) == 0 {
		err = errors.New("get session failed")
		return
	}
	//已经存在
	if HasOpenid(wxs.OpenId) {
		user := new(User) //TODO 查找session字段
		user.GetByOpenid(wxs.OpenId)
		if user.OpenId != wxs.OpenId {
			err = errors.New("select failed")
			return
		}
		if len(user.Session) == 0 {
			err = errors.New("session error")
			return
		}
		user.SessionKey = wxs.SessionKey //更新session_key
		if !user.UpdateSessionKey() {
			err = errors.New("update session_key failed")
			return
		}
		session = user.Session //暂时不重新生成
		return
	}
	session, err = GenSession()
	if err != nil {
		return
	}
	user := new(User)
	user.OpenId = wxs.OpenId
	user.SessionKey = wxs.SessionKey
	user.UnionId = wxs.UnionId
	user.Session = session
	//user.SessionTime = time.New()
	user.ID = bson.NewObjectId().Hex()
	if !user.Save() {
		err = errors.New("session save failed")
	}
	return
}

//GenSession 生成session
func GenSession() (string, error) {
	c := "cat /dev/urandom | od -x | tr -d ' ' | head -n 1"
	out, err := libs.ExecCmd(c)
	if err != nil {
		beego.Error("GenSession err: ", err)
	}
	return string(out), err
}

//VerifyUserInfo 校验用户信息,不包含 openid 等敏感信息
func VerifyUserInfo(arg *pb.CWxLogin, session string) (*User, error) {
	user := new(User)
	user.GetBySession(session)
	if len(user.Session) == 0 || len(user.ID) == 0 {
		return nil, errors.New("session invaild")
	}
	////验证sha1( rawData + sessionkey )
	sign := libs.Sha1Signature(arg.GetRawData(), user.SessionKey)
	if sign != arg.GetSignature() {
		return nil, errors.New("sign failed")
	}
	//验证敏感信息
	b, err := libs.DecryptWechatAppletUser(arg.GetEncryptedData(), arg.GetIv(), user.SessionKey)
	if err != nil {
		return nil, err
	}
	wxUserInfo, err := ParseUserInfo(b)
	if err != nil {
		return nil, err
	}
	if wxUserInfo.OpenId != user.OpenId {
		return nil, errors.New("openid error")
	}
	if wxUserInfo.Watermark.Appid != WX.Appid {
		return nil, errors.New("appid error")
	}
	if wxUserInfo.Watermark.Timestamp < time.Now().Unix() {
		return nil, errors.New("session expire")
	}
	user.NickName = wxUserInfo.NickName
	user.AvatarUrl = wxUserInfo.AvatarUrl
	user.City = wxUserInfo.City
	user.Country = wxUserInfo.Country
	user.Province = wxUserInfo.Province
	user.Gender = wxUserInfo.Gender
	if !user.Save() {
		beego.Error("VerifyUserInfo save error: ", user)
	}
	return user, nil
}
