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
	case *CGateData:
		b, err := msg.(*CGateData).Marshal()
		return 1005, b, err
	case *CPing:
		b, err := msg.(*CPing).Marshal()
		return 1006, b, err
	case *CPropData:
		b, err := msg.(*CPropData).Marshal()
		return 1007, b, err
	case *CGetCurrency:
		b, err := msg.(*CGetCurrency).Marshal()
		return 1008, b, err
	case *CShop:
		b, err := msg.(*CShop).Marshal()
		return 1011, b, err
	case *CBuy:
		b, err := msg.(*CBuy).Marshal()
		return 1012, b, err
	case *CTempShop:
		b, err := msg.(*CTempShop).Marshal()
		return 1013, b, err
	case *COverData:
		b, err := msg.(*COverData).Marshal()
		return 1014, b, err
	case *CCard:
		b, err := msg.(*CCard).Marshal()
		return 1015, b, err
	case *CLoginPrize:
		b, err := msg.(*CLoginPrize).Marshal()
		return 1016, b, err
	case *CUseProp:
		b, err := msg.(*CUseProp).Marshal()
		return 1017, b, err
	case *CStart:
		b, err := msg.(*CStart).Marshal()
		return 1018, b, err
	default:
		return 0, []byte{}, errors.New("unknown message")
	}
}