// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Unpack 解包消息
func Unpack(id uint32, b []byte) (interface{}, error) {
	switch id {
	case 1001:
		msg := new(CWxLogin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1002:
		msg := new(CLogin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1004:
		msg := new(CUserData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1005:
		msg := new(CGateData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1006:
		msg := new(CPing)
		err := msg.Unmarshal(b)
		return msg, err
	case 1007:
		msg := new(CPropData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1008:
		msg := new(CGetCurrency)
		err := msg.Unmarshal(b)
		return msg, err
	case 1011:
		msg := new(CShop)
		err := msg.Unmarshal(b)
		return msg, err
	case 1012:
		msg := new(CBuy)
		err := msg.Unmarshal(b)
		return msg, err
	case 1013:
		msg := new(COverData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1014:
		msg := new(CCard)
		err := msg.Unmarshal(b)
		return msg, err
	case 1015:
		msg := new(CLoginPrize)
		err := msg.Unmarshal(b)
		return msg, err
	case 1016:
		msg := new(CUseProp)
		err := msg.Unmarshal(b)
		return msg, err
	case 1017:
		msg := new(CStart)
		err := msg.Unmarshal(b)
		return msg, err
	default:
		return nil, errors.New("unknown message")
	}
}