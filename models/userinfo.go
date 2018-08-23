package models

import (
	"encoding/json"
)

//WxUserInfo 微信用户数据
type WxUserInfo struct {
	OpenId    string    `json:"openId"`
	NickName  string    `json:"nickName"`
	Gender    int32     `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarUrl string    `json:"avatarUrl"`
	UnionId   string    `json:"unionId"`
	Watermark Watermark `json:"watermark"`
}

type Watermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

//ParseUserInfo 解析微信用户数据
func ParseUserInfo(b []byte) (*WxUserInfo, error) {
	wxUserInfo := new(WxUserInfo)
	err := json.Unmarshal(b, wxUserInfo)
	if err != nil {
		return nil, err
	}
	return wxUserInfo, nil
}

/*
{
    "openId": "OPENID",
    "nickName": "NICKNAME",
    "gender": GENDER,
    "city": "CITY",
    "province": "PROVINCE",
    "country": "COUNTRY",
    "avatarUrl": "AVATARURL",
    "unionId": "UNIONID",
    "watermark":
    {
        "appid":"APPID",
        "timestamp":TIMESTAMP
    }
}
*/
