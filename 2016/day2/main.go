package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumber(line string, initial int) int {
	num := initial

	for _, c := range line {
		switch c {
		case 'U':
			if num > 3 {
				num -= 3
			}
		case 'D':
			if num < 7 {
				num += 3
			}
		case 'L':
			if num > 1 && (num-1)%3 != 0 {
				num -= 1
			}
		case 'R':
			if num < 9 && num%3 != 0 {
				num += 1
			}
		default:
			continue
		}
	}
	return num
}

func solve1(lines []string) string {
	code := make([]int, 0)
	// keypad := newKeypad(3)
	num := 5 // initially at starting position 5
	for _, line := range lines {
		num = getNumber(line, num)
		code = append(code, num)
	}

	codeString := ""
	for _, n := range code {
		codeString += strconv.Itoa(n)
	}
	return codeString
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error readig file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	for i, l := range lines {
		lines[i] = strings.TrimRight(l, "\r")
	}

	fmt.Println("Part 1:", solve1(lines))
}
