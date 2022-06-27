// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/log/eventpb/events.proto

package eventpb

import (
	fmt "fmt"
	github_com_cockroachdb_redact "github.com/cockroachdb/redact"
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

// CommonEventDetails contains the fields common to all events.
type CommonEventDetails struct {
	// The timestamp of the event. Expressed as nanoseconds since
	// the Unix epoch.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:",omitempty"`
	// The type of the event.
	EventType string `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:",omitempty" redact:"nonsensitive"`
}

func (m *CommonEventDetails) Reset()         { *m = CommonEventDetails{} }
func (m *CommonEventDetails) String() string { return proto.CompactTextString(m) }
func (*CommonEventDetails) ProtoMessage()    {}
func (*CommonEventDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_656955fd5b536468, []int{0}
}
func (m *CommonEventDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonEventDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CommonEventDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonEventDetails.Merge(m, src)
}
func (m *CommonEventDetails) XXX_Size() int {
	return m.Size()
}
func (m *CommonEventDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonEventDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CommonEventDetails proto.InternalMessageInfo

// CommonSQLEventDetails contains the fields common to all
// SQL events.
type CommonSQLEventDetails struct {
	// A normalized copy of the SQL statement that triggered the event.
	// The statement string contains a mix of sensitive and non-sensitive details (it is redactable).
	Statement github_com_cockroachdb_redact.RedactableString `protobuf:"bytes,1,opt,name=statement,proto3,customtype=github.com/cockroachdb/redact.RedactableString" json:",omitempty" redact:"mixed"`
	// The statement tag. This is separate from the statement string,
	// since the statement string can contain sensitive information. The
	// tag is guaranteed not to.
	Tag string `protobuf:"bytes,6,opt,name=tag,proto3" json:",omitempty" redact:"nonsensitive"`
	// The user account that triggered the event.
	// The special usernames `root` and `node` are not considered sensitive.
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:",omitempty" redact:"safeif:root|node"`
	// The primary object descriptor affected by the operation. Set to zero for operations
	// that don't affect descriptors.
	DescriptorID uint32 `protobuf:"varint,3,opt,name=descriptor_id,json=descriptorId,proto3" json:",omitempty"`
	// The application name for the session where the event was emitted.
	// This is included in the event to ease filtering of logging output
	// by application.
	// Application names starting with a dollar sign (`$`) are not considered sensitive.
	ApplicationName string `protobuf:"bytes,4,opt,name=application_name,json=applicationName,proto3" json:",omitempty" redact:"safeif:\$.*"`
	// The mapping of SQL placeholders to their values, for prepared statements.
	PlaceholderValues []string `protobuf:"bytes,5,rep,name=placeholder_values,json=placeholderValues,proto3" json:",omitempty"`
}

func (m *CommonSQLEventDetails) Reset()         { *m = CommonSQLEventDetails{} }
func (m *CommonSQLEventDetails) String() string { return proto.CompactTextString(m) }
func (*CommonSQLEventDetails) ProtoMessage()    {}
func (*CommonSQLEventDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_656955fd5b536468, []int{1}
}
func (m *CommonSQLEventDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonSQLEventDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CommonSQLEventDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonSQLEventDetails.Merge(m, src)
}
func (m *CommonSQLEventDetails) XXX_Size() int {
	return m.Size()
}
func (m *CommonSQLEventDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonSQLEventDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CommonSQLEventDetails proto.InternalMessageInfo

// CommonJobEventDetails contains the fields common to all job events.
type CommonJobEventDetails struct {
	// The ID of the job that triggered the event.
	JobID int64 `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:",omitempty"`
	// The type of the job that triggered the event.
	JobType string `protobuf:"bytes,2,opt,name=job_type,json=jobType,proto3" json:",omitempty" redact:"nonsensitive"`
	// A description of the job that triggered the event. Some jobs populate the
	// description with an approximate representation of the SQL statement run to
	// create the job.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:",omitempty"`
	// The user account that triggered the event.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:",omitempty"`
	// The object descriptors affected by the job. Set to zero for operations
	// that don't affect descriptors.
	DescriptorIDs []uint32 `protobuf:"varint,5,rep,packed,name=descriptor_ids,json=descriptorIds,proto3" json:",omitempty"`
	// The status of the job that triggered the event. This allows the job to
	// indicate which phase execution it is in when the event is triggered.
	Status string `protobuf:"bytes,6,opt,name=status,proto3" json:",omitempty" redact:"nonsensitive"`
}

func (m *CommonJobEventDetails) Reset()         { *m = CommonJobEventDetails{} }
func (m *CommonJobEventDetails) String() string { return proto.CompactTextString(m) }
func (*CommonJobEventDetails) ProtoMessage()    {}
func (*CommonJobEventDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_656955fd5b536468, []int{2}
}
func (m *CommonJobEventDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonJobEventDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CommonJobEventDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonJobEventDetails.Merge(m, src)
}
func (m *CommonJobEventDetails) XXX_Size() int {
	return m.Size()
}
func (m *CommonJobEventDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonJobEventDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CommonJobEventDetails proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CommonEventDetails)(nil), "cockroach.util.log.eventpb.CommonEventDetails")
	proto.RegisterType((*CommonSQLEventDetails)(nil), "cockroach.util.log.eventpb.CommonSQLEventDetails")
	proto.RegisterType((*CommonJobEventDetails)(nil), "cockroach.util.log.eventpb.CommonJobEventDetails")
}

func init() { proto.RegisterFile("util/log/eventpb/events.proto", fileDescriptor_656955fd5b536468) }

var fileDescriptor_656955fd5b536468 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6a, 0xdb, 0x4c,
	0x14, 0xc5, 0xad, 0x38, 0x71, 0x3e, 0xcd, 0x17, 0xa7, 0xed, 0x90, 0x50, 0x35, 0x50, 0xc9, 0x88,
	0xd2, 0xba, 0x6d, 0x90, 0x0b, 0x5d, 0x35, 0x50, 0x02, 0x8a, 0x1b, 0x70, 0x28, 0x85, 0x28, 0x25,
	0x8b, 0x52, 0x30, 0x23, 0x69, 0x22, 0x8f, 0x2b, 0xe9, 0x0a, 0xcd, 0xd8, 0xd4, 0xd0, 0x27, 0xe8,
	0xaa, 0x8f, 0xe5, 0x65, 0x96, 0xa1, 0x0b, 0x91, 0xca, 0xbb, 0x2c, 0xfb, 0x04, 0x45, 0x92, 0x1d,
	0xff, 0x51, 0x37, 0xe9, 0x6a, 0x04, 0xfa, 0x1d, 0xe6, 0xdc, 0x7b, 0x0e, 0x83, 0x1e, 0x0f, 0x04,
	0xf3, 0x5b, 0x3e, 0x78, 0x2d, 0x3a, 0xa4, 0xa1, 0x88, 0xec, 0xe2, 0xe4, 0x46, 0x14, 0x83, 0x00,
	0xbc, 0xe7, 0x80, 0xf3, 0x25, 0x06, 0xe2, 0xf4, 0x8c, 0x0c, 0x34, 0x7c, 0xf0, 0x8c, 0x29, 0xb8,
	0xb7, 0xe3, 0x81, 0x07, 0x39, 0xd6, 0xca, 0xbe, 0x0a, 0x85, 0xfe, 0x5d, 0x42, 0xf8, 0x08, 0x82,
	0x00, 0xc2, 0x77, 0x19, 0xd7, 0xa6, 0x82, 0x30, 0x9f, 0xe3, 0x7d, 0x24, 0x0b, 0x16, 0x50, 0x2e,
	0x48, 0x10, 0x29, 0x52, 0x43, 0x6a, 0x56, 0xcd, 0xed, 0x9b, 0x44, 0x43, 0xfb, 0x10, 0x30, 0x41,
	0x83, 0x48, 0x8c, 0xac, 0x39, 0x80, 0x8f, 0x11, 0xca, 0x6f, 0xe9, 0x8a, 0x51, 0x44, 0x95, 0xb5,
	0x86, 0xd4, 0x94, 0xcd, 0x67, 0xcb, 0xf8, 0xef, 0x44, 0xdb, 0x8d, 0xa9, 0x4b, 0x1c, 0x71, 0xa0,
	0x87, 0x10, 0x72, 0x1a, 0x72, 0x26, 0xd8, 0x90, 0xea, 0x96, 0x9c, 0x4b, 0x3f, 0x8e, 0x22, 0xaa,
	0x5f, 0x57, 0xd1, 0x6e, 0x61, 0xe6, 0xec, 0xf4, 0xfd, 0x92, 0x1f, 0x81, 0x64, 0x2e, 0x88, 0xa0,
	0x01, 0x0d, 0x45, 0xee, 0x47, 0x36, 0xcf, 0xc7, 0x89, 0x56, 0xf9, 0x99, 0x68, 0x86, 0xc7, 0x44,
	0x6f, 0x60, 0x1b, 0x0e, 0x04, 0xad, 0xdb, 0xf1, 0x5d, 0xbb, 0x55, 0xdc, 0x66, 0x58, 0xf9, 0x41,
	0x6c, 0x9f, 0x9e, 0x89, 0x98, 0x85, 0x5e, 0xc9, 0xd6, 0xf6, 0xcc, 0x56, 0xc0, 0xbe, 0x52, 0x57,
	0xb7, 0xe6, 0x17, 0xe1, 0x37, 0xa8, 0x2a, 0x88, 0xa7, 0xd4, 0xee, 0x36, 0x50, 0xa6, 0xc1, 0x87,
	0x68, 0x7d, 0xc0, 0x69, 0x3c, 0x5d, 0xc6, 0xcb, 0x92, 0xf6, 0xd1, 0x4c, 0xcb, 0xc9, 0x05, 0x65,
	0x17, 0x07, 0x31, 0x80, 0xf8, 0x16, 0x82, 0x4b, 0x75, 0x2b, 0x17, 0xe2, 0x23, 0x54, 0x77, 0x29,
	0x77, 0x62, 0x16, 0x09, 0x88, 0xbb, 0xcc, 0x55, 0xaa, 0x0d, 0xa9, 0x59, 0x37, 0xd5, 0x34, 0xd1,
	0xb6, 0xda, 0xb7, 0x3f, 0x3a, 0xed, 0x95, 0x54, 0xb6, 0xe6, 0xa2, 0x8e, 0x8b, 0x4f, 0xd1, 0x7d,
	0x12, 0x45, 0x3e, 0x73, 0x88, 0x60, 0x10, 0x76, 0x43, 0x12, 0x50, 0x65, 0x3d, 0x77, 0xf4, 0xb4,
	0xe4, 0x68, 0x67, 0xc5, 0xd1, 0xe7, 0x27, 0xc6, 0x0b, 0xdd, 0xba, 0xb7, 0xa0, 0xff, 0x40, 0x02,
	0x8a, 0xdf, 0x22, 0x1c, 0xf9, 0xc4, 0xa1, 0x3d, 0xf0, 0x5d, 0x1a, 0x77, 0x87, 0xc4, 0x1f, 0x50,
	0xae, 0x6c, 0x34, 0xaa, 0x4d, 0xb9, 0x54, 0x91, 0x07, 0x0b, 0xe4, 0x79, 0x0e, 0xea, 0x37, 0x6b,
	0xb3, 0x88, 0x4f, 0xc0, 0x5e, 0x8a, 0xd8, 0x40, 0xb5, 0x3e, 0xd8, 0xd9, 0xa4, 0x45, 0xdf, 0x1e,
	0xa6, 0x89, 0xb6, 0x71, 0x02, 0x76, 0x69, 0xc4, 0x8d, 0x3e, 0xd8, 0x1d, 0x17, 0x9b, 0xe8, 0xbf,
	0x8c, 0xff, 0x97, 0xca, 0x6d, 0xf6, 0xc1, 0xce, 0x0a, 0x87, 0x5f, 0xa1, 0xff, 0x67, 0xfb, 0x62,
	0x10, 0xe6, 0x2b, 0x2e, 0x4f, 0xb1, 0x88, 0x60, 0x7d, 0x9a, 0xeb, 0xfa, 0x5f, 0xd1, 0x22, 0xba,
	0x63, 0xb4, 0xbd, 0x14, 0x5d, 0xb1, 0x9e, 0xba, 0xa9, 0xa5, 0x89, 0x56, 0x5f, 0xcc, 0x8e, 0xaf,
	0xc8, 0xeb, 0x8b, 0xe1, 0x71, 0x7c, 0x88, 0x6a, 0x59, 0x17, 0x07, 0xfc, 0xae, 0x0d, 0x9c, 0xca,
	0xcc, 0xe7, 0xe3, 0x5f, 0x6a, 0x65, 0x9c, 0xaa, 0xd2, 0x65, 0xaa, 0x4a, 0x57, 0xa9, 0x2a, 0x5d,
	0xa7, 0xaa, 0xf4, 0x63, 0xa2, 0x56, 0x2e, 0x27, 0x6a, 0xe5, 0x6a, 0xa2, 0x56, 0x3e, 0x6d, 0x4e,
	0x5f, 0x07, 0xbb, 0x96, 0x3f, 0x07, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x86, 0x6e, 0x58,
	0x48, 0x61, 0x04, 0x00, 0x00,
}

func (m *CommonEventDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonEventDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonEventDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EventType) > 0 {
		i -= len(m.EventType)
		copy(dAtA[i:], m.EventType)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.EventType)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CommonSQLEventDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonSQLEventDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonSQLEventDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tag) > 0 {
		i -= len(m.Tag)
		copy(dAtA[i:], m.Tag)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Tag)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PlaceholderValues) > 0 {
		for iNdEx := len(m.PlaceholderValues) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PlaceholderValues[iNdEx])
			copy(dAtA[i:], m.PlaceholderValues[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.PlaceholderValues[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ApplicationName) > 0 {
		i -= len(m.ApplicationName)
		copy(dAtA[i:], m.ApplicationName)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ApplicationName)))
		i--
		dAtA[i] = 0x22
	}
	if m.DescriptorID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.DescriptorID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Statement) > 0 {
		i -= len(m.Statement)
		copy(dAtA[i:], m.Statement)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Statement)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommonJobEventDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonJobEventDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonJobEventDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DescriptorIDs) > 0 {
		dAtA2 := make([]byte, len(m.DescriptorIDs)*10)
		var j1 int
		for _, num := range m.DescriptorIDs {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintEvents(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.JobType) > 0 {
		i -= len(m.JobType)
		copy(dAtA[i:], m.JobType)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.JobType)))
		i--
		dAtA[i] = 0x12
	}
	if m.JobID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.JobID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommonEventDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovEvents(uint64(m.Timestamp))
	}
	l = len(m.EventType)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *CommonSQLEventDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Statement)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.DescriptorID != 0 {
		n += 1 + sovEvents(uint64(m.DescriptorID))
	}
	l = len(m.ApplicationName)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.PlaceholderValues) > 0 {
		for _, s := range m.PlaceholderValues {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	l = len(m.Tag)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *CommonJobEventDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.JobID != 0 {
		n += 1 + sovEvents(uint64(m.JobID))
	}
	l = len(m.JobType)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.DescriptorIDs) > 0 {
		l = 0
		for _, e := range m.DescriptorIDs {
			l += sovEvents(uint64(e))
		}
		n += 1 + sovEvents(uint64(l)) + l
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommonEventDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: CommonEventDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonEventDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *CommonSQLEventDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: CommonSQLEventDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonSQLEventDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statement = github_com_cockroachdb_redact.RedactableString(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DescriptorID", wireType)
			}
			m.DescriptorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DescriptorID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlaceholderValues", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlaceholderValues = append(m.PlaceholderValues, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *CommonJobEventDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: CommonJobEventDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonJobEventDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			m.JobID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvents
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.DescriptorIDs = append(m.DescriptorIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvents
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEvents
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEvents
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.DescriptorIDs) == 0 {
					m.DescriptorIDs = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEvents
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.DescriptorIDs = append(m.DescriptorIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field DescriptorIDs", wireType)
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)