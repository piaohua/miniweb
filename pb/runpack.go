// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Runpack 解包消息
func Runpack(id uint32, b []byte) (interface{}, error) {
	switch id {
	case 1501:
		msg := new(SWxLogin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1502:
		msg := new(SLoginOut)
		err := msg.Unmarshal(b)
		return msg, err
	case 1503:
		msg := new(SUserData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1504:
		msg := new(SPing)
		err := msg.Unmarshal(b)
		return msg, err
	default:
		return nil, errors.New("unknown message")
	}
}