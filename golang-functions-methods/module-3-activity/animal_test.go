package main

var animal = cow

func ExampleAnimal_Eat() {
	animal.Eat()

	// Output: grass
}

func ExampleAnimal_Move() {
	animal.Move()

	// Output: walk
}

func ExampleAnimal_Speak() {
	animal.Speak()

	// Output: moo
}
