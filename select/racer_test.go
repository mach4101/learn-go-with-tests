package main

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://www.vsbf.fun"
	fastURL := "http://www.baidu.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
