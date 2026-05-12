package main

import (
	"math"
)

func part2(lines []string) int {
	// find all unique routes in the input lines
	uniqueRoutes := findUniqueRoutes(lines)

	// map every 2 destinations to one distance
	routes := mapDistToRoute(lines, uniqueRoutes)

	// all unqiue permutations
	perms := permutations(uniqueRoutes)

	minDist := math.MaxInt

	for _, perm := range perms {
		distance := routeDistance(perm, routes)
		if distance < minDist {
			minDist = distance
		}
	}

	return minDist
}
