package clockoperating

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestTimetoRadian(t *testing.T) {
	testcases := []struct {
		Description string
		Input1      time.Duration
		Input2      time.Time
		Output      float64
	}{
		{
			Description: "Test Time to Radian 1 : 0 Second",
			Input1:      time.Second,
			Input2:      time.Date(2026, time.September, 00, 00, 00, 00, 00, time.UTC),
			Output:      0,
		},
		{
			Description: "Test Time to Radian 2 : 15 Second",
			Input1:      time.Second,
			Input2:      time.Date(2026, time.September, 00, 00, 00, 15, 00, time.UTC),
			Output:      math.Pi / 2,
		},
		{
			Description: "Test Time to Radian 3 : 30 Second",
			Input1:      time.Second,
			Input2:      time.Date(2026, time.September, 00, 00, 00, 30, 00, time.UTC),
			Output:      math.Pi,
		},
		{
			Description: "Test Time to Radian 4 : 45 Second",
			Input1:      time.Second,
			Input2:      time.Date(2026, time.September, 00, 00, 00, 45, 00, time.UTC),
			Output:      3 * math.Pi / 2,
		},
		{
			Description: "Test Time to Radian 5 : 15 Minutes",
			Input1:      time.Minute,
			Input2:      time.Date(2026, time.September, 00, 00, 15, 00, 00, time.UTC),
			Output:      math.Pi / 2,
		},
		{
			Description: "Test Time to Radian 6 : 21 Heure",
			Input1:      time.Hour,
			Input2:      time.Date(2026, time.September, 00, 21, 00, 00, 00, time.UTC),
			Output:      7 * math.Pi / 2,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Description, func(t *testing.T) {
			got := TimeToRadian(testcase.Input1, testcase.Input2)

			if testcase.Output != got {
				t.Errorf("Expected : %v is not Equal %v wanted", testcase.Output, got)
			}
		})
	}
}

func TestRadianToCoordinate(t *testing.T) {
	testcases := []struct {
		Description string
		Input1      float64
		Input2      float64
		Input3      Point
		Output      Point
	}{
		{
			Description: "Test Radian to Coordinate 1 : %f Radian, %f Radius, %v Offset",
			Input1:      0,
			Input2:      0,
			Input3:      Point{0, 0},
			Output:      Point{0, 0},
		},
		{
			Description: "Test Radian to Coordinate 2 : %f Radian, %f Radius, %v Offset",
			Input1:      0,
			Input2:      0,
			Input3:      Point{150, 150},
			Output:      Point{150, 150},
		},
		{
			Description: "Test Radian to Coordinate 3: %f Radian, %f Radius, %v Offset",
			Input1:      math.Pi / 2,
			Input2:      1,
			Input3:      Point{0, 0},
			Output:      Point{1, 0},
		},
		{
			Description: "Test Radian to Coordinate 3: %f Radian, %f Radius, %v Offset",
			Input1:      math.Pi,
			Input2:      1,
			Input3:      Point{0, 0},
			Output:      Point{0, 1},
		},
		{
			Description: "Test Radian to Coordinate 4: %f Radian, %f Radius, %v Offset",
			Input1:      math.Pi,
			Input2:      90,
			Input3:      Point{0, 0},
			Output:      Point{0, 90},
		},
		{
			Description: "Test Radian to Coordinate 4: %f Radian, %f Radius, %v Offset",
			Input1:      math.Pi,
			Input2:      90.0,
			Input3:      Point{150, 150},
			Output:      Point{150, 240},
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf(testcase.Description, testcase.Input1, testcase.Input2, testcase.Input3), func(t *testing.T) {
			got := RadiantoCoordinate(testcase.Input1, testcase.Input2, testcase.Input3)

			if testcase.Output != got {
				t.Errorf("Expected : %v is not Equal %v wanted", testcase.Output, got)
			}
		})
	}
}
