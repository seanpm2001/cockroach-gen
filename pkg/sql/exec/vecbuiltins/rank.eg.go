// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package vecbuiltins

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/sql/exec"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
)

type rankInitFields struct {
	input exec.Operator
	// distinctCol is the output column of the chain of ordered distinct
	// operators in which true will indicate that a new rank needs to be assigned
	// to the corresponding tuple.
	distinctCol     []bool
	outputColIdx    int
	partitionColIdx int
}

type rankDense_false_HasPartition_false_Op struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen. It
	// is used only in case of a regular rank function (i.e. not dense).
	rankIncrement int64
}

var _ exec.Operator = &rankDense_false_HasPartition_false_Op{}

func (r *rankDense_false_HasPartition_false_Op) Init() {
	r.input.Init()
	// RANK and DENSE_RANK start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *rankDense_false_HasPartition_false_Op) Next(ctx context.Context) coldata.Batch {
	batch := r.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if r.outputColIdx == batch.Width() {
		batch.AppendCol(types.Int64)
	} else if r.outputColIdx > batch.Width() {
		panic("unexpected: column outputColIdx is neither present nor the next to be appended")
	}
	rankCol := batch.ColVec(r.outputColIdx).Int64()
	sel := batch.Selection()
	// TODO(yuzefovich): template out sel vs non-sel cases.
	if sel != nil {
		for i := uint16(0); i < batch.Length(); i++ {
			if r.distinctCol[sel[i]] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[sel[i]] = r.rank
			} else {
				rankCol[sel[i]] = r.rank
				r.rankIncrement++
			}
		}
	} else {
		for i := uint16(0); i < batch.Length(); i++ {
			if r.distinctCol[i] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank
				r.rankIncrement++
			}
		}
	}
	return batch
}

type rankDense_false_HasPartition_true_Op struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen. It
	// is used only in case of a regular rank function (i.e. not dense).
	rankIncrement int64
}

var _ exec.Operator = &rankDense_false_HasPartition_true_Op{}

func (r *rankDense_false_HasPartition_true_Op) Init() {
	r.input.Init()
	// RANK and DENSE_RANK start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *rankDense_false_HasPartition_true_Op) Next(ctx context.Context) coldata.Batch {
	batch := r.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if r.partitionColIdx == batch.Width() {
		batch.AppendCol(types.Bool)
	} else if r.partitionColIdx > batch.Width() {
		panic("unexpected: column partitionColIdx is neither present nor the next to be appended")
	}
	partitionCol := batch.ColVec(r.partitionColIdx).Bool()

	if r.outputColIdx == batch.Width() {
		batch.AppendCol(types.Int64)
	} else if r.outputColIdx > batch.Width() {
		panic("unexpected: column outputColIdx is neither present nor the next to be appended")
	}
	rankCol := batch.ColVec(r.outputColIdx).Int64()
	sel := batch.Selection()
	// TODO(yuzefovich): template out sel vs non-sel cases.
	if sel != nil {
		for i := uint16(0); i < batch.Length(); i++ {
			if partitionCol[sel[i]] {
				r.rank = 1
				r.rankIncrement = 1
				rankCol[sel[i]] = 1
				continue
			}
			if r.distinctCol[sel[i]] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[sel[i]] = r.rank
			} else {
				rankCol[sel[i]] = r.rank
				r.rankIncrement++
			}
		}
	} else {
		for i := uint16(0); i < batch.Length(); i++ {
			if partitionCol[i] {
				r.rank = 1
				r.rankIncrement = 1
				rankCol[i] = 1
				continue
			}
			if r.distinctCol[i] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank
				r.rankIncrement++
			}
		}
	}
	return batch
}

type rankDense_true_HasPartition_false_Op struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen. It
	// is used only in case of a regular rank function (i.e. not dense).
	rankIncrement int64
}

var _ exec.Operator = &rankDense_true_HasPartition_false_Op{}

func (r *rankDense_true_HasPartition_false_Op) Init() {
	r.input.Init()
	// RANK and DENSE_RANK start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *rankDense_true_HasPartition_false_Op) Next(ctx context.Context) coldata.Batch {
	batch := r.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if r.outputColIdx == batch.Width() {
		batch.AppendCol(types.Int64)
	} else if r.outputColIdx > batch.Width() {
		panic("unexpected: column outputColIdx is neither present nor the next to be appended")
	}
	rankCol := batch.ColVec(r.outputColIdx).Int64()
	sel := batch.Selection()
	// TODO(yuzefovich): template out sel vs non-sel cases.
	if sel != nil {
		for i := uint16(0); i < batch.Length(); i++ {
			if r.distinctCol[sel[i]] {
				r.rank++
				rankCol[sel[i]] = r.rank
			} else {
				rankCol[sel[i]] = r.rank

			}
		}
	} else {
		for i := uint16(0); i < batch.Length(); i++ {
			if r.distinctCol[i] {
				r.rank++
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank

			}
		}
	}
	return batch
}

type rankDense_true_HasPartition_true_Op struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen. It
	// is used only in case of a regular rank function (i.e. not dense).
	rankIncrement int64
}

var _ exec.Operator = &rankDense_true_HasPartition_true_Op{}

func (r *rankDense_true_HasPartition_true_Op) Init() {
	r.input.Init()
	// RANK and DENSE_RANK start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *rankDense_true_HasPartition_true_Op) Next(ctx context.Context) coldata.Batch {
	batch := r.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if r.partitionColIdx == batch.Width() {
		batch.AppendCol(types.Bool)
	} else if r.partitionColIdx > batch.Width() {
		panic("unexpected: column partitionColIdx is neither present nor the next to be appended")
	}
	partitionCol := batch.ColVec(r.partitionColIdx).Bool()

	if r.outputColIdx == batch.Width() {
		batch.AppendCol(types.Int64)
	} else if r.outputColIdx > batch.Width() {
		panic("unexpected: column outputColIdx is neither present nor the next to be appended")
	}
	rankCol := batch.ColVec(r.outputColIdx).Int64()
	sel := batch.Selection()
	// TODO(yuzefovich): template out sel vs non-sel cases.
	if sel != nil {
		for i := uint16(0); i < batch.Length(); i++ {
			if partitionCol[sel[i]] {
				r.rank = 1
				r.rankIncrement = 1
				rankCol[sel[i]] = 1
				continue
			}
			if r.distinctCol[sel[i]] {
				r.rank++
				rankCol[sel[i]] = r.rank
			} else {
				rankCol[sel[i]] = r.rank

			}
		}
	} else {
		for i := uint16(0); i < batch.Length(); i++ {
			if partitionCol[i] {
				r.rank = 1
				r.rankIncrement = 1
				rankCol[i] = 1
				continue
			}
			if r.distinctCol[i] {
				r.rank++
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank

			}
		}
	}
	return batch
}
