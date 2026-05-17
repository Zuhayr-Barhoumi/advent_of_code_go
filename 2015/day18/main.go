package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	line := strings.Split(strings.TrimSpace(string(file)), "\n")

	fmt.Println("Part-1 — Number of on lights after 100 steps:", solve1(line))
}
