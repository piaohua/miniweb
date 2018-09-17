/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-08-23 10:47:22
 * Filename      : login.go
 * Description   : login handler
 * *******************************************************/

package models

import (
	"errors"
	"strconv"
	"time"

	"miniweb/libs"
	"miniweb/pb"

	"github.com/astaxie/beego"
)

//RunMode 运行方式
func RunMode() bool {
	runmode := beego.AppConfig.String("runmode")
	return runmode != "dev"
}

//GetSession 获取session
func GetSession(jscode, ip string) (session string, err error) {
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
	session, err = getNewSession()
	if err != nil {
		return
	}
	id, err := getNewID()
	if err != nil {
		return
	}
	user := new(User)
	user.OpenId = wxs.OpenId
	user.SessionKey = wxs.SessionKey
	user.UnionId = wxs.UnionId
	initUserLogin(id, ip, session, user)
	if !user.Save() {
		err = errors.New("session save failed")
	}
	return
}

//VerifyUserInfo 校验用户信息,不包含 openid 等敏感信息
func VerifyUserInfo(arg *pb.CWxLogin, session string) (*WxUserInfo, error) {
	sessionKey := GetSessionKey(session)
	if len(sessionKey) == 0 {
		return nil, errors.New("session invaild")
	}
	////验证sha1( rawData + sessionKey )
	sign := libs.Sha1Signature(arg.GetRawData(), sessionKey)
	if sign != arg.GetSignature() {
		return nil, errors.New("sign failed")
	}
	//验证敏感信息
	b, err := libs.DecryptWechatAppletUser(arg.GetEncryptedData(), sessionKey, arg.GetIv())
	if err != nil {
		beego.Error("VerifyUserInfo error: ", err)
		return nil, err
	}
	beego.Info("VerifyUserInfo userinfo: ", string(b))
	wxUserInfo, err := ParseUserInfo(b)
	if err != nil {
		beego.Error("VerifyUserInfo error: ", err)
		return nil, err
	}
	if wxUserInfo.Watermark.Appid != WX.Appid {
		return nil, errors.New("appid error")
	}
	if (time.Now().Unix() - wxUserInfo.Watermark.Timestamp) > 5 {
		return nil, errors.New("session expire")
	}
	return wxUserInfo, nil
}

//LoginUserInfo 登录获取玩家数据
func LoginUserInfo(wxUserInfo *WxUserInfo, session string,
	Type pb.LoginType) (*User, error) {
	if wxUserInfo == nil {
		return nil, errors.New("wxUserInfo error")
	}
	user := new(User)
	user.GetBySession(session)
	if wxUserInfo.OpenId != user.OpenId {
		return nil, errors.New("openid error")
	}
	switch Type {
	case pb.CODELOGIN:
		return user, nil
	case pb.WXLOGIN:
	}
	user.NickName = wxUserInfo.NickName
	user.AvatarUrl = wxUserInfo.AvatarUrl
	user.City = wxUserInfo.City
	user.Country = wxUserInfo.Country
	user.Province = wxUserInfo.Province
	user.Gender = wxUserInfo.Gender
	if !user.Save() {
		beego.Error("LoginUserInfo save error: ", user)
	}
	return user, nil
}

//GetSessionByCode 获取session(test)
func GetSessionByCode(jscode, ip string) (session string, err error) {
	//已经存在
	if HasOpenid(jscode) {
		user := new(User) //TODO 查找session字段
		user.GetByOpenid(jscode)
		if user.OpenId != jscode {
			err = errors.New("select failed")
			return
		}
		if len(user.Session) == 0 {
			err = errors.New("session error")
			return
		}
		session = user.Session //暂时不重新生成
		return
	}
	session, err = getNewSession()
	if err != nil {
		return
	}
	id, err := getNewID()
	if err != nil {
		return
	}
	user := new(User)
	user.OpenId = jscode
	initUserLogin(id, ip, session, user)
	if !user.Save() {
		err = errors.New("session save failed")
	}
	return
}

//VerifyUserLogin 登录校验(test)
func VerifyUserLogin(arg *pb.CLogin, session string) (*WxUserInfo, error) {
	openid := GetOpenid(session)
	if len(openid) == 0 {
		return nil, errors.New("session invaild")
	}
	sign := libs.Sha1Signature(strconv.Itoa(int(arg.GetTimestamp())), session)
	if sign != arg.GetSignature() {
		return nil, errors.New("sign failed")
	}
	if (time.Now().Unix() - arg.GetTimestamp()) > 10 {
		return nil, errors.New("session expire")
	}
	return &WxUserInfo{OpenId: openid}, nil
}

//initUserLogin login init
func initUserLogin(id, ip, session string, user *User) {
	user.ID = id
	user.RegistIP = ip
	user.Ctime = time.Now()
	user.Session = session
	//user.SessionTime = time.Now()
	user.Energy = 30
	user.Coin = 5000
	user.Diamond = 10
}
