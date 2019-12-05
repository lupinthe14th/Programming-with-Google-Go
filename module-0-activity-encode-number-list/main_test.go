package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  [][]int
}{
	{id: 1, input: 0, want: [][]int{{}}},
	{id: 2, input: 1, want: [][]int{{}, {0}}},
}

func TestF(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := f(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}

func Example_main() {
	main()

	// Output: [[] [0] [1] [0 0] [0 1] [1 0] [1 1] [0 0 0] [0 0 1] [0 1 0] [0 1 1]]
}

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(i)
	}
}
