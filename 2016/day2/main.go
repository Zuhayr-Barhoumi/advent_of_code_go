package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error readig file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	for i, l := range lines {
		lines[i] = strings.TrimRight(l, "\r")
	}

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}
