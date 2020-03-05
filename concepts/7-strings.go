package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicitly typed string"
	fmt.Printf("str1: %v:%T\n", str1, str1)

	var str2 string = "An explicitly typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str2))

	lValue := "Hello"
	uValue := "HELLO"

	// by default, the string comparison is case sensitive
	fmt.Println("Equal?", lValue == uValue)

	// using the method in strings package, we can perform a case non sensitive string comparison
	fmt.Println("Equal non case sensitive?", strings.EqualFold(lValue, uValue))

	// there are many more handy features like
	fmt.Println("Contains 'exp'?", strings.Contains(str2, "exp"))
}
