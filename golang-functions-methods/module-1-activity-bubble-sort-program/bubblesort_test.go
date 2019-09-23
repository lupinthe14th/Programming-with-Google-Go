package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var cases = []struct {
		id    int
		input []int
		want  []int
	}{
		{id: 1, input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{id: 2, input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{id: 3, input: []int{10, 9, 8, 7, -6, 5, 4, 3, 2, 1}, want: []int{-6, 1, 2, 3, 4, 5, 7, 8, 9, 10}},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			BubbleSort(tt.input)
			got := tt.input
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	var cases = []struct {
		id        int
		inputInts []int
		inputIdx  int
		want      []int
	}{
		{id: 1, inputInts: []int{1, 2, 3}, inputIdx: 0, want: []int{2, 1, 3}},
		{id: 2, inputInts: []int{1, 2, 3}, inputIdx: 1, want: []int{1, 3, 2}},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			Swap(tt.inputInts, tt.inputIdx)
			got := tt.inputInts
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
