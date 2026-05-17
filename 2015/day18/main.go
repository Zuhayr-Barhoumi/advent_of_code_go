package main

import (
	"fmt"
	"os"
	"strings"
)

type Light struct {
	X, Y int
	S    string
}

func solve1(lines []string) int {
	grid := make([][]Light, 100)

	// Init grid
	for i := range grid {
		grid[i] = make([]Light, 100)
	}

	// Fill grid with input
	for i, line := range lines {
		for j, r := range line {
			grid[i][j] = Light{X: i, Y: j, S: string(r)}
		}
	}

	return -1
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	fmt.Println("Part-1 — Number of on lights after 100 steps:", solve1(lines))
}
