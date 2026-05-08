package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hasForbiddenPair(line string) bool {
	forbiddenPairs := []string{"ab", "cd", "pq", "xy"}
	for _, pair := range forbiddenPairs {
		if strings.Contains(line, pair) {
			return true
		}
	}
	return false
}

func hasDoubled(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}
	return false
}

func containsVowels(line string) bool {
	count := 0
	for _, char := range line {
		if strings.ContainsRune("aeiou", char) {
			count += 1
		}
	}
	return count >= 3
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", file)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	naughtyStrings, niceStrings := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Printf("\n ——— \n")

		if hasForbiddenPair(line) {
			naughtyStrings += 1
			continue
		}
		if !hasDoubled(line) {
			naughtyStrings += 1
			continue
		}
		if !containsVowels(line) {
			naughtyStrings += 1
			continue
		}
		niceStrings += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	fmt.Printf("Part-1: Number of strings: %d\n", niceStrings)
}
