package main

import (
	"fmt"
	"stringutils"
)

func main() {
	fullName, _ := stringutils.FullName("Samarth", "Deyagond")
	fmt.Println(fullName)
}
