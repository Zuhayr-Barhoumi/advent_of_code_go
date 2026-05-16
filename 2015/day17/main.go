package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func getContainers(lines []string) []int {
	containers := make([]int, 0)
	for _, c := range lines {
		container, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal()
		}
		containers = append(containers, container)
	}

	return containers
}

func countCombos(containers []int, target int) int {
	// base cases
	if target == 0 {
		return 1
	} //good combo: perfect fit
	if target < 0 {
		return 0
	} // bad combo: containers don't fit
	if len(containers) == 0 {
		return 0
	} // no longer have containers + still havent reached target

	return countCombos(containers[1:], target-containers[0]) + countCombos(containers[1:], target)

}

func findMinCombo(containers []int, target int, containersUsed int, results *[]int) {
	if target == 0 {
		*results = append(*results, containersUsed)
		return
	}
	if target < 0 {
		return
	}
	if len(containers) == 0 {
		return
	}
	findMinCombo(containers[1:], target-containers[0], containersUsed+1, results) // include
	findMinCombo(containers[1:], target, containersUsed, results)                 // skip
}

func solve1(lines []string) int {
	containers := getContainers(lines)

	return countCombos(containers, 150)
}

func solve2(lines []string) int {
	containers := getContainers(lines)
	results := make([]int, 0)

	findMinCombo(containers, 150, 0, &results)

	minResult := math.MaxInt
	frequency := 0
	for _, r := range results {
		if r < minResult {
			minResult = r
		}
	}
	for _, r := range results {
		if r == minResult {
			frequency += 1
		}
	}
	return frequency
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	fmt.Println("Part-1: Number of container combinations to fit 150 liters:", solve1(lines))

	fmt.Println("Part-2: Number of ways to use Min Number of containers to fit 150 liters:", solve2(lines))
}
