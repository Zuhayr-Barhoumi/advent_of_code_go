package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func routeDistance(perm []string, routes map[[2]string]int) int {
	total := 0

	for i := 0; i < len(perm)-1; i++ {
		total += routes[[2]string{perm[i], perm[i+1]}]
	}
	return total
}

func part1(lines []string) int {
	// find all unique destinations in the input lines
	unique := []string{}

	// map every 2 destinations to one distance
	routes := make(map[[2]string]int)

	for _, line := range lines {
		// Split by whitespace
		words := strings.Fields(line)

		if !slices.Contains(unique, words[0]) {
			unique = append(unique, words[0])
		}
		if !slices.Contains(unique, words[2]) {
			unique = append(unique, words[2])
		}
	}

	// Store routes
	for _, line := range lines {
		words := strings.Fields(line)

		if len(words) > 5 {
			continue
		}

		dest1 := string(words[0])
		dest2 := string(words[2])

		distance, err := strconv.Atoi(words[4])
		if err != nil {
			continue
		}

		routes[[2]string{dest1, dest2}] = distance
		routes[[2]string{dest2, dest1}] = distance
	}

	perms := permutations(unique)
	minDist := math.MaxInt

	for _, perm := range perms {
		distance := routeDistance(perm, routes)
		if distance < minDist {
			minDist = distance
		}
	}

	return minDist
}
