package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Street struct {
	X, Y int
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	instructions := strings.Split(strings.TrimSpace(string(file)), ", ")
	distance := 0
	direction := "north"
	street := Street{}

	for _, current := range instructions {
		digits := strings.TrimLeft(current, "RL")
		blocks, err := strconv.Atoi(digits)
		if err != nil {
			log.Fatal(err)
		}
		if rune(current[0]) == 'R' {
			switch direction {
			case "north":
				direction = "east"
				street.X += blocks
			case "east":
				direction = "south"
				street.Y -= blocks
			case "south":
				direction = "west"
				street.X -= blocks
			case "west":
				direction = "north"
				street.Y += blocks
			}

		}
		if rune(current[0]) == 'L' {
			switch direction {
			case "north":
				direction = "west"
				street.X -= blocks
			case "west":
				direction = "south"
				street.Y -= blocks
			case "south":
				direction = "east"
				street.X += blocks
			case "east":
				direction = "north"
				street.Y += blocks

			}

		}
	}

	distance = int(math.Abs(float64(street.X)) + math.Abs(float64(street.Y)))
	fmt.Println("Part-1 — Number of blocks away:", distance)
}
