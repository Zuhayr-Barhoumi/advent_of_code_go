package main

import (
	"fmt"
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
	return "-1" + lines[0]
}
