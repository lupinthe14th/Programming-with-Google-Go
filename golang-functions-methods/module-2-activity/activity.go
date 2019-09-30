package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn is ...
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*math.Pow(t, 2))/2 + v0*t + s0
	}
	return fn
}

func main() {
	var a, v0, s0, t float64
	fmt.Print("Please enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Please enter initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Please enter initial displacement: ")
	fmt.Scan(&s0)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Print("Please enter how many seconds later: ")
	fmt.Scan(&t)
	fmt.Printf("%f seconds later displacement s: %f\n", t, fn(t))
	fmt.Print("Please enter how many seconds later: ")
	fmt.Scan(&t)
	fmt.Printf("%f seconds later displacement s: %f\n", t, fn(t))
}
