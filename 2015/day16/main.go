package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var tickerTape = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func getAunts(lines []string) []map[string]int {
	sues := make([]map[string]int, 0)

	for _, line := range lines {
		words := strings.Fields(line)

		prop1 := strings.TrimSuffix(words[2], ":")
		val1, err := strconv.Atoi(strings.TrimSuffix(words[3], ","))
		if err != nil {
			log.Fatal(err)
		}

		prop2 := strings.TrimSuffix(words[4], ":")
		val2, err := strconv.Atoi(strings.TrimSuffix(words[5], ","))
		if err != nil {
			log.Fatal(err)
		}

		prop3 := strings.TrimSuffix(words[6], ":")
		val3, err := strconv.Atoi(strings.TrimSuffix(words[7], ","))
		if err != nil {
			log.Fatal(err)
		}

		sues = append(sues, map[string]int{prop1: val1, prop2: val2, prop3: val3})
	}

	return sues
}

func solve1(lines []string) int {
	sues := getAunts(lines)
	for i, sue := range sues {
		match := true
		for prop, val := range sue {
			if tickerTape[prop] != val {
				match = false
				break
			}
		}
		if match {
			return i + 1
		}
	}
	return -1
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	fmt.Println("Part-1 — Aunt Sue Number:", solve1(lines))
}
