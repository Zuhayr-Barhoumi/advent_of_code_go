package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lookAndSay(input string) string {
	newSequence := strings.Builder{}

	digitCount := 1
	for i := 0; i < len(input); i++ {
		if i < len(input)-1 {
			if input[i] == input[i+1] {
				digitCount += 1
				continue
			}
		}
		newSequence.WriteString(strconv.Itoa(digitCount))

		newSequence.WriteString(string(input[i]))
		digitCount = 1
	}
	return newSequence.String()
}

func main() {
	input := "3113322113"

	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}

	fmt.Printf("Part 1: %d\n", len(input))

	input = "3113322113"

	for i := 0; i < 50; i++ {
		input = lookAndSay(input)
	}

	fmt.Printf("Part 2: %d\n", len(input))

}
