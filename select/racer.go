package racer

import (
	"net/http"
)

func Racer(a, b string) string {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
