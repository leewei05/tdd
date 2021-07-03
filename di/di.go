package di

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

/*
By using io.Writer, we can pass io.Stdout, bytes.Buffer type into Greet.
Because those types implement io.Writer with Write(p []byte) (n int, err error)
*/
