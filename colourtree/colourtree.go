package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var rainbow = []string{
	"\033[38;5;196m", // Red
	"\033[38;5;208m", // Orange
	"\033[38;5;226m", // Yellow
	"\033[38;5;46m",  // Green
	"\033[38;5;21m",  // Blue
	"\033[38;5;27m",  // Indigo
	"\033[38;5;129m", // Violet
}
var reset = "\033[0m"

func makeTree(curr string, level int) {
	cmd := exec.Command("ls", "-F", curr)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}
	if len(string(output)) == 0 {
		return
	}
	cleanedOutput := strings.TrimSpace(string(output))
	fileList := strings.SplitSeq(cleanedOutput, "\n")
	for file := range fileList {
		fmt.Println(rainbow[level%7] + strings.Repeat("-", level*5) + file + reset)
		if file[len(file)-1] == '/' {
			makeTree(curr+"/"+file, level+1)
		}
	}
}

func main() {
	makeTree(".", 0)
}
