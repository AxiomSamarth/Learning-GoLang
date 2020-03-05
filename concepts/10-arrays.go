package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"
	fmt.Println(colors)

	// print the array elements by indices
	fmt.Println(colors[1])

	// initialize the array with array items
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// we can also print length
	fmt.Println("Length of colors:", len(colors))
	fmt.Println("Length of numbers:", len(numbers))
}
