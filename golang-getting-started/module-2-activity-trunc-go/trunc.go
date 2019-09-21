package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Print("Please enter a floating point number: ")
	fmt.Scan(&num)
	fmt.Printf("%.0f", math.Trunc(num))
}
