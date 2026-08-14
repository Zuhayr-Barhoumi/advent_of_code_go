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
	streets := map[Street]bool{}
	lastRec := Street{0, 0}
	found := false
	visited := Street{0, 0}

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

		} else {
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

		if !found {
			visited = recordVisited(blocks, lastRec, direction, streets, &found)
		}
		lastRec = street
	}

	distance = int(math.Abs(float64(street.X)) + math.Abs(float64(street.Y)))
	firstRevisit := int(math.Abs(float64(visited.X))) + int(math.Abs(float64(visited.Y)))
	fmt.Println("Part-1 — Number of blocks away:", distance)
	fmt.Println("Part-2 — Distance to first location visisted twice:", firstRevisit)
}

func recordVisited(blocks int, lastRec Street, direction string, streets map[Street]bool, found *bool) Street {
	currentX := lastRec.X
	currentY := lastRec.Y
	visited := Street{0, 0}
	switch direction {
	case "north":
		for i := 1; i <= blocks; i++ {
			currentY += 1
			street := Street{lastRec.X, currentY}
			if !*found {
				if streets[street] {
					visited = street
					*found = true
				}
			}
			streets[street] = true
		}
	case "east":
		for i := 1; i <= blocks; i++ {
			currentX += 1
			street := Street{currentX, lastRec.Y}
			if !*found {
				if streets[street] {
					visited = street
					*found = true
				}
			}

			streets[street] = true
		}
	case "south":
		for i := 1; i <= blocks; i++ {
			currentY -= 1
			street := Street{lastRec.X, currentY}
			if !*found {
				if streets[street] {
					visited = street
					*found = true
				}
			}

			streets[street] = true
		}
	case "west":
		for i := 1; i <= blocks; i++ {
			currentX -= 1
			street := Street{currentX, lastRec.Y}
			if !*found {
				if streets[street] {
					visited = street
					*found = true
				}
			}

			streets[street] = true
		}
	}
	return visited
}
