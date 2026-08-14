package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func findTargetHashPrefix(target string) string {
	secretKey := "ckczppom"
	lowestN := 0
	for {
		lowest_n_string := strconv.Itoa((lowestN))
		combo := secretKey + lowest_n_string

		hashBytes := md5.Sum([]byte(combo)) // returns an array [16]byte

		fullHash := hex.EncodeToString(hashBytes[:])

		// fmt.Printf("—— Input: %s —— MD5 Hash: %s \n", combo, fullHash)

		if strings.HasPrefix(fullHash, target) {
			return lowest_n_string
		} else {
			lowestN += 1
		}
	}
}

func main() {
	target1 := "00000"
	target2 := "000000"

	part1Solution := findTargetHashPrefix(target1)
	part2Solution := findTargetHashPrefix(target2)

	fmt.Printf("Part-1: Found Lowest Number to 5 Zeros: %s\n", part1Solution)
	fmt.Printf("Part-2: Found Lowest Number to 5 Zeros: %s\n", part2Solution)
}
