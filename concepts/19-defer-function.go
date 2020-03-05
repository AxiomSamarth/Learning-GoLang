package main

import (
	"fmt"
)

func main() {

	// deferred statements are pushed to a stack and executed by popping one at a time
	// at the end of the function execution they are deferred in
	defer fmt.Println("This is deferred statement")
	fmt.Println("This is printed first")

	defer fmt.Println("This is statement 1")
	defer fmt.Println("This is statement 2")
	defer fmt.Println("This is statement 3")
	defer fmt.Println("This is statement 4")
}
