package main

import (
	"fmt"
	"os"
	"strings"
)

type House struct {
	X, Y int
}

type Grid[T any] struct {
	cells map[House]T
}

func NewGrid[T any]() *Grid[T] {
	return &Grid[T]{
		cells: make(map[House]T),
	}
}

// Set: Stores a value at x, y
func (g *Grid[T]) Set(x, y int, gifts T) {
	g.cells[House{X: x, Y: y}] = gifts
}

// Get: Retrievs a value at x, y
// Returns the value and a boolean indicating if the cell exists (is not empty)
func (g *Grid[T]) Get(x, y int) (T, bool) {
	val, ok := g.cells[House{X: x, Y: y}]
	return val, ok
}

// Len: returns the number of cells that currently hold a value
func (g *Grid[T]) Len() int {
	return len(g.cells)
}

// markHouse increments the present count for the given house
func markHouse(location House, grid *Grid[int]) {
	if gifts, ok := grid.Get(location.X, location.Y); ok {
		grid.Set(location.X, location.Y, gifts+1)
	} else {
		grid.Set(location.X, location.Y, 1)
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if string(data) == "" {
		fmt.Println("Empty file")
		return
	}
	moves := strings.Split(strings.TrimSpace(string(data)), "")

	location := House{X: 0, Y: 0}
	grid := NewGrid[int]()

	// Set the initial house gifts to 1
	grid.Set(location.X, location.Y, 1)

	for _, char := range moves {
		switch char {
		case "^":
			location.Y++
			markHouse(location, grid)
		case ">":
			location.X++
			markHouse(location, grid)
		case "<":
			location.X--
			markHouse(location, grid)
		case "v":
			location.Y--
			markHouse(location, grid)
		}
	}
	numberOfHousesVisited := grid.Len()
	fmt.Printf("Number of houses that received at least one gift: %d \n", numberOfHousesVisited)
	// 2572
	// for house, val := range grid.cells {
	// 	fmt.Printf(" —— House: %d, %d —— Gifts: %d\n", house.X, house.Y, val)
	// }
}
