package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args[0] is always the program name.
	// We check if len(os.Args) is greater than 1 to ensure an argument was passed.
	if len(os.Args) < 2 {
		fmt.Println("Please provide a name! Example: trial Alice")
		return
	}

	// os.Args[1] is the first argument passed after the program name
	name := os.Args[1]

	fmt.Printf("Hello %s\n", name)
}