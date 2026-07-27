package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat(5, "a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q is not equal to want %q", got, want)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat(3, "alo"))
	// Output: aloaloalo
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat(3, "x")
	}
}
