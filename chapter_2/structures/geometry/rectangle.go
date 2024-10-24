package geometry

import "math"

type Circle struct {
	X float64
	Y float64
	R float64
}

func CircleArea(c Circle) float64 {
	return math.Pi * c.R * c.R
}
