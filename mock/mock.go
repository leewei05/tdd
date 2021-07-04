package main

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

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

func (s *SpySleeper) Sleep() {
	s.Calls++
}
