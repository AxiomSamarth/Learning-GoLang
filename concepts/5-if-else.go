package main

import (
	"fmt"
	"math"
)

func isLessThanTen(x int64) bool {
	if x < 10 {
		return true
	}
	return false
}

// if can also be implemented in a shorthand way
// Variables declared by the statement are only in scope until the end of the if.

func pow(x, n, limit float64) float64 {
	if result := math.Pow(x, n); result < limit {
		return result
	}
	return limit
}

func main() {
	result := isLessThanTen(15)
	fmt.Println(result)

	result = isLessThanTen(5)
	fmt.Println(result)

	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}
