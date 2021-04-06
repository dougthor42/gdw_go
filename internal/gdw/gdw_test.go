package gdw

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const tolerance = 0.00001

func TestFlatLocation(t *testing.T) {
	tables := []struct {
		diameter float64
		want     float64
	}{
		{50, -23.7056196},
		{75, -35.8164473},
		{100, -47.2857008},
		{150, -69.2707550},
		{35, -17.5},
		{120, -60},
		{237.68, -118.84},
	}

	// Taken from the docs https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal
	// with help from https://dev.to/juliaferraioli/testing-in-go-testing-floating-point-numbers-4i0a
	// TODO: How to make this DRY?
	opt := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		if math.IsNaN(delta / mean) {
			return true
		}
		return (delta / mean) < tolerance
	})

	for _, tt := range tables {
		testName := fmt.Sprintf("%f", tt.diameter)
		t.Run(testName, func(t *testing.T) {
			got := FlatLocation(tt.diameter)
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("Got %f, wanted %f", got, tt.want)
			}
		})
	}
}

func TestMaxDistSqrd(t *testing.T) {
	tables := []struct {
		center Coord
		size   Size
		want   float64
	}{
		{Coord{0, 0}, Size{2, 2}, 2},
		{Coord{0, 0}, Size{6, 8}, 25},
		{Coord{0, 0}, Size{2, 36}, 325},
		{Coord{0, 0}, Size{0, 0}, 0},
		{Coord{0.5, 0.5}, Size{1, 1}, 2},
		{Coord{0, 0}, Size{3.14, 2.718}, 4.311781},
		{Coord{0, -10}, Size{3.14, 2.718}, 131.491781},
		{Coord{-10, 0}, Size{3.14, 2.718}, 135.711781},
		{Coord{-10, -10}, Size{3.14, 2.718}, 262.891781},
		{Coord{0, 10}, Size{3.14, 2.718}, 131.491781},
		{Coord{10, 0}, Size{3.14, 2.718}, 135.711781},
		{Coord{10, 10}, Size{3.14, 2.718}, 262.891781},
		{Coord{100000, 100000}, Size{2, 2}, 20000400002},
		{Coord{1000, 0}, Size{100, 0.00001}, 1102500},
	}

	// Taken from the docs https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal
	// with help from https://dev.to/juliaferraioli/testing-in-go-testing-floating-point-numbers-4i0a
	// TODO: How to make this DRY?
	opt := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		if math.IsNaN(delta / mean) {
			return true
		}
		return (delta / mean) < tolerance
	})

	for _, tt := range tables {
		testName := "foo"
		t.Run(testName, func(t *testing.T) {
			got := MaxDistSqrd(tt.center, tt.size)
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("Got %f, wanted %f.", got, tt.want)
			}
		})
	}
}

func TestExclusionRadSqrd(t *testing.T) {
	tables := []struct {
		dia  float64
		excl float64
		want float64
	}{
		{150, 5, 4900},
	}

	// Taken from the docs https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal
	// with help from https://dev.to/juliaferraioli/testing-in-go-testing-floating-point-numbers-4i0a
	// TODO: How to make this DRY?
	opt := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		if math.IsNaN(delta / mean) {
			return true
		}
		return (delta / mean) < tolerance
	})

	for _, tt := range tables {
		testName := "foo"
		t.Run(testName, func(t *testing.T) {
			got := ExclusionRadSqrd(tt.dia, tt.excl)
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("Got %f, wanted %f.", got, tt.want)
			}
		})
	}
}

var DummyWafer Wafer = Wafer{Size{5, 5}, Coord{0, 0}, 150, 4.5, 4.5, 70.2}

func TestDieState(t *testing.T) {
	tables := []struct {
		w    Wafer
		g    Grid
		want State
	}{
		{DummyWafer, Grid{21, 17}, StateOffWafer},
		{Wafer{Size{5, 5}, Coord{0, 0}, 150, 4.5, 4.5, 70.2}, Grid{30, 30}, StateProbe},
		{Wafer{Size{5, 5}, Coord{0, 0}, 150, 4.5, 4.5, 70.2}, Grid{28, 43}, StateFlatExclusion},
		{Wafer{Size{5, 5}, Coord{0, 0}, 150, 4.5, 4.5, 70.2}, Grid{31, 44}, StateFlat},
		{Wafer{Size{5, 5}, Coord{0, 0}, 150, 4.5, 4.5, 70.2}, Grid{40, 21}, StateExclusion},
	}

	for _, tt := range tables {
		testName := "foo"
		t.Run(testName, func(t *testing.T) {
			got := DieState(tt.w, tt.g)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Got %q, wanted %q.", got, tt.want)
			}
		})
	}
}

func TestMaxGrid(t *testing.T) {
	tables := []struct {
		dia, dist float64
		want      int
	}{
		{0, 0, 0},
		{50, 50, 2},
		{50, 10, 10},
	}

	for _, tt := range tables {
		testName := "foo"
		t.Run(testName, func(t *testing.T) {
			got := MaxGrid(tt.dia, tt.dist)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Got %d, wanted %d", got, tt.want)
			}
		})
	}
}

func TestCenterGrid(t *testing.T) {
	tables := []struct {
		maxGrid int
		offset  float64
		want    float64
	}{
		{0, 0, 0},
		{50, 50, 75},
		{50, 10, 35},
	}

	for _, tt := range tables {
		testName := "foo"
		t.Run(testName, func(t *testing.T) {
			got := CenterGrid(tt.maxGrid, tt.offset)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Got %f, wanted %f", got, tt.want)
			}
		})
	}
}

func TestGrossDiePerWafer(t *testing.T) {
	tables := []struct {
		name            string
		size            Size
		dia             int
		offsetXY        OffsetXY
		exclusion       float64
		scribeExclusion float64
		want            int
	}{
		{"ints", Size{5, 5}, 150, OffsetXY{OffsetEven, OffsetEven}, 5, 5, 546},
		{"floats", Size{5.0, 5.0}, 150, OffsetXY{OffsetEven, OffsetEven}, 5, 5, 546},
		{"t01", Size{3.34, 3.16}, 100, OffsetXY{OffsetEven, OffsetEven}, 5, 5, 548},
		{"t02-1", Size{2.43, 3.30}, 150, OffsetXY{OffsetEven, OffsetOdd}, 5, 4.5, 1814},
		{"t02-2", Size{2.43, 3.30}, 150, OffsetXY{OffsetEven, OffsetEven}, 5, 4.5, 1794},
		{"t02-3", Size{2.43, 3.30}, 150, OffsetXY{OffsetOdd, OffsetOdd}, 5, 4.5, 1800},
		{"t02-4", Size{3.34, 3.16}, 150, OffsetXY{OffsetOdd, OffsetEven}, 5, 4.5, 1804},
		{"t03", Size{4.34, 6.44}, 150, OffsetXY{OffsetEven, OffsetEven}, 5, 5, 484},
		{"t04", Size{1, 1}, 150, OffsetXY{OffsetEven, OffsetEven}, 5, 5, 14902},
		{"t05", Size{1, 1}, 200, OffsetXY{OffsetOdd, OffsetEven}, 5, 15, 27435},
		// {"t06", Size{2.9, 3.3}, 150, -1.65, 2.95, 5, 4.5, 1529},
		// {"t07", Size{2.69, 1.65}, 150, 1.345, 2.1, 5, 4.5, 3346},
		// {"t08", Size{4.4, 5.02}, 150, 0, -0.2, 5, 4.5, 648},
	}

	for _, tt := range tables {
		testName := tt.name
		t.Run(testName, func(t *testing.T) {
			dieList := GrossDiePerWafer(tt.size, tt.dia, tt.offsetXY, tt.exclusion, tt.scribeExclusion)

			// This test only verifies how many Probed die are
			// calculated.
			got := 0
			for _, d := range dieList {
				if d.state == StateProbe {
					got += 1
				}
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("Got %d, wanted %d", got, tt.want)
			}

		})
	}
}
