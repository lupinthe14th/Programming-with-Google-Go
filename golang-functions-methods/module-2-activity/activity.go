package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn returns a function to compute displacement
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*math.Pow(t, 2))/2 + v0*t + s0
	}
	return fn
}

func main() {
	var acc, vel, dis, t float64
	fmt.Print("Please enter acceleration: ")
	fmt.Scan(&acc)
	fmt.Print("Please enter initial velocity: ")
	fmt.Scan(&vel)
	fmt.Print("Please enter initial displacement: ")
	fmt.Scan(&dis)
	fn := GenDisplaceFn(acc, vel, dis)
	fmt.Print("Please enter elapsed: ")
	fmt.Scan(&t)
	fmt.Printf("Displacement would be %f after time %f with given acceleration %f, initial velocity %f and initial displacement %f\n", fn(t), t, acc, vel, dis)
}
