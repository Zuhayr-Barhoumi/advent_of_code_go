package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
	fmt.Println("Part 2:", solve(file))
}

func solve(file *os.File) int {
	scanner := bufio.NewScanner(file)

	totalRibbon := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			totalRibbon += calc(line)
		}
	}
	return totalRibbon
}
func calc(line string) int {
	dimensions := strings.Split(line, "x")
	l, _ := strconv.Atoi(dimensions[0])
	w, _ := strconv.Atoi(dimensions[1])
	h, _ := strconv.Atoi(dimensions[2])

	perimeter := findPerimeter(l, w, h)
	bow := l * w * h

	return perimeter + bow
	// 3842356
}

func findPerimeter(l, w, h int) int {
	return 2 * min(l+w, w+h, h+l)
}
