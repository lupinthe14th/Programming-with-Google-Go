package main

import (
	"fmt"
	"strings"
)

// Animal is each animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, locomotion, noise string
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a Cow) Speak() {
	fmt.Println(a.noise)
}

type Bird struct {
	food, locomotion, noise string
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a Bird) Speak() {
	fmt.Println(a.noise)
}

type Snake struct {
	food, locomotion, noise string
}

func (a Snake) Eat() {
	fmt.Println(a.food)
}

func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a Snake) Speak() {
	fmt.Println(a.noise)
}

func animalType(t string) Animal {
	var a Animal
	switch strings.ToLower(t) {
	case "cow":
		a = Cow{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		}
	case "bird":
		a = Bird{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		}

	case "snake":
		a = Snake{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		}
	}
	return a
}

func main() {
	var command, name, option string

	var animals = make(map[string]Animal)

	for {
		fmt.Print("> ")
		fmt.Scan(&command, &name, &option)
		switch command {
		case "newanimal":
			if animals[name] == nil {
				animals[name] = animalType(option)
				fmt.Println("Created it！")
			} else {
				fmt.Printf("%s is duplicate！\n", name)
			}
		case "query":
			var a Animal
			var animal = animals[name]
			a = animal
			switch strings.ToLower(option) {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			}
		}
	}
}
