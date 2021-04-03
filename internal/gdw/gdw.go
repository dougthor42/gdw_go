package gdw

import (
	"C"
	"math"
)

// Defined by SEMI M1-0302
var FlatLengths = map[float64]float64{
	50:  15.88,
	75:  22.22,
	100: 32.5,
	125: 42.5,
	150: 57.5,
}

type State int

const (
	StateOffWafer      State = iota
	StateFlat          State = iota
	StateExclusion     State = iota
	StateFlatExclusion State = iota
	StateScribe        State = iota
	StateProbe         State = iota
)

type Offset string

const (
	OffsetOdd  Offset = "odd"
	OffsetEven Offset = "even"
)

type Grid struct {
	X, Y int
}

// Go does not have classes, but you can define methods on types. Whaaaat!
//   func (w Wafer) ExclusionRadSqrd() float64 {}
//   w.ExclusionRadSqrd()
// https://tour.golang.org/methods/1
// But methods are just functions so you could also write:
//   func ExclusionRadSqrd(w Wafer) float64 {}
//   ExclusionRadSqrd(w)
func (g Grid) Row() int {
	return g.Y
}

func (g Grid) Column() int {
	return g.X
}

type Coord struct {
	X, Y float64
}

// Exactly the same as Coord, but named differently for sanity.
type Size struct {
	X, Y float64
}

// So, Go doesn't have nullable types. In order to make something "null",
// especially within a struct, use a pointer to a type.
// https://stackoverflow.com/a/51998383/1354930
type Wafer struct {
	dieSize       Size
	offset        Coord
	dia           float64
	exclusion     float64
	flatExclusion float64
	scribeY       float64
}

func (w Wafer) Radius() float64 {
	return w.dia / 2.0
}

// Exclusion Radius, squared.
func ExclusionRadSqrd(dia, excl float64) float64 {
	return math.Pow(dia/2.0, 2) + math.Pow(excl, 2) - (dia * excl)
}

// Return the flat's Y location WRT the wafer center for a given diameter.
func FlatLocation(dia float64) float64 {
	// Default case: the flat is at the wafer edge.
	var flatY float64
	flatY = -dia / 2.0

	// If the flat length is defined by SEMI-M1-0302, we can use that to
	// calculate the Y location.
	if val, ok := FlatLengths[dia]; ok {
		flatY = -math.Sqrt(math.Pow(dia/2, 2) - math.Pow(val*0.5, 2))
	}
	return flatY
}

// Calculate the distance to the furthest corner of a rectangle, squared.
func MaxDistSqrd(center Coord, size Size) float64 {
	halfX := size.X / 2.0
	halfY := size.Y / 2.0

	if center.X < 0 {
		halfX = -halfX
	}

	if center.Y < 0 {
		halfY = -halfY
	}

	dist := math.Pow(center.X+halfX, 2) + math.Pow(center.Y+halfY, 2)

	return dist
}

// Calculate how many die grids will cover the wafer.
// `dist` should be the X or Y die size.
func MaxGrid(dia float64, dist float64) int {
	return 2 * int(math.Ceil(dia/dist))
}

// Calculate the grid centerpoint.
func CenterGrid(maxGrid int, offset float64) float64 {
	return float64(maxGrid)/2.0 + offset
}

// Calculate the die state (on wafer, on edge, off wafer, etc.)
func DieState(w Wafer, grid Grid) State {
	// Calulate the die's center coordinates.
	dieCenterX := w.dieSize.X * (float64(grid.X) - CenterGrid(MaxGrid(w.dia, w.dieSize.X), w.offset.X))
	// Note: we have to reverse the Y coord (as pos. Y is down)
	dieCenterY := w.dieSize.Y * (CenterGrid(MaxGrid(w.dia, w.dieSize.Y), w.offset.Y) - float64(grid.Y))

	// Find the die's furthest point.
	center := Coord{dieCenterX, dieCenterY}
	dist := MaxDistSqrd(center, w.dieSize)

	// Determine the die's origin (lower-left corner).
	// This should be adjusted based on the plotting lib used.
	lowerLeftY := dieCenterY - (w.dieSize.Y / 2.0)

	// TODO: Can this be replaced with `switch`? I'm not sure, because
	// each case is a different check.
	if dist < math.Pow(w.Radius(), 2) {
		return StateOffWafer
	}
	if lowerLeftY < FlatLocation(w.dia) {
		return StateFlat
	}
	if dist > ExclusionRadSqrd(w.dia, w.exclusion) {
		return StateExclusion
	}
	if lowerLeftY < (FlatLocation(w.dia) + w.flatExclusion) {
		return StateFlatExclusion
	}
	if 1 < 0 {
		return StateScribe
	}

	// It passed all the checks, so it's within the probing region.
	return StateProbe
}
