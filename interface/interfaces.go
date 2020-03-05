package main

import (
	"fmt"
)

// Animal is an interface whose blueprint can be implemented by other Go structs in the package
type Animal interface {
	// inside an interface you write the function header only and other structs will implement them accordingly
	Speak() string
}

// Dog is a struct with Breed, Weight and Sound fields
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Cat is a struct with Sound field
type Cat struct {
	Sound string
}

// Cow is a struct with Sound field
type Cow struct {
	Sound string
}

// Speak is an exportable and implemented function of Animal interface for Dog struct
func (dog Dog) Speak() string {
	return dog.Sound
}

// Speak is an exportable and implemented function of Animal interface for Cat struct
func (cat Cat) Speak() string {
	return cat.Sound
}

// Speak is an exportable and implemented function of Animal interface for Cow struct
func (cow Cow) Speak() string {
	return cow.Sound
}

func main() {
	poodle := Animal(Dog{"Poodle", 15, "Bhow!"})
	fmt.Println(poodle)

	cat := Animal(Cat{"Meow!"})
	cow := Animal(Cow{"Moo!"})

	// call the method implemented by the interface
	fmt.Println(poodle.Speak())
	fmt.Println(cat.Speak())
	fmt.Println(cow.Speak())

	// you can create a slice of the structs that implements the interface
	animals := []Animal{poodle, cat, cow}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
