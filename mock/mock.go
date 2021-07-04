package main

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

type CountdownOperationsSpy struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

var (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(out io.Writer, spySleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		spySleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	spySleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
