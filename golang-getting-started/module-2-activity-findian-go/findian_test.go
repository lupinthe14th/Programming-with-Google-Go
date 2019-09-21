package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input string
	want  string
}{
	// Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
	{input: "ian", want: "Found!"},
	{input: "Ian", want: "Found!"},
	{input: "iuiygaygn", want: "Found!"},
	{input: "I d skd a efju N", want: "Found!"},
	{input: "ihhhhhn", want: "Not Found!"},
	{input: "ina", want: "Not Found!"},
	{input: "xian", want: "Not Found!"},
}

func TestFindian(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.input), func(t *testing.T) {
			actual := findian(tt.input)
			if actual != tt.want {
				t.Errorf("%v, want %v", actual, tt.want)
			}
		})
	}
}
