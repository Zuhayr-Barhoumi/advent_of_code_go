package main

import (
	"fmt"
	"os"
	"strconv"
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

func LightsCount(grid [][]Light, size int) int {
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

	s := strings.Split(string(file), "\n")
	// grid := NewGrid()

	// Create a 1000x1000 grid filled with zeros
	grid := make([][]Light, 1000) // 1000 rows
	for y := 0; y < 1000; y++ {
		grid[y] = make([]Light, 1000)
		for x := 0; x < 1000; x++ {
			grid[y][x] = Light{
				X:     x,
				Y:     y,
				State: 0,
			}
		}
	}

	for _, line := range s {
		if line == "" {
			fmt.Println("Empty file")
			continue
		}
		// if it's a toggle
		if strings.HasPrefix(line, "toggle") {
			words := strings.Split(line, " ")
			starting := words[1]
			ending := words[3]

			stX, stY := strings.Split(starting, ",")[0], strings.Split(starting, ",")[1]
			endX, endY := strings.Split(ending, ",")[0], strings.Split(ending, ",")[1]

			// Convert coords to ints
			startingX, _ := strconv.Atoi(stX)
			startingY, _ := strconv.Atoi(stY)
			endingX, _ := strconv.Atoi(endX)
			endingY, _ := strconv.Atoi(endY)

			toggleLights(grid, startingX, startingY, endingX, endingY)
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

			// Turn lights on
			if words[1] == "on" {
				turnLights(grid, startingX, startingY, endingX, endingY, 1)
			}

			// Turn lights off
			if words[1] == "off" {
				turnLights(grid, startingX, startingY, endingX, endingY, 0)
			}
		}
	}

	fmt.Printf("Part-1: Number of lights that are lit: %d\n", LightsCount(grid, 1000))
}
