package main

import (
	"fmt"
)

func main() {
	var num float64
	fmt.Println("Please enter a floating point number.")
	fmt.Scan(&num)
	fmt.Printf("first: %.0f\n", num)

	fmt.Println("Please enter a floating point number.")
	fmt.Scan(&num)
	fmt.Printf("second: %.0f\n", num)
}
