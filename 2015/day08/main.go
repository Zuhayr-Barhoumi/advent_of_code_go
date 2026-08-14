package main

import (
	"fmt"
	"os"
	"strings"
)

func solve(lines []string) int {
	memChars := 0
	codeChars := 0
	for _, l := range lines {
		memChars += countMemChars(l)
		codeChars += len(l)
	}
	return codeChars - memChars
}

func countMemChars(line string) int {
	count := 0
	for i := 1; i < len(line)-1; i++ {
		if line[i] == '\\' {
			if line[i+1] == 'x' {
				i += 3
			} else {
				i += 1
			}
		}
		count++
	}
	return count
}

func encode(line string) string {
	sb := strings.Builder{}

	sb.WriteByte('"') // Opening

	for i := 0; i < len(line); i++ {
		if line[i] == '\\' || line[i] == '"' {
			sb.WriteString("\\")
		}
		sb.WriteByte(line[i])
	}

	sb.WriteByte('"') // Closing

	return sb.String()
}

func solveEncoded(lines []string) int {
	encodedChars := 0
	codeChars := 0
	for _, l := range lines {
		codeChars += len(l)
		encodedLine := encode(l)
		encodedChars += len(encodedLine)
	}
	return encodedChars - codeChars
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	solution1 := solve(lines)
	solution2 := solveEncoded(lines)

	// 1371
	fmt.Println("Part 1:", solution1)
	// 2117
	fmt.Println("Part 2:", solution2)
}
