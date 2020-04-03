// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"context"
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execerror"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfra"
)

// Operator is a column vector operator that produces a Batch as output.
type Operator interface {
	// Init initializes this operator. Will be called once at operator setup
	// time. If an operator has an input operator, it's responsible for calling
	// Init on that input operator as well.
	// TODO(yuzefovich): we might need to clarify whether it is ok to call
	// Init() multiple times before the first call to Next(). It is possible to
	// hit the memory limit during Init(), and a disk-backed operator needs to
	// make sure that the input has been initialized. We could also in case that
	// Init() doesn't succeed for bufferingInMemoryOperator - which should only
	// happen when 'workmem' setting is too low - just bail, even if we have
	// disk spilling for that operator.
	Init()

	// Next returns the next Batch from this operator. Once the operator is
	// finished, it will return a Batch with length 0. Subsequent calls to
	// Next at that point will always return a Batch with length 0.
	//
	// Calling Next may invalidate the contents of the last Batch returned by
	// Next.
	// Canceling the provided context results in forceful termination of
	// execution.
	Next(context.Context) coldata.Batch

	execinfra.OpNode
}

// OperatorInitStatus indicates whether Init method has already been called on
// an Operator.
type OperatorInitStatus int

const (
	// OperatorNotInitialized indicates that Init has not been called yet.
	OperatorNotInitialized OperatorInitStatus = iota
	// OperatorInitialized indicates that Init has already been called.
	OperatorInitialized
)

// NonExplainable is a marker interface which identifies an Operator that
// should be omitted from the output of EXPLAIN (VEC). Note that VERBOSE
// explain option will override the omitting behavior.
type NonExplainable interface {
	// nonExplainableMarker is just a marker method. It should never be called.
	nonExplainableMarker()
}

// NewOneInputNode returns an execinfra.OpNode with a single Operator input.
func NewOneInputNode(input Operator) OneInputNode {
	return OneInputNode{input: input}
}

// OneInputNode is an execinfra.OpNode with a single Operator input.
type OneInputNode struct {
	input Operator
}

// ChildCount implements the execinfra.OpNode interface.
func (OneInputNode) ChildCount(verbose bool) int {
	return 1
}

// Child implements the execinfra.OpNode interface.
func (n OneInputNode) Child(nth int, verbose bool) execinfra.OpNode {
	if nth == 0 {
		return n.input
	}
	execerror.VectorizedInternalPanic(fmt.Sprintf("invalid index %d", nth))
	// This code is unreachable, but the compiler cannot infer that.
	return nil
}

// Input returns the single input of this OneInputNode as an Operator.
func (n OneInputNode) Input() Operator {
	return n.input
}

// ZeroInputNode is an execinfra.OpNode with no inputs.
type ZeroInputNode struct{}

// ChildCount implements the execinfra.OpNode interface.
func (ZeroInputNode) ChildCount(verbose bool) int {
	return 0
}

// Child implements the execinfra.OpNode interface.
func (ZeroInputNode) Child(nth int, verbose bool) execinfra.OpNode {
	execerror.VectorizedInternalPanic(fmt.Sprintf("invalid index %d", nth))
	// This code is unreachable, but the compiler cannot infer that.
	return nil
}

// newTwoInputNode returns an execinfra.OpNode with two Operator inputs.
func newTwoInputNode(inputOne, inputTwo Operator) twoInputNode {
	return twoInputNode{inputOne: inputOne, inputTwo: inputTwo}
}

type twoInputNode struct {
	inputOne Operator
	inputTwo Operator
}

func (twoInputNode) ChildCount(verbose bool) int {
	return 2
}

func (n *twoInputNode) Child(nth int, verbose bool) execinfra.OpNode {
	switch nth {
	case 0:
		return n.inputOne
	case 1:
		return n.inputTwo
	}
	execerror.VectorizedInternalPanic(fmt.Sprintf("invalid idx %d", nth))
	// This code is unreachable, but the compiler cannot infer that.
	return nil
}

// TODO(yuzefovich): audit all Operators to make sure that all internal memory
// is accounted for.

// InternalMemoryOperator is an interface that operators which use internal
// memory need to implement. "Internal memory" is defined as memory that is
// "private" to the operator and is not exposed to the outside; notably, it
// does *not* include any coldata.Batch'es and coldata.Vec's.
type InternalMemoryOperator interface {
	Operator
	// InternalMemoryUsage reports the internal memory usage (in bytes) of an
	// operator.
	InternalMemoryUsage() int
}

// resetter is an interface that operators can implement if they can be reset
// either for reusing (to keep the already allocated memory) or during tests.
type resetter interface {
	reset(ctx context.Context)
}

// resettableOperator is an Operator that can be reset.
type resettableOperator interface {
	Operator
	resetter
}

// IdempotentCloser is an object that releases resource on the first call to
// IdempotentClose but does nothing for any subsequent call.
type IdempotentCloser interface {
	IdempotentClose(ctx context.Context) error
}

// closerHelper is a simple helper that helps Operators implement
// IdempotentCloser. If close returns true, resources may be released, if it
// returns false, close has already been called.
// use.
type closerHelper struct {
	closed bool
}

// close marks the closerHelper as closed. If true is returned, this is the
// first call to close.
func (c *closerHelper) close() bool {
	if c.closed {
		return false
	}
	c.closed = true
	return true
}

type closableOperator interface {
	Operator
	IdempotentCloser
}

type noopOperator struct {
	OneInputNode
	NonExplainable
}

var _ Operator = &noopOperator{}

// NewNoop returns a new noop Operator.
func NewNoop(input Operator) Operator {
	return &noopOperator{OneInputNode: NewOneInputNode(input)}
}

func (n *noopOperator) Init() {
	n.input.Init()
}

func (n *noopOperator) Next(ctx context.Context) coldata.Batch {
	return n.input.Next(ctx)
}

func (n *noopOperator) reset(ctx context.Context) {
	if r, ok := n.input.(resetter); ok {
		r.reset(ctx)
	}
}

type zeroOperator struct {
	OneInputNode
	NonExplainable
}

var _ Operator = &zeroOperator{}

// NewZeroOp creates a new operator which just returns an empty batch.
func NewZeroOp(input Operator) Operator {
	return &zeroOperator{OneInputNode: NewOneInputNode(input)}
}

func (s *zeroOperator) Init() {
	s.input.Init()
}

func (s *zeroOperator) Next(ctx context.Context) coldata.Batch {
	// TODO(solon): Can we avoid calling Next on the input at all?
	s.input.Next(ctx)
	return coldata.ZeroBatch
}

type singleTupleNoInputOperator struct {
	ZeroInputNode
	NonExplainable
	batch  coldata.Batch
	nexted bool
}

var _ Operator = &singleTupleNoInputOperator{}

// NewSingleTupleNoInputOp creates a new Operator which returns a batch of
// length 1 with no actual columns on the first call to Next() and zero-length
// batches on all consecutive calls.
func NewSingleTupleNoInputOp(allocator *Allocator) Operator {
	return &singleTupleNoInputOperator{
		batch: allocator.NewMemBatchWithSize(nil /* types */, 1 /* size */),
	}
}

func (s *singleTupleNoInputOperator) Init() {
}

func (s *singleTupleNoInputOperator) Next(ctx context.Context) coldata.Batch {
	s.batch.ResetInternalBatch()
	if s.nexted {
		return coldata.ZeroBatch
	}
	s.nexted = true
	s.batch.SetLength(1)
	return s.batch
}

// feedOperator is used to feed an Operator chain with input by manually
// setting the next batch.
type feedOperator struct {
	ZeroInputNode
	NonExplainable
	batch coldata.Batch
}

func (feedOperator) Init() {}

func (o *feedOperator) Next(context.Context) coldata.Batch {
	return o.batch
}

var _ Operator = &feedOperator{}

// vectorTypeEnforcer is a utility Operator that on every call to Next
// enforces that non-zero length batch from the input has a vector of the
// desired type in the desired position. If the width of the batch is less than
// the desired position, a new vector will be appended; if the batch has a
// well-typed vector of an undesired type in the desired position, an error
// will occur.
//
// This Operator is designed to be planned as a wrapper on the input to a
// "projecting" Operator (such Operator that has a single column as its output
// and does not touch other columns by simply passing them along).
//
// The intended diagram is as follows:
//
//       original input                (with schema [t1, ..., tN])
//       --------------
//             |
//             ↓
//     vectorTypeEnforcer              (will enforce that tN+1 = outputType)
//     ------------------
//             |
//             ↓
//   "projecting" operator             (projects its output of type outputType
//   ---------------------              in column at position of N+1)
//
type vectorTypeEnforcer struct {
	OneInputNode
	NonExplainable

	allocator *Allocator
	typ       coltypes.T
	idx       int
}

var _ Operator = &vectorTypeEnforcer{}

func newVectorTypeEnforcer(allocator *Allocator, input Operator, typ coltypes.T, idx int) Operator {
	return &vectorTypeEnforcer{
		OneInputNode: NewOneInputNode(input),
		allocator:    allocator,
		typ:          typ,
		idx:          idx,
	}
}

func (e *vectorTypeEnforcer) Init() {
	e.input.Init()
}

func (e *vectorTypeEnforcer) Next(ctx context.Context) coldata.Batch {
	b := e.input.Next(ctx)
	if b.Length() == 0 {
		return b
	}
	e.allocator.maybeAppendColumn(b, e.typ, e.idx)
	return b
}

// batchSchemaPrefixEnforcer is similar to vectorTypeEnforcer in its purpose,
// but it enforces that the prefix of the columns of the non-zero length batch
// satisfies the desired schema. It needs to wrap the input to a "projecting"
// operator that internally uses other "projecting" operators (for example,
// caseOp and logical projection operators). This operator supports type
// schemas with coltypes.Unhandled type in which case in the corresponding
// position an "unknown" vector can be appended.
//
// NOTE: the type schema passed into batchSchemaPrefixEnforcer *must* include
// the output type of the Operator that the enforcer will be the input to.
type batchSchemaPrefixEnforcer struct {
	OneInputNode
	NonExplainable

	allocator *Allocator
	typs      []coltypes.T
}

var _ Operator = &batchSchemaPrefixEnforcer{}

func newBatchSchemaPrefixEnforcer(
	allocator *Allocator, input Operator, typs []coltypes.T,
) *batchSchemaPrefixEnforcer {
	return &batchSchemaPrefixEnforcer{
		OneInputNode: NewOneInputNode(input),
		allocator:    allocator,
		typs:         typs,
	}
}

func (e *batchSchemaPrefixEnforcer) Init() {
	e.input.Init()
}

func (e *batchSchemaPrefixEnforcer) Next(ctx context.Context) coldata.Batch {
	b := e.input.Next(ctx)
	if b.Length() == 0 {
		return b
	}
	for i, typ := range e.typs {
		e.allocator.maybeAppendColumn(b, typ, i)
	}
	return b
}
