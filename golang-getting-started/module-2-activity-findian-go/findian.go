package main

import (
	"fmt"
	"strings"
)

func findian(s string) string {
	ls := strings.ToLower(s)
	if strings.HasPrefix(ls, "i") && strings.Contains(ls, "a") && strings.HasSuffix(ls, "n") {
		return "Found!"
	}
	return "Not Found!"
}

func main() {
	var s string
	fmt.Printf("Please enter your string for processing: ")
	fmt.Scan(&s)
	fmt.Println(findian(s))
}
