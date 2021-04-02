package gdwcalc

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
	150: 75.5,
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
func FlatLocation(diameter float64) float64 {
	// We must cast diameter to a float64 to do the math we want.
	var dia float64 = float64(diameter)
	// Default case: the flat is at the wafer edge.
	var flatY float64
	flatY = -dia / 2

	// If the flat length is defined by SEMI-M1-0302, we can use that to
	// calculate the Y location.
	if val, ok := FlatLengths[diameter]; ok {
		flatY = -math.Sqrt(math.Pow(dia/2, 2) - math.Pow(val*0.5, 2))
	}
	return flatY

}
