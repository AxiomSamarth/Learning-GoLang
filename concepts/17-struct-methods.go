package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak prints the sound of the struct object
func (dog Dog) Speak() {
	fmt.Println("Speak function prints:", dog.Sound)
}

// ModifySpeak modifies the struct field Sound of the Dog struct. Also, notice that it accepts a pointer
func (dog *Dog) ModifySpeak() {
	dog.Sound = "Woof"
}

func main() {
	poodle := Dog{"Poodle", 15, "Bhow!"}
	fmt.Println(poodle)

	// print an attribute
	fmt.Println("Dog sound:", poodle.Sound)

	// invoke the struct method
	poodle.Speak()

	poodle.ModifySpeak()
	poodle.Speak()
}
