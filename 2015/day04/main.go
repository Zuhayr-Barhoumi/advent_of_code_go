package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	secret_key := "ckczppom"
	lowest_n := 0
	target := "00000"

	for {
		lowest_n_string := strconv.Itoa((lowest_n))
		combo := secret_key + lowest_n_string

		hashBytes := md5.Sum([]byte(combo)) // returns an array [16]byte

		fullHash := hex.EncodeToString(hashBytes[:])

		fmt.Printf("—— Input: %s —— MD5 Hash: %s \n", combo, fullHash)

		if strings.HasPrefix(fullHash, target) {
			fmt.Printf("Found Lowest Number to 5 Zeros: %s\n", lowest_n_string)
			break
		} else {
			lowest_n += 1
		}
	}
}
