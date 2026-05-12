package main

func part2(routes map[[2]string]int, perms [][]string) int {

	minDist := 0

	for _, perm := range perms {
		distance := routeDistance(perm, routes)
		if distance > minDist {
			minDist = distance
		}
	}

	return minDist
}
