package mock

import (
	"fmt"
	"io"
	"time"
)

var (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}
