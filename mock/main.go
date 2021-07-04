package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct{}

type Sleeper interface {
	Sleep()
}

var (
	countdownStart = 3
	finalWord      = "Go!"
)

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, spySleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		spySleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	spySleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
