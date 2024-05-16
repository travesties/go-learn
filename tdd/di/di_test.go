package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Travis")

	got := buffer.String()
	want := "Hello, Travis"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
