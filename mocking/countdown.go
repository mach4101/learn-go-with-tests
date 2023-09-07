package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const sleep = "sleep"
const write = "write"

// configuration of sleep
type ConfigurableSleeper struct {
	duration time.Duration
}

// record table
type CountdownOperationSpy struct {
	Calls []string
}

// interface
type Sleeper interface {
	Sleep()
}

func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

// implement Sleep interface
func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, "GO!")
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
