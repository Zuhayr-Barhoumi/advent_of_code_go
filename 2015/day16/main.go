package main

import (
	"fmt"
	"os"
	"strings"
)

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

	fmt.Println("Part-1 — Aunt Sue Number:", solve1(lines))
}
