package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sub(nums []int, c chan []int) {
	fmt.Println(nums)
	sort.Ints(nums)
	c <- nums
}

func sortRecord(nums []int) []int {
	var d int
	fmt.Println("Four sorted slices are as follows:")
	l := len(nums)
	d = l / 4
	c := make(chan []int, 4)
	go sub(nums[:d], c)
	go sub(nums[d:2*d], c)
	go sub(nums[2*d:3*d], c)
	go sub(nums[3*d:l], c)

	return merge(merge(<-c, <-c), merge(<-c, <-c))
}

func sortRecordN(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func merge(a, b []int) []int {
	ans := make([]int, len(a)+len(b))

	var i, j int = 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			ans[i+j] = a[i]
			i++
		} else {
			ans[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		ans[i+j] = a[i]
		i++

	}
	for j < len(b) {
		ans[i+j] = b[j]
		j++
	}

	return ans

}

func main() {
	fmt.Println("NOTE:")
	fmt.Println("1. All (including last) integers should be delimited by space")
	fmt.Println("2. Total count of integers should be divisible by 4")
	fmt.Println("Enter sequence of integers:")
	fmt.Println(">")
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp := strings.Split(readLine(reader), " ")

	var nums []int

	for i := 0; i < len(nTemp); i++ {
		n, err := strconv.ParseInt(nTemp[i], 10, 64)
		checkError(err)
		nums = append(nums, int(n))
	}
	ans := sortRecord(nums)
	fmt.Println("Sorted integer sequence is as follows:")
	fmt.Println(ans)
}

func readLine(r *bufio.Reader) string {
	b, _, err := r.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(b), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
