package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	// Introducing subtests. We can use t.Run to run multiple tests grouped
	// within a single test function wrapper.
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Travis", "")
		want := "Hello, Travis"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Travis", "Spanish")
		want := "Hola, Travis"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Travis", "French")
		want := "Bonjour, Travis"
		assertCorrectMessage(t, got, want)
	})
}

// For helper functions, it's a good idea to accept a testing.TB, which
// is an interface that *testing.T and *testing.B both satisfy, which allows
// you to call helper functions for a test or a benchmark.
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper needs to be called within a helper method to tell the test suite
	// that it is a helper method. By doing this, when a test fails the reported
	// line number will be in our function call rather than inside our test helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
