package main

import (
	"fmt"
	"sort"
)

func main() {

	// declare and intialize the maps with the make function
	states := make(map[string]string)
	fmt.Println(states)

	// populate the map with the key values
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"

	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	// to delete an item from the map, use the inbuilt delete function
	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"

	// to iterate over the maps you can use the below syntax and function range
	for key, value := range states {
		fmt.Println(key, ":", value)
	}

	// extract the keys from the maps as a slice
	keys := make([]string, len(states))
	i := 0

	// when the range is assigned to only one variable on the lieft hand side, it pics the key and ignores the value
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted:")

	// when the range is supplied with a slice, it returns an integer over the iteration that is same as the index of the slice
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
