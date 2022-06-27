// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: roachpb/internal_raft.proto

package roachpb

import (
	fmt "fmt"
	hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"
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

// RaftTruncatedState contains metadata about the truncated portion of the raft log.
// Raft requires access to the term of the last truncated log entry even after the
// rest of the entry has been discarded.
type RaftTruncatedState struct {
	// The highest index that has been removed from the log.
	Index uint64 `protobuf:"varint,1,opt,name=index" json:"index"`
	// The term corresponding to 'index'.
	Term uint64 `protobuf:"varint,2,opt,name=term" json:"term"`
}

func (m *RaftTruncatedState) Reset()         { *m = RaftTruncatedState{} }
func (m *RaftTruncatedState) String() string { return proto.CompactTextString(m) }
func (*RaftTruncatedState) ProtoMessage()    {}
func (*RaftTruncatedState) Descriptor() ([]byte, []int) {
	return fileDescriptor_91c3207013e7a8b4, []int{0}
}
func (m *RaftTruncatedState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaftTruncatedState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RaftTruncatedState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftTruncatedState.Merge(m, src)
}
func (m *RaftTruncatedState) XXX_Size() int {
	return m.Size()
}
func (m *RaftTruncatedState) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftTruncatedState.DiscardUnknown(m)
}

var xxx_messageInfo_RaftTruncatedState proto.InternalMessageInfo

// RangeTombstone contains information about a replica that has been deleted.
type RangeTombstone struct {
	NextReplicaID ReplicaID `protobuf:"varint,1,opt,name=next_replica_id,json=nextReplicaId,casttype=ReplicaID" json:"next_replica_id"`
}

func (m *RangeTombstone) Reset()         { *m = RangeTombstone{} }
func (m *RangeTombstone) String() string { return proto.CompactTextString(m) }
func (*RangeTombstone) ProtoMessage()    {}
func (*RangeTombstone) Descriptor() ([]byte, []int) {
	return fileDescriptor_91c3207013e7a8b4, []int{1}
}
func (m *RangeTombstone) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RangeTombstone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RangeTombstone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangeTombstone.Merge(m, src)
}
func (m *RangeTombstone) XXX_Size() int {
	return m.Size()
}
func (m *RangeTombstone) XXX_DiscardUnknown() {
	xxx_messageInfo_RangeTombstone.DiscardUnknown(m)
}

var xxx_messageInfo_RangeTombstone proto.InternalMessageInfo

// RaftSnapshotData is the payload of a raftpb.Snapshot. It contains a raw copy of
// all of the range's data and metadata, including the raft log, abort span, etc.
type RaftSnapshotData struct {
	KV []RaftSnapshotData_KeyValue `protobuf:"bytes,2,rep,name=KV" json:"KV"`
	// These are really raftpb.Entry, but we model them as raw bytes to avoid
	// roundtripping through memory.
	LogEntries [][]byte `protobuf:"bytes,3,rep,name=log_entries,json=logEntries" json:"log_entries,omitempty"`
}

func (m *RaftSnapshotData) Reset()         { *m = RaftSnapshotData{} }
func (m *RaftSnapshotData) String() string { return proto.CompactTextString(m) }
func (*RaftSnapshotData) ProtoMessage()    {}
func (*RaftSnapshotData) Descriptor() ([]byte, []int) {
	return fileDescriptor_91c3207013e7a8b4, []int{2}
}
func (m *RaftSnapshotData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaftSnapshotData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RaftSnapshotData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftSnapshotData.Merge(m, src)
}
func (m *RaftSnapshotData) XXX_Size() int {
	return m.Size()
}
func (m *RaftSnapshotData) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftSnapshotData.DiscardUnknown(m)
}

var xxx_messageInfo_RaftSnapshotData proto.InternalMessageInfo

type RaftSnapshotData_KeyValue struct {
	Key       []byte        `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value     []byte        `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Timestamp hlc.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp"`
}

func (m *RaftSnapshotData_KeyValue) Reset()         { *m = RaftSnapshotData_KeyValue{} }
func (m *RaftSnapshotData_KeyValue) String() string { return proto.CompactTextString(m) }
func (*RaftSnapshotData_KeyValue) ProtoMessage()    {}
func (*RaftSnapshotData_KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_91c3207013e7a8b4, []int{2, 0}
}
func (m *RaftSnapshotData_KeyValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaftSnapshotData_KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RaftSnapshotData_KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftSnapshotData_KeyValue.Merge(m, src)
}
func (m *RaftSnapshotData_KeyValue) XXX_Size() int {
	return m.Size()
}
func (m *RaftSnapshotData_KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftSnapshotData_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_RaftSnapshotData_KeyValue proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RaftTruncatedState)(nil), "cockroach.roachpb.RaftTruncatedState")
	proto.RegisterType((*RangeTombstone)(nil), "cockroach.roachpb.RangeTombstone")
	proto.RegisterType((*RaftSnapshotData)(nil), "cockroach.roachpb.RaftSnapshotData")
	proto.RegisterType((*RaftSnapshotData_KeyValue)(nil), "cockroach.roachpb.RaftSnapshotData.KeyValue")
}

func init() { proto.RegisterFile("roachpb/internal_raft.proto", fileDescriptor_91c3207013e7a8b4) }

var fileDescriptor_91c3207013e7a8b4 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xd6, 0x4a, 0x36, 0x75, 0xd6, 0x4e, 0xeb, 0x2e, 0x39, 0x08, 0x97, 0xae, 0x8c, 0x4f, 0x2e,
	0x14, 0x19, 0x72, 0xec, 0xad, 0x26, 0x85, 0x36, 0x86, 0x1e, 0x36, 0xc6, 0x87, 0x52, 0x30, 0x1b,
	0x79, 0x22, 0x8b, 0xac, 0x77, 0xc5, 0x7a, 0x54, 0x9c, 0xb7, 0xe8, 0x23, 0xe4, 0x31, 0xfa, 0x08,
	0x3e, 0xe6, 0x98, 0x93, 0x69, 0xe5, 0x4b, 0x9f, 0x21, 0xa7, 0xa2, 0x9f, 0xa4, 0xa6, 0xbd, 0xcd,
	0x7c, 0xdf, 0x37, 0xc3, 0x37, 0xdf, 0xd0, 0x57, 0xd6, 0xc8, 0x68, 0x99, 0x5e, 0x8e, 0x12, 0x8d,
	0x60, 0xb5, 0x54, 0x73, 0x2b, 0xaf, 0x30, 0x4c, 0xad, 0x41, 0xc3, 0x5e, 0x46, 0x26, 0xba, 0x2e,
	0x05, 0x61, 0x2d, 0xeb, 0xf9, 0x19, 0x26, 0x6a, 0xb4, 0x54, 0xd1, 0x08, 0x93, 0x15, 0xac, 0x51,
	0xae, 0xd2, 0x4a, 0xdc, 0x3b, 0x89, 0x4d, 0x6c, 0xca, 0x72, 0x54, 0x54, 0x15, 0x3a, 0x98, 0x52,
	0x26, 0xe4, 0x15, 0x4e, 0x6d, 0xa6, 0x23, 0x89, 0xb0, 0xb8, 0x40, 0x89, 0xc0, 0x7a, 0xb4, 0x99,
	0xe8, 0x05, 0x6c, 0x7c, 0xd2, 0x27, 0xc3, 0xc6, 0xb8, 0xb1, 0xdd, 0x05, 0x8e, 0xa8, 0x20, 0xe6,
	0xd3, 0x06, 0x82, 0x5d, 0xf9, 0xee, 0x01, 0x55, 0x22, 0xef, 0x5a, 0x3f, 0x6e, 0x03, 0xf2, 0xfb,
	0x36, 0x20, 0x83, 0xaf, 0xf4, 0xb9, 0x90, 0x3a, 0x86, 0xa9, 0x59, 0x5d, 0xae, 0xd1, 0x68, 0x60,
	0xe7, 0xf4, 0x85, 0x86, 0x0d, 0xce, 0x2d, 0xa4, 0x2a, 0x89, 0xe4, 0x3c, 0x59, 0x94, 0xbb, 0x9b,
	0xe3, 0x41, 0xb1, 0x20, 0xdf, 0x05, 0xc7, 0x9f, 0x61, 0x83, 0xa2, 0x62, 0x3f, 0x9d, 0x3d, 0xec,
	0x82, 0xa3, 0xa7, 0x46, 0x1c, 0xeb, 0x03, 0x6e, 0x31, 0x78, 0x20, 0xb4, 0x5b, 0x98, 0xbe, 0xd0,
	0x32, 0x5d, 0x2f, 0x0d, 0x9e, 0x49, 0x94, 0xec, 0x23, 0x75, 0x27, 0x33, 0xdf, 0xed, 0x7b, 0xc3,
	0xf6, 0xe9, 0xdb, 0xf0, 0xbf, 0x60, 0xc2, 0x7f, 0x07, 0xc2, 0x09, 0xdc, 0xcc, 0xa4, 0xca, 0x60,
	0x4c, 0x6b, 0x07, 0xee, 0x64, 0x26, 0xdc, 0xc9, 0x8c, 0x05, 0xb4, 0xad, 0x4c, 0x3c, 0x07, 0x8d,
	0x36, 0x81, 0xb5, 0xef, 0xf5, 0xbd, 0x61, 0x47, 0x50, 0x65, 0xe2, 0x0f, 0x15, 0xd2, 0xcb, 0x68,
	0xeb, 0x71, 0x98, 0x75, 0xa9, 0x77, 0x0d, 0x37, 0xe5, 0x2d, 0x1d, 0x51, 0x94, 0xec, 0x84, 0x36,
	0xbf, 0x15, 0x54, 0x19, 0x50, 0x47, 0x54, 0x0d, 0x7b, 0x4f, 0x8f, 0x9e, 0x1e, 0xe2, 0x7b, 0x7d,
	0x32, 0x6c, 0x9f, 0xbe, 0x3e, 0x70, 0x59, 0x7c, 0x2d, 0x5c, 0xaa, 0x28, 0x9c, 0x3e, 0x8a, 0xea,
	0x64, 0xff, 0x4e, 0x9d, 0x37, 0x5a, 0xa4, 0xeb, 0x8e, 0xdf, 0x6c, 0x7f, 0x71, 0x67, 0x9b, 0x73,
	0x72, 0x97, 0x73, 0x72, 0x9f, 0x73, 0xf2, 0x33, 0xe7, 0xe4, 0xfb, 0x9e, 0x3b, 0x77, 0x7b, 0xee,
	0xdc, 0xef, 0xb9, 0xf3, 0xe5, 0x59, 0x7d, 0xf2, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xde,
	0x36, 0x90, 0x3c, 0x02, 0x00, 0x00,
}

func (this *RaftTruncatedState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RaftTruncatedState)
	if !ok {
		that2, ok := that.(RaftTruncatedState)
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
	if this.Index != that1.Index {
		return false
	}
	if this.Term != that1.Term {
		return false
	}
	return true
}
func (m *RaftTruncatedState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftTruncatedState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaftTruncatedState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintInternalRaft(dAtA, i, uint64(m.Term))
	i--
	dAtA[i] = 0x10
	i = encodeVarintInternalRaft(dAtA, i, uint64(m.Index))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *RangeTombstone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RangeTombstone) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RangeTombstone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintInternalRaft(dAtA, i, uint64(m.NextReplicaID))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *RaftSnapshotData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftSnapshotData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaftSnapshotData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LogEntries) > 0 {
		for iNdEx := len(m.LogEntries) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.LogEntries[iNdEx])
			copy(dAtA[i:], m.LogEntries[iNdEx])
			i = encodeVarintInternalRaft(dAtA, i, uint64(len(m.LogEntries[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.KV) > 0 {
		for iNdEx := len(m.KV) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KV[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintInternalRaft(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *RaftSnapshotData_KeyValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftSnapshotData_KeyValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaftSnapshotData_KeyValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintInternalRaft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Value != nil {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintInternalRaft(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if m.Key != nil {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintInternalRaft(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInternalRaft(dAtA []byte, offset int, v uint64) int {
	offset -= sovInternalRaft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedRaftTruncatedState(r randyInternalRaft, easy bool) *RaftTruncatedState {
	this := &RaftTruncatedState{}
	this.Index = uint64(uint64(r.Uint32()))
	this.Term = uint64(uint64(r.Uint32()))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyInternalRaft interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneInternalRaft(r randyInternalRaft) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringInternalRaft(r randyInternalRaft) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneInternalRaft(r)
	}
	return string(tmps)
}
func randUnrecognizedInternalRaft(r randyInternalRaft, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldInternalRaft(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldInternalRaft(dAtA []byte, r randyInternalRaft, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateInternalRaft(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateInternalRaft(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateInternalRaft(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateInternalRaft(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateInternalRaft(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateInternalRaft(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateInternalRaft(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *RaftTruncatedState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovInternalRaft(uint64(m.Index))
	n += 1 + sovInternalRaft(uint64(m.Term))
	return n
}

func (m *RangeTombstone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovInternalRaft(uint64(m.NextReplicaID))
	return n
}

func (m *RaftSnapshotData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.KV) > 0 {
		for _, e := range m.KV {
			l = e.Size()
			n += 1 + l + sovInternalRaft(uint64(l))
		}
	}
	if len(m.LogEntries) > 0 {
		for _, b := range m.LogEntries {
			l = len(b)
			n += 1 + l + sovInternalRaft(uint64(l))
		}
	}
	return n
}

func (m *RaftSnapshotData_KeyValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = len(m.Key)
		n += 1 + l + sovInternalRaft(uint64(l))
	}
	if m.Value != nil {
		l = len(m.Value)
		n += 1 + l + sovInternalRaft(uint64(l))
	}
	l = m.Timestamp.Size()
	n += 1 + l + sovInternalRaft(uint64(l))
	return n
}

func sovInternalRaft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInternalRaft(x uint64) (n int) {
	return sovInternalRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftTruncatedState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: RaftTruncatedState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftTruncatedState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInternalRaft
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
func (m *RangeTombstone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: RangeTombstone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeTombstone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextReplicaID", wireType)
			}
			m.NextReplicaID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextReplicaID |= ReplicaID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInternalRaft
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
func (m *RaftSnapshotData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: RaftSnapshotData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftSnapshotData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KV", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
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
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInternalRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KV = append(m.KV, RaftSnapshotData_KeyValue{})
			if err := m.KV[len(m.KV)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogEntries", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInternalRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogEntries = append(m.LogEntries, make([]byte, postIndex-iNdEx))
			copy(m.LogEntries[len(m.LogEntries)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInternalRaft
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
func (m *RaftSnapshotData_KeyValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: KeyValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInternalRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInternalRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
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
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInternalRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInternalRaft
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
func skipInternalRaft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInternalRaft
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
					return 0, ErrIntOverflowInternalRaft
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
					return 0, ErrIntOverflowInternalRaft
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
				return 0, ErrInvalidLengthInternalRaft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInternalRaft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInternalRaft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInternalRaft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInternalRaft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInternalRaft = fmt.Errorf("proto: unexpected end of group")
)