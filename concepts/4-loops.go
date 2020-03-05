package main

import (
	"fmt"
)

func main() {

	// GoLang has only one loop construct - the 'for' loop
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	// We can skip the init and post condition of the loop and still
	// like a charm. Just like C's while loop
	i := 0
	for i < 10 {
		fmt.Print(i)
		i++
	}
}
