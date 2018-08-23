// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: game_code.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ErrCode int32

const (
	OK          ErrCode = 0
	RegistFaild ErrCode = 1
	LoginFaild  ErrCode = 2
)

var ErrCode_name = map[int32]string{
	0: "OK",
	1: "RegistFaild",
	2: "LoginFaild",
}
var ErrCode_value = map[string]int32{
	"OK":          0,
	"RegistFaild": 1,
	"LoginFaild":  2,
}

func (ErrCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameCode, []int{0} }

func init() {
	proto.RegisterEnum("pb.ErrCode", ErrCode_name, ErrCode_value)
}
func (x ErrCode) String() string {
	s, ok := ErrCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("game_code.proto", fileDescriptorGameCode) }

var fileDescriptorGameCode = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4f, 0xcc, 0x4d,
	0x8d, 0x4f, 0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0xd2,
	0x32, 0xe2, 0x62, 0x77, 0x2d, 0x2a, 0x72, 0xce, 0x4f, 0x49, 0x15, 0x62, 0xe3, 0x62, 0xf2, 0xf7,
	0x16, 0x60, 0x10, 0xe2, 0xe7, 0xe2, 0x0e, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x71, 0x4b, 0xcc, 0xcc,
	0x49, 0x11, 0x60, 0x14, 0xe2, 0xe3, 0xe2, 0xf2, 0xc9, 0x4f, 0xcf, 0xcc, 0x83, 0xf0, 0x99, 0x9c,
	0x74, 0x2e, 0x3c, 0x94, 0x63, 0xb8, 0xf1, 0x50, 0x8e, 0xe1, 0xc3, 0x43, 0x39, 0xc6, 0x86, 0x47,
	0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83,
	0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x43,
	0x12, 0x1b, 0xd8, 0x32, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0x50, 0x39, 0x5b, 0x7f,
	0x00, 0x00, 0x00,
}
