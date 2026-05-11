package main

import (
	"fmt"
	"os"
	"strings"
)

func solve(lines []string) int {
	memChars := 0
	codeChars := 0
	for _, l := range lines {
		memChars += countMemChars(l)
		codeChars += len(l)
	}
	return codeChars - memChars
}

func countMemChars(line string) int {
	count := 0
	for i := 1; i < len(line)-1; i++ {
		if line[i] == '\\' {
			if line[i+1] == 'x' {
				i += 3
			} else {
				i += 1
			}
		}
		count++
	}
	return count
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	solution1 := solve(lines)

	// 1371
	fmt.Println("Part 1:", solution1)
}
