// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: actor_node.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	actor_node.proto
	game_code.proto
	game_login.proto
	game_play.proto
	game_pub.proto
	game_shop.proto
	game_type.proto
	game_user.proto

It has these top-level messages:
	ServeStart
	ServeStarted
	ServeStop
	ServeStoped
	ServeClose
	Tick
	CWxLogin
	SWxLogin
	SLoginOut
	CLogin
	SLogin
	CStart
	SStart
	COverData
	SOverData
	CCard
	SCard
	CLoginPrize
	SLoginPrize
	CUseProp
	SUseProp
	UserData
	GateData
	PropData
	Shop
	LoginPrize
	CShop
	SShop
	CBuy
	SBuy
	CPing
	SPing
	CUserData
	SUserData
	CGateData
	SGateData
	CPropData
	SPropData
	CGetCurrency
	SGetCurrency
	SPushProp
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// 启动服务
type ServeStart struct {
	Message string `protobuf:"bytes,1,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
}

func (m *ServeStart) Reset()                    { *m = ServeStart{} }
func (*ServeStart) ProtoMessage()               {}
func (*ServeStart) Descriptor() ([]byte, []int) { return fileDescriptorActorNode, []int{0} }

func (m *ServeStart) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ServeStarted struct {
	Message string `protobuf:"bytes,1,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
}

func (m *ServeStarted) Reset()                    { *m = ServeStarted{} }
func (*ServeStarted) ProtoMessage()               {}
func (*ServeStarted) Descriptor() ([]byte, []int) { return fileDescriptorActorNode, []int{1} }

func (m *ServeStarted) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 停止服务
type ServeStop struct {
	Message string `protobuf:"bytes,1,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
}

func (m *ServeStop) Reset()                    { *m = ServeStop{} }
func (*ServeStop) ProtoMessage()               {}
func (*ServeStop) Descriptor() ([]byte, []int) { return fileDescriptorActorNode, []int{2} }

func (m *ServeStop) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ServeStoped struct {
	Message string `protobuf:"bytes,1,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
}

func (m *ServeStoped) Reset()                    { *m = ServeStoped{} }
func (*ServeStoped) ProtoMessage()               {}
func (*ServeStoped) Descriptor() ([]byte, []int) { return fileDescriptorActorNode, []int{3} }

func (m *ServeStoped) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 关闭服务
type ServeClose struct {
	Message string `protobuf:"bytes,1,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
}

func (m *ServeClose) Reset()                    { *m = ServeClose{} }
func (*ServeClose) ProtoMessage()               {}
func (*ServeClose) Descriptor() ([]byte, []int) { return fileDescriptorActorNode, []int{4} }

func (m *ServeClose) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 时钟嘀嗒
type Tick struct {
}

func (m *Tick) Reset()                    { *m = Tick{} }
func (*Tick) ProtoMessage()               {}
func (*Tick) Descriptor() ([]byte, []int) { return fileDescriptorActorNode, []int{5} }

func init() {
	proto.RegisterType((*ServeStart)(nil), "pb.ServeStart")
	proto.RegisterType((*ServeStarted)(nil), "pb.ServeStarted")
	proto.RegisterType((*ServeStop)(nil), "pb.ServeStop")
	proto.RegisterType((*ServeStoped)(nil), "pb.ServeStoped")
	proto.RegisterType((*ServeClose)(nil), "pb.ServeClose")
	proto.RegisterType((*Tick)(nil), "pb.Tick")
}
func (this *ServeStart) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServeStart)
	if !ok {
		that2, ok := that.(ServeStart)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *ServeStarted) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServeStarted)
	if !ok {
		that2, ok := that.(ServeStarted)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *ServeStop) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServeStop)
	if !ok {
		that2, ok := that.(ServeStop)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *ServeStoped) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServeStoped)
	if !ok {
		that2, ok := that.(ServeStoped)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *ServeClose) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServeClose)
	if !ok {
		that2, ok := that.(ServeClose)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *Tick) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Tick)
	if !ok {
		that2, ok := that.(Tick)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *ServeStart) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.ServeStart{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ServeStarted) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.ServeStarted{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ServeStop) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.ServeStop{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ServeStoped) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.ServeStoped{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ServeClose) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.ServeClose{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Tick) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&pb.Tick{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringActorNode(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ServeStart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServeStart) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintActorNode(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *ServeStarted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServeStarted) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintActorNode(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *ServeStop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServeStop) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintActorNode(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *ServeStoped) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServeStoped) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintActorNode(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *ServeClose) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServeClose) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintActorNode(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *Tick) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tick) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintActorNode(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ServeStart) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovActorNode(uint64(l))
	}
	return n
}

func (m *ServeStarted) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovActorNode(uint64(l))
	}
	return n
}

func (m *ServeStop) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovActorNode(uint64(l))
	}
	return n
}

func (m *ServeStoped) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovActorNode(uint64(l))
	}
	return n
}

func (m *ServeClose) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovActorNode(uint64(l))
	}
	return n
}

func (m *Tick) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovActorNode(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozActorNode(x uint64) (n int) {
	return sovActorNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ServeStart) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServeStart{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServeStarted) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServeStarted{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServeStop) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServeStop{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServeStoped) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServeStoped{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServeClose) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServeClose{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Tick) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Tick{`,
		`}`,
	}, "")
	return s
}
func valueToStringActorNode(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ServeStart) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServeStart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServeStart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthActorNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActorNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServeStarted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServeStarted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServeStarted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthActorNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActorNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServeStop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServeStop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServeStop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthActorNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActorNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServeStoped) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServeStoped: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServeStoped: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthActorNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActorNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServeClose) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServeClose: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServeClose: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthActorNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActorNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Tick) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Tick: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tick: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipActorNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipActorNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActorNode
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowActorNode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowActorNode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthActorNode
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowActorNode
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipActorNode(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthActorNode = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActorNode   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("actor_node.proto", fileDescriptorActorNode) }

var fileDescriptorActorNode = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4c, 0x2e, 0xc9,
	0x2f, 0x8a, 0xcf, 0xcb, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0x52, 0xe3, 0xe2, 0x0a, 0x4e, 0x2d, 0x2a, 0x4b, 0x0d, 0x2e, 0x49, 0x2c, 0x2a, 0x11, 0x92,
	0xe0, 0x62, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x62, 0xcf, 0x85, 0x70, 0x95, 0x34, 0xb8, 0x78, 0x10, 0xea, 0x52, 0x53, 0xf0, 0xa8, 0x54, 0xe5,
	0xe2, 0x84, 0xaa, 0xcc, 0x2f, 0xc0, 0xa3, 0x4c, 0x9d, 0x8b, 0x1b, 0xae, 0x0c, 0xaf, 0x79, 0x30,
	0x17, 0x3a, 0xe7, 0xe4, 0x17, 0xa7, 0xe2, 0x51, 0xc7, 0xc6, 0xc5, 0x12, 0x92, 0x99, 0x9c, 0xed,
	0xa4, 0x73, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c,
	0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18,
	0x92, 0xd8, 0xc0, 0x41, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x58, 0x72, 0x6a, 0x3c, 0x1e,
	0x01, 0x00, 0x00,
}
