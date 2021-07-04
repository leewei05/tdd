package main

import (
	"os"
	"time"
)

type DefaultSleeper struct{}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
