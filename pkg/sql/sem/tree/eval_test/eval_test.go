// Copyright 2015 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package eval_test

import (
	"context"
	gosql "database/sql"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/settings/cluster"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/colbuilder"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfra"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfrapb"
	"github.com/cockroachdb/cockroach/pkg/sql/parser"
	"github.com/cockroachdb/cockroach/pkg/sql/rowexec"
	_ "github.com/cockroachdb/cockroach/pkg/sql/sem/builtins"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/testutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/datadriven"
	"github.com/stretchr/testify/require"
)

func TestEval(t *testing.T) {
	defer leaktest.AfterTest(t)()
	ctx := context.Background()
	evalCtx := tree.NewTestingEvalContext(cluster.MakeTestingClusterSettings())
	defer evalCtx.Stop(ctx)

	walk := func(t *testing.T, getExpr func(*testing.T, *datadriven.TestData) string) {
		datadriven.Walk(t, filepath.Join("../testdata", "eval"), func(t *testing.T, path string) {
			datadriven.RunTest(t, path, func(t *testing.T, d *datadriven.TestData) string {
				if d.Cmd != "eval" {
					t.Fatalf("unsupported command %s", d.Cmd)
				}
				return getExpr(t, d) + "\n"
			})
		})
	}

	// The opt and no-opt tests don't do an end-to-end SQL test. Do that
	// here by executing a SELECT. In order to make the output be the same
	// we have to also figure out what the expected output type is so we
	// can correctly format the datum.
	t.Run("sql", func(t *testing.T) {
		s, sqlDB, _ := serverutils.StartServer(t, base.TestServerArgs{})
		defer s.Stopper().Stop(ctx)

		walk(t, func(t *testing.T, d *datadriven.TestData) string {
			var res gosql.NullString
			if err := sqlDB.QueryRow(fmt.Sprintf("SELECT (%s)::STRING", d.Input)).Scan(&res); err != nil {
				return strings.TrimPrefix(err.Error(), "pq: ")
			}
			if !res.Valid {
				return "NULL"
			}

			// We have a non-null result. We can't just return
			// res.String here because these strings don't
			// match the datum.String() representations. For
			// example, a bitarray has a res.String of something
			// like `1001001` but the datum representation is
			// `B'1001001'`. Thus we have to parse res.String (a
			// SQL result) back into a datum and return that.

			expr, err := parser.ParseExpr(d.Input)
			if err != nil {
				t.Fatal(err)
			}
			// expr.TypeCheck to avoid constant folding.
			semaCtx := tree.MakeSemaContext()
			typedExpr, err := expr.TypeCheck(ctx, &semaCtx, types.Any)
			if err != nil {
				// An error here should have been found above by QueryRow.
				t.Fatal(err)
			}

			switch typedExpr.ResolvedType().Family() {
			case types.TupleFamily:
				// ParseAndRequireString doesn't handle tuples, so we have to convert them ourselves.
				var datums tree.Datums
				// Fetch the original expression's tuple values.
				tuple := typedExpr.(*tree.Tuple)
				for i, s := range strings.Split(res.String[1:len(res.String)-1], ",") {
					if s == "" {
						continue
					}
					// Figure out the type of the tuple value.
					expr, err := tuple.Exprs[i].TypeCheck(ctx, &semaCtx, types.Any)
					if err != nil {
						t.Fatal(err)
					}
					// Now parse the new string as the expected type.
					datum, err := tree.ParseAndRequireString(expr.ResolvedType(), s, evalCtx)
					if err != nil {
						t.Errorf("%s: %s", err, s)
						return err.Error()
					}
					datums = append(datums, datum)
				}
				return tree.NewDTuple(typedExpr.ResolvedType(), datums...).String()
			}
			datum, err := tree.ParseAndRequireString(typedExpr.ResolvedType(), res.String, evalCtx)
			if err != nil {
				t.Errorf("%s: %s", err, res.String)
				return err.Error()
			}
			return datum.String()
		})
	})

	t.Run("vectorized", func(t *testing.T) {
		walk(t, func(t *testing.T, d *datadriven.TestData) string {
			if d.Input == "B'11111111111111111111111110000101'::int4" {
				// Skip this test: https://github.com/cockroachdb/cockroach/pull/40790#issuecomment-532597294.
				return strings.TrimSpace(d.Expected)
			}
			flowCtx := &execinfra.FlowCtx{
				EvalCtx: evalCtx,
			}
			memMonitor := execinfra.NewTestMemMonitor(ctx, cluster.MakeTestingClusterSettings())
			defer memMonitor.Stop(ctx)
			acc := memMonitor.MakeBoundAccount()
			defer acc.Close(ctx)
			expr, err := parser.ParseExpr(d.Input)
			require.NoError(t, err)
			if _, ok := expr.(*tree.RangeCond); ok {
				// RangeCond gets normalized to comparison expressions and its Eval
				// method returns an error, so skip it for execution.
				return strings.TrimSpace(d.Expected)
			}
			semaCtx := tree.MakeSemaContext()
			typedExpr, err := expr.TypeCheck(ctx, &semaCtx, types.Any)
			if err != nil {
				// Skip this test as it's testing an expected error which would be
				// caught before execution.
				return strings.TrimSpace(d.Expected)
			}
			typs := []*types.T{typedExpr.ResolvedType()}

			// inputTyps has no relation to the actual expression result type. Used
			// for generating a batch.
			inputTyps := []*types.T{types.Int}

			batchesReturned := 0
			args := colexec.NewColOperatorArgs{
				Spec: &execinfrapb.ProcessorSpec{
					Input: []execinfrapb.InputSyncSpec{{
						Type:        execinfrapb.InputSyncSpec_UNORDERED,
						ColumnTypes: inputTyps,
					}},
					Core: execinfrapb.ProcessorCoreUnion{
						Noop: &execinfrapb.NoopCoreSpec{},
					},
					Post: execinfrapb.PostProcessSpec{
						RenderExprs: []execinfrapb.Expression{{Expr: d.Input}},
					},
				},
				Inputs: []colexecbase.Operator{
					&colexecbase.CallbackOperator{
						NextCb: func(_ context.Context) coldata.Batch {
							if batchesReturned > 0 {
								return coldata.ZeroBatch
							}
							// It doesn't matter what types we create the input batch with.
							batch := coldata.NewMemBatch(inputTyps, coldata.StandardColumnFactory)
							batch.SetLength(1)
							batchesReturned++
							return batch
						},
					},
				},
				StreamingMemAccount: &acc,
				// Unsupported post processing specs are wrapped and run through the
				// row execution engine.
				ProcessorConstructor: rowexec.NewProcessor,
			}
			args.TestingKnobs.UseStreamingMemAccountForBuffering = true
			result, err := colbuilder.NewColOperator(ctx, flowCtx, args)
			if testutils.IsError(err, "unsupported type") {
				// Skip this test as execution is not supported by the vectorized
				// engine.
				return strings.TrimSpace(d.Expected)
			} else {
				require.NoError(t, err)
			}

			mat, err := colexec.NewMaterializer(
				flowCtx,
				0, /* processorID */
				result.Op,
				typs,
				nil, /* output */
				result.MetadataSources,
				nil, /* toClose */
				nil, /* outputStatsToTrace */
				nil, /* cancelFlow */
			)
			require.NoError(t, err)

			var (
				row  sqlbase.EncDatumRow
				meta *execinfrapb.ProducerMetadata
			)
			ctx = mat.Start(ctx)
			row, meta = mat.Next()
			if meta != nil {
				if meta.Err != nil {
					return fmt.Sprint(meta.Err)
				}
				t.Fatalf("unexpected metadata: %+v", meta)
			}
			if row == nil {
				t.Fatal("unexpected end of input")
			}
			return row[0].Datum.String()
		})
	})
}
