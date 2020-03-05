package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("day", dow)

	result := ""

	switch dow {
	case 1:
		result = "It is Sunday"

	case 7:
		result = "It is Saturday"

	default:
		result = "It is a weekday"
	}

	fmt.Println("Day", dow, ",", result)

}
