package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Part-1: Number of nice strings: %d\n", part1(input))
	fmt.Printf("Part-2: Number of nice strings: %d\n", part2(input))
}
