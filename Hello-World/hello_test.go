package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Test 1 : Hello Func with one string parameter : Name", func(t *testing.T) {
		got := Hello("Bruno", "")
		want := "Hello, Bruno"

		assertCorrectMessage(t, want, got)
	})

	t.Run("Test 2 : Hello func with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, want, got)
	})

	t.Run("Test 3 : Hello func with two string parameters : name + language", func(t *testing.T) {
		got := Hello("Bruno", "Spanish")
		want := "Holla, Bruno"

		assertCorrectMessage(t, want, got)
	})
}

func assertCorrectMessage(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
