package main

func part2(routes map[[2]string]int, perms [][]string) int {

	maxDist := 0

	for _, perm := range perms {
		distance := routeDistance(perm, routes)
		if distance > maxDist {
			maxDist = distance
		}
	}

	return maxDist
}
