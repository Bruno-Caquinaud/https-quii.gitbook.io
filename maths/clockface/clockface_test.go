package clockface

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Bruno-Caquinaud/https-quii.gitbook.io/maths/clock"
	"github.com/Bruno-Caquinaud/https-quii.gitbook.io/maths/clockoperating"
)

func TestSVGRendering(t *testing.T) {
	type SVG struct {
		XMLName xml.Name `xml:"svg"`
		Text    string   `xml:",chardata"`
		Xmlns   string   `xml:"xmlns,attr"`
		Width   string   `xml:"width,attr"`
		Height  string   `xml:"height,attr"`
		ViewBox string   `xml:"viewBox,attr"`
		Version string   `xml:"version,attr"`
		Circle  struct {
			Text  string `xml:",chardata"`
			Cx    string `xml:"cx,attr"`
			Cy    string `xml:"cy,attr"`
			R     string `xml:"r,attr"`
			Style string `xml:"style,attr"`
		} `xml:"circle"`
		Line []struct {
			Text  string `xml:",chardata"`
			X1    string `xml:"x1,attr"`
			Y1    string `xml:"y1,attr"`
			X2    string `xml:"x2,attr"`
			Y2    string `xml:"y2,attr"`
			Style string `xml:"style,attr"`
		} `xml:"line"`
	}

	testcases := []struct {
		Description string
		Input1      clock.Clock
		Output      SVG
	}{
		{
			Description: "Test SVG Rendering 1 :",
			Input1: clock.Clock{
				Center: clockoperating.Point{X: 150, Y: 150},
				Radius: 90,
				Second: clockoperating.Point{X: 240, Y: 150},
				Minute: clockoperating.Point{X: 150, Y: 240},
				Hour:   clockoperating.Point{X: 60, Y: 150}},
			Output: SVG{XMLName: xml.Name{Space: "http://www.w3.org/2000/svg", Local: "svg"},
				Text:    "",
				Xmlns:   "http://www.w3.org/2000/svg",
				Width:   "100%",
				Height:  "100%",
				ViewBox: "0 0 300 300",
				Version: "2.0",
				Circle: struct {
					Text  string "xml:\",chardata\""
					Cx    string "xml:\"cx,attr\""
					Cy    string "xml:\"cy,attr\""
					R     string "xml:\"r,attr\""
					Style string "xml:\"style,attr\""
				}{
					Text: "", Cx: "150", Cy: "150", R: "90", Style: "fill:#999;stroke:#fa7;stroke-width:5px;",
				},
				Line: []struct {
					Text  string "xml:\",chardata\""
					X1    string "xml:\"x1,attr\""
					Y1    string "xml:\"y1,attr\""
					X2    string "xml:\"x2,attr\""
					Y2    string "xml:\"y2,attr\""
					Style string "xml:\"style,attr\""
				}{
					{Text: "", X1: "150", Y1: "150", X2: "120", Y2: "150", Style: "fill:none;stroke:#f00;stroke-width:3px;"},
					{Text: "", X1: "150", Y1: "150", X2: "150", Y2: "195", Style: "fill:none;stroke:#0f0;stroke-width:3px;"},
					{Text: "", X1: "150", Y1: "150", X2: "240", Y2: "150", Style: "fill:none;stroke:#00f;stroke-width:3px;"},
				},
			},
		},
		{
			Description: "Test SVG Rendering 2 :",
			Input1: clock.Clock{
				Center: clockoperating.Point{X: 150, Y: 150},
				Radius: 90,
				Second: clockoperating.Point{X: 150, Y: 60},
				Minute: clockoperating.Point{X: 150, Y: 60},
				Hour:   clockoperating.Point{X: 150, Y: 60}},
			Output: SVG{XMLName: xml.Name{Space: "http://www.w3.org/2000/svg", Local: "svg"},
				Text:    "",
				Xmlns:   "http://www.w3.org/2000/svg",
				Width:   "100%",
				Height:  "100%",
				ViewBox: "0 0 300 300",
				Version: "2.0",
				Circle: struct {
					Text  string "xml:\",chardata\""
					Cx    string "xml:\"cx,attr\""
					Cy    string "xml:\"cy,attr\""
					R     string "xml:\"r,attr\""
					Style string "xml:\"style,attr\""
				}{
					Text: "", Cx: "150", Cy: "150", R: "90", Style: "fill:#999;stroke:#fa7;stroke-width:5px;",
				},
				Line: []struct {
					Text  string "xml:\",chardata\""
					X1    string "xml:\"x1,attr\""
					Y1    string "xml:\"y1,attr\""
					X2    string "xml:\"x2,attr\""
					Y2    string "xml:\"y2,attr\""
					Style string "xml:\"style,attr\""
				}{
					{Text: "", X1: "150", Y1: "150", X2: "150", Y2: "120", Style: "fill:none;stroke:#f00;stroke-width:3px;"},
					{Text: "", X1: "150", Y1: "150", X2: "150", Y2: "105", Style: "fill:none;stroke:#0f0;stroke-width:3px;"},
					{Text: "", X1: "150", Y1: "150", X2: "150", Y2: "60", Style: "fill:none;stroke:#00f;stroke-width:3px;"},
				},
			},
		},
	}

	for _, testcase := range testcases {

		t.Run(testcase.Description, func(t *testing.T) {
			svgImage := SvgRendering(testcase.Input1)
			svgStruct := SVG{}
			xml.Unmarshal([]byte(svgImage), &svgStruct)

			if !reflect.DeepEqual(svgStruct, testcase.Output) {
				t.Errorf("Image are not equal, Got: %#v, Expected: %#v", svgStruct, testcase.Output)
			}
		})
	}
}
