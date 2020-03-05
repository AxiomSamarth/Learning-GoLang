package main

import (
	"fmt"
)

func add(x float64, y float64) float64 {
	/* 	this is how we comment multiline comment
	in GoLang */
	return x + y
}

// this is one line comment
func multiple(a, b string) (string, string) {
	return a, b
}

// this is the entry point for the GoLang programs
func main() {
	var num1, num2 float64 = 10.5, 15.5
	fmt.Println(add(num1, num2))

	var w1, w2 string = "Hi", "Ted"
	fmt.Println(multiple(w1, w2))

	fmt.Print("This is how we print ", w1, " and ", w2, "\n")
	fmt.Printf("This is how we print %s and %s\n", w1, w2)
	fmt.Println("This is how we print", w1, "and", w2)

	// this is the preferred way of declaring variables by developers of GoLang
	anInteger := 100
	fmt.Println(anInteger)
}
