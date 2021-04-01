package gdwcalc

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello World"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestFlatLocation(t *testing.T) {
	tables := []struct {
		diameter float64
		want     float64
	}{
		{50, -23.70562},
		{75, -35.8164473},
		{100, -47.2857008},
		{150, -69.2707550},
		{35, -17.5},
		{120, -60},
		{237.68, -118.84},
	}

	for _, tt := range tables {
		testName := fmt.Sprintf("%f", tt.diameter)
		t.Run(testName, func(t *testing.T) {
			got := FlatLocation(tt.diameter)
			if got != tt.want {
				t.Errorf("Got %f, wanted %f", got, tt.want)
			}
		})
	}
}
