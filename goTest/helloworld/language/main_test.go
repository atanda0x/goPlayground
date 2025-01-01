package main

import "testing"

func TestGreet(t *testing.T) {
	t.Run("Hello", func(t *testing.T) {
		got := Hello("Atanda")
		want := "Hello, Atanda"

		AssertCorrectMessage(t, got, want)
	})

	t.Run("Empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		AssertCorrectMessage(t, got, want)
	})
}

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
