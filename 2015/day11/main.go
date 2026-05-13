package main

import "fmt"

func main() {
	/*
		RULES:
			- Include >= 3 letters increasing (tripples: "abc", "bcd",... "xyz")
			- Not contain i, o or l
			- Include >= 2 different pairs non-overlapping (example: "aa", "bb", "zz")
	*/

	input := "cqjxjnds"
	nextPassword := getNextValid(input)
	fmt.Println("Part-1:", nextPassword)

	nextPassword = getNextValid(nextPassword)
	fmt.Println("Part-2:", nextPassword)

}

func getNextValid(input string) string {
	pwd := []byte(input)

	for {
		incrementOnce(pwd)
		if validPassword(pwd) {
			return string(pwd)
		}
	}
}

func validPassword(pwd []byte) bool {
	return ruleOneCheck(pwd) && ruleTwoCheck(pwd) && ruleThreeCheck(pwd)
}

func incrementOnce(pwd []byte) {
	for i := len(pwd) - 1; i >= 0; i-- {
		if pwd[i] != 'z' {
			pwd[i] += 1
			break
		}
		pwd[i] = 'a'
	}
}

func ruleOneCheck(pwd []byte) bool {
	check := false
	for i := 0; i < len(pwd)-2; i++ {
		if pwd[i]+1 == pwd[i+1] {
			if pwd[i]+2 == pwd[i+2] {
				check = true
				break
			}
		}
	}
	return check
}

func ruleTwoCheck(pwd []byte) bool {
	check := true
	for _, c := range pwd {
		if c == 'i' || c == 'l' || c == 'o' {
			check = false
			break
		}
	}
	return check
}

func ruleThreeCheck(pwd []byte) bool {
	var lastPairChar byte
	count := 0

	for i := 0; i < len(pwd)-1; i++ {
		if pwd[i] != lastPairChar {
			if pwd[i] == pwd[i+1] {
				lastPairChar = pwd[i]
				count++
				if count > 1 {
					// We got 2 pairs
					break
				}
				if i+2 < len(pwd) {
					// skip the next letter since pair is found
					i += 1
				} else {
					// reached end — no second pair
					break
				}
			}
		}

	}
	return count >= 2
}
