// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package exec

import (
	"fmt"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/types/conv"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	semtypes "github.com/cockroachdb/cockroach/pkg/sql/types"
)

// EncDatumRowsToColVec converts one column from EncDatumRows to a column
// vector. columnIdx is the 0-based index of the column in the EncDatumRows.
func EncDatumRowsToColVec(
	rows sqlbase.EncDatumRows,
	vec coldata.Vec,
	columnIdx int,
	columnType *semtypes.T,
	alloc *sqlbase.DatumAlloc,
) error {

	switch columnType.Family() {
	case semtypes.StringFamily:

		col := vec.Bytes()
		datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.([]byte)
			}
		}
	case semtypes.DateFamily:

		col := vec.Int64()
		datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int64)
			}
		}
	case semtypes.DecimalFamily:

		col := vec.Decimal()
		datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(apd.Decimal)
			}
		}
	case semtypes.FloatFamily:
		switch columnType.Width() {
		case 64:

			col := vec.Float64()
			datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
			for i := range rows {
				row := rows[i]
				if row[columnIdx].Datum == nil {
					if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := row[columnIdx].Datum
				if datum == tree.DNull {
					vec.Nulls().SetNull(uint16(i))
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}
					col[i] = v.(float64)
				}
			}
		case 32:

			col := vec.Float64()
			datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
			for i := range rows {
				row := rows[i]
				if row[columnIdx].Datum == nil {
					if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := row[columnIdx].Datum
				if datum == tree.DNull {
					vec.Nulls().SetNull(uint16(i))
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}
					col[i] = v.(float64)
				}
			}
		default:
			panic(fmt.Sprintf("unsupported width %d for column type %s", columnType.Width(), columnType.String()))
		}
	case semtypes.BoolFamily:

		col := vec.Bool()
		datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(bool)
			}
		}
	case semtypes.IntFamily:
		switch columnType.Width() {
		case 64:

			col := vec.Int64()
			datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
			for i := range rows {
				row := rows[i]
				if row[columnIdx].Datum == nil {
					if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := row[columnIdx].Datum
				if datum == tree.DNull {
					vec.Nulls().SetNull(uint16(i))
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}
					col[i] = v.(int64)
				}
			}
		case 16:

			col := vec.Int16()
			datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
			for i := range rows {
				row := rows[i]
				if row[columnIdx].Datum == nil {
					if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := row[columnIdx].Datum
				if datum == tree.DNull {
					vec.Nulls().SetNull(uint16(i))
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}
					col[i] = v.(int16)
				}
			}
		case 32:

			col := vec.Int32()
			datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
			for i := range rows {
				row := rows[i]
				if row[columnIdx].Datum == nil {
					if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := row[columnIdx].Datum
				if datum == tree.DNull {
					vec.Nulls().SetNull(uint16(i))
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}
					col[i] = v.(int32)
				}
			}
		default:
			panic(fmt.Sprintf("unsupported width %d for column type %s", columnType.Width(), columnType.String()))
		}
	case semtypes.OidFamily:

		col := vec.Int64()
		datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int64)
			}
		}
	case semtypes.BytesFamily:

		col := vec.Bytes()
		datumToPhysicalFn := conv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.([]byte)
			}
		}
	default:
		panic(fmt.Sprintf("unsupported column type %s", columnType.String()))
	}
	return nil
}
