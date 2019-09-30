package main

import (
	"fmt"
	"testing"
)

type initial struct {
	a, v0, s0 float64
}

var cases = []struct {
	id   int
	init initial
	t    []float64
	want []float64
}{
	{id: 1, init: initial{a: 10, v0: 2, s0: 1}, t: []float64{0, 3, 5}, want: []float64{1, 52, 136}},
}

func TestGen(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			fn := GenDisplaceFn(tt.init.a, tt.init.v0, tt.init.s0)
			for i, at := range tt.t {
				got := fn(at)
				if got != tt.want[i] {
					t.Errorf("%f, want: %f", got, tt.want[i])
				}
			}
		})
	}
}
