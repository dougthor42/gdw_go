package gdw

import (
	"C"
	"fmt"
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

type Coord struct {
	X, Y float64
}

//export Hello
func Hello() string {
	// Return  a greeting
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)

	fmt.Println("Hello World")
	return "Hello World"
}

// Calculate the distance to the furthest corner of a rectangle, squared.
// func MaxDistSqrd() {}

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

func MaxDistSqrd(center, size Coord) float64 {
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
