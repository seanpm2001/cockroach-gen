// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/catalog/descpb/locking.proto

package descpb

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

// ScanLockingStrength controls the row-level locking mode used by scans.
//
// Typically, SQL scans read sequential keys from the key-value layer without
// acquiring any locks. This means that two scans by different transactions will
// not conflict and cause one of the two transactions to block the other. This
// is usually desirable, as it increases concurrency between readers.
//
// However, there are cases where a SQL scan would like to acquire locks on each
// of the keys that it reads to more carefully control concurrent access to the
// data that it reads. The prototypical example of this is a scan that is used
// to fetch the initial value of a row that its transction intends to later
// update. In this case, it would be beneficial to acquire a lock on the row
// during the initial scan instead of waiting until the mutation to acquire a
// lock. This prevents the row from being modified between the scan and the
// mutation. It also prevents situations that can lead to deadlocks.
//
// Locking modes have differing levels of strength, growing from "weakest" to
// "strongest" in the order that the variants are presented in the enumeration.
// The "stronger" a locking mode, the more protection it provides for the lock
// holder but the more restrictive it is to concurrent transactions attempting
// to access the same keys.
//
// The following matrix presents the compatibility of locking strengths with one
// another.
//
//  +-------------------+---------------+-----------+-------------------+------------+
//  |                   | FOR_KEY_SHARE | FOR_SHARE | FOR_NO_KEY_UPDATE | FOR_UPDATE |
//  +-------------------+---------------+-----------+-------------------+------------+
//  | FOR_KEY_SHARE     |               |           |                   |      X     |
//  +-------------------+---------------+-----------+-------------------+------------+
//  | FOR_SHARE         |               |           |         X         |      X     |
//  +-------------------+---------------+-----------+-------------------+------------+
//  | FOR_NO_KEY_UPDATE |               |     X     |         X         |      X     |
//  +-------------------+---------------+-----------+-------------------+------------+
//  | FOR_UPDATE        |       X       |     X     |         X         |      X     |
//  +-------------------+---------------+-----------+-------------------+------------+
//
// A transaction can hold conflicting locks on the same row, but two different
// transactions can never hold conflicting locks on the same row. Once acquired,
// a lock is held until the end of the transaction.
type ScanLockingStrength int32

const (
	// FOR_NONE represents the default - no row-level locking.
	ScanLockingStrength_FOR_NONE ScanLockingStrength = 0
	// FOR_KEY_SHARE represents the FOR KEY SHARE row-level locking mode.
	//
	// The mode behaves similarly to FOR SHARE, except that the lock is weaker:
	// SELECT FOR UPDATE is blocked, but not SELECT FOR NO KEY UPDATE. A
	// key-shared lock blocks other transactions from performing DELETE or any
	// UPDATE that changes the key values, but not other UPDATE, and neither does
	// it prevent SELECT FOR NO KEY UPDATE, SELECT FOR SHARE, or SELECT FOR KEY
	// SHARE.
	//
	// The locking mode was introduced into Postgres as an alternative to FOR
	// SHARE to improve concurrency between foreign key validation scans, which
	// acquire FOR KEY SHARE locks, and UPDATEs to existing rows, which acquire
	// FOR NO KEY UPDATE locks.
	//
	// NOTE: FOR_KEY_SHARE is currently ignored. No locks are acquired.
	ScanLockingStrength_FOR_KEY_SHARE ScanLockingStrength = 1
	// FOR_SHARE represents the FOR SHARE row-level locking mode.
	//
	// The mode behaves similarly to FOR NO KEY UPDATE, except that it acquires a
	// shared lock rather than exclusive lock on each retrieved row. A shared lock
	// blocks other transactions from performing UPDATE, DELETE, SELECT FOR UPDATE
	// or SELECT FOR NO KEY UPDATE on these rows, but it does not prevent them
	// from performing SELECT FOR SHARE or SELECT FOR KEY SHARE.
	//
	// NOTE: FOR_SHARE is currently ignored. No locks are acquired.
	ScanLockingStrength_FOR_SHARE ScanLockingStrength = 2
	// FOR_NO_KEY_UPDATE represents the FOR NO KEY UPDATE row-level locking mode.
	//
	// The mode behaves similarly to FOR UPDATE, except that the lock acquired is
	// weaker: this lock will not block SELECT FOR KEY SHARE commands that attempt
	// to acquire a lock on the same rows. This lock mode is also acquired by any
	// UPDATE that does not acquire a FOR UPDATE lock.
	//
	// The locking mode was introduced into Postgres as an alternative to FOR
	// UDPATE to improve concurrency between foreign key validation scans, which
	// acquire FOR KEY SHARE locks, and UPDATEs to existing rows, which acquire
	// FOR NO KEY UPDATE locks.
	//
	// NOTE: FOR_NO_KEY_UPDATE is currently promoted to FOR_UPDATE.
	ScanLockingStrength_FOR_NO_KEY_UPDATE ScanLockingStrength = 3
	// FOR_UPDATE represents the FOR UPDATE row-level locking mode.
	//
	// The mode causes the rows retrieved by the scan to be locked as though for
	// update. This prevents them from being locked, modified or deleted by other
	// transactions until the current transaction ends. That is, other
	// transactions that attempt UPDATE, DELETE, SELECT FOR UPDATE, SELECT FOR NO
	// KEY UPDATE, SELECT FOR SHARE or SELECT FOR KEY SHARE of these rows will be
	// blocked until the current transaction ends. Conversely, SELECT FOR UPDATE
	// will wait for a concurrent transaction that has run any of those commands
	// on the same row, and will then lock and return the updated row (or no row,
	// if the row was deleted).
	//
	// NOTE: FOR_UPDATE is currently implemented by acquiring lock.Exclusive locks
	// on each key scanned.
	ScanLockingStrength_FOR_UPDATE ScanLockingStrength = 4
)

var ScanLockingStrength_name = map[int32]string{
	0: "FOR_NONE",
	1: "FOR_KEY_SHARE",
	2: "FOR_SHARE",
	3: "FOR_NO_KEY_UPDATE",
	4: "FOR_UPDATE",
}

var ScanLockingStrength_value = map[string]int32{
	"FOR_NONE":          0,
	"FOR_KEY_SHARE":     1,
	"FOR_SHARE":         2,
	"FOR_NO_KEY_UPDATE": 3,
	"FOR_UPDATE":        4,
}

func (x ScanLockingStrength) Enum() *ScanLockingStrength {
	p := new(ScanLockingStrength)
	*p = x
	return p
}

func (x ScanLockingStrength) String() string {
	return proto.EnumName(ScanLockingStrength_name, int32(x))
}

func (x *ScanLockingStrength) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ScanLockingStrength_value, data, "ScanLockingStrength")
	if err != nil {
		return err
	}
	*x = ScanLockingStrength(value)
	return nil
}

func (ScanLockingStrength) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_484351bd13f4281f, []int{0}
}

// LockingWaitPolicy controls the policy used for handling conflicting locks
// held by other active transactions when attempting to lock rows due to FOR
// UPDATE/SHARE clauses (i.e. it represents the NOWAIT and SKIP LOCKED options).
type ScanLockingWaitPolicy int32

const (
	// BLOCK represents the default - wait for the lock to become available.
	ScanLockingWaitPolicy_BLOCK ScanLockingWaitPolicy = 0
	// SKIP_LOCKED represents SKIP LOCKED - skip rows that can't be locked.
	//
	// NOTE: SKIP_LOCKED is not currently implemented and does not make it out of
	// the SQL optimizer without throwing an error.
	ScanLockingWaitPolicy_SKIP_LOCKED ScanLockingWaitPolicy = 1
	// ERROR represents NOWAIT - raise an error if a row cannot be locked.
	ScanLockingWaitPolicy_ERROR ScanLockingWaitPolicy = 2
)

var ScanLockingWaitPolicy_name = map[int32]string{
	0: "BLOCK",
	1: "SKIP_LOCKED",
	2: "ERROR",
}

var ScanLockingWaitPolicy_value = map[string]int32{
	"BLOCK":       0,
	"SKIP_LOCKED": 1,
	"ERROR":       2,
}

func (x ScanLockingWaitPolicy) Enum() *ScanLockingWaitPolicy {
	p := new(ScanLockingWaitPolicy)
	*p = x
	return p
}

func (x ScanLockingWaitPolicy) String() string {
	return proto.EnumName(ScanLockingWaitPolicy_name, int32(x))
}

func (x *ScanLockingWaitPolicy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ScanLockingWaitPolicy_value, data, "ScanLockingWaitPolicy")
	if err != nil {
		return err
	}
	*x = ScanLockingWaitPolicy(value)
	return nil
}

func (ScanLockingWaitPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_484351bd13f4281f, []int{1}
}

func init() {
	proto.RegisterEnum("cockroach.sql.sqlbase.ScanLockingStrength", ScanLockingStrength_name, ScanLockingStrength_value)
	proto.RegisterEnum("cockroach.sql.sqlbase.ScanLockingWaitPolicy", ScanLockingWaitPolicy_name, ScanLockingWaitPolicy_value)
}

func init() { proto.RegisterFile("sql/catalog/descpb/locking.proto", fileDescriptor_484351bd13f4281f) }

var fileDescriptor_484351bd13f4281f = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x2e, 0xcc, 0xd1,
	0x4f, 0x4e, 0x2c, 0x49, 0xcc, 0xc9, 0x4f, 0xd7, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0x48, 0xd2, 0xcf,
	0xc9, 0x4f, 0xce, 0xce, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4d, 0xce,
	0x4f, 0xce, 0x2e, 0xca, 0x4f, 0x4c, 0xce, 0xd0, 0x2b, 0x2e, 0xcc, 0x01, 0xe1, 0xa4, 0xc4, 0xe2,
	0x54, 0xad, 0x1c, 0x2e, 0xe1, 0xe0, 0xe4, 0xc4, 0x3c, 0x1f, 0x88, 0xda, 0xe0, 0x92, 0xa2, 0xd4,
	0xbc, 0xf4, 0x92, 0x0c, 0x21, 0x1e, 0x2e, 0x0e, 0x37, 0xff, 0xa0, 0x78, 0x3f, 0x7f, 0x3f, 0x57,
	0x01, 0x06, 0x21, 0x41, 0x2e, 0x5e, 0x10, 0xcf, 0xdb, 0x35, 0x32, 0x3e, 0xd8, 0xc3, 0x31, 0xc8,
	0x55, 0x80, 0x51, 0x88, 0x97, 0x8b, 0x13, 0x24, 0x04, 0xe1, 0x32, 0x09, 0x89, 0x72, 0x09, 0x42,
	0xd4, 0x83, 0x15, 0x85, 0x06, 0xb8, 0x38, 0x86, 0xb8, 0x0a, 0x30, 0x0b, 0xf1, 0x71, 0x71, 0x81,
	0x84, 0xa1, 0x7c, 0x16, 0x2d, 0x3b, 0x2e, 0x51, 0x24, 0xdb, 0xc2, 0x13, 0x33, 0x4b, 0x02, 0xf2,
	0x73, 0x32, 0x93, 0x2b, 0x85, 0x38, 0xb9, 0x58, 0x9d, 0x7c, 0xfc, 0x9d, 0xbd, 0x05, 0x18, 0x84,
	0xf8, 0xb9, 0xb8, 0x83, 0xbd, 0x3d, 0x03, 0xe2, 0x41, 0x5c, 0x57, 0x17, 0x01, 0x46, 0x90, 0x9c,
	0x6b, 0x50, 0x90, 0x7f, 0x90, 0x00, 0x93, 0x93, 0xc6, 0x89, 0x87, 0x72, 0x0c, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0x78, 0xe3, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0x1b, 0xc4, 0xf3, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x9e, 0x04, 0xcd, 0x11, 0x01, 0x00, 0x00,
}