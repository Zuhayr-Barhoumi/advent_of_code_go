package main

import (
	"fmt"
	"strconv"
	"strings"
)

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func part1(lines []string) string {
	// find all unique destitions in the input lines
	unique := []string{}

	// map every 2 destinations to one distance
	routes := make(map[[2]string]int)

	for _, line := range lines {
		// Split by whitespace
		words := strings.Fields(line)

		if !contains(unique, words[0]) {
			unique = append(unique, words[0])
		}
		if !contains(unique, words[2]) {
			unique = append(unique, words[2])
		}

	}

	fmt.Println("Unique destinations:", unique)

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
	}
	for k, w := range routes {
		fmt.Println("Routes:", k, w)
	}
	return "-1" + lines[0]
}
