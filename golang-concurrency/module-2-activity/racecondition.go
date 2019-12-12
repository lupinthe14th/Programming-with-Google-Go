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
