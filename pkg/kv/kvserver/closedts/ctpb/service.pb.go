// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv/kvserver/closedts/ctpb/service.proto

package ctpb

import (
	context "context"
	fmt "fmt"
	github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Update contains information about (the advancement of) closed timestamps for
// ranges with leases on the sender node. Updates are of two types: snapshots
// and incrementals. Snapshots are stand-alone messages, explicitly containing
// state about a bunch of ranges. Incrementals are deltas since the previous
// message (the previous message can be a snapshot or another incremental); they
// contain info about which new ranges are included in the info provided, which
// ranges are removed, and how the closed timestamps for different categories of
// ranges advanced. Ranges communicated by a previous message and not touched by
// an incremental are "implicitly" referenced by the incremental. In order to
// properly handle incrementals, the recipient maintains a "stream's state": the
// group of ranges that can be implicitly referenced by the next message.
type Update struct {
	// node_id identifies the sending node.
	NodeID github_com_cockroachdb_cockroach_pkg_roachpb.NodeID `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.NodeID" json:"node_id,omitempty"`
	// seq_num identifies this update across all updates produced by a node. The
	// sequence is reset when the node restarts, so a recipient can only count on
	// it increasing within a single PushUpdates stream.
	//
	// All messages have sequence numbers, including snapshots. A snapshot can be
	// applied on top of any state (i.e. it can be applied after having skipped
	// messages); its sequence number tells the recipient what incremental message
	// it should expect afterwards.
	SeqNum SeqNum `protobuf:"varint,2,opt,name=seq_num,json=seqNum,proto3,casttype=SeqNum" json:"seq_num,omitempty"`
	// snapshot indicates whether this message is standalone, or whether it's just
	// a delta since the messages with the previous seq_num. A snapshot
	// re-initializes all of the recipient's state. The first message on a stream
	// is always a snapshot. Afterwards, there could be others if the sender is
	// temporarily slowed down or if the stream experience network problems and
	// some incremental messages are dropped  (although generally we expect that
	// to result in a stream failing and a new one being established).
	Snapshot         bool                 `protobuf:"varint,3,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	ClosedTimestamps []Update_GroupUpdate `protobuf:"bytes,4,rep,name=closed_timestamps,json=closedTimestamps,proto3" json:"closed_timestamps"`
	// removed contains the set of ranges that are no longer tracked on this
	// stream. The closed timestamps in this message and future messages no longer
	// apply to these removed ranges.
	//
	// The field will be empty if snapshot is true, as a snapshot message implies
	// that all ranges not present in the snapshot's added_or_updated list are no
	// longer tracked.
	Removed        []github_com_cockroachdb_cockroach_pkg_roachpb.RangeID `protobuf:"varint,5,rep,packed,name=removed,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID" json:"removed,omitempty"`
	AddedOrUpdated []Update_RangeUpdate                                   `protobuf:"bytes,6,rep,name=added_or_updated,json=addedOrUpdated,proto3" json:"added_or_updated"`
}

func (m *Update) Reset()      { *m = Update{} }
func (*Update) ProtoMessage() {}
func (*Update) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0044b737e1a7700, []int{0}
}
func (m *Update) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Update) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Update) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update.Merge(m, src)
}
func (m *Update) XXX_Size() int {
	return m.Size()
}
func (m *Update) XXX_DiscardUnknown() {
	xxx_messageInfo_Update.DiscardUnknown(m)
}

var xxx_messageInfo_Update proto.InternalMessageInfo

// closed_timestamps represents the timestamps that are being closed for each
// group of ranges, with a group being represented by its policy.
//
// The recipient is supposed to forward the closed timestamps of the affected
// ranges to these values. Upon receiving one of these updates, the recipient
// should generally not assume that it hasn't been informed of a higher closed
// timestamp for any range in particular - races between this side-transport
// and the regular Raft transport are possible, as are races between two
// side-transport streams for an outgoing and incoming leaseholder.
type Update_GroupUpdate struct {
	Policy          roachpb.RangeClosedTimestampPolicy `protobuf:"varint,1,opt,name=policy,proto3,enum=cockroach.roachpb.RangeClosedTimestampPolicy" json:"policy,omitempty"`
	ClosedTimestamp hlc.Timestamp                      `protobuf:"bytes,2,opt,name=closed_timestamp,json=closedTimestamp,proto3" json:"closed_timestamp"`
}

func (m *Update_GroupUpdate) Reset()         { *m = Update_GroupUpdate{} }
func (m *Update_GroupUpdate) String() string { return proto.CompactTextString(m) }
func (*Update_GroupUpdate) ProtoMessage()    {}
func (*Update_GroupUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0044b737e1a7700, []int{0, 0}
}
func (m *Update_GroupUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Update_GroupUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Update_GroupUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_GroupUpdate.Merge(m, src)
}
func (m *Update_GroupUpdate) XXX_Size() int {
	return m.Size()
}
func (m *Update_GroupUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_Update_GroupUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_Update_GroupUpdate proto.InternalMessageInfo

// added_or_updated contains the set of ranges that are either being added to
// the tracked ranges set with a given (lai, policy) or updated within the
// tracked range set with a new (lai, policy). All future updates on the
// stream are applicable to these ranges until they are removed, either
// explicitly by being included in a future removed set or implicitly by not
// being included in the added_or_updated field of a future snapshot.
type Update_RangeUpdate struct {
	RangeID github_com_cockroachdb_cockroach_pkg_roachpb.RangeID `protobuf:"varint,1,opt,name=range_id,json=rangeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID" json:"range_id,omitempty"`
	LAI     LAI                                                  `protobuf:"varint,2,opt,name=lai,proto3,casttype=LAI" json:"lai,omitempty"`
	Policy  roachpb.RangeClosedTimestampPolicy                   `protobuf:"varint,3,opt,name=policy,proto3,enum=cockroach.roachpb.RangeClosedTimestampPolicy" json:"policy,omitempty"`
}

func (m *Update_RangeUpdate) Reset()         { *m = Update_RangeUpdate{} }
func (m *Update_RangeUpdate) String() string { return proto.CompactTextString(m) }
func (*Update_RangeUpdate) ProtoMessage()    {}
func (*Update_RangeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0044b737e1a7700, []int{0, 1}
}
func (m *Update_RangeUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Update_RangeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Update_RangeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_RangeUpdate.Merge(m, src)
}
func (m *Update_RangeUpdate) XXX_Size() int {
	return m.Size()
}
func (m *Update_RangeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_Update_RangeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_Update_RangeUpdate proto.InternalMessageInfo

type Response struct {
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0044b737e1a7700, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Update)(nil), "cockroach.kv.kvserver.ctupdate.Update")
	proto.RegisterType((*Update_GroupUpdate)(nil), "cockroach.kv.kvserver.ctupdate.Update.GroupUpdate")
	proto.RegisterType((*Update_RangeUpdate)(nil), "cockroach.kv.kvserver.ctupdate.Update.RangeUpdate")
	proto.RegisterType((*Response)(nil), "cockroach.kv.kvserver.ctupdate.Response")
}

func init() {
	proto.RegisterFile("kv/kvserver/closedts/ctpb/service.proto", fileDescriptor_b0044b737e1a7700)
}

var fileDescriptor_b0044b737e1a7700 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xf5, 0xfe, 0x9c, 0x3a, 0xd1, 0x46, 0xbf, 0x52, 0x56, 0x1c, 0xac, 0x08, 0xec, 0xa8, 0x48,
	0xc5, 0x17, 0xbc, 0x28, 0xe5, 0x80, 0xb8, 0x35, 0x05, 0xa1, 0x4a, 0xa8, 0x54, 0x6e, 0xb9, 0xf4,
	0x40, 0xb4, 0xf1, 0xae, 0x1c, 0x2b, 0xb6, 0x77, 0xeb, 0x5d, 0x5b, 0xe2, 0x5b, 0x70, 0xe4, 0xc8,
	0x81, 0x0f, 0x53, 0x71, 0xea, 0x09, 0xf5, 0x64, 0x81, 0xfb, 0x2d, 0x72, 0x42, 0xfe, 0x93, 0x90,
	0xf6, 0x00, 0x54, 0x3d, 0xed, 0xec, 0x4e, 0xe6, 0xcd, 0x7b, 0x6f, 0x26, 0x86, 0x4f, 0xe6, 0x39,
	0x9e, 0xe7, 0x92, 0xa5, 0x39, 0x4b, 0xb1, 0x1f, 0x71, 0xc9, 0xa8, 0x92, 0xd8, 0x57, 0x62, 0x8a,
	0xab, 0xc7, 0xd0, 0x67, 0xae, 0x48, 0xb9, 0xe2, 0xc8, 0xf2, 0xb9, 0x3f, 0x4f, 0x39, 0xf1, 0x67,
	0xee, 0x3c, 0x77, 0x97, 0x25, 0xae, 0xaf, 0x32, 0x41, 0x89, 0x62, 0x03, 0x54, 0xe7, 0xc4, 0x14,
	0x53, 0xa2, 0x48, 0x53, 0x33, 0x30, 0x33, 0x15, 0x46, 0x78, 0x16, 0xf9, 0x58, 0x85, 0x31, 0x93,
	0x8a, 0xc4, 0xa2, 0xcd, 0x3c, 0x08, 0x78, 0xc0, 0xeb, 0x10, 0x57, 0x51, 0xfb, 0xfa, 0x30, 0xe0,
	0x3c, 0x88, 0x18, 0x26, 0x22, 0xc4, 0x24, 0x49, 0xb8, 0x22, 0x2a, 0xe4, 0x89, 0x6c, 0xb2, 0xdb,
	0xdf, 0x0c, 0x68, 0xbc, 0xaf, 0x9b, 0xa1, 0x53, 0xd8, 0x4d, 0x38, 0x65, 0x93, 0x90, 0x9a, 0x60,
	0x08, 0x9c, 0x8d, 0xf1, 0x5e, 0x59, 0xd8, 0xc6, 0x21, 0xa7, 0xec, 0xe0, 0xd5, 0xa2, 0xb0, 0x77,
	0x83, 0x50, 0xcd, 0xb2, 0xa9, 0xeb, 0xf3, 0x18, 0xaf, 0x68, 0xd3, 0xe9, 0xef, 0x18, 0x8b, 0x79,
	0x80, 0x5b, 0xc2, 0x6e, 0x53, 0xe6, 0x19, 0x15, 0xe2, 0x01, 0x45, 0x8f, 0x61, 0x57, 0xb2, 0xb3,
	0x49, 0x92, 0xc5, 0xe6, 0x7f, 0x43, 0xe0, 0xe8, 0x63, 0xb8, 0x28, 0x6c, 0xe3, 0x98, 0x9d, 0x1d,
	0x66, 0xb1, 0x67, 0xc8, 0xfa, 0x44, 0x03, 0xd8, 0x93, 0x09, 0x11, 0x72, 0xc6, 0x95, 0xa9, 0x0f,
	0x81, 0xd3, 0xf3, 0x56, 0x77, 0xc4, 0xe0, 0xfd, 0xc6, 0xc8, 0xc9, 0x4a, 0xb5, 0x34, 0x3b, 0x43,
	0xdd, 0xe9, 0x8f, 0x46, 0xee, 0x9f, 0x5d, 0x74, 0x1b, 0x7d, 0xee, 0x9b, 0x94, 0x67, 0xa2, 0x89,
	0xc7, 0x9d, 0xf3, 0xc2, 0xd6, 0xbc, 0xad, 0x06, 0xf2, 0x64, 0x85, 0x88, 0x3c, 0xd8, 0x4d, 0x59,
	0xcc, 0x73, 0x46, 0xcd, 0x8d, 0xa1, 0xee, 0x6c, 0x8c, 0x5f, 0x2c, 0x0a, 0xfb, 0xf9, 0xad, 0x94,
	0x7b, 0x24, 0x09, 0x2a, 0xe9, 0x4b, 0x20, 0x34, 0x85, 0x5b, 0x84, 0x52, 0x46, 0x27, 0x3c, 0x9d,
	0x34, 0x8c, 0xa8, 0x69, 0xdc, 0x8a, 0x79, 0x0d, 0x79, 0x8d, 0xf9, 0x66, 0x8d, 0xf8, 0x2e, 0x6d,
	0x1e, 0xe9, 0xe0, 0x2b, 0x80, 0xfd, 0x35, 0x7d, 0xe8, 0x35, 0x34, 0x04, 0x8f, 0x42, 0xff, 0x63,
	0x3d, 0xca, 0xcd, 0xd1, 0xd3, 0xb5, 0x4e, 0xd7, 0x88, 0xee, 0x5f, 0x77, 0xe0, 0xa8, 0x2e, 0xf2,
	0xda, 0x62, 0x74, 0x08, 0xb7, 0x6e, 0xba, 0x5e, 0xcf, 0xaf, 0x3f, 0x7a, 0xb4, 0x06, 0x58, 0x2d,
	0xa4, 0x3b, 0x8b, 0x7c, 0x77, 0x05, 0xd3, 0xb2, 0xbc, 0x77, 0xc3, 0xdf, 0xc1, 0x77, 0x00, 0xfb,
	0x6b, 0x62, 0xd0, 0x07, 0xd8, 0x4b, 0xab, 0xeb, 0x72, 0xe7, 0x3a, 0xe3, 0xfd, 0xb2, 0xb0, 0xbb,
	0xad, 0x85, 0x77, 0xb0, 0xbe, 0x0e, 0x28, 0x1a, 0x42, 0x3d, 0x22, 0x61, 0xbb, 0x72, 0x9b, 0x65,
	0x61, 0xeb, 0x6f, 0xf7, 0x0e, 0x16, 0xcd, 0xe1, 0x55, 0xa9, 0x35, 0xa3, 0xf4, 0x3b, 0x18, 0xf5,
	0xb2, 0xf3, 0xf9, 0x8b, 0xad, 0x6d, 0x43, 0xd8, 0xf3, 0x98, 0x14, 0x3c, 0x91, 0x6c, 0x94, 0xc3,
	0xff, 0x8f, 0x43, 0xca, 0x4e, 0x52, 0x92, 0x48, 0xc1, 0xd3, 0x6a, 0x83, 0xfb, 0x47, 0x99, 0x9c,
	0x35, 0xca, 0x25, 0xda, 0xf9, 0xb7, 0xd9, 0x0f, 0x9c, 0xbf, 0xfd, 0x6e, 0xd9, 0x71, 0x5b, 0x73,
	0xc0, 0x33, 0x30, 0xde, 0x39, 0xff, 0x69, 0x69, 0xe7, 0xa5, 0x05, 0x2e, 0x4a, 0x0b, 0x5c, 0x96,
	0x16, 0xf8, 0x51, 0x5a, 0xe0, 0xd3, 0x95, 0xa5, 0x5d, 0x5c, 0x59, 0xda, 0xe5, 0x95, 0xa5, 0x9d,
	0x76, 0xaa, 0x0f, 0xd1, 0xd4, 0xa8, 0xff, 0xff, 0xbb, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbe,
	0xe6, 0x5c, 0x73, 0xac, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SideTransportClient is the client API for SideTransport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SideTransportClient interface {
	PushUpdates(ctx context.Context, opts ...grpc.CallOption) (SideTransport_PushUpdatesClient, error)
}

type sideTransportClient struct {
	cc *grpc.ClientConn
}

func NewSideTransportClient(cc *grpc.ClientConn) SideTransportClient {
	return &sideTransportClient{cc}
}

func (c *sideTransportClient) PushUpdates(ctx context.Context, opts ...grpc.CallOption) (SideTransport_PushUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SideTransport_serviceDesc.Streams[0], "/cockroach.kv.kvserver.ctupdate.SideTransport/PushUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &sideTransportPushUpdatesClient{stream}
	return x, nil
}

type SideTransport_PushUpdatesClient interface {
	Send(*Update) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type sideTransportPushUpdatesClient struct {
	grpc.ClientStream
}

func (x *sideTransportPushUpdatesClient) Send(m *Update) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sideTransportPushUpdatesClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SideTransportServer is the server API for SideTransport service.
type SideTransportServer interface {
	PushUpdates(SideTransport_PushUpdatesServer) error
}

// UnimplementedSideTransportServer can be embedded to have forward compatible implementations.
type UnimplementedSideTransportServer struct {
}

func (*UnimplementedSideTransportServer) PushUpdates(srv SideTransport_PushUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method PushUpdates not implemented")
}

func RegisterSideTransportServer(s *grpc.Server, srv SideTransportServer) {
	s.RegisterService(&_SideTransport_serviceDesc, srv)
}

func _SideTransport_PushUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SideTransportServer).PushUpdates(&sideTransportPushUpdatesServer{stream})
}

type SideTransport_PushUpdatesServer interface {
	Send(*Response) error
	Recv() (*Update, error)
	grpc.ServerStream
}

type sideTransportPushUpdatesServer struct {
	grpc.ServerStream
}

func (x *sideTransportPushUpdatesServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sideTransportPushUpdatesServer) Recv() (*Update, error) {
	m := new(Update)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SideTransport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.kv.kvserver.ctupdate.SideTransport",
	HandlerType: (*SideTransportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PushUpdates",
			Handler:       _SideTransport_PushUpdates_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "kv/kvserver/closedts/ctpb/service.proto",
}

func (m *Update) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Update) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Update) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AddedOrUpdated) > 0 {
		for iNdEx := len(m.AddedOrUpdated) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AddedOrUpdated[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Removed) > 0 {
		dAtA2 := make([]byte, len(m.Removed)*10)
		var j1 int
		for _, num1 := range m.Removed {
			num := uint64(num1)
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
		i = encodeVarintService(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClosedTimestamps) > 0 {
		for iNdEx := len(m.ClosedTimestamps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClosedTimestamps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Snapshot {
		i--
		if m.Snapshot {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.SeqNum != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.SeqNum))
		i--
		dAtA[i] = 0x10
	}
	if m.NodeID != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.NodeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Update_GroupUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Update_GroupUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Update_GroupUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ClosedTimestamp.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintService(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Policy != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Policy))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Update_RangeUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Update_RangeUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Update_RangeUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Policy != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Policy))
		i--
		dAtA[i] = 0x18
	}
	if m.LAI != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.LAI))
		i--
		dAtA[i] = 0x10
	}
	if m.RangeID != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.RangeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Update) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NodeID != 0 {
		n += 1 + sovService(uint64(m.NodeID))
	}
	if m.SeqNum != 0 {
		n += 1 + sovService(uint64(m.SeqNum))
	}
	if m.Snapshot {
		n += 2
	}
	if len(m.ClosedTimestamps) > 0 {
		for _, e := range m.ClosedTimestamps {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	if len(m.Removed) > 0 {
		l = 0
		for _, e := range m.Removed {
			l += sovService(uint64(e))
		}
		n += 1 + sovService(uint64(l)) + l
	}
	if len(m.AddedOrUpdated) > 0 {
		for _, e := range m.AddedOrUpdated {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	return n
}

func (m *Update_GroupUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Policy != 0 {
		n += 1 + sovService(uint64(m.Policy))
	}
	l = m.ClosedTimestamp.Size()
	n += 1 + l + sovService(uint64(l))
	return n
}

func (m *Update_RangeUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RangeID != 0 {
		n += 1 + sovService(uint64(m.RangeID))
	}
	if m.LAI != 0 {
		n += 1 + sovService(uint64(m.LAI))
	}
	if m.Policy != 0 {
		n += 1 + sovService(uint64(m.Policy))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Update) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: Update: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Update: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeqNum", wireType)
			}
			m.SeqNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SeqNum |= SeqNum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snapshot", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
			m.Snapshot = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosedTimestamps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClosedTimestamps = append(m.ClosedTimestamps, Update_GroupUpdate{})
			if err := m.ClosedTimestamps[len(m.ClosedTimestamps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType == 0 {
				var v github_com_cockroachdb_cockroach_pkg_roachpb.RangeID
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowService
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Removed = append(m.Removed, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowService
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
					return ErrInvalidLengthService
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthService
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
				if elementCount != 0 && len(m.Removed) == 0 {
					m.Removed = make([]github_com_cockroachdb_cockroach_pkg_roachpb.RangeID, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v github_com_cockroachdb_cockroach_pkg_roachpb.RangeID
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowService
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Removed = append(m.Removed, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Removed", wireType)
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedOrUpdated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddedOrUpdated = append(m.AddedOrUpdated, Update_RangeUpdate{})
			if err := m.AddedOrUpdated[len(m.AddedOrUpdated)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *Update_GroupUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GroupUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			m.Policy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Policy |= roachpb.RangeClosedTimestampPolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosedTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClosedTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *Update_RangeUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: RangeUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeID", wireType)
			}
			m.RangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeID |= github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LAI", wireType)
			}
			m.LAI = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LAI |= LAI(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			m.Policy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Policy |= roachpb.RangeClosedTimestampPolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
