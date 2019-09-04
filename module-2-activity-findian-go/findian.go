package main

import (
	"fmt"
	"strings"
)

func findian(s string) {
	if strings.ContainsAny(s, "ian") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func main() {
	var found, notFound string
	fmt.Println("Please enter a string containing the characters 'i', 'a', and 'n'.")
	fmt.Scan(&found)
	findian(found)
	fmt.Println("Please enter a string not containing the characters 'i', 'a', and 'n'.")
	fmt.Scan(&notFound)
	findian(notFound)
}
