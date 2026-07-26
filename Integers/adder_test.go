package intergers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Test 1 : Add with two integers parameters", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("sum %d not equal to expected %d", sum, expected)
		}
	})
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}
