package array_slice

import (
	"testing"
	"slices"
	"reflect"
)

func TestSum(t * testing.T) {

	numbers := []int{1, 2, 3, 4, 5}
	want := Sum(numbers)
	expected := 15

	if want != expected {
		t.Errorf("want %d not equal to expected %d", want, expected)
	}
}

func TestSumAll(t * testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}

	expected := SumAll(numbers1, numbers2)
	want := []int{6, 15}

	if !slices.Equal(expected, want) {
		t.Errorf("expected %v != want %v", expected, want)
	} 
}

func TestSumAllTails(t *testing.T) {
	checksum := func(t * testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	
	t.Run("Test 1 : Sum Tails", func(t * testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checksum(t, got, want)
	})

	t.Run("Test 2 : Sum empty Tails", func(t * testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}

		checksum(t, got, want)
	})
}