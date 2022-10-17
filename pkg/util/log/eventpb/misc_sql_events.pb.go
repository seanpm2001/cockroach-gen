// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/log/eventpb/misc_sql_events.proto

package eventpb

import (
	fmt "fmt"
	logpb "github.com/cockroachdb/cockroach/pkg/util/log/logpb"
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

// SetClusterSetting is recorded when a cluster setting is changed.
type SetClusterSetting struct {
	logpb.CommonEventDetails `protobuf:"bytes,1,opt,name=common,proto3,embedded=common" json:""`
	CommonSQLEventDetails    `protobuf:"bytes,2,opt,name=sql,proto3,embedded=sql" json:""`
	// The name of the affected cluster setting.
	SettingName string `protobuf:"bytes,3,opt,name=setting_name,json=settingName,proto3" json:",omitempty" redact:"nonsensitive"`
	// The new value of the cluster setting.
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:",omitempty"`
}

func (m *SetClusterSetting) Reset()         { *m = SetClusterSetting{} }
func (m *SetClusterSetting) String() string { return proto.CompactTextString(m) }
func (*SetClusterSetting) ProtoMessage()    {}
func (*SetClusterSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef37ab124d19540a, []int{0}
}
func (m *SetClusterSetting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetClusterSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SetClusterSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetClusterSetting.Merge(m, src)
}
func (m *SetClusterSetting) XXX_Size() int {
	return m.Size()
}
func (m *SetClusterSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_SetClusterSetting.DiscardUnknown(m)
}

var xxx_messageInfo_SetClusterSetting proto.InternalMessageInfo

// SetTenantClusterSetting is recorded when a cluster setting override
// is changed, either for another tenant or for all tenants.
type SetTenantClusterSetting struct {
	logpb.CommonEventDetails `protobuf:"bytes,1,opt,name=common,proto3,embedded=common" json:""`
	CommonSQLEventDetails    `protobuf:"bytes,2,opt,name=sql,proto3,embedded=sql" json:""`
	// The name of the affected cluster setting.
	SettingName string `protobuf:"bytes,3,opt,name=setting_name,json=settingName,proto3" json:",omitempty" redact:"nonsensitive"`
	// The new value of the cluster setting.
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:",omitempty"`
	// The target Tenant ID. Empty if targeting all tenants.
	TenantId uint64 `protobuf:"varint,5,opt,name=tenant_id,json=tenantId,proto3" json:",omitempty"`
	// Whether the override applies to all tenants.
	AllTenants bool `protobuf:"varint,6,opt,name=all_tenants,json=allTenants,proto3" json:",omitempty"`
}

func (m *SetTenantClusterSetting) Reset()         { *m = SetTenantClusterSetting{} }
func (m *SetTenantClusterSetting) String() string { return proto.CompactTextString(m) }
func (*SetTenantClusterSetting) ProtoMessage()    {}
func (*SetTenantClusterSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef37ab124d19540a, []int{1}
}
func (m *SetTenantClusterSetting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetTenantClusterSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SetTenantClusterSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTenantClusterSetting.Merge(m, src)
}
func (m *SetTenantClusterSetting) XXX_Size() int {
	return m.Size()
}
func (m *SetTenantClusterSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTenantClusterSetting.DiscardUnknown(m)
}

var xxx_messageInfo_SetTenantClusterSetting proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SetClusterSetting)(nil), "cockroach.util.log.eventpb.SetClusterSetting")
	proto.RegisterType((*SetTenantClusterSetting)(nil), "cockroach.util.log.eventpb.SetTenantClusterSetting")
}

func init() {
	proto.RegisterFile("util/log/eventpb/misc_sql_events.proto", fileDescriptor_ef37ab124d19540a)
}

var fileDescriptor_ef37ab124d19540a = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x92, 0xcd, 0xaa, 0xd3, 0x40,
	0x1c, 0xc5, 0x33, 0xfd, 0xb2, 0x9d, 0x16, 0xc1, 0xa0, 0x18, 0x02, 0x4e, 0x42, 0x90, 0x1a, 0x51,
	0x12, 0xd4, 0x9d, 0xcb, 0x56, 0x17, 0x4a, 0x11, 0xda, 0xb8, 0x72, 0x13, 0xa6, 0xe9, 0x10, 0x07,
	0x27, 0x33, 0x6d, 0x66, 0x5a, 0xf0, 0x15, 0x5c, 0xf9, 0x14, 0x3e, 0x4b, 0x97, 0x5d, 0x76, 0x15,
	0x34, 0xdd, 0x75, 0xe9, 0x13, 0x5c, 0xf2, 0x71, 0x2f, 0xf7, 0xd2, 0xfb, 0xf1, 0x02, 0x77, 0x37,
	0x33, 0xe7, 0x9c, 0x1f, 0xff, 0x39, 0xfc, 0xe1, 0x70, 0xad, 0x28, 0xf3, 0x99, 0x88, 0x7d, 0xb2,
	0x21, 0x5c, 0x2d, 0xe7, 0x7e, 0x42, 0x65, 0x14, 0xca, 0x15, 0x0b, 0xcb, 0x07, 0xe9, 0x2d, 0x53,
	0xa1, 0x84, 0x6e, 0x46, 0x22, 0xfa, 0x91, 0x0a, 0x1c, 0x7d, 0xf7, 0x8a, 0x84, 0xc7, 0x44, 0xec,
	0xd5, 0x09, 0xf3, 0x71, 0x2c, 0x62, 0x51, 0xda, 0xfc, 0xe2, 0x54, 0x25, 0xcc, 0x67, 0x27, 0xe4,
	0xcb, 0x40, 0xd3, 0xbc, 0x90, 0x99, 0x88, 0xcf, 0xc5, 0x4a, 0x73, 0xfe, 0x34, 0xe0, 0xa3, 0x80,
	0xa8, 0x31, 0x5b, 0x4b, 0x45, 0xd2, 0x80, 0x28, 0x45, 0x79, 0xac, 0x4f, 0x60, 0x27, 0x12, 0x49,
	0x22, 0xb8, 0x01, 0x6c, 0xe0, 0xf6, 0xdf, 0x0e, 0xbd, 0x6b, 0x66, 0x1a, 0x97, 0x8e, 0x8f, 0x05,
	0xec, 0x03, 0x51, 0x98, 0x32, 0x39, 0x1a, 0x6c, 0x33, 0x4b, 0xdb, 0x65, 0x16, 0x38, 0x66, 0x96,
	0x36, 0xab, 0x19, 0xfa, 0x14, 0x36, 0xe5, 0x8a, 0x19, 0x8d, 0x12, 0xf5, 0xc6, 0xbb, 0xf9, 0x7b,
	0x35, 0x32, 0x98, 0x4e, 0x6e, 0xa1, 0x16, 0x2c, 0xfd, 0x33, 0x1c, 0xc8, 0x6a, 0xd6, 0x90, 0xe3,
	0x84, 0x18, 0x4d, 0x1b, 0xb8, 0xbd, 0xd1, 0x8b, 0x63, 0x66, 0xc1, 0xd7, 0x22, 0xa1, 0x8a, 0x24,
	0x4b, 0xf5, 0xf3, 0x7f, 0x66, 0x3d, 0x49, 0xc9, 0x02, 0x47, 0xea, 0xbd, 0xc3, 0x05, 0x97, 0x84,
	0x4b, 0xaa, 0xe8, 0x86, 0x38, 0xb3, 0x7e, 0x1d, 0xfe, 0x82, 0x13, 0xa2, 0x3f, 0x87, 0xed, 0x0d,
	0x66, 0x6b, 0x62, 0xb4, 0x4a, 0xc8, 0xc3, 0xab, 0x90, 0x59, 0x25, 0x3a, 0xbf, 0x9a, 0xf0, 0x69,
	0x40, 0xd4, 0x57, 0xc2, 0x31, 0xbf, 0xaf, 0xeb, 0xae, 0xba, 0xf4, 0x57, 0xb0, 0xa7, 0xca, 0xaa,
	0x42, 0xba, 0x30, 0xda, 0x36, 0x70, 0x5b, 0x27, 0xce, 0x6e, 0x65, 0xf8, 0xb4, 0xd0, 0x7d, 0xd8,
	0xc7, 0x8c, 0x85, 0xd5, 0x5d, 0x1a, 0x1d, 0x1b, 0xb8, 0xdd, 0x13, 0x3b, 0xc4, 0x8c, 0x55, 0xed,
	0xcb, 0xd1, 0xcb, 0xed, 0x3f, 0xa4, 0x6d, 0x73, 0x04, 0x76, 0x39, 0x02, 0xfb, 0x1c, 0x81, 0xbf,
	0x39, 0x02, 0xbf, 0x0f, 0x48, 0xdb, 0x1d, 0x90, 0xb6, 0x3f, 0x20, 0xed, 0xdb, 0x83, 0xba, 0xa3,
	0x79, 0xa7, 0xdc, 0xf3, 0x77, 0x67, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xf2, 0x97, 0xa6, 0x7e,
	0x03, 0x00, 0x00,
}

func (m *SetClusterSetting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetClusterSetting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetClusterSetting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintMiscSqlEvents(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SettingName) > 0 {
		i -= len(m.SettingName)
		copy(dAtA[i:], m.SettingName)
		i = encodeVarintMiscSqlEvents(dAtA, i, uint64(len(m.SettingName)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.CommonSQLEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMiscSqlEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.CommonEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMiscSqlEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *SetTenantClusterSetting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetTenantClusterSetting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetTenantClusterSetting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllTenants {
		i--
		if m.AllTenants {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.TenantId != 0 {
		i = encodeVarintMiscSqlEvents(dAtA, i, uint64(m.TenantId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintMiscSqlEvents(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SettingName) > 0 {
		i -= len(m.SettingName)
		copy(dAtA[i:], m.SettingName)
		i = encodeVarintMiscSqlEvents(dAtA, i, uint64(len(m.SettingName)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.CommonSQLEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMiscSqlEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.CommonEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMiscSqlEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMiscSqlEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovMiscSqlEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SetClusterSetting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommonEventDetails.Size()
	n += 1 + l + sovMiscSqlEvents(uint64(l))
	l = m.CommonSQLEventDetails.Size()
	n += 1 + l + sovMiscSqlEvents(uint64(l))
	l = len(m.SettingName)
	if l > 0 {
		n += 1 + l + sovMiscSqlEvents(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMiscSqlEvents(uint64(l))
	}
	return n
}

func (m *SetTenantClusterSetting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommonEventDetails.Size()
	n += 1 + l + sovMiscSqlEvents(uint64(l))
	l = m.CommonSQLEventDetails.Size()
	n += 1 + l + sovMiscSqlEvents(uint64(l))
	l = len(m.SettingName)
	if l > 0 {
		n += 1 + l + sovMiscSqlEvents(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMiscSqlEvents(uint64(l))
	}
	if m.TenantId != 0 {
		n += 1 + sovMiscSqlEvents(uint64(m.TenantId))
	}
	if m.AllTenants {
		n += 2
	}
	return n
}

func sovMiscSqlEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMiscSqlEvents(x uint64) (n int) {
	return sovMiscSqlEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SetClusterSetting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMiscSqlEvents
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
			return fmt.Errorf("proto: SetClusterSetting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetClusterSetting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiscSqlEvents
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
				return ErrInvalidLengthMiscSqlEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMiscSqlEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonSQLEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiscSqlEvents
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
				return ErrInvalidLengthMiscSqlEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMiscSqlEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonSQLEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettingName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiscSqlEvents
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
				return ErrInvalidLengthMiscSqlEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMiscSqlEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SettingName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiscSqlEvents
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
				return ErrInvalidLengthMiscSqlEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMiscSqlEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMiscSqlEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMiscSqlEvents
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
func (m *SetTenantClusterSetting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMiscSqlEvents
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
			return fmt.Errorf("proto: SetTenantClusterSetting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetTenantClusterSetting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiscSqlEvents
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
				return ErrInvalidLengthMiscSqlEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMiscSqlEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonSQLEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiscSqlEvents
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
				return ErrInvalidLengthMiscSqlEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMiscSqlEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonSQLEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettingName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiscSqlEvents
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
				return ErrInvalidLengthMiscSqlEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMiscSqlEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SettingName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiscSqlEvents
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
				return ErrInvalidLengthMiscSqlEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMiscSqlEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TenantId", wireType)
			}
			m.TenantId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiscSqlEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TenantId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllTenants", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiscSqlEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllTenants = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMiscSqlEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMiscSqlEvents
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
func skipMiscSqlEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMiscSqlEvents
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
					return 0, ErrIntOverflowMiscSqlEvents
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
					return 0, ErrIntOverflowMiscSqlEvents
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
				return 0, ErrInvalidLengthMiscSqlEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMiscSqlEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMiscSqlEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMiscSqlEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMiscSqlEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMiscSqlEvents = fmt.Errorf("proto: unexpected end of group")
)
