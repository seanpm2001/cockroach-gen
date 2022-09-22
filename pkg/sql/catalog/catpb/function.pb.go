// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/catalog/catpb/function.proto

package catpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Function_Volatility int32

const (
	Function_UNKNOWN_VOLATILITY Function_Volatility = 0
	Function_VOLATILE           Function_Volatility = 1
	Function_IMMUTABLE          Function_Volatility = 2
	Function_STABLE             Function_Volatility = 3
)

var Function_Volatility_name = map[int32]string{
	0: "UNKNOWN_VOLATILITY",
	1: "VOLATILE",
	2: "IMMUTABLE",
	3: "STABLE",
}

var Function_Volatility_value = map[string]int32{
	"UNKNOWN_VOLATILITY": 0,
	"VOLATILE":           1,
	"IMMUTABLE":          2,
	"STABLE":             3,
}

func (x Function_Volatility) Enum() *Function_Volatility {
	p := new(Function_Volatility)
	*p = x
	return p
}

func (x Function_Volatility) String() string {
	return proto.EnumName(Function_Volatility_name, int32(x))
}

func (x *Function_Volatility) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Function_Volatility_value, data, "Function_Volatility")
	if err != nil {
		return err
	}
	*x = Function_Volatility(value)
	return nil
}

func (Function_Volatility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0, 0}
}

type Function_NullInputBehavior int32

const (
	Function_UNKNOWN_NULL_INPUT_BEHAVIOR Function_NullInputBehavior = 0
	Function_CALLED_ON_NULL_INPUT        Function_NullInputBehavior = 1
	Function_RETURNS_NULL_ON_NULL_INPUT  Function_NullInputBehavior = 2
	Function_STRICT                      Function_NullInputBehavior = 3
)

var Function_NullInputBehavior_name = map[int32]string{
	0: "UNKNOWN_NULL_INPUT_BEHAVIOR",
	1: "CALLED_ON_NULL_INPUT",
	2: "RETURNS_NULL_ON_NULL_INPUT",
	3: "STRICT",
}

var Function_NullInputBehavior_value = map[string]int32{
	"UNKNOWN_NULL_INPUT_BEHAVIOR": 0,
	"CALLED_ON_NULL_INPUT":        1,
	"RETURNS_NULL_ON_NULL_INPUT":  2,
	"STRICT":                      3,
}

func (x Function_NullInputBehavior) Enum() *Function_NullInputBehavior {
	p := new(Function_NullInputBehavior)
	*p = x
	return p
}

func (x Function_NullInputBehavior) String() string {
	return proto.EnumName(Function_NullInputBehavior_name, int32(x))
}

func (x *Function_NullInputBehavior) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Function_NullInputBehavior_value, data, "Function_NullInputBehavior")
	if err != nil {
		return err
	}
	*x = Function_NullInputBehavior(value)
	return nil
}

func (Function_NullInputBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0, 1}
}

type Function_Language int32

const (
	Function_UNKNOWN_LANGUAGE Function_Language = 0
	Function_SQL              Function_Language = 1
)

var Function_Language_name = map[int32]string{
	0: "UNKNOWN_LANGUAGE",
	1: "SQL",
}

var Function_Language_value = map[string]int32{
	"UNKNOWN_LANGUAGE": 0,
	"SQL":              1,
}

func (x Function_Language) Enum() *Function_Language {
	p := new(Function_Language)
	*p = x
	return p
}

func (x Function_Language) String() string {
	return proto.EnumName(Function_Language_name, int32(x))
}

func (x *Function_Language) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Function_Language_value, data, "Function_Language")
	if err != nil {
		return err
	}
	*x = Function_Language(value)
	return nil
}

func (Function_Language) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0, 2}
}

type Function_Arg_Class int32

const (
	Function_Arg_UNKNOWN_ARG_CLASS Function_Arg_Class = 0
	Function_Arg_IN                Function_Arg_Class = 1
	Function_Arg_OUT               Function_Arg_Class = 2
	Function_Arg_IN_OUT            Function_Arg_Class = 3
	Function_Arg_VARIADIC          Function_Arg_Class = 4
)

var Function_Arg_Class_name = map[int32]string{
	0: "UNKNOWN_ARG_CLASS",
	1: "IN",
	2: "OUT",
	3: "IN_OUT",
	4: "VARIADIC",
}

var Function_Arg_Class_value = map[string]int32{
	"UNKNOWN_ARG_CLASS": 0,
	"IN":                1,
	"OUT":               2,
	"IN_OUT":            3,
	"VARIADIC":          4,
}

func (x Function_Arg_Class) Enum() *Function_Arg_Class {
	p := new(Function_Arg_Class)
	*p = x
	return p
}

func (x Function_Arg_Class) String() string {
	return proto.EnumName(Function_Arg_Class_name, int32(x))
}

func (x *Function_Arg_Class) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Function_Arg_Class_value, data, "Function_Arg_Class")
	if err != nil {
		return err
	}
	*x = Function_Arg_Class(value)
	return nil
}

func (Function_Arg_Class) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0, 0, 0}
}

// Function contains a few enum types of function properties
type Function struct {
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0}
}
func (m *Function) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(m, src)
}
func (m *Function) XXX_Size() int {
	return m.Size()
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

type Function_Arg struct {
}

func (m *Function_Arg) Reset()         { *m = Function_Arg{} }
func (m *Function_Arg) String() string { return proto.CompactTextString(m) }
func (*Function_Arg) ProtoMessage()    {}
func (*Function_Arg) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0, 0}
}
func (m *Function_Arg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Function_Arg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Function_Arg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function_Arg.Merge(m, src)
}
func (m *Function_Arg) XXX_Size() int {
	return m.Size()
}
func (m *Function_Arg) XXX_DiscardUnknown() {
	xxx_messageInfo_Function_Arg.DiscardUnknown(m)
}

var xxx_messageInfo_Function_Arg proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cockroach.sql.catalog.catpb.Function_Volatility", Function_Volatility_name, Function_Volatility_value)
	proto.RegisterEnum("cockroach.sql.catalog.catpb.Function_NullInputBehavior", Function_NullInputBehavior_name, Function_NullInputBehavior_value)
	proto.RegisterEnum("cockroach.sql.catalog.catpb.Function_Language", Function_Language_name, Function_Language_value)
	proto.RegisterEnum("cockroach.sql.catalog.catpb.Function_Arg_Class", Function_Arg_Class_name, Function_Arg_Class_value)
	proto.RegisterType((*Function)(nil), "cockroach.sql.catalog.catpb.Function")
	proto.RegisterType((*Function_Arg)(nil), "cockroach.sql.catalog.catpb.Function.Arg")
}

func init() { proto.RegisterFile("sql/catalog/catpb/function.proto", fileDescriptor_e00fc847fc4944a7) }

var fileDescriptor_e00fc847fc4944a7 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x86, 0xed, 0x86, 0x8d, 0xf2, 0x09, 0x24, 0xcf, 0x2a, 0x08, 0x75, 0x92, 0x41, 0x39, 0x41,
	0x9c, 0xb4, 0xd7, 0xe0, 0x66, 0xa1, 0x58, 0xb8, 0xce, 0xc8, 0x4f, 0x11, 0x9c, 0x44, 0x26, 0x1a,
	0x59, 0x84, 0x15, 0x77, 0xf9, 0x41, 0x82, 0xab, 0xe0, 0xb2, 0x76, 0x38, 0x71, 0xb4, 0x43, 0x48,
	0x6f, 0x04, 0x25, 0x69, 0x25, 0x38, 0xf2, 0x6b, 0xfb, 0xf9, 0x9e, 0x57, 0xb2, 0xe1, 0x65, 0x7d,
	0x63, 0x96, 0x99, 0x6e, 0xb4, 0xb1, 0x79, 0xbf, 0xee, 0x3e, 0x2f, 0xbf, 0xb4, 0x65, 0xd6, 0x14,
	0xb6, 0x5c, 0xec, 0x2a, 0xdb, 0x58, 0x7a, 0x9e, 0xd9, 0xec, 0x6b, 0x65, 0x75, 0x76, 0xbd, 0xa8,
	0x6f, 0xcc, 0xe2, 0xc0, 0x2e, 0x06, 0x76, 0x3e, 0xcb, 0x6d, 0x6e, 0x07, 0x6e, 0xd9, 0xa7, 0x71,
	0xc4, 0xfd, 0x35, 0x81, 0xe9, 0x9b, 0x83, 0x65, 0x7e, 0x09, 0x0e, 0xaf, 0x72, 0x57, 0xc0, 0x89,
	0x67, 0x74, 0x5d, 0xd3, 0xa7, 0x70, 0x96, 0xa8, 0x77, 0x2a, 0xf8, 0xa0, 0x52, 0x1e, 0xae, 0x53,
	0x4f, 0xf2, 0x28, 0x22, 0x88, 0x9e, 0xc2, 0x44, 0x28, 0x82, 0xe9, 0x43, 0x70, 0x82, 0x24, 0x26,
	0x13, 0x0a, 0x70, 0x2a, 0x54, 0xda, 0x67, 0x87, 0x3e, 0x86, 0xe9, 0x96, 0x87, 0x82, 0x5f, 0x08,
	0x8f, 0x3c, 0x70, 0x37, 0x00, 0x5b, 0x6b, 0x74, 0x53, 0x98, 0xa2, 0xf9, 0x4e, 0x9f, 0x01, 0x3d,
	0xfa, 0xb6, 0x81, 0xe4, 0xb1, 0x90, 0x22, 0xfe, 0x48, 0xd0, 0x30, 0x33, 0xee, 0x7d, 0x82, 0xe9,
	0x13, 0x78, 0x24, 0x36, 0x9b, 0x24, 0xe6, 0x2b, 0xe9, 0x8f, 0xf2, 0x68, 0xcc, 0x8e, 0xfb, 0x03,
	0xce, 0x54, 0x6b, 0x8c, 0x28, 0x77, 0x6d, 0xb3, 0xba, 0xba, 0xd6, 0xdf, 0x0a, 0x5b, 0xd1, 0x17,
	0x70, 0x7e, 0xb4, 0xaa, 0x44, 0xca, 0x54, 0xa8, 0xcb, 0x24, 0x4e, 0x57, 0xfe, 0x5b, 0xbe, 0x15,
	0x41, 0x48, 0x10, 0x7d, 0x0e, 0x33, 0x8f, 0x4b, 0xe9, 0x5f, 0xa4, 0xc1, 0xbf, 0x08, 0xc1, 0x94,
	0xc1, 0x3c, 0xf4, 0xe3, 0x24, 0x54, 0xd1, 0x78, 0xfe, 0xff, 0xfd, 0xa1, 0x3b, 0x14, 0x5e, 0x4c,
	0x1c, 0xf7, 0x35, 0x4c, 0xa5, 0x2e, 0xf3, 0x56, 0xe7, 0x57, 0x74, 0x06, 0xe4, 0x58, 0x29, 0xb9,
	0x5a, 0x27, 0x7c, 0xed, 0x13, 0xd4, 0xbf, 0x47, 0xf4, 0x5e, 0x12, 0xbc, 0x7a, 0x75, 0xfb, 0x87,
	0xa1, 0xdb, 0x8e, 0xe1, 0xbb, 0x8e, 0xe1, 0xfb, 0x8e, 0xe1, 0xdf, 0x1d, 0xc3, 0x3f, 0xf7, 0x0c,
	0xdd, 0xed, 0x19, 0xba, 0xdf, 0x33, 0xf4, 0xe9, 0x64, 0xf8, 0x93, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x26, 0x00, 0x88, 0xd3, 0x01, 0x00, 0x00,
}

func (m *Function) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Function) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Function) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Function_Arg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Function_Arg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Function_Arg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintFunction(dAtA []byte, offset int, v uint64) int {
	offset -= sovFunction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Function) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Function_Arg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovFunction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFunction(x uint64) (n int) {
	return sovFunction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Function) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Function: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Function: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFunction
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
func (m *Function_Arg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Arg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Arg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFunction
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
func skipFunction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFunction
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
					return 0, ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFunction
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
			if length < 0 {
				return 0, ErrInvalidLengthFunction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFunction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFunction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFunction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFunction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFunction = fmt.Errorf("proto: unexpected end of group")
)
