package main

import (
	"fmt"
	"os"
	"strings"
)

type ingredient struct {
	capacity, durability, flavor, texture, calories int
}

func solve1(lines []string) int {
	return -1
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	fmt.Println("Part-1 — Total Score:", solve1(lines))
}
