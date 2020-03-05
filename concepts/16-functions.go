package main

import (
	"fmt"
)

func main() {
	doSomething()

	sum := addValues(10, 50)
	fmt.Println(sum)

	sum = addAllValues(15, 30, 12, 79)
	fmt.Println(sum)

	MyFullName, _ := fullName("Samarth", "Deyagond")
	fmt.Println(MyFullName)

	MyFullName, _ = fullNameNakedReturn("Teddy", "Winters")
	fmt.Println(MyFullName)
}

func doSomething() {
	fmt.Println("Doing something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

// this is how we can send arbitrary number of parameters to the function
func addAllValues(values ...int) int {
	sum := 0
	fmt.Printf("%T\n", values)
	for index := range values {
		sum += values[index]
		index++
	}
	return sum
}

// returning multiple values from the function
func fullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

// if you simply want to return from the function and want the values to be reflected then below is the way
func fullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
