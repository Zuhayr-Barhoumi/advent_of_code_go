package main

import (
	"fmt"
	"os"
	"strings"
)

func permutations(locations []string) [][]string {
	// base case
	if len(locations) == 0 {
		return [][]string{{}}
	}

	result := [][]string{}

	for i, l := range locations {
		remaining := make([]string, 0, len(locations)-1)
		remaining = append(remaining, locations[:i]...)
		remaining = append(remaining, locations[i+1:]...)
		for _, perm := range permutations(remaining) {
			result = append(result, append([]string{l}, perm...))
		}
	}

	return result
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace((string(file))), "\n")

	fmt.Println("Part-1:", part1(lines))
}
