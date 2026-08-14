package main

import (
	"strconv"
	"strings"
)

func IncreaseState(grid [][]Light, x, y, amount int) {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		grid[y][x].State += (amount)
	}
}

func increase(grid [][]Light, startingX, startingY, endingX, endingY, amount int) {
	for y := startingY; y <= endingY; y++ {
		for x := startingX; x <= endingX; x++ {
			IncreaseState(grid, x, y, amount)
		}
	}
}

func DecreaseState(grid [][]Light, x, y, amount int) {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		grid[y][x].State = max(grid[y][x].State-amount, 0)
	}
}

func decrease(grid [][]Light, startingX, startingY, endingX, endingY, amount int) {
	for y := startingY; y <= endingY; y++ {
		for x := startingX; x <= endingX; x++ {
			DecreaseState(grid, x, y, amount)
		}
	}
}

func countBrightness(grid [][]Light) int {
	size := 1000
	count := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			count += grid[y][x].State

		}
	}
	return count
}

func part2(grid [][]Light, s []string) int {
	for _, line := range s {
		// if it's a toggle
		if strings.HasPrefix(line, "toggle") {
			words := strings.Fields(line)
			starting := words[1]
			ending := words[3]

			stX, stY := strings.Split(starting, ",")[0], strings.Split(starting, ",")[1]
			endX, endY := strings.Split(ending, ",")[0], strings.Split(ending, ",")[1]

			// Convert coords to ints
			startingX, _ := strconv.Atoi(stX)
			startingY, _ := strconv.Atoi(stY)
			endingX, _ := strconv.Atoi(endX)
			endingY, _ := strconv.Atoi(endY)

			// Toggle increases brightness by 2
			increase(grid, startingX, startingY, endingX, endingY, 2)
		}

		// if it's "turn" instruction
		if strings.HasPrefix(line, "turn") {
			words := strings.Split(line, " ")
			starting := words[2]
			ending := words[4]

			stX, stY := strings.Split(starting, ",")[0], strings.Split(starting, ",")[1]
			endX, endY := strings.Split(ending, ",")[0], strings.Split(ending, ",")[1]

			// Convert coords to ints
			startingX, _ := strconv.Atoi(stX)
			startingY, _ := strconv.Atoi(stY)
			endingX, _ := strconv.Atoi(endX)
			endingY, _ := strconv.Atoi(endY)

			// Turn on increases brightness by 1
			if words[1] == "on" {
				increase(grid, startingX, startingY, endingX, endingY, 1)
			}

			// Turn off decreases brightness by 1
			if words[1] == "off" {
				decrease(grid, startingX, startingY, endingX, endingY, 1)
			}
		}
	}

	return countBrightness(grid)
}
