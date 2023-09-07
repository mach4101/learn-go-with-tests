package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aTime := measureResponseTime(a)
	bTime := measureResponseTime(b)

	if aTime < bTime {
		return a
	} else {
		return b
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
