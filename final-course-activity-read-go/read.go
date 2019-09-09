package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type User struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

func main() {
	var input string
	fmt.Print("Please enter input file name:")
	fmt.Scan(&input)
	log.Print(input)
	fp, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 41)
	var users []User
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(string(line), " ")
		if !isPrefix {
			users = append(users, User{Fname: lines[0], Lname: lines[1]})
		}
	}

	u, err := json.Marshal(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(u))
}
