package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Please enter name: ")
	name := bufio.NewScanner(os.Stdin)
	name.Scan()
	fmt.Print("Please enter address: ")
	address := bufio.NewScanner(os.Stdin)
	address.Scan()
	user := map[string]string{
		"name":    name.Text(),
		"address": address.Text(),
	}
	blob, _ := json.Marshal(user)
	fmt.Println(string(blob))
}
