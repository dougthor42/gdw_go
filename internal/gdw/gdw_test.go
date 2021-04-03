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
		size   Coord
		want   float64
	}{
		{Coord{0, 0}, Coord{2, 2}, 2},
		{Coord{0, 0}, Coord{6, 8}, 25},
		{Coord{0, 0}, Coord{2, 36}, 325},
		{Coord{0, 0}, Coord{0, 0}, 0},
		{Coord{0.5, 0.5}, Coord{1, 1}, 2},
		{Coord{0, 0}, Coord{3.14, 2.718}, 4.311781},
		{Coord{0, -10}, Coord{3.14, 2.718}, 131.491781},
		{Coord{-10, 0}, Coord{3.14, 2.718}, 135.711781},
		{Coord{-10, -10}, Coord{3.14, 2.718}, 262.891781},
		{Coord{0, 10}, Coord{3.14, 2.718}, 131.491781},
		{Coord{10, 0}, Coord{3.14, 2.718}, 135.711781},
		{Coord{10, 10}, Coord{3.14, 2.718}, 262.891781},
		{Coord{100000, 100000}, Coord{2, 2}, 20000400002},
		{Coord{1000, 0}, Coord{100, 0.00001}, 1102500},
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


func TestDieState(t *testing.T) {
	tables := []struct {
		w Wafer
		g Grid
		want State
	}{
		{Wafer{}, Grid{}, StateOffWafer},
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
