package strucMethodInterface

import (
	"testing"
)

func TestPerimeter(t * testing.T) {
	t.Run("Test 1 : Perimeter with values", func(t * testing.T) {
		rectangle := Rectangle{123.5, 21.0}
		got := rectangle.Perimeter()
		want := 289.0

		if got != want {
			t.Errorf("got %.2f != want %.2f", got, want)
		}
	})
}

func TestArea(t * testing.T) {

	areaTestCases := []struct {
		name string
		shape Shape
		expectedArea float64
	}{
		{name : "Rectangle", shape : Rectangle{2.5, 10.0}, expectedArea : 25.0},
		{name : "Circle", shape : Circle{10.0}, expectedArea : 314.1592653589793},
		{name : "Triangle", shape : Triangle{12, 6}, expectedArea : 36.0},
	}

	checkArea := func(t * testing.T, shape Shape, expectedArea float64) {

		got := shape.Area()

		if got != expectedArea {
			t.Errorf("%#v, got %.2f != want %.2f", shape, got, expectedArea)
		}
	}

	for _, testCase := range areaTestCases {
		t.Run(testCase.name, func(t * testing.T) {
			 checkArea(t, testCase.shape, testCase.expectedArea)
		})
	}
}