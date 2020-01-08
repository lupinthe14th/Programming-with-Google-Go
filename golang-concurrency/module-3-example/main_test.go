package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input []int
	want  int
}

var cases = []Case{
	{input: []int{}, want: 0},
	{input: []int{1, 2, 3, 4}, want: 24},
	{input: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 40320},
}

func TestProduct(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := product(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func TestSingleProduct(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := singleProduct(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func BenchmarkProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		product(cases[1].input)
	}
}

func BenchmarkSingeProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleProduct(cases[1].input)
	}
}
