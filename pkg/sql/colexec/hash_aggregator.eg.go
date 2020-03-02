// Code generated by execgen; DO NOT EDIT.
// Copyright 2020 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"bytes"
	"fmt"
	"math"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execerror"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

// match takes a selection vector and compares it against the values of the key
// of its aggregation function. It returns a selection vector representing the
// unmatched tuples and a boolean to indicate whether or not there are any
// matching tuples. It directly writes the result of matched tuples into the
// selection vector of 'b' and sets the length of the batch to the number of
// matching tuples. match also takes a diff boolean slice for internal use.
// This slice need to be allocated to be at at least as big as sel and set to
// all false. diff will be reset to all false when match returns. This is to
// avoid additional slice allocation.
// NOTE: the return vector will reuse the memory allocated for the selection
//       vector.
func (v hashAggFuncs) match(
	sel []int,
	b coldata.Batch,
	keyCols []uint32,
	keyTypes []coltypes.T,
	keyMapping coldata.Batch,
	diff []bool,
) (bool, []int) {
	// We want to directly write to the selection vector to avoid extra
	// allocation.
	b.SetSelection(true)
	matched := b.Selection()
	matched = matched[:0]

	aggKeyIdx := v.keyIdx

	for keyIdx, colIdx := range keyCols {
		lhs := keyMapping.ColVec(keyIdx)
		lhsHasNull := lhs.MaybeHasNulls()
		lhsNull := lhs.Nulls().NullAt(v.keyIdx)

		rhs := b.ColVec(int(colIdx))
		rhsHasNull := rhs.MaybeHasNulls()

		keyTyp := keyTypes[keyIdx]

		switch keyTyp {
		case coltypes.Bool:
			lhsCol := lhs.Bool()
			rhsCol := rhs.Bool()
			if lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					diffNull := (lhsNull != rhs.Nulls().NullAt(rowIdx))
					if diffNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						if !lhsVal && rhsVal {
							cmpResult = -1
						} else if lhsVal && !rhsVal {
							cmpResult = 1
						} else {
							cmpResult = 0
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if lhsHasNull && !rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					if lhsNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						if !lhsVal && rhsVal {
							cmpResult = -1
						} else if lhsVal && !rhsVal {
							cmpResult = 1
						} else {
							cmpResult = 0
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if !lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						if !lhsVal && rhsVal {
							cmpResult = -1
						} else if lhsVal && !rhsVal {
							cmpResult = 1
						} else {
							cmpResult = 0
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						if !lhsVal && rhsVal {
							cmpResult = -1
						} else if lhsVal && !rhsVal {
							cmpResult = 1
						} else {
							cmpResult = 0
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			}
		case coltypes.Bytes:
			lhsCol := lhs.Bytes()
			rhsCol := rhs.Bytes()
			if lhsHasNull && rhsHasNull {
				lhsVal := lhsCol.Get(aggKeyIdx)

				for selIdx, rowIdx := range sel {
					diffNull := (lhsNull != rhs.Nulls().NullAt(rowIdx))
					if diffNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol.Get(rowIdx)

					var cmp bool

					{
						var cmpResult int
						cmpResult = bytes.Compare(lhsVal, rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if lhsHasNull && !rhsHasNull {
				lhsVal := lhsCol.Get(aggKeyIdx)

				for selIdx, rowIdx := range sel {
					if lhsNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol.Get(rowIdx)

					var cmp bool

					{
						var cmpResult int
						cmpResult = bytes.Compare(lhsVal, rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if !lhsHasNull && rhsHasNull {
				lhsVal := lhsCol.Get(aggKeyIdx)

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol.Get(rowIdx)

					var cmp bool

					{
						var cmpResult int
						cmpResult = bytes.Compare(lhsVal, rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else {
				lhsVal := lhsCol.Get(aggKeyIdx)

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol.Get(rowIdx)

					var cmp bool

					{
						var cmpResult int
						cmpResult = bytes.Compare(lhsVal, rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			}
		case coltypes.Decimal:
			lhsCol := lhs.Decimal()
			rhsCol := rhs.Decimal()
			if lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					diffNull := (lhsNull != rhs.Nulls().NullAt(rowIdx))
					if diffNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int
						cmpResult = tree.CompareDecimals(&lhsVal, &rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if lhsHasNull && !rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					if lhsNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int
						cmpResult = tree.CompareDecimals(&lhsVal, &rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if !lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int
						cmpResult = tree.CompareDecimals(&lhsVal, &rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int
						cmpResult = tree.CompareDecimals(&lhsVal, &rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			}
		case coltypes.Int16:
			lhsCol := lhs.Int16()
			rhsCol := rhs.Int16()
			if lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					diffNull := (lhsNull != rhs.Nulls().NullAt(rowIdx))
					if diffNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if lhsHasNull && !rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					if lhsNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if !lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			}
		case coltypes.Int32:
			lhsCol := lhs.Int32()
			rhsCol := rhs.Int32()
			if lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					diffNull := (lhsNull != rhs.Nulls().NullAt(rowIdx))
					if diffNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if lhsHasNull && !rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					if lhsNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if !lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			}
		case coltypes.Int64:
			lhsCol := lhs.Int64()
			rhsCol := rhs.Int64()
			if lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					diffNull := (lhsNull != rhs.Nulls().NullAt(rowIdx))
					if diffNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if lhsHasNull && !rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					if lhsNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if !lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := int64(lhsVal), int64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else {
								cmpResult = 0
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			}
		case coltypes.Float64:
			lhsCol := lhs.Float64()
			rhsCol := rhs.Float64()
			if lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					diffNull := (lhsNull != rhs.Nulls().NullAt(rowIdx))
					if diffNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := float64(lhsVal), float64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else if a == b {
								cmpResult = 0
							} else if math.IsNaN(a) {
								if math.IsNaN(b) {
									cmpResult = 0
								} else {
									cmpResult = -1
								}
							} else {
								cmpResult = 1
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if lhsHasNull && !rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					if lhsNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := float64(lhsVal), float64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else if a == b {
								cmpResult = 0
							} else if math.IsNaN(a) {
								if math.IsNaN(b) {
									cmpResult = 0
								} else {
									cmpResult = -1
								}
							} else {
								cmpResult = 1
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if !lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := float64(lhsVal), float64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else if a == b {
								cmpResult = 0
							} else if math.IsNaN(a) {
								if math.IsNaN(b) {
									cmpResult = 0
								} else {
									cmpResult = -1
								}
							} else {
								cmpResult = 1
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						{
							a, b := float64(lhsVal), float64(rhsVal)
							if a < b {
								cmpResult = -1
							} else if a > b {
								cmpResult = 1
							} else if a == b {
								cmpResult = 0
							} else if math.IsNaN(a) {
								if math.IsNaN(b) {
									cmpResult = 0
								} else {
									cmpResult = -1
								}
							} else {
								cmpResult = 1
							}
						}

						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			}
		case coltypes.Timestamp:
			lhsCol := lhs.Timestamp()
			rhsCol := rhs.Timestamp()
			if lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					diffNull := (lhsNull != rhs.Nulls().NullAt(rowIdx))
					if diffNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						if lhsVal.Before(rhsVal) {
							cmpResult = -1
						} else if rhsVal.Before(lhsVal) {
							cmpResult = 1
						} else {
							cmpResult = 0
						}
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if lhsHasNull && !rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					if lhsNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						if lhsVal.Before(rhsVal) {
							cmpResult = -1
						} else if rhsVal.Before(lhsVal) {
							cmpResult = 1
						} else {
							cmpResult = 0
						}
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if !lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						if lhsVal.Before(rhsVal) {
							cmpResult = -1
						} else if rhsVal.Before(lhsVal) {
							cmpResult = 1
						} else {
							cmpResult = 0
						}
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int

						if lhsVal.Before(rhsVal) {
							cmpResult = -1
						} else if rhsVal.Before(lhsVal) {
							cmpResult = 1
						} else {
							cmpResult = 0
						}
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			}
		case coltypes.Interval:
			lhsCol := lhs.Interval()
			rhsCol := rhs.Interval()
			if lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					diffNull := (lhsNull != rhs.Nulls().NullAt(rowIdx))
					if diffNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int
						cmpResult = lhsVal.Compare(rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if lhsHasNull && !rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {
					if lhsNull {
						diff[selIdx] = true
						continue
					}

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int
						cmpResult = lhsVal.Compare(rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else if !lhsHasNull && rhsHasNull {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int
						cmpResult = lhsVal.Compare(rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			} else {
				lhsVal := lhsCol[aggKeyIdx]

				for selIdx, rowIdx := range sel {

					rhsVal := rhsCol[rowIdx]

					var cmp bool

					{
						var cmpResult int
						cmpResult = lhsVal.Compare(rhsVal)
						cmp = cmpResult != 0
					}

					diff[selIdx] = diff[selIdx] || cmp
				}

			}
		default:
			execerror.VectorizedInternalPanic(fmt.Sprintf("unhandled type %d", keyTyp))
		}
	}

	remaining := sel[:0]
	anyMatched := false

	for selIdx, isDiff := range diff {
		if isDiff {
			remaining = append(remaining, sel[selIdx])
		} else {
			matched = append(matched, sel[selIdx])
		}
	}

	if len(matched) > 0 {
		b.SetLength(len(matched))
		anyMatched = true
	}

	// Reset diff slice back to all false.
	for n := 0; n < len(diff); n += copy(diff, zeroBoolColumn) {
	}

	return anyMatched, remaining
}
