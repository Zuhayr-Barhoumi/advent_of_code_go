package main

func part2(moves []string) int {
	santaLocation := House{X: 0, Y: 0}
	roboLocation := House{X: 0, Y: 0}

	grid := NewGrid[int]()

	// Set the initial house gifts to 1
	grid.Set(0, 0, 2)

	for i, char := range moves {
		location := &santaLocation
		if i%2 == 1 {
			location = &roboLocation
		}

		// Apply movement
		switch char {
		case "^":
			location.Y++
		case ">":
			location.X++
		case "<":
			location.X--
		case "v":
			location.Y--
		default:
			continue
		}
		markHouse(*location, grid)
	}
	numberOfHousesVisited := grid.Len()
	return numberOfHousesVisited
	// 2631
}
