package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("mach")
	want := "Hello, mach"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
