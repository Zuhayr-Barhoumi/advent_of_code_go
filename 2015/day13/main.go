package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func happinessChange(perm []string, happinessMap map[[2]string]int) int {
	total := 0

	for i := 0; i < len(perm); i++ {
		total += happinessMap[[2]string{perm[i], perm[(i+1)%len(perm)]}]
		total += happinessMap[[2]string{perm[(i+1)%len(perm)], perm[i]}]
	}

	return total
}

func getHappinessValuesMap(lines []string) map[[2]string]int {
	happinessMap := make(map[[2]string]int)

	for _, line := range lines {
		words := strings.Fields(line)

		first := words[0]
		second := words[10]

		// remove the dot at the end
		second = strings.TrimSuffix(second, ".")

		value, err := strconv.Atoi(words[3])
		if err != nil {
			fmt.Println("Error converting string to int", err)
			continue
		}
		if slices.Contains(words, "lose") {
			value = -value
		}
		happinessMap[[2]string{first, second}] = value
	}
	return happinessMap
}

func getUnique(lines []string) []string {
	unique := []string{}
	for _, line := range lines {
		words := strings.Fields(line)
		if !slices.Contains(unique, words[0]) {
			unique = append(unique, words[0])
		}
	}
	return unique
}

func permutations(uniqueVisitors []string) [][]string {
	// base case
	if len(uniqueVisitors) == 0 {
		return [][]string{{}}
	}
	result := [][]string{}

	for i, v := range uniqueVisitors {
		remaining := make([]string, 0, len(uniqueVisitors)-1)
		remaining = append(remaining, uniqueVisitors[:i]...)
		remaining = append(remaining, uniqueVisitors[i+1:]...)
		for _, perm := range permutations(remaining) {
			result = append(result, append([]string{v}, perm...))
		}
	}

	return result
}

func solve(lines []string) int {
	uniqueVisitors := getUnique(lines)

	happinessValues := getHappinessValuesMap(lines)

	perms := permutations(uniqueVisitors)

	maxChange := math.MinInt

	for _, perm := range perms {
		change := happinessChange(perm, happinessValues)
		if change > maxChange {
			maxChange = change
		}
	}

	return maxChange
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error readin file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	fmt.Println("Part 1- Optimal Total Change in Happiness:", solve(lines))
}
