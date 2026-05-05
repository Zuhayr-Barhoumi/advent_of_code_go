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
		return
	}
	defer file.Close()
	fmt.Println("Part 1:", solve(file))
}

func solve(file *os.File) int {
	scanner := bufio.NewScanner(file)

	totalWrappingPaper := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			totalWrappingPaper += calc(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return totalWrappingPaper
}

func calc(line string) int {
	dimensions := strings.Split(line, "x")
	l, _ := strconv.Atoi(dimensions[0])
	w, _ := strconv.Atoi(dimensions[1])
	h, _ := strconv.Atoi(dimensions[2])

	area := findArea(l, w, h)

	smallestSide := findSmallestSide(l, w, h)

	return area + smallestSide
}

func findArea(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

func findSmallestSide(l, w, h int) int {
	return min(l*w, w*h, h*l)
	// 1606483
}
