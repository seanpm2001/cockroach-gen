// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv/kvserver/concurrency/poison/policy.proto

package poison

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

// Policy determines how a request will react to encountering a poisoned
// latch. A poisoned latch is a latch for which the holder is unable to make
// progress. That is, waiters of this latch should not expect to be able to
// acquire this latch "for some time"; in practice this is the case of an
// unavailable Replica.
//
// The name is inspired by Rust's mutexes, which undergo poisoning[^1] when a
// thread panics while holding the mutex.
//
// [^1]: https://doc.rust-lang.org/std/sync/struct.Mutex.html#poisoning
type Policy int32

const (
	// Policy_Wait instructs a request to return an error upon encountering
	// a poisoned latch.
	Policy_Wait Policy = 0
	// Policy_Error instructs a request to return an error upon encountering
	// a poisoned latch.
	Policy_Error Policy = 1
)

var Policy_name = map[int32]string{
	0: "Wait",
	1: "Error",
}

var Policy_value = map[string]int32{
	"Wait":  0,
	"Error": 1,
}

func (x Policy) String() string {
	return proto.EnumName(Policy_name, int32(x))
}

func (Policy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_43c896794a9a5973, []int{0}
}

func init() {
	proto.RegisterEnum("cockroach.kv.kvserver.concurrency.poison.Policy", Policy_name, Policy_value)
}

func init() {
	proto.RegisterFile("kv/kvserver/concurrency/poison/policy.proto", fileDescriptor_43c896794a9a5973)
}

var fileDescriptor_43c896794a9a5973 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xce, 0x2e, 0xd3, 0xcf,
	0x2e, 0x2b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x4f, 0xce, 0xcf, 0x4b, 0x2e, 0x2d, 0x2a, 0x4a,
	0xcd, 0x4b, 0xae, 0xd4, 0x2f, 0xc8, 0xcf, 0x2c, 0xce, 0xcf, 0xd3, 0x2f, 0xc8, 0xcf, 0xc9, 0x4c,
	0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x48, 0xce, 0x4f, 0xce, 0x2e, 0xca, 0x4f,
	0x4c, 0xce, 0xd0, 0xcb, 0x2e, 0xd3, 0x83, 0x69, 0xd3, 0x43, 0xd2, 0xa6, 0x07, 0xd1, 0xa6, 0x25,
	0xcb, 0xc5, 0x16, 0x00, 0xd6, 0x29, 0xc4, 0xc1, 0xc5, 0x12, 0x9e, 0x98, 0x59, 0x22, 0xc0, 0x20,
	0xc4, 0xc9, 0xc5, 0xea, 0x5a, 0x54, 0x94, 0x5f, 0x24, 0xc0, 0xe8, 0xa4, 0x71, 0xe2, 0xa1, 0x1c,
	0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0xde, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xc5,
	0x06, 0x31, 0x28, 0x89, 0x0d, 0x6c, 0xb3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x84, 0xfd, 0x8c,
	0x5d, 0xa8, 0x00, 0x00, 0x00,
}
