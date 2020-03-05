package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	// working with integer math
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Sum:", intSum)

	// working with float math
	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum without precision", floatSum)
	fmt.Printf("Float sum: %.10g\n", floatSum)

	// math package has sub package called big which helps us deal with big numbers
	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(23.5)
	b2.SetFloat64(65.1)
	b3.SetFloat64(76.3)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("Big sum: %.10g\n", &bigSum)

	// GoLang's math package also has some constant value
	radius := 15.5
	circumference := radius * 2 * math.Pi
	fmt.Println("Circumference of the circle", circumference)
}
