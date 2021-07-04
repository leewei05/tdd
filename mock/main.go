package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type Sleeper interface {
	Sleep()
}

var (
	countdownStart = 3
	finalWord      = "Go!"
)

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
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

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
