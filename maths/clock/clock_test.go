package clock

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/Bruno-Caquinaud/https-quii.gitbook.io/maths/clockoperating"
)

func TestComputeCoordinate(t *testing.T) {
	testcases := []struct {
		Description string
		Input1      *Clock
		Input2      time.Time
		Output      Clock
	}{
		{
			Description: "Test Compute Coordinate 1 : Clock %v, Time %v",
			Input1:      &Clock{Center: clockoperating.Point{X: 0, Y: 0}, Radius: 0},
			Input2:      time.Date(2026, time.September, 13, 06, 30, 45, 00, time.UTC),
			Output: Clock{Second: clockoperating.Point{X: 0, Y: 0},
				Minute: clockoperating.Point{X: 0, Y: 0},
				Hour:   clockoperating.Point{X: 0, Y: 0}},
		},
		{
			Description: "Test Compute Coordinate 2 :  Clock %v, Time %v",
			Input1:      &Clock{Center: clockoperating.Point{X: 0, Y: 0}, Radius: 1},
			Input2:      time.Date(2026, time.September, 13, 6, 30, 45, 00, time.UTC),
			Output: Clock{Radius: 1,
				Second: clockoperating.Point{X: -1, Y: 0},
				Minute: clockoperating.Point{X: 0, Y: 1},
				Hour:   clockoperating.Point{X: 0, Y: 1}},
		},
		{
			Description: "Test Compute Coordinate 3 :  Clock %v, Time %v",
			Input1:      &Clock{Center: clockoperating.Point{X: 150, Y: 150}, Radius: 90},
			Input2:      time.Date(2026, time.September, 13, 21, 30, 15, 00, time.UTC),
			Output: Clock{Center: clockoperating.Point{X: 150, Y: 150},
				Radius: 90,
				Second: clockoperating.Point{X: 240, Y: 150},
				Minute: clockoperating.Point{X: 150, Y: 240},
				Hour:   clockoperating.Point{X: 60, Y: 150}},
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf(testcase.Description, testcase.Input1, testcase.Input2), func(t *testing.T) {
			testcase.Input1.ComputeCoordinates(testcase.Input2)

			if !reflect.DeepEqual(testcase.Output, *testcase.Input1) {
				t.Errorf("Expected : %v is not Equal %v wanted", testcase.Output, *testcase.Input1)
			}
		})
	}
}
