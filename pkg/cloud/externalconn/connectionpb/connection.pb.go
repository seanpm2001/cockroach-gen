// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cloud/externalconn/connectionpb/connection.proto

package connectionpb

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

type ConnectionProvider int32

const (
	ConnectionProvider_Unknown ConnectionProvider = 0
	// External Storage providers.
	ConnectionProvider_nodelocal     ConnectionProvider = 1
	ConnectionProvider_s3            ConnectionProvider = 4
	ConnectionProvider_userfile      ConnectionProvider = 5
	ConnectionProvider_gs            ConnectionProvider = 6
	ConnectionProvider_azure_storage ConnectionProvider = 7
	// KMS providers.
	ConnectionProvider_gcp_kms ConnectionProvider = 2
	ConnectionProvider_aws_kms ConnectionProvider = 8
	// Sink providers.
	ConnectionProvider_kafka        ConnectionProvider = 3
	ConnectionProvider_http         ConnectionProvider = 9
	ConnectionProvider_https        ConnectionProvider = 10
	ConnectionProvider_sql          ConnectionProvider = 11
	ConnectionProvider_webhookhttp  ConnectionProvider = 12
	ConnectionProvider_webhookhttps ConnectionProvider = 13
	ConnectionProvider_gcpubsub     ConnectionProvider = 14
)

var ConnectionProvider_name = map[int32]string{
	0:  "Unknown",
	1:  "nodelocal",
	4:  "s3",
	5:  "userfile",
	6:  "gs",
	7:  "azure_storage",
	2:  "gcp_kms",
	8:  "aws_kms",
	3:  "kafka",
	9:  "http",
	10: "https",
	11: "sql",
	12: "webhookhttp",
	13: "webhookhttps",
	14: "gcpubsub",
}

var ConnectionProvider_value = map[string]int32{
	"Unknown":       0,
	"nodelocal":     1,
	"s3":            4,
	"userfile":      5,
	"gs":            6,
	"azure_storage": 7,
	"gcp_kms":       2,
	"aws_kms":       8,
	"kafka":         3,
	"http":          9,
	"https":         10,
	"sql":           11,
	"webhookhttp":   12,
	"webhookhttps":  13,
	"gcpubsub":      14,
}

func (x ConnectionProvider) String() string {
	return proto.EnumName(ConnectionProvider_name, int32(x))
}

func (ConnectionProvider) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f2ab8ae20df1e701, []int{0}
}

// ConnectionType is the type of the External Connection object.
type ConnectionType int32

const (
	TypeUnspecified ConnectionType = 0
	TypeStorage     ConnectionType = 1
	TypeKMS         ConnectionType = 2
)

var ConnectionType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "STORAGE",
	2: "KMS",
}

var ConnectionType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"STORAGE":     1,
	"KMS":         2,
}

func (x ConnectionType) String() string {
	return proto.EnumName(ConnectionType_name, int32(x))
}

func (ConnectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f2ab8ae20df1e701, []int{1}
}

// SimpleURI encapsulates the information that represents an External Connection
// object that only relies on a URI to connect.
type SimpleURI struct {
	URI string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (m *SimpleURI) Reset()         { *m = SimpleURI{} }
func (m *SimpleURI) String() string { return proto.CompactTextString(m) }
func (*SimpleURI) ProtoMessage()    {}
func (*SimpleURI) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2ab8ae20df1e701, []int{0}
}
func (m *SimpleURI) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimpleURI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SimpleURI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleURI.Merge(m, src)
}
func (m *SimpleURI) XXX_Size() int {
	return m.Size()
}
func (m *SimpleURI) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleURI.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleURI proto.InternalMessageInfo

// ConnectionsDetails is the byte representation of the resource represented by
// an External Connection object.
type ConnectionDetails struct {
	Provider ConnectionProvider `protobuf:"varint,1,opt,name=provider,proto3,enum=cockroach.cloud.externalconn.connectionpb.ConnectionProvider" json:"provider,omitempty"`
	// Types that are valid to be assigned to Details:
	//
	//	*ConnectionDetails_SimpleURI
	Details isConnectionDetails_Details `protobuf_oneof:"details"`
}

func (m *ConnectionDetails) Reset()         { *m = ConnectionDetails{} }
func (m *ConnectionDetails) String() string { return proto.CompactTextString(m) }
func (*ConnectionDetails) ProtoMessage()    {}
func (*ConnectionDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2ab8ae20df1e701, []int{1}
}
func (m *ConnectionDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectionDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ConnectionDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionDetails.Merge(m, src)
}
func (m *ConnectionDetails) XXX_Size() int {
	return m.Size()
}
func (m *ConnectionDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionDetails proto.InternalMessageInfo

type isConnectionDetails_Details interface {
	isConnectionDetails_Details()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ConnectionDetails_SimpleURI struct {
	SimpleURI *SimpleURI `protobuf:"bytes,2,opt,name=simple_uri,json=simpleUri,proto3,oneof" json:"simple_uri,omitempty"`
}

func (*ConnectionDetails_SimpleURI) isConnectionDetails_Details() {}

func (m *ConnectionDetails) GetDetails() isConnectionDetails_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *ConnectionDetails) GetSimpleURI() *SimpleURI {
	if x, ok := m.GetDetails().(*ConnectionDetails_SimpleURI); ok {
		return x.SimpleURI
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConnectionDetails) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConnectionDetails_SimpleURI)(nil),
	}
}

func init() {
	proto.RegisterEnum("cockroach.cloud.externalconn.connectionpb.ConnectionProvider", ConnectionProvider_name, ConnectionProvider_value)
	proto.RegisterEnum("cockroach.cloud.externalconn.connectionpb.ConnectionType", ConnectionType_name, ConnectionType_value)
	proto.RegisterType((*SimpleURI)(nil), "cockroach.cloud.externalconn.connectionpb.SimpleURI")
	proto.RegisterType((*ConnectionDetails)(nil), "cockroach.cloud.externalconn.connectionpb.ConnectionDetails")
}

func init() {
	proto.RegisterFile("cloud/externalconn/connectionpb/connection.proto", fileDescriptor_f2ab8ae20df1e701)
}

var fileDescriptor_f2ab8ae20df1e701 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x49, 0x1a, 0xc7, 0xaf, 0xf3, 0xe7, 0x7a, 0x74, 0x80, 0x08, 0xb9, 0x51, 0x85,
	0x50, 0xe9, 0xe0, 0xa0, 0x96, 0x95, 0x81, 0xb4, 0x01, 0xa2, 0xaa, 0x50, 0xd9, 0xf5, 0x00, 0x4b,
	0xe4, 0xd8, 0x17, 0xc7, 0xb2, 0xeb, 0x33, 0x3e, 0x9b, 0x00, 0x9f, 0x00, 0x65, 0xe2, 0x0b, 0x64,
	0xe2, 0xcb, 0x74, 0xec, 0x82, 0xd4, 0xa9, 0x80, 0xf3, 0x45, 0xd0, 0x39, 0x28, 0x89, 0xc4, 0x42,
	0x17, 0xeb, 0xde, 0xe7, 0xde, 0xe7, 0xa7, 0xe7, 0x91, 0x0f, 0x9e, 0x3a, 0x21, 0xcb, 0xdc, 0x2e,
	0xfd, 0x94, 0xd2, 0x24, 0xb2, 0x43, 0x87, 0x45, 0x51, 0x57, 0x7c, 0xa8, 0x93, 0xfa, 0x2c, 0x8a,
	0x47, 0x1b, 0x83, 0x1e, 0x27, 0x2c, 0x65, 0xe4, 0x89, 0xc3, 0x9c, 0x20, 0x61, 0xb6, 0x33, 0xd1,
	0x0b, 0xaf, 0xbe, 0xe9, 0xd5, 0x37, 0xbd, 0xed, 0x1d, 0x8f, 0x79, 0xac, 0x70, 0x75, 0xc5, 0x69,
	0x09, 0xd8, 0x7b, 0x0c, 0x8a, 0xe9, 0x5f, 0xc6, 0x21, 0xb5, 0x8c, 0x01, 0x79, 0x00, 0xe5, 0x2c,
	0xf1, 0xef, 0xa3, 0x0e, 0xda, 0x57, 0x7a, 0x72, 0x7e, 0xbb, 0x5b, 0xb6, 0x8c, 0x81, 0x21, 0xb4,
	0xbd, 0x9f, 0x08, 0xb6, 0x8f, 0x57, 0xb8, 0x13, 0x9a, 0xda, 0x7e, 0xc8, 0xc9, 0x3b, 0xa8, 0xc5,
	0x09, 0xfb, 0xe8, 0xbb, 0x34, 0x29, 0x5c, 0xcd, 0xc3, 0xe7, 0xfa, 0x7f, 0x27, 0xd2, 0xd7, 0xbc,
	0xf3, 0xbf, 0x10, 0x63, 0x85, 0x23, 0x2e, 0x00, 0x2f, 0x82, 0x0d, 0x45, 0xa4, 0x52, 0x07, 0xed,
	0xab, 0x87, 0xcf, 0xee, 0x00, 0x5f, 0xb5, 0xea, 0x35, 0xf2, 0xdb, 0xdd, 0x75, 0xc9, 0xd7, 0x92,
	0xa1, 0x2c, 0xc1, 0x56, 0xe2, 0xf7, 0x14, 0x90, 0xdd, 0x65, 0x97, 0x83, 0x1f, 0x08, 0xc8, 0xbf,
	0x89, 0x88, 0x0a, 0xb2, 0x15, 0x05, 0x11, 0x9b, 0x46, 0x58, 0x22, 0x0d, 0x50, 0x22, 0xe6, 0xd2,
	0x90, 0x39, 0x76, 0x88, 0x11, 0xa9, 0x42, 0x89, 0x1f, 0xe1, 0x0a, 0xa9, 0x43, 0x2d, 0xe3, 0x34,
	0x19, 0xfb, 0x21, 0xc5, 0x5b, 0x42, 0xf5, 0x38, 0xae, 0x92, 0x6d, 0x68, 0xd8, 0x5f, 0xb2, 0x84,
	0x0e, 0x79, 0xca, 0x12, 0xdb, 0xa3, 0x58, 0x16, 0x30, 0xcf, 0x89, 0x87, 0xc1, 0x25, 0xc7, 0x25,
	0x31, 0xd8, 0x53, 0x5e, 0x0c, 0x35, 0xa2, 0xc0, 0x56, 0x60, 0x8f, 0x03, 0x1b, 0x97, 0x49, 0x0d,
	0x2a, 0x93, 0x34, 0x8d, 0xb1, 0x22, 0x44, 0x71, 0xe2, 0x18, 0x88, 0x0c, 0x65, 0xfe, 0x21, 0xc4,
	0x2a, 0x69, 0x81, 0x3a, 0xa5, 0xa3, 0x09, 0x63, 0x41, 0xb1, 0x54, 0x27, 0x18, 0xea, 0x1b, 0x02,
	0xc7, 0x0d, 0x11, 0xc7, 0x73, 0xe2, 0x6c, 0xc4, 0xb3, 0x11, 0x6e, 0x1e, 0xc4, 0xd0, 0x5c, 0xd7,
	0xba, 0xf8, 0x1c, 0x53, 0xf2, 0x08, 0x54, 0xeb, 0x8d, 0x79, 0xde, 0x3f, 0x1e, 0xbc, 0x1c, 0xf4,
	0x4f, 0xb0, 0xd4, 0xbe, 0x37, 0x9b, 0x77, 0x5a, 0xe2, 0xca, 0x8a, 0x78, 0x4c, 0x1d, 0x7f, 0xec,
	0x53, 0x97, 0x3c, 0x04, 0xd9, 0xbc, 0x78, 0x6b, 0xbc, 0x78, 0xd5, 0xc7, 0xa8, 0xdd, 0x9a, 0xcd,
	0x3b, 0xaa, 0xd8, 0x30, 0x97, 0x5d, 0xc8, 0x0e, 0x94, 0x4f, 0xcf, 0x4c, 0x5c, 0x6a, 0xab, 0xb3,
	0x79, 0x47, 0x16, 0x37, 0xa7, 0x67, 0x66, 0xbb, 0xf2, 0xf5, 0xbb, 0x26, 0xf5, 0xf4, 0xab, 0xdf,
	0x9a, 0x74, 0x95, 0x6b, 0xe8, 0x3a, 0xd7, 0xd0, 0x4d, 0xae, 0xa1, 0x5f, 0xb9, 0x86, 0xbe, 0x2d,
	0x34, 0xe9, 0x7a, 0xa1, 0x49, 0x37, 0x0b, 0x4d, 0x7a, 0x5f, 0xdf, 0xfc, 0x55, 0xa3, 0x6a, 0xf1,
	0x14, 0x8f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x03, 0x28, 0x88, 0x42, 0xff, 0x02, 0x00, 0x00,
}

func (m *SimpleURI) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleURI) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimpleURI) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.URI) > 0 {
		i -= len(m.URI)
		copy(dAtA[i:], m.URI)
		i = encodeVarintConnection(dAtA, i, uint64(len(m.URI)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConnectionDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectionDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectionDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Details != nil {
		{
			size := m.Details.Size()
			i -= size
			if _, err := m.Details.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Provider != 0 {
		i = encodeVarintConnection(dAtA, i, uint64(m.Provider))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ConnectionDetails_SimpleURI) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectionDetails_SimpleURI) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SimpleURI != nil {
		{
			size, err := m.SimpleURI.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConnection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintConnection(dAtA []byte, offset int, v uint64) int {
	offset -= sovConnection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SimpleURI) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.URI)
	if l > 0 {
		n += 1 + l + sovConnection(uint64(l))
	}
	return n
}

func (m *ConnectionDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Provider != 0 {
		n += 1 + sovConnection(uint64(m.Provider))
	}
	if m.Details != nil {
		n += m.Details.Size()
	}
	return n
}

func (m *ConnectionDetails_SimpleURI) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SimpleURI != nil {
		l = m.SimpleURI.Size()
		n += 1 + l + sovConnection(uint64(l))
	}
	return n
}

func sovConnection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConnection(x uint64) (n int) {
	return sovConnection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SimpleURI) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnection
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
			return fmt.Errorf("proto: SimpleURI: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleURI: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConnection
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
func (m *ConnectionDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnection
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
			return fmt.Errorf("proto: ConnectionDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectionDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			m.Provider = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Provider |= ConnectionProvider(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SimpleURI", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SimpleURI{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Details = &ConnectionDetails_SimpleURI{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConnection
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
func skipConnection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConnection
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
					return 0, ErrIntOverflowConnection
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
					return 0, ErrIntOverflowConnection
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
				return 0, ErrInvalidLengthConnection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConnection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConnection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConnection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConnection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConnection = fmt.Errorf("proto: unexpected end of group")
)
