package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func mapRouteToDistance(lines []string) map[[2]string]int {
	routes := make(map[[2]string]int)

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

	return routes
}

func routeDistance(perm []string, routes map[[2]string]int) int {
	total := 0

	for i := 0; i < len(perm)-1; i++ {
		total += routes[[2]string{perm[i], perm[i+1]}]
	}
	return total
}

func findUniqueRoutes(lines []string) []string {
	unique := []string{}
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
	return unique
}

func permutations(locations []string) [][]string {
	// base case
	if len(locations) == 0 {
		return [][]string{{}}
	}

	result := [][]string{}

	for i, l := range locations {
		remaining := make([]string, 0, len(locations)-1)
		remaining = append(remaining, locations[:i]...)
		remaining = append(remaining, locations[i+1:]...)
		for _, perm := range permutations(remaining) {
			result = append(result, append([]string{l}, perm...))
		}
	}

	return result
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace((string(file))), "\n")

	// find all unique routes in the input lines
	uniqueRoutes := findUniqueRoutes(lines)

	// map every 2 locations to one distance
	routes := mapRouteToDistance(lines)

	// all unqiue permutations
	perms := permutations(uniqueRoutes)

	fmt.Println("Part-1:", part1(routes, perms))
	fmt.Println("Part-2:", part2(routes, perms))
}
