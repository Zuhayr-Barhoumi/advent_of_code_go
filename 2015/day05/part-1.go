package main

import "strings"

func isNice(s string) bool {
	for _, pair := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, pair) {
			return false
		}
	}

	vowels := 0
	for _, c := range s {
		if strings.ContainsRune("aeiou", c) {
			vowels++
		}
	}
	if vowels < 3 {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}
func part1(input []byte) int {
	nice := 0
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if isNice(line) {
			nice++
		}
	}

	return nice
}
