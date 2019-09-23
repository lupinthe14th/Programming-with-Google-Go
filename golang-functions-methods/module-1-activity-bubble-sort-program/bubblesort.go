package main

import (
	"errors"
	"fmt"
)

// BubbleSort function is modify the slice so that the elements in sorted order.
func BubbleSort(num []int) {
	for i := len(num); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			for num[j] > num[j+1] {
				Swap(num, j)
			}
		}
	}
}

// Swap is which swaps the position of two adjacent elements in the slice.
// The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
func Swap(num []int, idx int) {
	if idx > len(num) {
		panic(errors.New("No solutions"))
	}
	tmp := num[idx+1]
	num[idx+1] = num[idx]
	num[idx] = tmp
}

func main() {
	var numbers []int

	for i := 0; i < 10; i++ {
		var x int
		fmt.Print("Please enter input 10 integer: ")
		fmt.Scan(&x)
		numbers = append(numbers, x)
	}
	BubbleSort(numbers)
	for i, number := range numbers {
		if i == len(numbers) {
			fmt.Println(number)
		}
		fmt.Printf("%d ", number)
	}
}
