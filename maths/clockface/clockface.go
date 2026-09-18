package clockface

import (
	"fmt"
	"strings"

	"github.com/Bruno-Caquinaud/https-quii.gitbook.io/maths/clock"
	"github.com/Bruno-Caquinaud/https-quii.gitbook.io/maths/clockoperating"
)

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const circle = `<circle cx="%d" cy="%d" r="%d" style="fill:%s;stroke:%s;stroke-width:5px;"/>`
const line = `<line x1="%d" y1="%d" x2="%d" y2="%d" style="fill:none;stroke:%s;stroke-width:3px;"/>`
const svgEnd = `</svg>`

func addLine(start, end clockoperating.Point, factor int, color string) string {
	return fmt.Sprintf(line, start.X, start.Y, start.X+(end.X-start.X)/factor, start.Y+(end.Y-start.Y)/factor, color)
}

func addCircle(center clockoperating.Point, radius int, color1, color2 string) string {
	return fmt.Sprintf(circle, center.X, center.Y, radius, color1, color2)
}

func SvgRendering(clock clock.Clock) string {
	var svgBuilder strings.Builder

	hourHand := addLine(clock.Center, clock.Hour, 3, "#f00")
	minuteHand := addLine(clock.Center, clock.Minute, 2, "#0f0")
	secondHand := addLine(clock.Center, clock.Second, 1, "#00f")
	bezel := addCircle(clock.Center, clock.Radius, "#999", "#fa7")

	svgBuilder.WriteString(svgStart)
	svgBuilder.WriteString(bezel)
	svgBuilder.WriteString(hourHand)
	svgBuilder.WriteString(minuteHand)
	svgBuilder.WriteString(secondHand)
	svgBuilder.WriteString(svgEnd)

	return svgBuilder.String()
}
