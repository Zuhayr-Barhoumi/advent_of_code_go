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
	fmt.Println("Part 1:", solve(string(input)))
}

func solve(input string) int {
	s := strings.TrimSpace(input)
	n := 0

	for _, char := range s {
		if (char == '(') {
			n++
		} else {
			n--
		}
		fmt.Printf("Rune: %c\n", char)
	}

	return n
}
