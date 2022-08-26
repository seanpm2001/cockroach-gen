// Copyright 2015 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package metric

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	_ "github.com/cockroachdb/cockroach/pkg/util/log" // for flags
	"github.com/kr/pretty"
	"github.com/prometheus/client_golang/prometheus"
	prometheusgo "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"
)

func testMarshal(t *testing.T, m json.Marshaler, exp string) {
	if b, err := m.MarshalJSON(); err != nil || !bytes.Equal(b, []byte(exp)) {
		t.Fatalf("unexpected: err=%v\nbytes=%s\nwanted=%s\nfor:\n%+v", err, b, exp, m)
	}
}

var emptyMetadata = Metadata{Name: ""}

func TestGauge(t *testing.T) {
	g := NewGauge(emptyMetadata)
	g.Update(10)
	if v := g.Value(); v != 10 {
		t.Fatalf("unexpected value: %d", v)
	}
	var wg sync.WaitGroup
	for i := int64(0); i < 10; i++ {
		wg.Add(2)
		go func(i int64) { g.Inc(i); wg.Done() }(i)
		go func(i int64) { g.Dec(i); wg.Done() }(i)
	}
	wg.Wait()
	if v := g.Value(); v != 10 {
		t.Fatalf("unexpected value: %d", v)
	}
	testMarshal(t, g, "10")
}

func TestFunctionalGauge(t *testing.T) {
	valToReturn := int64(10)
	g := NewFunctionalGauge(emptyMetadata, func() int64 { return valToReturn })
	if v := g.Value(); v != 10 {
		t.Fatalf("unexpected value: %d", v)
	}
	valToReturn = 15
	if v := g.Value(); v != 15 {
		t.Fatalf("unexpected value: %d", v)
	}
}

func TestGaugeFloat64(t *testing.T) {
	g := NewGaugeFloat64(emptyMetadata)
	g.Update(10.4)
	if v := g.Value(); v != 10.4 {
		t.Fatalf("unexpected value: %f", v)
	}
	testMarshal(t, g, "10.4")

	var wg sync.WaitGroup
	for i := int64(0); i < 10; i++ {
		wg.Add(2)
		go func(i int64) { g.Inc(float64(i)); wg.Done() }(i)
		go func(i int64) { g.Dec(float64(i)); wg.Done() }(i)
	}
	wg.Wait()
	if v := g.Value(); math.Abs(v-10.4) > 0.001 {
		t.Fatalf("unexpected value: %g", v)
	}
}

func TestCounter(t *testing.T) {
	c := NewCounter(emptyMetadata)
	c.Inc(90)
	if v := c.Count(); v != 90 {
		t.Fatalf("unexpected value: %d", v)
	}

	testMarshal(t, c, "90")
}

func setNow(d time.Duration) {
	now = func() time.Time {
		return time.Time{}.Add(d)
	}
}

func TestHistogramPrometheus(t *testing.T) {
	u := func(v int) *uint64 {
		n := uint64(v)
		return &n
	}

	f := func(v int) *float64 {
		n := float64(v)
		return &n
	}

	h := NewHistogram(Metadata{}, time.Hour, 10, 1)
	h.RecordValue(1)
	h.RecordValue(5)
	h.RecordValue(5)
	h.RecordValue(10)
	h.RecordValue(15000) // counts as 10
	act := *h.ToPrometheusMetric().Histogram

	expSum := float64(1*1 + 2*5 + 2*10)

	exp := prometheusgo.Histogram{
		SampleCount: u(5),
		SampleSum:   &expSum,
		Bucket: []*prometheusgo.Bucket{
			{CumulativeCount: u(1), UpperBound: f(1)},
			{CumulativeCount: u(3), UpperBound: f(5)},
			{CumulativeCount: u(5), UpperBound: f(10)},
		},
	}

	if !reflect.DeepEqual(act, exp) {
		t.Fatalf("expected differs from actual: %s", pretty.Diff(exp, act))
	}
}

func TestHistogramV2(t *testing.T) {
	u := func(v int) *uint64 {
		n := uint64(v)
		return &n
	}

	f := func(v int) *float64 {
		n := float64(v)
		return &n
	}

	h := NewHistogramV2(
		Metadata{},
		time.Hour,
		prometheus.HistogramOpts{
			Namespace:   "",
			Subsystem:   "",
			Name:        "",
			Help:        "",
			ConstLabels: nil,
			Buckets: []float64{
				1.0,
				5.0,
				10.0,
				25.0,
				100.0,
			},
		},
	)

	// should return 0 if no observations are made
	require.Equal(t, 0.0, h.ValueAtQuantileWindowed(0))

	// 200 is intentionally set us the first value to verify that the function
	// does not return NaN or Inf.
	measurements := []int64{200, 0, 4, 5, 10, 20, 25, 30, 40, 90}
	var expSum float64
	for i, m := range measurements {
		h.RecordValue(m)
		if i == 0 {
			require.Equal(t, 0.0, h.ValueAtQuantileWindowed(0))
			require.Equal(t, 100.0, h.ValueAtQuantileWindowed(99))
		}
		expSum += float64(m)
	}

	act := *h.ToPrometheusMetric().Histogram
	exp := prometheusgo.Histogram{
		SampleCount: u(len(measurements)),
		SampleSum:   &expSum,
		Bucket: []*prometheusgo.Bucket{
			{CumulativeCount: u(1), UpperBound: f(1)},
			{CumulativeCount: u(3), UpperBound: f(5)},
			{CumulativeCount: u(4), UpperBound: f(10)},
			{CumulativeCount: u(6), UpperBound: f(25)},
			{CumulativeCount: u(9), UpperBound: f(100)},
			// NB: 200 is greater than the largest defined bucket so prometheus
			// puts it in an implicit bucket with +Inf as the upper bound.
		},
	}

	if !reflect.DeepEqual(act, exp) {
		t.Fatalf("expected differs from actual: %s", pretty.Diff(exp, act))
	}

	require.Equal(t, 0.0, h.ValueAtQuantileWindowed(0))
	require.Equal(t, 1.0, h.ValueAtQuantileWindowed(10))
	require.Equal(t, 17.5, h.ValueAtQuantileWindowed(50))
	require.Equal(t, 75.0, h.ValueAtQuantileWindowed(80))
	require.Equal(t, 100.0, h.ValueAtQuantileWindowed(99.99))
}

// TestHistogramBuckets is used to generate additional prometheus buckets to be
// used with HistogramV2. Please include obs-inf in the review process of new
// buckets.
func TestHistogramBuckets(t *testing.T) {
	verifyAndPrint := func(t *testing.T, exp, act []float64) {
		t.Helper()
		var buf strings.Builder
		for idx, f := range exp {
			if idx == 0 {
				fmt.Fprintf(&buf, "// Generated via %s.", t.Name())
			}
			fmt.Fprintf(&buf, "\n%f, // %s", f, time.Duration(f))
		}
		t.Logf("%s", &buf)
		require.InDeltaSlice(t, exp, act, 1 /* delta */, "Please update the bucket boundaries for %s", t.Name())
	}
	t.Run("IOLatencyBuckets", func(t *testing.T) {
		exp := prometheus.ExponentialBucketsRange(10e3, 10e9, 15)
		verifyAndPrint(t, exp, IOLatencyBuckets)
	})

	t.Run("NetworkLatencyBuckets", func(t *testing.T) {
		exp := prometheus.ExponentialBucketsRange(500e3, 1e9, 15)
		verifyAndPrint(t, exp, NetworkLatencyBuckets)
	})
}

func TestNewHistogramV2Rotate(t *testing.T) {
	defer TestingSetNow(nil)()
	setNow(0)

	h := NewHistogramV2(emptyMetadata, 10*time.Second, prometheus.HistogramOpts{Buckets: nil})
	for i := 0; i < 4; i++ {
		// Windowed histogram is initially empty.
		h.Inspect(func(interface{}) {}) // triggers ticking
		var m prometheusgo.Metric
		require.NoError(t, h.Windowed().Write(&m))
		require.Zero(t, *m.Histogram.SampleSum)
		// But cumulative histogram has history (if i > 0).
		require.EqualValues(t, i, *h.ToPrometheusMetric().Histogram.SampleCount)

		// Add a measurement and verify it's there.
		{
			h.RecordValue(12345)
			f := float64(12345)
			var m prometheusgo.Metric
			require.NoError(t, h.Windowed().Write(&m))
			require.Equal(t, *m.Histogram.SampleSum, f)
		}
		// Tick. This rotates the histogram.
		setNow(time.Duration(i+1) * 10 * time.Second)
		// Go to beginning.
	}
}

func TestHistogramRotate(t *testing.T) {
	defer TestingSetNow(nil)()
	setNow(0)
	duration := histWrapNum * time.Second
	h := NewHistogram(emptyMetadata, duration, 1000+10*histWrapNum, 3)
	var cur time.Duration
	for i := 0; i < 3*histWrapNum; i++ {
		v := int64(10 * i)
		h.RecordValue(v)
		cur += time.Second
		setNow(cur)
		cur, windowDuration := h.Windowed()
		if windowDuration != duration {
			t.Fatalf("window changed: is %s, should be %s", windowDuration, duration)
		}

		// When i == histWrapNum-1, we expect the entry from i==0 to move out
		// of the window (since we rotated for the histWrapNum'th time).
		expMin := int64((1 + i - (histWrapNum - 1)) * 10)
		if expMin < 0 {
			expMin = 0
		}

		if min := cur.Min(); min != expMin {
			t.Fatalf("%d: unexpected minimum %d, expected %d", i, min, expMin)
		}

		if max, expMax := cur.Max(), v; max != expMax {
			t.Fatalf("%d: unexpected maximum %d, expected %d", i, max, expMax)
		}
	}
}
