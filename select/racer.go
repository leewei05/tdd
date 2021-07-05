package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondsTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool
func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func(u string) {
		http.Get(u)
		close(ch)
	}(url)

	return ch
}
