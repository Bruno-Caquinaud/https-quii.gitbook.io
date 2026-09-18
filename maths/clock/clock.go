package clock

import (
	"time"

	"github.com/Bruno-Caquinaud/https-quii.gitbook.io/maths/clockoperating"
)

type Clock struct {
	Center clockoperating.Point
	Radius int

	Second clockoperating.Point
	Minute clockoperating.Point
	Hour   clockoperating.Point
}

func (c *Clock) ComputeCoordinates(tm time.Time) {
	timeDisplayed := []time.Duration{time.Second, time.Minute, time.Hour}

	for _, timeUnit := range timeDisplayed {
		radian := clockoperating.TimeToRadian(timeUnit, tm)
		coordinates := clockoperating.RadiantoCoordinate(radian, float64(c.Radius), c.Center)

		switch timeUnit {
		case time.Second:
			c.Second = coordinates
		case time.Minute:
			c.Minute = coordinates
		case time.Hour:
			c.Hour = coordinates
		}
	}
}
