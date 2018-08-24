package models

import (
	"fmt"

	"miniweb/libs"

	"github.com/astaxie/beego"
)

// OAUTH2PAGE oauth2鉴权
const (
	Code2PAGE = "https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code"
)

var (
	WX *WeiXin
)

//WeiXin 微信服务
type WeiXin struct {
	Appid  string
	Secret string
}

func init() {
	appid := beego.AppConfig.String("wx.appid")
	secret := beego.AppConfig.String("wx.secret")
	WX = &WeiXin{
		Appid:  appid,
		Secret: secret,
	}
}

// Jscode2Session code换session
func (s *WeiXin) Jscode2Session(code string) (wxs *WxSession, err error) {
	url := fmt.Sprintf(Code2PAGE, s.Appid, s.Secret, code)
	wxs = new(WxSession)
	err = libs.GetJson(url, wxs)

	if wxs.Error() != nil {
		err = wxs.Error()
	}
	return
}

//WxSession 登录获取微信JSON数据
type WxSession struct {
	WxErr
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

//WxErr 通用错误
type WxErr struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (w *WxErr) Error() error {
	if w.ErrCode != 0 {
		return fmt.Errorf("err: errcode=%v , errmsg=%v", w.ErrCode, w.ErrMsg)
	}
	return nil
}
