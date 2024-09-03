package clockface

import (
	"math"
	"time"
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip - so the clock rotates clockwise
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         // translate - move our circle into position
	return p
}

func SecondHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}
