// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Rpacket 打包消息
func Rpacket(msg interface{}) (uint32, []byte, error) {
	switch msg.(type) {
	case *CWxLogin:
		b, err := msg.(*CWxLogin).Marshal()
		return 1001, b, err
	case *CLogin:
		b, err := msg.(*CLogin).Marshal()
		return 1002, b, err
	case *CUserData:
		b, err := msg.(*CUserData).Marshal()
		return 1004, b, err
	case *CGameData:
		b, err := msg.(*CGameData).Marshal()
		return 1005, b, err
	case *CPing:
		b, err := msg.(*CPing).Marshal()
		return 1006, b, err
	default:
		return 0, []byte{}, errors.New("unknown message")
	}
}