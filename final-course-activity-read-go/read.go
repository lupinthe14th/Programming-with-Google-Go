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

type Person struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

func main() {
	var input string
	fmt.Print("Please enter input file name:")
	fmt.Scan(&input)
	log.Print(input)
	fd, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	reader := bufio.NewReaderSize(fd, 1024)
	var persons []Person
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
			p := Person{Fname: lines[0], Lname: lines[1]}
			if len(p.Fname) > 20 || len(p.Lname) > 20 {
				err := fmt.Sprintf("Error: over size 20 characters: %+v", p)
				log.Print(err)
				continue
			}
			persons = append(persons, p)
		}
	}

	p, err := json.Marshal(&persons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(p))
}
