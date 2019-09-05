package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input int
	want  string
}{
	{input: 1, want: "1"},
	{input: 2, want: "2"},
	{input: 3, want: "Fizz"},
	{input: 4, want: "4"},
	{input: 5, want: "Buzz"},
	{input: 6, want: "Fizz"},
	{input: 7, want: "7"},
	{input: 8, want: "8"},
	{input: 9, want: "Fizz"},
	{input: 10, want: "Buzz"},
	{input: 11, want: "11"},
	{input: 12, want: "Fizz"},
	{input: 13, want: "13"},
	{input: 14, want: "14"},
	{input: 15, want: "FizzBuzz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			actual := fizzBuzz(tt.input)
			if actual != tt.want {
				t.Errorf("%v, want %v", actual, tt.want)
			}
		})
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzz(i)
	}
}
