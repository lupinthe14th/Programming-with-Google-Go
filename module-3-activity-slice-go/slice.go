package main

import (
	"fmt"
	"os"
	"os/signal"
	"sort"
)

var numList []int

func slice(n int) []int {
	numList = append(numList, n)
	sort.Ints(numList)
	return numList
}

// IsContains tells whether nl contains n.
func IsContains(nl []int, n int) bool {
	for _, i := range nl {
		if i == n {
			fmt.Println("Is contains")
			return true
		}
	}
	return false
}

func main() {
	go func() {
		for {
			var num int
			fmt.Print("Please entering distinct integers: ")
			fmt.Scan(&num)
			if !IsContains(numList, num) {
				nl := slice(num)
				if len(nl) > 2 {
					fmt.Println(nl)
				}
			}
		}
	}()

	fmt.Println("wont exit enter Ctr-c")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
}
