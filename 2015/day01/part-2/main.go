package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the entire input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Part 2:", solve(string(input)))
}

func solve(input string) int {
	s := strings.TrimSpace(input)
	l := 0

	for i, char := range s {
		if char == '(' {
			l++
		} else if char == ')' {
			l--
		}
		if l == -1 {
			return i + 1
		}
	}
	return -1
}
