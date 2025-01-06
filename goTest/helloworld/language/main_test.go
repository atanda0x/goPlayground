package main

import "testing"

func TestGreet(t *testing.T) {
	t.Run("Hello", func(t *testing.T) {
		got := Hello("Atanda", "English")
		want := "Hello, Atanda"

		AssertCorrectMessage(t, got, want)
	})

	t.Run("Empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		AssertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Nafiu", "Spanish")
		want := "Hola, Nafiu"

		AssertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("0x", "French")
		want := "Bonjour, 0x"

		AssertCorrectMessage(t, got, want)
	})
}

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
