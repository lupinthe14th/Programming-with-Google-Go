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
	a := make([]int, 0)
	for _, v := range string([]rune(ans)[1:]) {
		s, _ := strconv.Atoi(string(v))
		a = append(a, s)
	}
	return a
}

func f(c int) [][]int {
	matrix := make([][]int, 0, c)
	for i := range matrix {
		matrix[i] = make([]int, 0, c)
	}

	matrix = append(matrix, []int{})
	if c > 0 {
		for i := 1; i <= c; i++ {
			matrix = append(matrix, s(i))
		}
	}
	return matrix
}

func main() {
	fmt.Print(f(10))
}
