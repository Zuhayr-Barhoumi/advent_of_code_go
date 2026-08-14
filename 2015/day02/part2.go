package main

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		if line != "" {
			total += calcRibbon(line)
		}
	}
	return total
}

func calcRibbon(line string) int {
	l, w, h := parseDimensions(line)
	return 2*min(l+w, w+h, h+l) + l*w*h
}
