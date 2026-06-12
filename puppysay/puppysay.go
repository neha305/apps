package main

import (
	"fmt"
	"os"
	"strings"
)

func printMsg(msg string) {
	x := len(msg)
	topBorder := strings.Repeat("_", x+2)
	downBorder := strings.Repeat("-", x+2)
	fmt.Println(" " + topBorder)
	fmt.Println("< " + msg + " >")
	fmt.Println(" " + downBorder)
}

func printPuppy() {
	puppy := []string{
		"    \\ ",
		"     \\   __      _",
		"      \\o'')}____//",
		"        `_/      )",
		"         (_(_/-(_/",
	}
	fmt.Println(strings.Join(puppy, "\n"))
}

func puppySay(msg string) {
	printMsg(msg)
	printPuppy()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a name! Example: trial Alice")
		return
	}

	msg := strings.Join(os.Args[1:], " ")
	puppySay(msg)
}
