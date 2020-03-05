package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var resp string

	// fmt.Print("Enter some text:")
	// fmt.Scanln(&resp)
	// fmt.Println(resp)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	resp, _ = reader.ReadString('\n')
	fmt.Println(resp)

	fmt.Print("Enter a number: ")
	resp, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(resp), 64)
	if err == nil {
		fmt.Println("The number is", f)
	}
}
