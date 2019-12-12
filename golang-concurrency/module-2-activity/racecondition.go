// # What's race condition?
// A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.
// Most of the time, A race condition shows up in what's called a data race, where one concurrent operation attempts to read a variable while at some undermined time another concurrent operation is attempting to write to the same variable.
// - Katherine Cox-Buday (2017) Concurrency in Go, O'Reilly Media, Inc.
package main

import (
	"fmt"
)

func main() {
	x := 1

	go func() {
		x++
	}()

	go func() {
		x--
	}()

	fmt.Println(x)
}
