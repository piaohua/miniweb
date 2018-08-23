package models

//SessionResult 3rd_session结果
type SessionResult struct {
	Session string `json:"session"`
	WsAddr  string `json:"wsaddr"`
	WxErr
}
