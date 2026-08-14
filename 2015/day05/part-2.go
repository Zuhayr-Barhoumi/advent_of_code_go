package main

import "strings"

func hasRepeatedPair(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		// take a pair and compare it with the rest of the letters (substring)
		pair := s[i : i+2]
		if strings.Contains(s[i+2:], pair) {
			return true
		}
	}
	return false
}

func isNice2(s string) bool {
	if !hasRepeatedPair(s) {
		return false
	}

	// has repeated leter with a gap
	for i := 0; i < len(s)-2; i++ {
		if s[i+2] == s[i] {
			return true
		}

	}
	return false
}
func part2(input []byte) int {
	nice := 0
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if isNice2(line) {
			nice++
		}
	}

	return nice
}
