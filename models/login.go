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
	user.Session = session
	//user.SessionTime = time.Now()
	user.ID = id
	user.RegistIP = ip
	user.Ctime = time.Now()
	if !user.Save() {
		err = errors.New("session save failed")
	}
	return
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
	user.Session = session
	//user.SessionTime = time.Now()
	user.ID = id
	user.RegistIP = ip
	user.Ctime = time.Now()
	if !user.Save() {
		err = errors.New("session save failed")
	}
	return
}

//VerifyUserLogin 登录校验(test)
func VerifyUserLogin(arg *pb.CLogin, session string) (*User, error) {
	user := new(User)
	user.GetBySession(session)
	if len(user.Session) == 0 || len(user.ID) == 0 {
		return nil, errors.New("session invaild")
	}
	sign := libs.Sha1Signature(strconv.Itoa(int(arg.GetTimestamp())), session)
	if sign != arg.GetSignature() {
		return nil, errors.New("sign failed")
	}
	if (time.Now().Unix() - arg.GetTimestamp()) > 10 {
		return nil, errors.New("session expire")
	}
	return user, nil
}
