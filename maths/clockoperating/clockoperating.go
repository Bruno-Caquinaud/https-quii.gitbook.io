package clockoperating

import (
	"math"
	"time"
)

type Point struct {
	X int
	Y int
}

func TimeToRadian(unit time.Duration, tm time.Time) (radian float64) {
	timeBase, timeAmount := 0.0, 0.0

	switch unit {
	case time.Second:
		timeAmount = float64(tm.Second())
		timeBase = 60.0
	case time.Minute:
		timeAmount = float64(tm.Minute())
		timeBase = 60.0
	case time.Hour:
		timeAmount = float64(tm.Hour())
		timeBase = 12
	}

	return (2 * math.Pi / timeBase) * timeAmount
}

func RadiantoCoordinate(radian float64, radius float64, offset Point) Point {
	x := float64(offset.X) + math.Sin(radian)*float64(radius)
	y := float64(offset.Y) - math.Cos(radian)*float64(radius)

	return Point{int(x), int(y)}
}
