// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: game_shop.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SBuy_BuyStatus int32

const (
	BuySuccess SBuy_BuyStatus = 0
	BuyFailed  SBuy_BuyStatus = 1
)

var SBuy_BuyStatus_name = map[int32]string{
	0: "BuySuccess",
	1: "BuyFailed",
}
var SBuy_BuyStatus_value = map[string]int32{
	"BuySuccess": 0,
	"BuyFailed":  1,
}

func (SBuy_BuyStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameShop, []int{3, 0} }

// 商城
type CShop struct {
}

func (m *CShop) Reset()                    { *m = CShop{} }
func (*CShop) ProtoMessage()               {}
func (*CShop) Descriptor() ([]byte, []int) { return fileDescriptorGameShop, []int{0} }

type SShop struct {
	List  []*Shop `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
	Error ErrCode `protobuf:"varint,2,opt,name=error,proto3,enum=pb.ErrCode" json:"error,omitempty"`
}

func (m *SShop) Reset()                    { *m = SShop{} }
func (*SShop) ProtoMessage()               {}
func (*SShop) Descriptor() ([]byte, []int) { return fileDescriptorGameShop, []int{1} }

func (m *SShop) GetList() []*Shop {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *SShop) GetError() ErrCode {
	if m != nil {
		return m.Error
	}
	return OK
}

// 商城购买
type CBuy struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *CBuy) Reset()                    { *m = CBuy{} }
func (*CBuy) ProtoMessage()               {}
func (*CBuy) Descriptor() ([]byte, []int) { return fileDescriptorGameShop, []int{2} }

func (m *CBuy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SBuy struct {
	Status SBuy_BuyStatus `protobuf:"varint,1,opt,name=status,proto3,enum=pb.SBuy_BuyStatus" json:"status,omitempty"`
	Error  ErrCode        `protobuf:"varint,2,opt,name=error,proto3,enum=pb.ErrCode" json:"error,omitempty"`
}

func (m *SBuy) Reset()                    { *m = SBuy{} }
func (*SBuy) ProtoMessage()               {}
func (*SBuy) Descriptor() ([]byte, []int) { return fileDescriptorGameShop, []int{3} }

func (m *SBuy) GetStatus() SBuy_BuyStatus {
	if m != nil {
		return m.Status
	}
	return BuySuccess
}

func (m *SBuy) GetError() ErrCode {
	if m != nil {
		return m.Error
	}
	return OK
}

// 商城
type CTempShop struct {
	Type   GateType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.GateType" json:"type,omitempty"`
	Gateid int32    `protobuf:"varint,2,opt,name=gateid,proto3" json:"gateid,omitempty"`
}

func (m *CTempShop) Reset()                    { *m = CTempShop{} }
func (*CTempShop) ProtoMessage()               {}
func (*CTempShop) Descriptor() ([]byte, []int) { return fileDescriptorGameShop, []int{4} }

func (m *CTempShop) GetType() GateType {
	if m != nil {
		return m.Type
	}
	return GATE_TYPE0
}

func (m *CTempShop) GetGateid() int32 {
	if m != nil {
		return m.Gateid
	}
	return 0
}

type STempShop struct {
	List  []*Shop `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
	Error ErrCode `protobuf:"varint,2,opt,name=error,proto3,enum=pb.ErrCode" json:"error,omitempty"`
}

func (m *STempShop) Reset()                    { *m = STempShop{} }
func (*STempShop) ProtoMessage()               {}
func (*STempShop) Descriptor() ([]byte, []int) { return fileDescriptorGameShop, []int{5} }

func (m *STempShop) GetList() []*Shop {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *STempShop) GetError() ErrCode {
	if m != nil {
		return m.Error
	}
	return OK
}

func init() {
	proto.RegisterType((*CShop)(nil), "pb.CShop")
	proto.RegisterType((*SShop)(nil), "pb.SShop")
	proto.RegisterType((*CBuy)(nil), "pb.CBuy")
	proto.RegisterType((*SBuy)(nil), "pb.SBuy")
	proto.RegisterType((*CTempShop)(nil), "pb.CTempShop")
	proto.RegisterType((*STempShop)(nil), "pb.STempShop")
	proto.RegisterEnum("pb.SBuy_BuyStatus", SBuy_BuyStatus_name, SBuy_BuyStatus_value)
}
func (x SBuy_BuyStatus) String() string {
	s, ok := SBuy_BuyStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *CShop) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CShop)
	if !ok {
		that2, ok := that.(CShop)
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
func (this *SShop) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SShop)
	if !ok {
		that2, ok := that.(SShop)
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
	if len(this.List) != len(that1.List) {
		return false
	}
	for i := range this.List {
		if !this.List[i].Equal(that1.List[i]) {
			return false
		}
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *CBuy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CBuy)
	if !ok {
		that2, ok := that.(CBuy)
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
	if this.Id != that1.Id {
		return false
	}
	return true
}
func (this *SBuy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SBuy)
	if !ok {
		that2, ok := that.(SBuy)
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
	if this.Status != that1.Status {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *CTempShop) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CTempShop)
	if !ok {
		that2, ok := that.(CTempShop)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Gateid != that1.Gateid {
		return false
	}
	return true
}
func (this *STempShop) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*STempShop)
	if !ok {
		that2, ok := that.(STempShop)
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
	if len(this.List) != len(that1.List) {
		return false
	}
	for i := range this.List {
		if !this.List[i].Equal(that1.List[i]) {
			return false
		}
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *CShop) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&pb.CShop{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SShop) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.SShop{")
	if this.List != nil {
		s = append(s, "List: "+fmt.Sprintf("%#v", this.List)+",\n")
	}
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CBuy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.CBuy{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SBuy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.SBuy{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CTempShop) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.CTempShop{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Gateid: "+fmt.Sprintf("%#v", this.Gateid)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *STempShop) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.STempShop{")
	if this.List != nil {
		s = append(s, "List: "+fmt.Sprintf("%#v", this.List)+",\n")
	}
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGameShop(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CShop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CShop) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SShop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SShop) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, msg := range m.List {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGameShop(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Error != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGameShop(dAtA, i, uint64(m.Error))
	}
	return i, nil
}

func (m *CBuy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CBuy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGameShop(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func (m *SBuy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SBuy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGameShop(dAtA, i, uint64(m.Status))
	}
	if m.Error != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGameShop(dAtA, i, uint64(m.Error))
	}
	return i, nil
}

func (m *CTempShop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CTempShop) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGameShop(dAtA, i, uint64(m.Type))
	}
	if m.Gateid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGameShop(dAtA, i, uint64(m.Gateid))
	}
	return i, nil
}

func (m *STempShop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *STempShop) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, msg := range m.List {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGameShop(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Error != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGameShop(dAtA, i, uint64(m.Error))
	}
	return i, nil
}

func encodeVarintGameShop(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CShop) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *SShop) Size() (n int) {
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovGameShop(uint64(l))
		}
	}
	if m.Error != 0 {
		n += 1 + sovGameShop(uint64(m.Error))
	}
	return n
}

func (m *CBuy) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGameShop(uint64(l))
	}
	return n
}

func (m *SBuy) Size() (n int) {
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovGameShop(uint64(m.Status))
	}
	if m.Error != 0 {
		n += 1 + sovGameShop(uint64(m.Error))
	}
	return n
}

func (m *CTempShop) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovGameShop(uint64(m.Type))
	}
	if m.Gateid != 0 {
		n += 1 + sovGameShop(uint64(m.Gateid))
	}
	return n
}

func (m *STempShop) Size() (n int) {
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovGameShop(uint64(l))
		}
	}
	if m.Error != 0 {
		n += 1 + sovGameShop(uint64(m.Error))
	}
	return n
}

func sovGameShop(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGameShop(x uint64) (n int) {
	return sovGameShop(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CShop) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CShop{`,
		`}`,
	}, "")
	return s
}
func (this *SShop) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SShop{`,
		`List:` + strings.Replace(fmt.Sprintf("%v", this.List), "Shop", "Shop", 1) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CBuy) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CBuy{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SBuy) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SBuy{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CTempShop) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CTempShop{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Gateid:` + fmt.Sprintf("%v", this.Gateid) + `,`,
		`}`,
	}, "")
	return s
}
func (this *STempShop) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&STempShop{`,
		`List:` + strings.Replace(fmt.Sprintf("%v", this.List), "Shop", "Shop", 1) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGameShop(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CShop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGameShop
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
			return fmt.Errorf("proto: CShop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CShop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGameShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGameShop
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
func (m *SShop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGameShop
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
			return fmt.Errorf("proto: SShop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SShop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGameShop
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &Shop{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			m.Error = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Error |= (ErrCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGameShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGameShop
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
func (m *CBuy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGameShop
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
			return fmt.Errorf("proto: CBuy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CBuy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameShop
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
				return ErrInvalidLengthGameShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGameShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGameShop
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
func (m *SBuy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGameShop
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
			return fmt.Errorf("proto: SBuy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SBuy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (SBuy_BuyStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			m.Error = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Error |= (ErrCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGameShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGameShop
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
func (m *CTempShop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGameShop
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
			return fmt.Errorf("proto: CTempShop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CTempShop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (GateType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gateid", wireType)
			}
			m.Gateid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gateid |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGameShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGameShop
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
func (m *STempShop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGameShop
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
			return fmt.Errorf("proto: STempShop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: STempShop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGameShop
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &Shop{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			m.Error = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGameShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Error |= (ErrCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGameShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGameShop
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
func skipGameShop(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGameShop
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
					return 0, ErrIntOverflowGameShop
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
					return 0, ErrIntOverflowGameShop
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
				return 0, ErrInvalidLengthGameShop
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGameShop
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
				next, err := skipGameShop(dAtA[start:])
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
	ErrInvalidLengthGameShop = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGameShop   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("game_shop.proto", fileDescriptorGameShop) }

var fileDescriptorGameShop = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xb1, 0x6e, 0xb3, 0x30,
	0x14, 0x85, 0x31, 0x3f, 0xe4, 0x2f, 0x37, 0x2d, 0x8d, 0x3c, 0x44, 0x51, 0x54, 0x59, 0x94, 0x29,
	0x8a, 0x2a, 0x86, 0xf4, 0x0d, 0x40, 0x69, 0x3b, 0x74, 0x82, 0xec, 0x15, 0x04, 0x2b, 0x41, 0x4a,
	0x64, 0x0b, 0xcc, 0xc0, 0xd6, 0xbe, 0x41, 0x1f, 0xa3, 0x8f, 0xd2, 0x31, 0x63, 0xc7, 0xe2, 0x2e,
	0x1d, 0xf3, 0x08, 0x95, 0x0d, 0xca, 0x5c, 0xa9, 0x1b, 0xe7, 0x7c, 0x87, 0x7b, 0xec, 0x6b, 0xb8,
	0xdc, 0xa4, 0x7b, 0xfa, 0x54, 0x6d, 0x19, 0x0f, 0x78, 0xc9, 0x04, 0xc3, 0x26, 0xcf, 0xa6, 0x9d,
	0xb9, 0x66, 0x39, 0xed, 0xcc, 0xa9, 0xab, 0x0d, 0x5e, 0x67, 0xbd, 0xee, 0x02, 0xa2, 0xe1, 0x7d,
	0xc0, 0xff, 0x0f, 0x76, 0x94, 0x6c, 0x19, 0xf7, 0x1f, 0xc0, 0x4e, 0xd4, 0x07, 0xbe, 0x02, 0x6b,
	0x57, 0x54, 0x62, 0x82, 0xbc, 0x7f, 0xb3, 0xe1, 0xe2, 0x2c, 0xe0, 0x59, 0xa0, 0xfc, 0x58, 0xbb,
	0xf8, 0x1a, 0x6c, 0x5a, 0x96, 0xac, 0x9c, 0x98, 0x1e, 0x9a, 0xb9, 0x8b, 0xa1, 0xc2, 0xcb, 0xb2,
	0x8c, 0x58, 0x4e, 0xe3, 0x8e, 0xf8, 0x63, 0xb0, 0xa2, 0xb0, 0x6e, 0xb0, 0x0b, 0x66, 0x91, 0x4f,
	0x90, 0x87, 0x66, 0x4e, 0x6c, 0x16, 0xb9, 0xff, 0x82, 0xc0, 0x4a, 0x14, 0x98, 0xc3, 0xa0, 0x12,
	0xa9, 0xa8, 0x2b, 0x0d, 0xdd, 0x05, 0xd6, 0x1d, 0x61, 0xdd, 0x04, 0x61, 0xdd, 0x24, 0x9a, 0xc4,
	0x7d, 0xe2, 0x37, 0x7d, 0x73, 0x70, 0x4e, 0xff, 0x61, 0x17, 0x40, 0x89, 0x7a, 0xbd, 0xa6, 0x55,
	0x35, 0x32, 0xf0, 0x85, 0x86, 0x77, 0x69, 0xb1, 0xa3, 0xf9, 0x08, 0xf9, 0x4b, 0x70, 0xa2, 0x15,
	0xdd, 0x73, 0x7d, 0x53, 0x0f, 0x2c, 0xb5, 0x89, 0xfe, 0x14, 0xe7, 0x6a, 0xf4, 0x7d, 0x2a, 0xe8,
	0xaa, 0xe1, 0x34, 0xd6, 0x04, 0x8f, 0x61, 0xb0, 0x49, 0x05, 0x2d, 0x72, 0x5d, 0x6f, 0xc7, 0xbd,
	0xf2, 0x1f, 0xc1, 0x49, 0x4e, 0x63, 0xfe, 0xba, 0xb0, 0xf0, 0xe6, 0xd0, 0x12, 0xe3, 0xa3, 0x25,
	0xc6, 0xb1, 0x25, 0xe8, 0x59, 0x12, 0xf4, 0x26, 0x09, 0x7a, 0x97, 0x04, 0x1d, 0x24, 0x41, 0x9f,
	0x92, 0xa0, 0x6f, 0x49, 0x8c, 0xa3, 0x24, 0xe8, 0xf5, 0x8b, 0x18, 0xd9, 0x40, 0x3f, 0xdc, 0xed,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0x73, 0xe5, 0xf3, 0x01, 0x02, 0x00, 0x00,
}
