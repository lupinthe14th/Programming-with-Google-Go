package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"strconv"
)

func fizzBuzz(num int) string {
	if math.Mod(float64(num), 3) == 0 && math.Mod(float64(num), 5) == 0 {
		return "FizzBuzz"
	} else if math.Mod(float64(num), 3) == 0 {
		return "Fizz"
	} else if math.Mod(float64(num), 5) == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(num)
	}
}

func main() {
	go func() {
		for {
			var x int
			fmt.Println("Please enter a intager or Ctr-c")
			fmt.Scan(&x)
			fmt.Println(fizzBuzz(x))
		}
	}()

	doneC := make(chan os.Signal, 1)
	signal.Notify(doneC, os.Interrupt)
	<-doneC
}
