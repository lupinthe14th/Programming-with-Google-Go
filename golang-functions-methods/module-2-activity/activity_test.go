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
	t    float64
	want float64
}{
	{id: 1, init: initial{a: 10, v0: 2, s0: 1}, t: 0, want: 1},
	{id: 2, init: initial{a: 10, v0: 2, s0: 1}, t: 3, want: 52},
	{id: 3, init: initial{a: 10, v0: 2, s0: 1}, t: 5, want: 136},
}

func TestGen(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			fn := GenDisplaceFn(tt.init.a, tt.init.v0, tt.init.s0)
			got := fn(tt.t)
			if got != tt.want {
				t.Errorf("%f, want: %f", got, tt.want)
			}
		})
	}
}
