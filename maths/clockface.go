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

// Seconds

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

// Minutes

func minuteHandPoint(t time.Time) Point {
	return AngleToPoint(minutesInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Minute())))
}

// WHERE DO THESE GO?

func AngleInRadians(t time.Time) float64 {
	// Here we're converting the seconds to radians and dividing by 60 to get the one minute adjustment
	// for the minute hand.  Not dividing by 60 gives us the angle of the second hand.
	return (secondsInRadians(t) / 60) +
		// Then we divide 30 by the minute we're rendering so we know how many radians we're dealing with
		// and then divide PI by the result to get the actual radians that our angle is at, just like we
		// do for the second hand calculation.
		(math.Pi / (30 / float64(t.Minute())))
}

func AngleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
