package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// append new array element
	colors = append(colors, "Purple")
	fmt.Println(colors)

	// remove the first array element
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 17
	numbers[1] = 2
	numbers[2] = 13
	fmt.Println(numbers, cap(numbers))

	// sort the array of numbers
	sort.Ints(numbers)
	fmt.Println(numbers)
}
