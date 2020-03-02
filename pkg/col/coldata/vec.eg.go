// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package coldata

import (
	"fmt"
	"time"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
)

func (m *memColumn) Append(args SliceArgs) {
	switch args.ColType {
	case coltypes.Bool:
		fromCol := args.Src.Bool()
		toCol := m.Bool()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:args.DestIdx]
			for _, selIdx := range sel {
				val := fromCol[selIdx]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Bytes:
		fromCol := args.Src.Bytes()
		toCol := m.Bytes()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol.AppendSlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// We need to truncate toCol before appending to it, so in case of Bytes,
			// we append an empty slice.
			toCol.AppendSlice(toCol, args.DestIdx, 0, 0)
			// We will be getting all values below to be appended, regardless of
			// whether the value is NULL. It is possible that Bytes' invariant of
			// non-decreasing offsets on the source is currently not maintained, so
			// we explicitly enforce it.
			maxIdx := 0
			for _, selIdx := range sel {
				if selIdx > maxIdx {
					maxIdx = selIdx
				}
			}
			fromCol.UpdateOffsetsToBeNonDecreasing(maxIdx + 1)
			for _, selIdx := range sel {
				val := fromCol.Get(selIdx)
				toCol.AppendVal(val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Decimal:
		fromCol := args.Src.Decimal()
		toCol := m.Decimal()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			{
				__desiredCap := args.DestIdx + args.SrcEndIdx - args.SrcStartIdx
				if cap(toCol) >= __desiredCap {
					toCol = toCol[:__desiredCap]
				} else {
					__prevCap := cap(toCol)
					__capToAllocate := __desiredCap
					if __capToAllocate < 2*__prevCap {
						__capToAllocate = 2 * __prevCap
					}
					__new_slice := make([]apd.Decimal, __desiredCap, __capToAllocate)
					copy(__new_slice, toCol[:args.DestIdx])
					toCol = __new_slice
				}
				__src_slice := fromCol[args.SrcStartIdx:args.SrcEndIdx]
				__dst_slice := toCol[args.DestIdx:]
				for __i := range __src_slice {
					__dst_slice[__i].Set(&__src_slice[__i])
				}
			}
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:args.DestIdx]
			for _, selIdx := range sel {
				val := fromCol[selIdx]
				toCol = append(toCol, apd.Decimal{})
				toCol[len(toCol)-1].Set(&val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Int16:
		fromCol := args.Src.Int16()
		toCol := m.Int16()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:args.DestIdx]
			for _, selIdx := range sel {
				val := fromCol[selIdx]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Int32:
		fromCol := args.Src.Int32()
		toCol := m.Int32()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:args.DestIdx]
			for _, selIdx := range sel {
				val := fromCol[selIdx]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Int64:
		fromCol := args.Src.Int64()
		toCol := m.Int64()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:args.DestIdx]
			for _, selIdx := range sel {
				val := fromCol[selIdx]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Float64:
		fromCol := args.Src.Float64()
		toCol := m.Float64()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:args.DestIdx]
			for _, selIdx := range sel {
				val := fromCol[selIdx]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Timestamp:
		fromCol := args.Src.Timestamp()
		toCol := m.Timestamp()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:args.DestIdx]
			for _, selIdx := range sel {
				val := fromCol[selIdx]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	case coltypes.Interval:
		fromCol := args.Src.Interval()
		toCol := m.Interval()
		// NOTE: it is unfortunate that we always append whole slice without paying
		// attention to whether the values are NULL. However, if we do start paying
		// attention, the performance suffers dramatically, so we choose to copy
		// over "actual" as well as "garbage" values.
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			toCol = toCol[0:args.DestIdx]
			for _, selIdx := range sel {
				val := fromCol[selIdx]
				toCol = append(toCol, val)
			}
		}
		m.nulls.set(args)
		m.col = toCol
	default:
		panic(fmt.Sprintf("unhandled type %s", args.ColType))
	}
}

func (m *memColumn) Copy(args CopySliceArgs) {
	if !args.SelOnDest {
		// We're about to overwrite this entire range, so unset all the nulls.
		m.Nulls().UnsetNullRange(args.DestIdx, args.DestIdx+(args.SrcEndIdx-args.SrcStartIdx))
	}
	// } else {
	// SelOnDest indicates that we're applying the input selection vector as a lens
	// into the output vector as well. We'll set the non-nulls by hand below.
	// }

	switch args.ColType {
	case coltypes.Bool:
		fromCol := args.Src.Bool()
		toCol := m.Bool()
		if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull(selIdx)
						} else {
							v := fromCol[selIdx]
							m.nulls.UnsetNull(selIdx)
							toCol[selIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[selIdx] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol[selIdx]
							toCol[i+args.DestIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[i+args.DestIdx] = v
				}
			}
			return
		}
		// No Sel.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		m.nulls.set(args.SliceArgs)
	case coltypes.Bytes:
		fromCol := args.Src.Bytes()
		toCol := m.Bytes()
		if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull(selIdx)
						} else {
							v := fromCol.Get(selIdx)
							m.nulls.UnsetNull(selIdx)
							toCol.Set(selIdx, v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol.Get(selIdx)
					toCol.Set(selIdx, v)
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol.Get(selIdx)
							toCol.Set(i+args.DestIdx, v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol.Get(selIdx)
					toCol.Set(i+args.DestIdx, v)
				}
			}
			return
		}
		// No Sel.
		toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
		m.nulls.set(args.SliceArgs)
	case coltypes.Decimal:
		fromCol := args.Src.Decimal()
		toCol := m.Decimal()
		if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull(selIdx)
						} else {
							v := fromCol[selIdx]
							m.nulls.UnsetNull(selIdx)
							toCol[selIdx].Set(&v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[selIdx].Set(&v)
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol[selIdx]
							toCol[i+args.DestIdx].Set(&v)
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[i+args.DestIdx].Set(&v)
				}
			}
			return
		}
		// No Sel.
		{
			__tgt_slice := toCol[args.DestIdx:]
			__src_slice := fromCol[args.SrcStartIdx:args.SrcEndIdx]
			for __i := range __src_slice {
				__tgt_slice[__i].Set(&__src_slice[__i])
			}
		}
		m.nulls.set(args.SliceArgs)
	case coltypes.Int16:
		fromCol := args.Src.Int16()
		toCol := m.Int16()
		if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull(selIdx)
						} else {
							v := fromCol[selIdx]
							m.nulls.UnsetNull(selIdx)
							toCol[selIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[selIdx] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol[selIdx]
							toCol[i+args.DestIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[i+args.DestIdx] = v
				}
			}
			return
		}
		// No Sel.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		m.nulls.set(args.SliceArgs)
	case coltypes.Int32:
		fromCol := args.Src.Int32()
		toCol := m.Int32()
		if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull(selIdx)
						} else {
							v := fromCol[selIdx]
							m.nulls.UnsetNull(selIdx)
							toCol[selIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[selIdx] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol[selIdx]
							toCol[i+args.DestIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[i+args.DestIdx] = v
				}
			}
			return
		}
		// No Sel.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		m.nulls.set(args.SliceArgs)
	case coltypes.Int64:
		fromCol := args.Src.Int64()
		toCol := m.Int64()
		if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull(selIdx)
						} else {
							v := fromCol[selIdx]
							m.nulls.UnsetNull(selIdx)
							toCol[selIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[selIdx] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol[selIdx]
							toCol[i+args.DestIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[i+args.DestIdx] = v
				}
			}
			return
		}
		// No Sel.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		m.nulls.set(args.SliceArgs)
	case coltypes.Float64:
		fromCol := args.Src.Float64()
		toCol := m.Float64()
		if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull(selIdx)
						} else {
							v := fromCol[selIdx]
							m.nulls.UnsetNull(selIdx)
							toCol[selIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[selIdx] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol[selIdx]
							toCol[i+args.DestIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[i+args.DestIdx] = v
				}
			}
			return
		}
		// No Sel.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		m.nulls.set(args.SliceArgs)
	case coltypes.Timestamp:
		fromCol := args.Src.Timestamp()
		toCol := m.Timestamp()
		if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull(selIdx)
						} else {
							v := fromCol[selIdx]
							m.nulls.UnsetNull(selIdx)
							toCol[selIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[selIdx] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol[selIdx]
							toCol[i+args.DestIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[i+args.DestIdx] = v
				}
			}
			return
		}
		// No Sel.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		m.nulls.set(args.SliceArgs)
	case coltypes.Interval:
		fromCol := args.Src.Interval()
		toCol := m.Interval()
		if args.Sel != nil {
			sel := args.Sel
			if args.SelOnDest {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							// Remove an unused warning in some cases.
							_ = i
							m.nulls.SetNull(selIdx)
						} else {
							v := fromCol[selIdx]
							m.nulls.UnsetNull(selIdx)
							toCol[selIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[selIdx] = v
				}
			} else {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol[selIdx]
							toCol[i+args.DestIdx] = v
						}
					}
					return
				}
				// No Nulls.
				for i := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					selIdx := sel[args.SrcStartIdx+i]
					v := fromCol[selIdx]
					toCol[i+args.DestIdx] = v
				}
			}
			return
		}
		// No Sel.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		m.nulls.set(args.SliceArgs)
	default:
		panic(fmt.Sprintf("unhandled type %s", args.ColType))
	}
}

func (m *memColumn) Window(colType coltypes.T, start int, end int) Vec {
	switch colType {
	case coltypes.Bool:
		col := m.Bool()
		return &memColumn{
			t:     colType,
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Bytes:
		col := m.Bytes()
		return &memColumn{
			t:     colType,
			col:   col.Window(start, end),
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Decimal:
		col := m.Decimal()
		return &memColumn{
			t:     colType,
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Int16:
		col := m.Int16()
		return &memColumn{
			t:     colType,
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Int32:
		col := m.Int32()
		return &memColumn{
			t:     colType,
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Int64:
		col := m.Int64()
		return &memColumn{
			t:     colType,
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Float64:
		col := m.Float64()
		return &memColumn{
			t:     colType,
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Timestamp:
		col := m.Timestamp()
		return &memColumn{
			t:     colType,
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Interval:
		col := m.Interval()
		return &memColumn{
			t:     colType,
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) PrettyValueAt(colIdx int, colType coltypes.T) string {
	if m.nulls.NullAt(colIdx) {
		return "NULL"
	}
	switch colType {
	case coltypes.Bool:
		col := m.Bool()
		v := col[colIdx]
		return fmt.Sprintf("%v", v)
	case coltypes.Bytes:
		col := m.Bytes()
		v := col.Get(colIdx)
		return fmt.Sprintf("%v", v)
	case coltypes.Decimal:
		col := m.Decimal()
		v := col[colIdx]
		return fmt.Sprintf("%v", v)
	case coltypes.Int16:
		col := m.Int16()
		v := col[colIdx]
		return fmt.Sprintf("%v", v)
	case coltypes.Int32:
		col := m.Int32()
		v := col[colIdx]
		return fmt.Sprintf("%v", v)
	case coltypes.Int64:
		col := m.Int64()
		v := col[colIdx]
		return fmt.Sprintf("%v", v)
	case coltypes.Float64:
		col := m.Float64()
		v := col[colIdx]
		return fmt.Sprintf("%v", v)
	case coltypes.Timestamp:
		col := m.Timestamp()
		v := col[colIdx]
		return fmt.Sprintf("%v", v)
	case coltypes.Interval:
		col := m.Interval()
		v := col[colIdx]
		return fmt.Sprintf("%v", v)
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

// SetValueAt is an inefficient helper to set the value in a Vec when the type
// is unknown.
func SetValueAt(v Vec, elem interface{}, rowIdx int, colType coltypes.T) {
	switch colType {
	case coltypes.Bool:
		target := v.Bool()
		newVal := elem.(bool)
		target[rowIdx] = newVal
	case coltypes.Bytes:
		target := v.Bytes()
		newVal := elem.([]byte)
		target.Set(rowIdx, newVal)
	case coltypes.Decimal:
		target := v.Decimal()
		newVal := elem.(apd.Decimal)
		target[rowIdx].Set(&newVal)
	case coltypes.Int16:
		target := v.Int16()
		newVal := elem.(int16)
		target[rowIdx] = newVal
	case coltypes.Int32:
		target := v.Int32()
		newVal := elem.(int32)
		target[rowIdx] = newVal
	case coltypes.Int64:
		target := v.Int64()
		newVal := elem.(int64)
		target[rowIdx] = newVal
	case coltypes.Float64:
		target := v.Float64()
		newVal := elem.(float64)
		target[rowIdx] = newVal
	case coltypes.Timestamp:
		target := v.Timestamp()
		newVal := elem.(time.Time)
		target[rowIdx] = newVal
	case coltypes.Interval:
		target := v.Interval()
		newVal := elem.(duration.Duration)
		target[rowIdx] = newVal
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

// GetValueAt is an inefficient helper to get the value in a Vec when the type
// is unknown.
func GetValueAt(v Vec, rowIdx int, colType coltypes.T) interface{} {
	switch colType {
	case coltypes.Bool:
		target := v.Bool()
		return target[rowIdx]
	case coltypes.Bytes:
		target := v.Bytes()
		return target.Get(rowIdx)
	case coltypes.Decimal:
		target := v.Decimal()
		return target[rowIdx]
	case coltypes.Int16:
		target := v.Int16()
		return target[rowIdx]
	case coltypes.Int32:
		target := v.Int32()
		return target[rowIdx]
	case coltypes.Int64:
		target := v.Int64()
		return target[rowIdx]
	case coltypes.Float64:
		target := v.Float64()
		return target[rowIdx]
	case coltypes.Timestamp:
		target := v.Timestamp()
		return target[rowIdx]
	case coltypes.Interval:
		target := v.Interval()
		return target[rowIdx]
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}
