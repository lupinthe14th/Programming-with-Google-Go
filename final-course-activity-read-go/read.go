package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
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
	var tmp []byte
	var users []User
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		tmp = append(tmp, line...)
		log.Print(string(tmp))
		if !isPrefix {
			users = append(users, User{Fname: string(tmp[0]), Lname: string(tmp[1])})
		}
	}

	u, err := json.Marshal(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(u))
}
