package main

import (
	"math"
)

func part1(routes map[[2]string]int, perms [][]string) int {

	minDist := math.MaxInt

	for _, perm := range perms {
		distance := routeDistance(perm, routes)
		if distance < minDist {
			minDist = distance
		}
	}

	return minDist
}
