package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type Person struct {
	Fname string `json:"fname" validate:"max=20"`
	Lname string `json:"lname" validate:"max=20"`
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

	reader := bufio.NewReaderSize(fp, 1024)
	var persons []Person
	validate := validator.New()
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
			if err := validate.Struct(p); err != nil {
				log.Print(err)
				log.Printf("person: %+v", p)
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
