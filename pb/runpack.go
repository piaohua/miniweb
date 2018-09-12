// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Runpack 解包消息
func Runpack(id uint32, b []byte) (interface{}, error) {
	switch id {
	case 1001:
		msg := new(SWxLogin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1002:
		msg := new(SLogin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1003:
		msg := new(SLoginOut)
		err := msg.Unmarshal(b)
		return msg, err
	case 1004:
		msg := new(SUserData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1005:
		msg := new(SGateData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1006:
		msg := new(SPing)
		err := msg.Unmarshal(b)
		return msg, err
	case 1007:
		msg := new(SPropData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1008:
		msg := new(SGetCurrency)
		err := msg.Unmarshal(b)
		return msg, err
	case 1010:
		msg := new(SPushProp)
		err := msg.Unmarshal(b)
		return msg, err
	case 1011:
		msg := new(SShop)
		err := msg.Unmarshal(b)
		return msg, err
	case 1012:
		msg := new(SBuy)
		err := msg.Unmarshal(b)
		return msg, err
	case 1013:
		msg := new(STempShop)
		err := msg.Unmarshal(b)
		return msg, err
	case 1014:
		msg := new(SOverData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1015:
		msg := new(SCard)
		err := msg.Unmarshal(b)
		return msg, err
	case 1016:
		msg := new(SLoginPrize)
		err := msg.Unmarshal(b)
		return msg, err
	case 1017:
		msg := new(SUseProp)
		err := msg.Unmarshal(b)
		return msg, err
	case 1018:
		msg := new(SStart)
		err := msg.Unmarshal(b)
		return msg, err
	case 1019:
		msg := new(SShareInfo)
		err := msg.Unmarshal(b)
		return msg, err
	case 1020:
		msg := new(SInviteInfo)
		err := msg.Unmarshal(b)
		return msg, err
	case 1021:
		msg := new(SShare)
		err := msg.Unmarshal(b)
		return msg, err
	case 1022:
		msg := new(SInvite)
		err := msg.Unmarshal(b)
		return msg, err
	default:
		return nil, errors.New("unknown message")
	}
}