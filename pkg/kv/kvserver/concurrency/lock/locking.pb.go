// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv/kvserver/concurrency/lock/locking.proto

package lock

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Strength represents the different locking modes that determine how key-values
// can be accessed by concurrent transactions.
//
// Locking modes apply to locks that are held with a per-key granularity. It is
// up to users of the key-value layer to decide on which keys to acquire locks
// for when imposing structure that can span multiple keys, such as SQL rows
// (see column families and secondary indexes).
//
// Locking modes have differing levels of strength, growing from "weakest" to
// "strongest" in the order that the variants are presented in the enumeration.
// The "stronger" a locking mode, the more protection it provides for the lock
// holder but the more restrictive it is to concurrent transactions attempting
// to access the same keys.
//
// Compatibility Matrix
//
// The following matrix presents the compatibility of locking strengths with one
// another. A cell with an X means that the two strengths are incompatible with
// each other and that they can not both be held on a given key by different
// transactions, concurrently. A cell without an X means that the two strengths
// are compatible with each other and that they can be held on a given key by
// different transactions, concurrently.
//
//  +-----------+-----------+-----------+-----------+-----------+
//  |           |   None    |  Shared   |  Upgrade  | Exclusive |
//  +-----------+-----------+-----------+-----------+-----------+
//  | None      |           |           |           |     X^†   |
//  +-----------+-----------+-----------+-----------+-----------+
//  | Shared    |           |           |     X     |     X     |
//  +-----------+-----------+-----------+-----------+-----------+
//  | Upgrade   |           |     X     |     X     |     X     |
//  +-----------+-----------+-----------+-----------+-----------+
//  | Exclusive |     X^†   |     X     |     X     |     X     |
//  +-----------+-----------+-----------+-----------+-----------+
//
// [†] reads under optimistic concurrency control in CockroachDB only conflict
// with Exclusive locks if the read's timestamp is equal to or greater than the
// lock's timestamp. If the read's timestamp is below the Exclusive lock's
// timestamp then the two are compatible.
type Strength int32

const (
	// None represents the absence of a lock or the intention to acquire locks.
	// It corresponds to the behavior of transactions performing key-value reads
	// under optimistic concurrency control. No locks are acquired on the keys
	// read by these requests when they evaluate. However, the reads do respect
	// Exclusive locks already held by other transactions at timestamps equal to
	// or less than their read timestamp.
	//
	// Optimistic concurrency control (OCC) can improve performance under some
	// workloads because it avoids the need to perform any locking during reads.
	// This can increase the amount of concurrency that the system can permit
	// between ongoing transactions. However, OCC does mandate a read validation
	// phase if/when transactions need to commit at a different timestamp than
	// they performed all reads at. CockroachDB calls this a "read refresh",
	// which is implemented by the txnSpanRefresher. If a read refresh fails due
	// to new key-value writes that invalidate what was previously read,
	// transactions are forced to restart. See the comment on txnSpanRefresher
	// for more.
	None Strength = 0
	// Shared (S) locks are used by read-only operations and allow concurrent
	// transactions to read under pessimistic concurrency control. Shared locks
	// are compatible with each other but are not compatible with Upgrade or
	// Exclusive locks. This means that multiple transactions can hold a Shared
	// lock on the same key at the same time, but no other transaction can
	// modify the key at the same time. A holder of a Shared lock on a key is
	// only permitted to read the key's value while the lock is held.
	//
	// Share locks are currently unused, as all KV reads are currently performed
	// optimistically (see None).
	Shared Strength = 1
	// Upgrade (U) locks are a hybrid of Shared and Exclusive locks which are
	// used to prevent a common form of deadlock. When a transaction intends to
	// modify existing KVs, it is often the case that it reads the KVs first and
	// then attempts to modify them. Under pessimistic concurrency control, this
	// would correspond to first acquiring a Shared lock on the keys and then
	// converting the lock to an Exclusive lock when modifying the keys. If two
	// transactions were to acquire the Shared lock initially and then attempt
	// to update the keys concurrently, both transactions would get stuck
	// waiting for the other to release its Shared lock and a deadlock would
	// occur. To resolve the deadlock, one of the two transactions would need to
	// be aborted.
	//
	// To avoid this potential deadlock problem, an Upgrade lock can be used in
	// place of a Shared lock. Upgrade locks are not compatible with any other
	// form of locking. As with Shared locks, the lock holder of an Upgrade lock
	// on a key is only allowed to read from the key while the lock is held.
	// This resolves the deadlock scenario presented above because only one of
	// the transactions would have been able to acquire an Upgrade lock at a
	// time while reading the initial state of the KVs. This means that the
	// Shared-to-Exclusive lock upgrade would never need to wait on another
	// transaction to release its locks.
	//
	// Under pure pessimistic concurrency control, an Upgrade lock is equivalent
	// to an Exclusive lock. However, unlike with Exclusive locks, reads under
	// optimistic concurrency control do not conflict with Upgrade locks. This
	// is because a transaction can only hold an Upgrade lock on keys that it
	// has not yet modified. This improves concurrency between read and write
	// transactions compared to if the writing transaction had immediately
	// acquired an Exclusive lock.
	//
	// The trade-off here is twofold. First, if the Upgrade lock holder does
	// convert its lock on a key to an Exclusive lock after an optimistic read
	// has observed the state of the key, the transaction that performed the
	// optimistic read may be unable to perform a successful read refresh if it
	// attempts to refresh to a timestamp at or past the timestamp of the lock
	// conversion. Second, the optimistic reads permitted while the Upgrade lock
	// is held will bump the timestamp cache. This may result in the Upgrade
	// lock holder being forced to increase its write timestamp when converting
	// to an Exclusive lock, which in turn may force it to restart if its read
	// refresh fails.
	Upgrade Strength = 2
	// Exclusive (X) locks are used by read-write and read-only operations and
	// provide a transaction with exclusive access to a key. When an Exclusive
	// lock is held by a transaction on a given key, no other transaction can
	// read from or write to that key. The lock holder is free to read from and
	// write to the key as frequently as it would like.
	Exclusive Strength = 3
)

var Strength_name = map[int32]string{
	0: "None",
	1: "Shared",
	2: "Upgrade",
	3: "Exclusive",
}

var Strength_value = map[string]int32{
	"None":      0,
	"Shared":    1,
	"Upgrade":   2,
	"Exclusive": 3,
}

func (x Strength) String() string {
	return proto.EnumName(Strength_name, int32(x))
}

func (Strength) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_828a6b2fccefea6b, []int{0}
}

// Durability represents the different durability properties of a lock acquired
// by a transaction. Durability levels provide varying degrees of survivability,
// often in exchange for the cost of lock acquisition.
type Durability int32

const (
	// Replicated locks are held on at least a quorum of Replicas in a Range.
	// They are slower to acquire and release than Unreplicated locks because
	// updating them requires both cross-node coordination and interaction with
	// durable storage. In exchange, Replicated locks provide a guarantee of
	// survivability across lease transfers, leaseholder crashes, and other
	// forms of failure events. They will remain available as long as their
	// Range remains available and they will never be lost.
	Replicated Durability = 0
	// Unreplicated locks are held only on a single Replica in a Range, which is
	// typically the leaseholder. Unreplicated locks are very fast to acquire
	// and release because they are held in memory or on fast local storage and
	// require no cross-node coordination to update. In exchange, Unreplicated
	// locks provide no guarantee of survivability across lease transfers or
	// leaseholder crashes. They should therefore be thought of as best-effort
	// and should not be relied upon for correctness.
	Unreplicated Durability = 1
)

var Durability_name = map[int32]string{
	0: "Replicated",
	1: "Unreplicated",
}

var Durability_value = map[string]int32{
	"Replicated":   0,
	"Unreplicated": 1,
}

func (x Durability) String() string {
	return proto.EnumName(Durability_name, int32(x))
}

func (Durability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_828a6b2fccefea6b, []int{1}
}

// WaitPolicy specifies the behavior of a request when it encounters conflicting
// locks held by other active transactions. The default behavior is to block
// until the conflicting lock is released, but other policies can make sense in
// special situations.
type WaitPolicy int32

const (
	// Block indicates that if a request encounters a conflicting lock held by
	// another active transaction, it should wait for the conflicting lock to be
	// released before proceeding.
	WaitPolicy_Block WaitPolicy = 0
	// Error indicates that if a request encounters a conflicting lock held by
	// another active transaction, it should raise an error instead of blocking.
	// If the request encounters a conflicting lock that was abandoned by an
	// inactive transaction, which is likely due to a transaction coordinator
	// crash, the lock is removed and no error is raised.
	WaitPolicy_Error WaitPolicy = 1
	// SkipLocked indicates that if a request encounters a conflicting lock held
	// by another transaction while scanning, it should skip over the key that is
	// locked instead of blocking and later acquiring a lock on that key. The
	// locked key will not be included in the scan result.
	WaitPolicy_SkipLocked WaitPolicy = 2
)

var WaitPolicy_name = map[int32]string{
	0: "Block",
	1: "Error",
	2: "SkipLocked",
}

var WaitPolicy_value = map[string]int32{
	"Block":      0,
	"Error":      1,
	"SkipLocked": 2,
}

func (x WaitPolicy) String() string {
	return proto.EnumName(WaitPolicy_name, int32(x))
}

func (WaitPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_828a6b2fccefea6b, []int{2}
}

func init() {
	proto.RegisterEnum("cockroach.kv.kvserver.concurrency.lock.Strength", Strength_name, Strength_value)
	proto.RegisterEnum("cockroach.kv.kvserver.concurrency.lock.Durability", Durability_name, Durability_value)
	proto.RegisterEnum("cockroach.kv.kvserver.concurrency.lock.WaitPolicy", WaitPolicy_name, WaitPolicy_value)
}

func init() {
	proto.RegisterFile("kv/kvserver/concurrency/lock/locking.proto", fileDescriptor_828a6b2fccefea6b)
}

var fileDescriptor_828a6b2fccefea6b = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0xbf, 0x4a, 0x03, 0x41,
	0x10, 0x06, 0xf0, 0xdd, 0x18, 0x63, 0x32, 0xfe, 0x61, 0x59, 0xac, 0x52, 0x6c, 0x99, 0x22, 0xc5,
	0x1d, 0xa8, 0x4f, 0x10, 0x4c, 0x27, 0x22, 0x86, 0x20, 0xd8, 0x5d, 0xf6, 0x96, 0xcb, 0xb2, 0xc7,
	0xce, 0x31, 0xd9, 0x1c, 0xe6, 0x0d, 0x2c, 0x7d, 0x07, 0x5f, 0x26, 0x65, 0xca, 0x94, 0x7a, 0x79,
	0x11, 0xd9, 0x53, 0xd1, 0x66, 0xf8, 0x18, 0xf8, 0x7d, 0xcc, 0xc0, 0xd8, 0xd5, 0xa9, 0xab, 0x57,
	0x86, 0x6a, 0x43, 0xa9, 0x46, 0xaf, 0xd7, 0x44, 0xc6, 0xeb, 0x4d, 0x5a, 0xa2, 0x76, 0xed, 0xb0,
	0xbe, 0x48, 0x2a, 0xc2, 0x80, 0x72, 0xa4, 0x51, 0x3b, 0xc2, 0x4c, 0x2f, 0x13, 0x57, 0x27, 0xbf,
	0x2a, 0xf9, 0xa7, 0x92, 0x08, 0x86, 0x97, 0x05, 0x16, 0xd8, 0x92, 0x34, 0xa6, 0x6f, 0x3d, 0x9e,
	0x40, 0x7f, 0x16, 0xc8, 0xf8, 0x22, 0x2c, 0x65, 0x1f, 0xba, 0xf7, 0xe8, 0x8d, 0x60, 0x12, 0xa0,
	0x37, 0x5b, 0x66, 0x64, 0x72, 0xc1, 0xe5, 0x29, 0x9c, 0xcc, 0xab, 0x82, 0xb2, 0xdc, 0x88, 0x8e,
	0x3c, 0x87, 0xc1, 0xf4, 0x45, 0x97, 0xeb, 0x95, 0xad, 0x8d, 0x38, 0x1a, 0x76, 0x5f, 0xdf, 0x15,
	0x1b, 0xdf, 0x00, 0xdc, 0xae, 0x29, 0x5b, 0xd8, 0xd2, 0x86, 0x8d, 0xbc, 0x00, 0x78, 0x34, 0x55,
	0x69, 0x75, 0x16, 0x4c, 0x2e, 0x98, 0x14, 0x70, 0x36, 0xf7, 0xf4, 0xb7, 0xe1, 0x3f, 0xea, 0x0a,
	0xe0, 0x29, 0xb3, 0xe1, 0x01, 0x4b, 0xab, 0x37, 0x72, 0x00, 0xc7, 0x93, 0x78, 0xa6, 0x60, 0x31,
	0x4e, 0x89, 0x90, 0x04, 0x8f, 0x5d, 0x33, 0x67, 0xab, 0x3b, 0xd4, 0xce, 0xe4, 0xa2, 0x33, 0x19,
	0x6d, 0x3f, 0x15, 0xdb, 0x36, 0x8a, 0xef, 0x1a, 0xc5, 0xf7, 0x8d, 0xe2, 0x1f, 0x8d, 0xe2, 0x6f,
	0x07, 0xc5, 0x76, 0x07, 0xc5, 0xf6, 0x07, 0xc5, 0x9e, 0xbb, 0xb1, 0x64, 0xd1, 0x6b, 0x9f, 0xbb,
	0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xff, 0xaf, 0xc5, 0xb1, 0x48, 0x01, 0x00, 0x00,
}
