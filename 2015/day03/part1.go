package main

func part1(moves []string) int {
	/* Create a 2D grid
	 * Set the starting location as visited (0, 0)
	 * Take the next move:
		* north (^) add 1 to row which is the first value,
		* south (v) subtract 1 from the row,
		* east (>) add 1 to the column which is the second value,
		* west subtract 1 from the column(<)
		* After any of those moves mark the cell as visited once more
	*/

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
	return numberOfHousesVisited
	// fmt.Printf("Number of houses that received at least one gift: %d \n", numberOfHousesVisited)

	// 2572
	// for house, val := range grid.cells {
	// 	fmt.Printf(" —— House: %d, %d —— Gifts: %d\n", house.X, house.Y, val)
	// }

}
