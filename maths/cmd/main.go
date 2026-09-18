package main

import (
	"fmt"
	"time"

	"github.com/Bruno-Caquinaud/https-quii.gitbook.io/maths/clock"
	"github.com/Bruno-Caquinaud/https-quii.gitbook.io/maths/clockface"
	"github.com/Bruno-Caquinaud/https-quii.gitbook.io/maths/clockoperating"
)

func main() {
	tm := time.Date(2022, 12, 13, 18, 15, 45, 00, time.UTC)
	clk := clock.Clock{
		Center: clockoperating.Point{X: 150, Y: 150},
		Radius: 90,
	}

	clk.ComputeCoordinates(tm)
	svg := clockface.SvgRendering(clk)
	fmt.Printf("%s", svg)
}
