package main

import (
	"fmt"
	"os"
	"strings"
)

type Light struct {
	X     int
	Y     int
	State int
}

func ToggleState(grid [][]Light, x, y int) {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		grid[y][x].State = 1 - grid[y][x].State // If state is 0 -> 1 - 0 = 1 & If state is 1 -> 1 - 1 = 0
	}
}

func toggleLights(grid [][]Light, startingX, startingY, endingX, endingY int) {
	for y := startingY; y <= endingY; y++ {
		for x := startingX; x <= endingX; x++ {
			ToggleState(grid, x, y)
		}
	}
}

func SetState(grid [][]Light, x, y, state int) {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		grid[y][x].State = state
	}
}

func turnLights(grid [][]Light, startingX, startingY, endingX, endingY, state int) {
	for y := startingY; y <= endingY; y++ {
		for x := startingX; x <= endingX; x++ {
			SetState(grid, x, y, state)
		}
	}
}

func GetState(grid [][]Light, x, y int) int {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		return grid[y][x].State
	}
	return -1
}

func countLit(grid [][]Light) int {
	size := 1000
	count := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if GetState(grid, x, y) == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	s := strings.Split(strings.TrimSpace(string(file)), "\n")

	// Create a 1000x1000 grid
	grid1 := make([][]Light, 1000) // 1000 rows
	for i := range grid1 {
		grid1[i] = make([]Light, 1000)
	}
	grid2 := make([][]Light, 1000) // 1000 rows
	for i := range grid2 {
		grid2[i] = make([]Light, 1000)
	}

	// 400410
	fmt.Printf("Part-1: Number of lights that are lit: %d\n", part1(grid1, s))
	fmt.Printf("Part-2: Total brightness level: %d\n", part2(grid2, s))
}
