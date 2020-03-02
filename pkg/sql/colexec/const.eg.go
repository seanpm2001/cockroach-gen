// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"context"
	"time"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
	"github.com/pkg/errors"
)

// NewConstOp creates a new operator that produces a constant value constVal of
// type t at index outputIdx.
func NewConstOp(
	allocator *Allocator, input Operator, t coltypes.T, constVal interface{}, outputIdx int,
) (Operator, error) {
	switch t {
	case coltypes.Bool:
		return &constBoolOp{
			OneInputNode: NewOneInputNode(input),
			allocator:    allocator,
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(bool),
		}, nil
	case coltypes.Bytes:
		return &constBytesOp{
			OneInputNode: NewOneInputNode(input),
			allocator:    allocator,
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.([]byte),
		}, nil
	case coltypes.Decimal:
		return &constDecimalOp{
			OneInputNode: NewOneInputNode(input),
			allocator:    allocator,
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(apd.Decimal),
		}, nil
	case coltypes.Int16:
		return &constInt16Op{
			OneInputNode: NewOneInputNode(input),
			allocator:    allocator,
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(int16),
		}, nil
	case coltypes.Int32:
		return &constInt32Op{
			OneInputNode: NewOneInputNode(input),
			allocator:    allocator,
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(int32),
		}, nil
	case coltypes.Int64:
		return &constInt64Op{
			OneInputNode: NewOneInputNode(input),
			allocator:    allocator,
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(int64),
		}, nil
	case coltypes.Float64:
		return &constFloat64Op{
			OneInputNode: NewOneInputNode(input),
			allocator:    allocator,
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(float64),
		}, nil
	case coltypes.Timestamp:
		return &constTimestampOp{
			OneInputNode: NewOneInputNode(input),
			allocator:    allocator,
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(time.Time),
		}, nil
	case coltypes.Interval:
		return &constIntervalOp{
			OneInputNode: NewOneInputNode(input),
			allocator:    allocator,
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(duration.Duration),
		}, nil
	default:
		return nil, errors.Errorf("unsupported const type %s", t)
	}
}

type constBoolOp struct {
	OneInputNode

	allocator *Allocator
	typ       coltypes.T
	outputIdx int
	constVal  bool
}

func (c constBoolOp) Init() {
	c.input.Init()
}

func (c constBoolOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	c.allocator.MaybeAddColumn(batch, c.typ, c.outputIdx)
	vec := batch.ColVec(c.outputIdx)
	col := vec.Bool()
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				col = col[0:n]
				for i := range col {
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constBytesOp struct {
	OneInputNode

	allocator *Allocator
	typ       coltypes.T
	outputIdx int
	constVal  []byte
}

func (c constBytesOp) Init() {
	c.input.Init()
}

func (c constBytesOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	c.allocator.MaybeAddColumn(batch, c.typ, c.outputIdx)
	vec := batch.ColVec(c.outputIdx)
	col := vec.Bytes()
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col.Set(i, c.constVal)
				}
			} else {
				col = col
				_ = 0
				_ = n
				for i := 0; i < n; i++ {
					col.Set(i, c.constVal)
				}
			}
		},
	)
	return batch
}

type constDecimalOp struct {
	OneInputNode

	allocator *Allocator
	typ       coltypes.T
	outputIdx int
	constVal  apd.Decimal
}

func (c constDecimalOp) Init() {
	c.input.Init()
}

func (c constDecimalOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	c.allocator.MaybeAddColumn(batch, c.typ, c.outputIdx)
	vec := batch.ColVec(c.outputIdx)
	col := vec.Decimal()
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i].Set(&c.constVal)
				}
			} else {
				col = col[0:n]
				for i := range col {
					col[i].Set(&c.constVal)
				}
			}
		},
	)
	return batch
}

type constInt16Op struct {
	OneInputNode

	allocator *Allocator
	typ       coltypes.T
	outputIdx int
	constVal  int16
}

func (c constInt16Op) Init() {
	c.input.Init()
}

func (c constInt16Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	c.allocator.MaybeAddColumn(batch, c.typ, c.outputIdx)
	vec := batch.ColVec(c.outputIdx)
	col := vec.Int16()
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				col = col[0:n]
				for i := range col {
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constInt32Op struct {
	OneInputNode

	allocator *Allocator
	typ       coltypes.T
	outputIdx int
	constVal  int32
}

func (c constInt32Op) Init() {
	c.input.Init()
}

func (c constInt32Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	c.allocator.MaybeAddColumn(batch, c.typ, c.outputIdx)
	vec := batch.ColVec(c.outputIdx)
	col := vec.Int32()
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				col = col[0:n]
				for i := range col {
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constInt64Op struct {
	OneInputNode

	allocator *Allocator
	typ       coltypes.T
	outputIdx int
	constVal  int64
}

func (c constInt64Op) Init() {
	c.input.Init()
}

func (c constInt64Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	c.allocator.MaybeAddColumn(batch, c.typ, c.outputIdx)
	vec := batch.ColVec(c.outputIdx)
	col := vec.Int64()
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				col = col[0:n]
				for i := range col {
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constFloat64Op struct {
	OneInputNode

	allocator *Allocator
	typ       coltypes.T
	outputIdx int
	constVal  float64
}

func (c constFloat64Op) Init() {
	c.input.Init()
}

func (c constFloat64Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	c.allocator.MaybeAddColumn(batch, c.typ, c.outputIdx)
	vec := batch.ColVec(c.outputIdx)
	col := vec.Float64()
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				col = col[0:n]
				for i := range col {
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constTimestampOp struct {
	OneInputNode

	allocator *Allocator
	typ       coltypes.T
	outputIdx int
	constVal  time.Time
}

func (c constTimestampOp) Init() {
	c.input.Init()
}

func (c constTimestampOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	c.allocator.MaybeAddColumn(batch, c.typ, c.outputIdx)
	vec := batch.ColVec(c.outputIdx)
	col := vec.Timestamp()
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				col = col[0:n]
				for i := range col {
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constIntervalOp struct {
	OneInputNode

	allocator *Allocator
	typ       coltypes.T
	outputIdx int
	constVal  duration.Duration
}

func (c constIntervalOp) Init() {
	c.input.Init()
}

func (c constIntervalOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	c.allocator.MaybeAddColumn(batch, c.typ, c.outputIdx)
	vec := batch.ColVec(c.outputIdx)
	col := vec.Interval()
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				col = col[0:n]
				for i := range col {
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

// NewConstNullOp creates a new operator that produces a constant (untyped) NULL
// value at index outputIdx.
func NewConstNullOp(allocator *Allocator, input Operator, outputIdx int, typ coltypes.T) Operator {
	return &constNullOp{
		OneInputNode: NewOneInputNode(input),
		allocator:    allocator,
		outputIdx:    outputIdx,
		typ:          typ,
	}
}

type constNullOp struct {
	OneInputNode
	allocator *Allocator
	outputIdx int
	typ       coltypes.T
}

var _ Operator = &constNullOp{}

func (c constNullOp) Init() {
	c.input.Init()
}

func (c constNullOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	c.allocator.MaybeAddColumn(batch, c.typ, c.outputIdx)

	col := batch.ColVec(c.outputIdx)
	nulls := col.Nulls()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			nulls.SetNull(i)
		}
	} else {
		nulls.SetNulls()
	}
	return batch
}
