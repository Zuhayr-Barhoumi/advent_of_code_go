package main

import (
	"fmt"
	"os"
	"strings"
)

type Cell struct {
	X, Y int
	S    string
}

func nextState(currentState string, neighborsOn int) string {
	if currentState == "#" {
		if neighborsOn == 2 || neighborsOn == 3 {
			return "#"
		}
		return "."
	} else {
		if neighborsOn == 3 {
			return "#"
		}
		return "."
	}
}

func countNeighborsOn(grid [][]Cell, x, y int) int {
	count := 0
	for ix := -1; ix <= 1; ix++ {
		for iy := -1; iy <= 1; iy++ {
			if ix == 0 && iy == 0 {
				continue
			}
			neighborX, neighborY := x+ix, y+iy
			if getState(grid, neighborX, neighborY, 100) == "#" {
				count++
			}
		}
	}
	return count
}

func countOn(grid [][]Cell, size int) int {
	count := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x].S == "#" {
				count += 1
			}
		}
	}
	return count
}

func getState(grid [][]Cell, x, y, size int) string {
	if y >= 0 && y < size && x >= 0 && x < size {
		return grid[y][x].S
	}
	return "."
}

func simulate(grid [][]Cell, nSteps, size int) int {
	for step := 0; step < nSteps; step++ {
		nextGrid := make([][]Cell, size)
		for i := range nextGrid {
			nextGrid[i] = make([]Cell, size)
		}
		// calc state for each cell
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				neighborsOn := countNeighborsOn(grid, x, y)
				currentState := getState(grid, x, y, size)
				nextGrid[y][x] = Cell{
					X: x,
					Y: y,
					S: nextState(currentState, neighborsOn),
				}
			}
		}
		grid = nextGrid
	}

	return countOn(grid, size)
}

func solve1(lines []string) int {

	grid := make([][]Cell, 100)

	// Init grid
	for i := range grid {
		grid[i] = make([]Cell, 100)
	}

	// Fill grid with input
	for i, line := range lines {
		for j, r := range line {
			grid[i][j] = Cell{X: j, Y: i, S: string(r)}
		}
	}

	return simulate(grid, 100, 100)
}

func main() {
	/*
		* If the cell is on (#):

			* Stays on if it has 2 or 3 live neighbors.
			* Turns off otherwise.

		* If the cell is off (.):

			* Turns on if it has exactly 3 live neighbors.
			* Stays off otherwise.
	*/
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	fmt.Println("Part-1 — Number of on lights after 100 steps:", solve1(lines))
}
