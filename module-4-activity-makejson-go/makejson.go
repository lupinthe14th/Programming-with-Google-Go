package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// PrintPerson is a function that print a Person JSON object.
func PrintPerson(name, address string) error {
	p := map[string]string{
		"name":    name,
		"address": address,
	}
	barr, err := json.Marshal(p)
	if err != nil {
		return err
	}
	fmt.Println(string(barr))
	return nil
}

func main() {
	fmt.Print("Please enter name: ")
	name := bufio.NewScanner(os.Stdin)
	name.Scan()
	fmt.Print("Please enter address: ")
	address := bufio.NewScanner(os.Stdin)
	address.Scan()
	if err := PrintPerson(name.Text(), address.Text()); err != nil {
		log.Fatal(err)
	}
}
