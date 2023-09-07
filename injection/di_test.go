package injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "mach")

	got := buffer.String()
	want := "Hello, mach"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
