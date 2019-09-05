package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
)

func fizzBuzz(num int) {
	if math.Mod(float64(num), 3) == 0 && math.Mod(float64(num), 5) == 0 {
		fmt.Println("FizzBuzz")
	} else if math.Mod(float64(num), 3) == 0 {
		fmt.Println("Fizz")
	} else if math.Mod(float64(num), 5) == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(num)
	}
}

func main() {
	go func() {
		for {
			var x int
			fmt.Println("Please enter a intager or Ctr-c")
			fmt.Scan(&x)
			fizzBuzz(x)
		}
	}()

	doneC := make(chan os.Signal, 1)
	signal.Notify(doneC, os.Interrupt)
	<-doneC
}
