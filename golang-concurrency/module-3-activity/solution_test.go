package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

type Case struct {
	input, want []int
}

var cases = []Case{
	{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
	{input: []int{12, 4, 2, 13, 10, 0, 3, 11, 7, 5, 15, 1, 9, 14, 6, 8}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
	{input: []int{12, 4, 2, 13, 10, 0, 3, 11, 7, 5, 15, 1, 9, 14, 6, 8, 16, 17, 18, 19}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}},
}

func TestSortRecord(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sortRecord(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}

func TestSortRecordN(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sortRecordN(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}

func Example_main() {
	c, _ := ioutil.ReadFile("./input.csv")
	inr, inw, _ := os.Pipe()
	inw.Write(c)
	inw.Close()
	orgStdin := os.Stdin

	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Unordered output:
	// NOTE:
	// 1. All (including last) integers should be delimited by space
	// 2. Total count of integers should be divisible by 4
	// Enter sequence of integers:
	// >
	// Four sorted slices are as follows:
	// [12 4 2 13]
	// [10 0 3 11]
	// [7 5 15 1]
	// [9 14 6 8]
	// Sorted integer sequence is as follows:
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
}

func TestCheckError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	err := errors.New("Error")
	checkError(err)
}

func BenchmarkSortRecord(b *testing.B) {
	for i := 4; i < b.N; i++ {
		sortRecord(rand.Perm(i))
	}
}

func BenchmarkSortRecordN(b *testing.B) {
	for i := 4; i < b.N; i++ {
		sortRecordN(rand.Perm(i))
	}
}
