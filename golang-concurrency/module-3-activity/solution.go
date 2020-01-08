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
	sort.Ints(nums)
	c <- nums
}

func sortRecord(nums []int) []int {
	var ans []int
	var d int
	l := len(nums)
	d = l / 4
	c := make(chan []int, 4)
	go sub(nums[:d], c)
	go sub(nums[d:2*d], c)
	go sub(nums[2*d:3*d], c)
	go sub(nums[3*d:l], c)

	for i := 0; i < 4; i++ {
		tmp := <-c
		ans = append(ans, tmp...)
	}
	sort.Ints(ans)
	return ans
}

func sortRecordN(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp := strings.Split(readLine(reader), " ")

	var nums []int

	for i := 0; i < len(nTemp); i++ {
		n, err := strconv.ParseInt(nTemp[i], 10, 64)
		checkError(err)
		nums = append(nums, int(n))
	}
	fmt.Println(sortRecord(nums))
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
