package main

import (
	"fmt"
	"sort"
	"strconv"
)

var numList = make([]int, 3)

func add(i, n int) []int {
	if i <= 2 {
		numList[i] = n
	} else {
		numList = append(numList, n)
	}
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
	i := 0
	for {
		var s string
		fmt.Print("Please entering distinct integers: ")
		fmt.Scan(&s)
		if s == "X" {
			break
		}
		num, _ := strconv.Atoi(s)
		if !IsContains(numList, num) {
			nl := add(i, num)
			if i >= 2 {
				sort.Ints(nl)
				fmt.Println(nl)
			}
			i++
		}
	}
}
