package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Test 1 : Hello Func with string parameter", func(t *testing.T) {
		got := Hello("Bruno")
		want := "Hello, Bruno"

		assertCorrectMessage(t, want, got)
	})

	t.Run("Test 2 : Hello func with empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Worl"

		assertCorrectMessage(t, want, got)
	})
}

func assertCorrectMessage(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
