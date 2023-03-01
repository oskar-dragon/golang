package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Oskar")

	got := buffer.String()
	want := "Hello, Oskar"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}