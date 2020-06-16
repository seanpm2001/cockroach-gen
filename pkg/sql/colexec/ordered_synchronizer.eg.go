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
	"container/heap"
	"context"
	"fmt"
	"time"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfra"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfrapb"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
	"github.com/cockroachdb/cockroach/pkg/util/encoding"
)

// OrderedSynchronizer receives rows from multiple inputs and produces a single
// stream of rows, ordered according to a set of columns. The rows in each input
// stream are assumed to be ordered according to the same set of columns.
type OrderedSynchronizer struct {
	allocator             *colmem.Allocator
	inputs                []SynchronizerInput
	ordering              sqlbase.ColumnOrdering
	typs                  []*types.T
	canonicalTypeFamilies []types.Family

	// inputBatches stores the current batch for each input.
	inputBatches []coldata.Batch
	// inputIndices stores the current index into each input batch.
	inputIndices []int
	// heap is a min heap which stores indices into inputBatches. The "current
	// value" of ith input batch is the tuple at inputIndices[i] position of
	// inputBatches[i] batch. If an input is fully exhausted, it will be removed
	// from heap.
	heap []int
	// comparators stores one comparator per ordering column.
	comparators []vecComparator
	output      coldata.Batch
	outNulls    []*coldata.Nulls
	// In order to reduce the number of interface conversions, we will get access
	// to the underlying slice for the output vectors and will use them directly.
	outBoolCols      [][]bool
	outBytesCols     []*coldata.Bytes
	outDecimalCols   [][]apd.Decimal
	outInt16Cols     [][]int16
	outInt32Cols     [][]int32
	outInt64Cols     [][]int64
	outFloat64Cols   [][]float64
	outTimestampCols [][]time.Time
	outIntervalCols  [][]duration.Duration
	outDatumCols     []coldata.DatumVec
	// outColsMap contains the positions of the corresponding vectors in the
	// slice for the same types. For example, if we have an output batch with
	// types = [Int64, Int64, Bool, Bytes, Bool, Int64], then outColsMap will be
	//                      [0, 1, 0, 0, 1, 2]
	//                       ^  ^  ^  ^  ^  ^
	//                       |  |  |  |  |  |
	//                       |  |  |  |  |  3rd among all Int64's
	//                       |  |  |  |  2nd among all Bool's
	//                       |  |  |  1st among all Bytes's
	//                       |  |  1st among all Bool's
	//                       |  2nd among all Int64's
	//                       1st among all Int64's
	outColsMap []int
}

var _ colexecbase.Operator = &OrderedSynchronizer{}

// ChildCount implements the execinfrapb.OpNode interface.
func (o *OrderedSynchronizer) ChildCount(verbose bool) int {
	return len(o.inputs)
}

// Child implements the execinfrapb.OpNode interface.
func (o *OrderedSynchronizer) Child(nth int, verbose bool) execinfra.OpNode {
	return o.inputs[nth].Op
}

// NewOrderedSynchronizer creates a new OrderedSynchronizer.
func NewOrderedSynchronizer(
	allocator *colmem.Allocator,
	inputs []SynchronizerInput,
	typs []*types.T,
	ordering sqlbase.ColumnOrdering,
) (*OrderedSynchronizer, error) {
	return &OrderedSynchronizer{
		allocator:             allocator,
		inputs:                inputs,
		ordering:              ordering,
		typs:                  typs,
		canonicalTypeFamilies: typeconv.ToCanonicalTypeFamilies(typs),
	}, nil
}

// Next is part of the Operator interface.
func (o *OrderedSynchronizer) Next(ctx context.Context) coldata.Batch {
	if o.inputBatches == nil {
		o.inputBatches = make([]coldata.Batch, len(o.inputs))
		o.heap = make([]int, 0, len(o.inputs))
		for i := range o.inputs {
			o.inputBatches[i] = o.inputs[i].Op.Next(ctx)
			o.updateComparators(i)
			if o.inputBatches[i].Length() > 0 {
				o.heap = append(o.heap, i)
			}
		}
		heap.Init(o)
	}
	o.output.ResetInternalBatch()
	outputIdx := 0
	o.allocator.PerformOperation(o.output.ColVecs(), func() {
		for outputIdx < coldata.BatchSize() {
			if o.Len() == 0 {
				// All inputs exhausted.
				break
			}

			minBatch := o.heap[0]
			// Copy the min row into the output.
			batch := o.inputBatches[minBatch]
			srcRowIdx := o.inputIndices[minBatch]
			if sel := batch.Selection(); sel != nil {
				srcRowIdx = sel[srcRowIdx]
			}
			for i := range o.typs {
				vec := batch.ColVec(i)
				if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(srcRowIdx) {
					o.outNulls[i].SetNull(outputIdx)
				} else {
					switch o.canonicalTypeFamilies[i] {
					case types.BoolFamily:
						switch o.typs[i].Width() {
						case -1:
						default:
							srcCol := vec.Bool()
							outCol := o.outBoolCols[o.outColsMap[i]]
							v := srcCol[srcRowIdx]
							outCol[outputIdx] = v
						}
					case types.BytesFamily:
						switch o.typs[i].Width() {
						case -1:
						default:
							srcCol := vec.Bytes()
							outCol := o.outBytesCols[o.outColsMap[i]]
							v := srcCol.Get(srcRowIdx)
							outCol.Set(outputIdx, v)
						}
					case types.DecimalFamily:
						switch o.typs[i].Width() {
						case -1:
						default:
							srcCol := vec.Decimal()
							outCol := o.outDecimalCols[o.outColsMap[i]]
							v := srcCol[srcRowIdx]
							outCol[outputIdx].Set(&v)
						}
					case types.IntFamily:
						switch o.typs[i].Width() {
						case 16:
							srcCol := vec.Int16()
							outCol := o.outInt16Cols[o.outColsMap[i]]
							v := srcCol[srcRowIdx]
							outCol[outputIdx] = v
						case 32:
							srcCol := vec.Int32()
							outCol := o.outInt32Cols[o.outColsMap[i]]
							v := srcCol[srcRowIdx]
							outCol[outputIdx] = v
						case -1:
						default:
							srcCol := vec.Int64()
							outCol := o.outInt64Cols[o.outColsMap[i]]
							v := srcCol[srcRowIdx]
							outCol[outputIdx] = v
						}
					case types.FloatFamily:
						switch o.typs[i].Width() {
						case -1:
						default:
							srcCol := vec.Float64()
							outCol := o.outFloat64Cols[o.outColsMap[i]]
							v := srcCol[srcRowIdx]
							outCol[outputIdx] = v
						}
					case types.TimestampTZFamily:
						switch o.typs[i].Width() {
						case -1:
						default:
							srcCol := vec.Timestamp()
							outCol := o.outTimestampCols[o.outColsMap[i]]
							v := srcCol[srcRowIdx]
							outCol[outputIdx] = v
						}
					case types.IntervalFamily:
						switch o.typs[i].Width() {
						case -1:
						default:
							srcCol := vec.Interval()
							outCol := o.outIntervalCols[o.outColsMap[i]]
							v := srcCol[srcRowIdx]
							outCol[outputIdx] = v
						}
					case typeconv.DatumVecCanonicalTypeFamily:
						switch o.typs[i].Width() {
						case -1:
						default:
							srcCol := vec.Datum()
							outCol := o.outDatumCols[o.outColsMap[i]]
							v := srcCol.Get(srcRowIdx)
							outCol.Set(outputIdx, v)
						}
					default:
						colexecerror.InternalError(fmt.Sprintf("unhandled type %s", o.typs[i].String()))
					}
				}
			}

			// Advance the input batch, fetching a new batch if necessary.
			if o.inputIndices[minBatch]+1 < o.inputBatches[minBatch].Length() {
				o.inputIndices[minBatch]++
			} else {
				o.inputBatches[minBatch] = o.inputs[minBatch].Op.Next(ctx)
				o.inputIndices[minBatch] = 0
				o.updateComparators(minBatch)
			}
			if o.inputBatches[minBatch].Length() == 0 {
				heap.Remove(o, 0)
			} else {
				heap.Fix(o, 0)
			}

			outputIdx++
		}
	})

	o.output.SetLength(outputIdx)
	return o.output
}

// Init is part of the Operator interface.
func (o *OrderedSynchronizer) Init() {
	o.inputIndices = make([]int, len(o.inputs))
	o.output = o.allocator.NewMemBatch(o.typs)
	o.outNulls = make([]*coldata.Nulls, len(o.typs))
	o.outColsMap = make([]int, len(o.typs))
	for i, outVec := range o.output.ColVecs() {
		o.outNulls[i] = outVec.Nulls()
		switch typeconv.TypeFamilyToCanonicalTypeFamily(o.typs[i].Family()) {
		case types.BoolFamily:
			switch o.typs[i].Width() {
			case -1:
			default:
				o.outColsMap[i] = len(o.outBoolCols)
				o.outBoolCols = append(o.outBoolCols, outVec.Bool())
			}
		case types.BytesFamily:
			switch o.typs[i].Width() {
			case -1:
			default:
				o.outColsMap[i] = len(o.outBytesCols)
				o.outBytesCols = append(o.outBytesCols, outVec.Bytes())
			}
		case types.DecimalFamily:
			switch o.typs[i].Width() {
			case -1:
			default:
				o.outColsMap[i] = len(o.outDecimalCols)
				o.outDecimalCols = append(o.outDecimalCols, outVec.Decimal())
			}
		case types.IntFamily:
			switch o.typs[i].Width() {
			case 16:
				o.outColsMap[i] = len(o.outInt16Cols)
				o.outInt16Cols = append(o.outInt16Cols, outVec.Int16())
			case 32:
				o.outColsMap[i] = len(o.outInt32Cols)
				o.outInt32Cols = append(o.outInt32Cols, outVec.Int32())
			case -1:
			default:
				o.outColsMap[i] = len(o.outInt64Cols)
				o.outInt64Cols = append(o.outInt64Cols, outVec.Int64())
			}
		case types.FloatFamily:
			switch o.typs[i].Width() {
			case -1:
			default:
				o.outColsMap[i] = len(o.outFloat64Cols)
				o.outFloat64Cols = append(o.outFloat64Cols, outVec.Float64())
			}
		case types.TimestampTZFamily:
			switch o.typs[i].Width() {
			case -1:
			default:
				o.outColsMap[i] = len(o.outTimestampCols)
				o.outTimestampCols = append(o.outTimestampCols, outVec.Timestamp())
			}
		case types.IntervalFamily:
			switch o.typs[i].Width() {
			case -1:
			default:
				o.outColsMap[i] = len(o.outIntervalCols)
				o.outIntervalCols = append(o.outIntervalCols, outVec.Interval())
			}
		case typeconv.DatumVecCanonicalTypeFamily:
			switch o.typs[i].Width() {
			case -1:
			default:
				o.outColsMap[i] = len(o.outDatumCols)
				o.outDatumCols = append(o.outDatumCols, outVec.Datum())
			}
		default:
			colexecerror.InternalError(fmt.Sprintf("unhandled type %s", o.typs[i]))
		}
	}
	for i := range o.inputs {
		o.inputs[i].Op.Init()
	}
	o.comparators = make([]vecComparator, len(o.ordering))
	for i := range o.ordering {
		typ := o.typs[o.ordering[i].ColIdx]
		o.comparators[i] = GetVecComparator(typ, len(o.inputs))
	}
}

func (o *OrderedSynchronizer) DrainMeta(ctx context.Context) []execinfrapb.ProducerMetadata {
	var bufferedMeta []execinfrapb.ProducerMetadata
	for _, input := range o.inputs {
		bufferedMeta = append(bufferedMeta, input.MetadataSources.DrainMeta(ctx)...)
	}
	return bufferedMeta
}

func (o *OrderedSynchronizer) compareRow(batchIdx1 int, batchIdx2 int) int {
	batch1 := o.inputBatches[batchIdx1]
	batch2 := o.inputBatches[batchIdx2]
	valIdx1 := o.inputIndices[batchIdx1]
	valIdx2 := o.inputIndices[batchIdx2]
	if sel := batch1.Selection(); sel != nil {
		valIdx1 = sel[valIdx1]
	}
	if sel := batch2.Selection(); sel != nil {
		valIdx2 = sel[valIdx2]
	}
	for i := range o.ordering {
		info := o.ordering[i]
		res := o.comparators[i].compare(batchIdx1, batchIdx2, valIdx1, valIdx2)
		if res != 0 {
			switch d := info.Direction; d {
			case encoding.Ascending:
				return res
			case encoding.Descending:
				return -res
			default:
				colexecerror.InternalError(fmt.Sprintf("unexpected direction value %d", d))
			}
		}
	}
	return 0
}

// updateComparators should be run whenever a new batch is fetched. It updates
// all the relevant vectors in o.comparators.
func (o *OrderedSynchronizer) updateComparators(batchIdx int) {
	batch := o.inputBatches[batchIdx]
	if batch.Length() == 0 {
		return
	}
	for i := range o.ordering {
		vec := batch.ColVec(o.ordering[i].ColIdx)
		o.comparators[i].setVec(batchIdx, vec)
	}
}

// Len is part of heap.Interface and is only meant to be used internally.
func (o *OrderedSynchronizer) Len() int {
	return len(o.heap)
}

// Less is part of heap.Interface and is only meant to be used internally.
func (o *OrderedSynchronizer) Less(i, j int) bool {
	return o.compareRow(o.heap[i], o.heap[j]) < 0
}

// Swap is part of heap.Interface and is only meant to be used internally.
func (o *OrderedSynchronizer) Swap(i, j int) {
	o.heap[i], o.heap[j] = o.heap[j], o.heap[i]
}

// Push is part of heap.Interface and is only meant to be used internally.
func (o *OrderedSynchronizer) Push(x interface{}) {
	o.heap = append(o.heap, x.(int))
}

// Pop is part of heap.Interface and is only meant to be used internally.
func (o *OrderedSynchronizer) Pop() interface{} {
	x := o.heap[len(o.heap)-1]
	o.heap = o.heap[:len(o.heap)-1]
	return x
}
