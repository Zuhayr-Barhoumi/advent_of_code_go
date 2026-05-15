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

func calcDistance(r Reindeer) int {
	completeCycles := TIME / (r.flightDuration + r.restDuration)
	leftOverTime := TIME % (r.flightDuration + r.restDuration)

	distance := (completeCycles * r.flightDuration * r.speed) + (min(leftOverTime, r.flightDuration) * r.speed)

	return distance
}

func getDistancesMap(reindeers []Reindeer) map[Reindeer]int {
	reinDistance := make(map[Reindeer]int)

	for _, r := range reindeers {
		distance := calcDistance(r)
		reinDistance[r] = distance
	}
	return reinDistance
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
	distancesMap := getDistancesMap(reindeers)
	winningDeerDistance := 0

	for i := 0; i < len(distancesMap); i++ {
		// fmt.Printf("%v %v\n", reindeers[i], distancesMap[reindeers[i]])
		if distancesMap[reindeers[i]] > winningDeerDistance {
			winningDeerDistance = distancesMap[reindeers[i]]
		}
	}
	return winningDeerDistance
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	fmt.Println("Part-1:", solve1(lines))
}
