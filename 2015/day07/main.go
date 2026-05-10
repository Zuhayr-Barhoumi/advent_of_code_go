package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	lines := strings.Split(strings.TrimSpace((string(file))), "\n")

	wires := make(map[string]string)
	cache := make(map[string]int)

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		wire := parts[1]
		instruction := parts[0]
		wires[wire] = instruction
	}

	fmt.Println("Part 1:", evaluate("a", wires, cache))

}

func evaluate(wire string, wires map[string]string, cache map[string]int) int {
	// return the value if it's cached
	if val, exists := cache[wire]; exists {
		return val
	}
	// return wire value as is if it's actually a plain number
	if val, err := strconv.Atoi(wire); err == nil {
		return val
	}

	// otherwise get the instructino for the wire from the wires map
	instruction := wires[wire]
	words := strings.Fields((instruction))

	var result int
	// evaluate based on instruction shape
	// length 1 : "123", 2: "Not x", 3: "x AND y"
	switch {
	case len(words) == 1:
		result = evaluate(words[0], wires, cache)
	case len(words) == 2:
		result = ^evaluate(words[1], wires, cache) & 0xFFFF // Go's ^ operator flips all bits producing a large (64 bit) negative number? The mask keeps only the lower 16 bits
	case words[1] == "AND":
		result = evaluate(words[0], wires, cache) & evaluate(words[2], wires, cache)
	case words[1] == "OR":
		result = evaluate(words[0], wires, cache) | evaluate(words[2], wires, cache)
	case words[1] == "LSHIFT":
		shift, _ := strconv.Atoi(words[2])
		result = evaluate(words[0], wires, cache) << shift
	case words[1] == "RSHIFT":
		shift, _ := strconv.Atoi(words[2])
		result = evaluate(words[0], wires, cache) >> shift
	}

	cache[wire] = result
	return result
}
