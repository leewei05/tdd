package main

import (
	"os"
	"tdd/mock"
)

func main() {
	mock.Countdown(os.Stdout)
}
