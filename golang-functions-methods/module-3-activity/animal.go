package main

import (
	"fmt"
)

// Animal is each animal data structure
type Animal struct {
	food, locomotion, noise string
}

var cow = Animal{
	food:       "grass",
	locomotion: "walk",
	noise:      "moo",
}

var bird = Animal{
	food:       "worms",
	locomotion: "fly",
	noise:      "peep",
}

var snake = Animal{
	food:       "mice",
	locomotion: "slither",
	noise:      "hsss",
}

// Eat is print the animal's food
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move is print the animal's locomotion
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak is print the animal's spoken sound
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var name, action string
	var animal Animal
	for {
		fmt.Print(">")
		fmt.Scan(&name, &action)
		switch name {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		}
		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}

	}
}
