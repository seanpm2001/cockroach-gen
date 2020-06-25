// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"fmt"
	"time"

	"github.com/cockroachdb/apd/v2"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
)

// EncDatumRowsToColVec converts one column from EncDatumRows to a column
// vector. columnIdx is the 0-based index of the column in the EncDatumRows.
func EncDatumRowsToColVec(
	allocator *colmem.Allocator,
	rows sqlbase.EncDatumRows,
	vec coldata.Vec,
	columnIdx int,
	t *types.T,
	alloc *sqlbase.DatumAlloc,
) error {
	var err error
	allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			switch typeconv.TypeFamilyToCanonicalTypeFamily(t.Family()) {
			case types.BoolFamily:
				switch t.Width() {
				case -1:
				default:
					col := vec.Bool()
					datumToPhysicalFn := GetDatumToPhysicalFn(t)
					var v interface{}
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(t, alloc); err != nil {
								return
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(i)
						} else {
							v, err = datumToPhysicalFn(datum)
							if err != nil {
								return
							}

							castV := v.(bool)
							col[i] = castV
						}
					}
				}
			case types.BytesFamily:
				switch t.Width() {
				case -1:
				default:
					col := vec.Bytes()
					datumToPhysicalFn := GetDatumToPhysicalFn(t)
					var v interface{}
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(t, alloc); err != nil {
								return
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(i)
						} else {
							v, err = datumToPhysicalFn(datum)
							if err != nil {
								return
							}

							castV := v.([]byte)
							col.Set(i, castV)
						}
					}
				}
			case types.DecimalFamily:
				switch t.Width() {
				case -1:
				default:
					col := vec.Decimal()
					datumToPhysicalFn := GetDatumToPhysicalFn(t)
					var v interface{}
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(t, alloc); err != nil {
								return
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(i)
						} else {
							v, err = datumToPhysicalFn(datum)
							if err != nil {
								return
							}

							castV := v.(apd.Decimal)
							col[i].Set(&castV)
						}
					}
				}
			case types.IntFamily:
				switch t.Width() {
				case 16:
					col := vec.Int16()
					datumToPhysicalFn := GetDatumToPhysicalFn(t)
					var v interface{}
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(t, alloc); err != nil {
								return
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(i)
						} else {
							v, err = datumToPhysicalFn(datum)
							if err != nil {
								return
							}

							castV := v.(int16)
							col[i] = castV
						}
					}
				case 32:
					col := vec.Int32()
					datumToPhysicalFn := GetDatumToPhysicalFn(t)
					var v interface{}
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(t, alloc); err != nil {
								return
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(i)
						} else {
							v, err = datumToPhysicalFn(datum)
							if err != nil {
								return
							}

							castV := v.(int32)
							col[i] = castV
						}
					}
				case -1:
				default:
					col := vec.Int64()
					datumToPhysicalFn := GetDatumToPhysicalFn(t)
					var v interface{}
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(t, alloc); err != nil {
								return
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(i)
						} else {
							v, err = datumToPhysicalFn(datum)
							if err != nil {
								return
							}

							castV := v.(int64)
							col[i] = castV
						}
					}
				}
			case types.FloatFamily:
				switch t.Width() {
				case -1:
				default:
					col := vec.Float64()
					datumToPhysicalFn := GetDatumToPhysicalFn(t)
					var v interface{}
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(t, alloc); err != nil {
								return
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(i)
						} else {
							v, err = datumToPhysicalFn(datum)
							if err != nil {
								return
							}

							castV := v.(float64)
							col[i] = castV
						}
					}
				}
			case types.TimestampTZFamily:
				switch t.Width() {
				case -1:
				default:
					col := vec.Timestamp()
					datumToPhysicalFn := GetDatumToPhysicalFn(t)
					var v interface{}
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(t, alloc); err != nil {
								return
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(i)
						} else {
							v, err = datumToPhysicalFn(datum)
							if err != nil {
								return
							}

							castV := v.(time.Time)
							col[i] = castV
						}
					}
				}
			case types.IntervalFamily:
				switch t.Width() {
				case -1:
				default:
					col := vec.Interval()
					datumToPhysicalFn := GetDatumToPhysicalFn(t)
					var v interface{}
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(t, alloc); err != nil {
								return
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(i)
						} else {
							v, err = datumToPhysicalFn(datum)
							if err != nil {
								return
							}

							castV := v.(duration.Duration)
							col[i] = castV
						}
					}
				}
			case typeconv.DatumVecCanonicalTypeFamily:
				switch t.Width() {
				case -1:
				default:
					col := vec.Datum()
					datumToPhysicalFn := GetDatumToPhysicalFn(t)
					var v interface{}
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(t, alloc); err != nil {
								return
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(i)
						} else {
							v, err = datumToPhysicalFn(datum)
							if err != nil {
								return
							}

							castV := v.(interface{})
							col.Set(i, castV)
						}
					}
				}
			default:
				colexecerror.InternalError(fmt.Sprintf("unsupported type %s", t))
			}
		},
	)
	return err
}
