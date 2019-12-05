package main

import (
	"fmt"
	"strconv"
)

func s(num int) []int {
	num++
	var ans string
	for num > 0 {
		ans = strconv.Itoa((num % 2)) + ans
		num /= 2
	}

	a := make([]int, len(ans)-1)
	for i, v := range string([]rune(ans)[1:]) {
		s, _ := strconv.Atoi(string(v))
		a[i] = s
	}
	return a
}

func f(c int) [][]int {
	matrix := make([][]int, c+1)

	matrix[0] = []int{}
	if c > 0 {
		for i := 1; i <= c; i++ {
			matrix[i] = s(i)
		}
	}
	return matrix
}

func main() {
	fmt.Print(f(10))
}
