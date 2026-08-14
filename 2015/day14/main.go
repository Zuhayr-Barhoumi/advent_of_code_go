package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TIME = 2503

type Reindeer struct {
	name                                string
	speed, flightDuration, restDuration int
}

func calcScores(reindeers []Reindeer) []int {
	scores := make([]int, len(reindeers))
	for i := 1; i <= TIME; i++ {
		leadDistance := 0
		for _, reindeer := range reindeers {
			distance := calcDistance(reindeer, i)
			if distance > leadDistance {
				leadDistance = distance
			}
		}
		for j, reindeer := range reindeers {
			distance := calcDistance(reindeer, i)
			if distance == leadDistance {
				scores[j]++
			}
		}
	}
	return scores
}

func calcDistance(r Reindeer, time int) int {
	completeCycles := time / (r.flightDuration + r.restDuration)

	leftOverTime := time % (r.flightDuration + r.restDuration)

	// Distance from complete cycles
	distance := completeCycles*r.flightDuration*r.speed + (min(leftOverTime, r.flightDuration) * r.speed)

	return distance
}

func getReindeers(lines []string) []Reindeer {
	reindeers := []Reindeer{}
	for _, line := range lines {
		words := strings.Fields(line)

		n := string(words[0])
		s, _ := strconv.Atoi(words[3])
		f, _ := strconv.Atoi(words[6])
		r, _ := strconv.Atoi(words[13])

		reindeer := Reindeer{
			name:           n,
			speed:          s,
			flightDuration: f,
			restDuration:   r,
		}

		reindeers = append(reindeers, reindeer)
	}
	return reindeers
}

func solve1(lines []string) int {
	reindeers := getReindeers(lines)

	winningDistance := 0
	for _, r := range reindeers {
		distance := calcDistance(r, TIME)
		if distance > winningDistance {
			winningDistance = distance
		}
	}

	return winningDistance
}

func solve2(lines []string) int {
	reindeers := getReindeers(lines)
	scores := calcScores(reindeers)

	winningScore := 0
	for _, score := range scores {
		if score > winningScore {
			winningScore = score
		}
	}
	return winningScore
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	fmt.Println("Part-1:", solve1(lines))
	fmt.Println("Part-2:", solve2(lines))
}
