package main

import "strings"

func newKeypad() [5][5]byte {
	keypad := [5][5]byte{
		{0, 0, '1', 0, 0},
		{0, '2', '3', '4', 0},
		{'5', '6', '7', '8', '9'},
		{0, 'A', 'B', 'C', 0},
		{0, 0, 'D', 0, 0},
	}

	return keypad
}

func getButton(line string, row, col int, keypad [5][5]byte) (int, int, byte) {
	r, c := row, col
	for _, ch := range line {
		rn, cn := r, c
		switch ch {
		case 'U':
			rn--
		case 'D':
			rn++
		case 'L':
			cn--
		case 'R':
			cn++
		}
		if rn >= 0 && rn < 5 && cn >= 0 && cn < 5 && keypad[rn][cn] != 0 {
			r, c = rn, cn
		}
	}
	return r, c, keypad[r][c]
}

func solve2(lines []string) string {
	var btn byte
	var sb strings.Builder

	keypad := newKeypad()
	row := 2
	col := 0

	for _, line := range lines {
		row, col, btn = getButton(line, row, col, keypad)
		sb.WriteByte(btn)
	}

	return sb.String()
}

func part2(lines []string) string {
	return string(solve2(lines))
}
