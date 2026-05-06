package main

import (
	"strconv"
	"strings"
)

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		if line != "" {
			total += calcWrapping(line)
		}
	}
	return total
}

func calcWrapping(line string) int {
	l, w, h := parseDimensions(line)
	return 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
}

func parseDimensions(line string) (int, int, int) {
	d := strings.Split(line, "x")
	l, _ := strconv.Atoi(d[0])
	w, _ := strconv.Atoi(d[1])
	h, _ := strconv.Atoi(d[2])
	return l, w, h
}
