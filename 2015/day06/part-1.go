package main

import (
	"strconv"
	"strings"
)

func part1(grid [][]Light, s []string) int {
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

	// 400410
	return countLit(grid)
}
