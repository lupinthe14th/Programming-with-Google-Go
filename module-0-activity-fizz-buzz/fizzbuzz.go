package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

func fizzBuzz(num uint64) string {
	fizz := num % 3
	log.Print(fizz)
	buzz := num % 5
	log.Print(buzz)
	if fizz == 0 && buzz == 0 {
		return "FizzBuzz"
	} else if fizz == 0 {
		return "Fizz"
	} else if buzz == 0 {
		return "Buzz"
	} else {
		return fmt.Sprint(num)
	}
}

func main() {
	go func() {
		for {
			var x uint64
			fmt.Println("Please enter a intager or Ctr-c: ")
			fmt.Scan(&x)
			fmt.Println(fizzBuzz(x))
		}
	}()

	doneC := make(chan os.Signal, 1)
	signal.Notify(doneC, os.Interrupt)
	<-doneC
}
