package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go!"

	// creating and writing to a file
	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("All done with the file handling. Wrote %v characters.\n", ln)

	// read a file
	result, err := ioutil.ReadFile("./fromString.txt")
	checkError(err)
	content = string(result)
	fmt.Println("Read the file successfully. Content:", content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
