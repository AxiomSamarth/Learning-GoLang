package main

import "fmt"

func main() {
	var p *int

	var v int = 100
	p = &v

	if p != nil {
		fmt.Println("Value of p", *p)
	} else {
		fmt.Println("p is nil")
	}

	// the pointer concept is same as that in C programming langauge
}
