// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Package lock provides type definitions for locking-related concepts used by
// concurrency control in the key-value layer.
package lock

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
//
type Strength uint32

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
	None Strength = iota

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
	Shared

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
	// form of locking. As with Shared locks, the lock holder of a Shared lock
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
	Upgrade

	// Exclusive (X) locks are used by read-write and read-only operations and
	// provide a transaction with exclusive access to a key. When an Exclusive
	// lock is held by a transaction on a given key, no other transaction can
	// read from or write to that key. The lock holder is free to read from and
	// write to the key as frequently as it would like.
	Exclusive
)

// Silence unused warnings.
var _ = None
var _ = Shared
var _ = Upgrade
var _ = Exclusive

// Durability represents the different durability properties of a lock acquired
// by a transaction. Durability levels provide varying degrees of survivability,
// often in exchange for the cost of lock acquisition.
type Durability uint32

const (
	_ Durability = iota

	// Unreplicated locks are held only on a single Replica in a Range, which is
	// typically the leaseholder. Unreplicated locks are very fast to acquire
	// and release because they are held in memory or on fast local storage and
	// require no cross-node coordination to update. In exchange, Unreplicated
	// locks provide no guarantee of survivability across lease transfers or
	// leaseholder crashes. They should therefore be thought of as best-effort
	// and should not be relied upon for correctness.
	Unreplicated

	// Replicated locks are held on at least a quorum of Replicas in a Range.
	// They are slower to acquire and release than Unreplicated locks because
	// updating them requires both cross-node coordination and interaction with
	// durable storage. In exchange, Replicated locks provide a guarantee of
	// survivability across lease transfers, leaseholder crashes, and other
	// forms of failure events. They will remain available as long as their
	// Range remains available and they will never be lost.
	Replicated
)

// Silence unused warnings.
var _ = Unreplicated
var _ = Replicated
