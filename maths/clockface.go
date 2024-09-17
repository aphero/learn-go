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
	return angleToPoint(secondsInRadians(t))
}

// Minutes

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) +
		(math.Pi / (6 / float64(t.Hour()%12)))
}

// WHERE DO THESE GO?

// func AngleInRadians(t time.Time) float64 {
// 	// Here we're converting the seconds to radians and dividing by 60 to get the one minute adjustment
// 	// for the minute hand.  Not dividing by 60 gives us the angle of the second hand.
// 	return (secondsInRadians(t) / 60) +
// 		// Then we divide 30 by the minute we're rendering so we know how many radians we're dealing with
// 		// and then divide PI by the result to get the actual radians that our angle is at, just like we
// 		// do for the second hand calculation.
// 		(math.Pi / (30 / float64(t.Minute())))
// }

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
