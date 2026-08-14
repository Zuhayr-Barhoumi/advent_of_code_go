package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

func getIngredients(lines []string) []ingredient {
	ingredients := []ingredient{}

	for _, line := range lines {
		words := strings.Fields(line)
		name := strings.TrimSuffix(string(words[0]), ":")
		capacity, _ := strconv.Atoi(strings.TrimSuffix(words[2], ","))
		durability, _ := strconv.Atoi(strings.TrimSuffix(words[4], ","))
		flavor, _ := strconv.Atoi(strings.TrimSuffix(words[6], ","))
		texture, _ := strconv.Atoi(strings.TrimSuffix(words[8], ","))
		calories, _ := strconv.Atoi(strings.TrimSuffix(words[10], ","))

		ing := ingredient{
			name,
			capacity,
			durability,
			flavor,
			texture,
			calories,
		}

		ingredients = append(ingredients, ing)
	}
	return ingredients
}

func calcRecipeScore(ingredients []ingredient, recipe []int) (int, int) {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	calories := 0

	for i := range ingredients {
		capacity += ingredients[i].capacity * recipe[i]
		durability += ingredients[i].durability * recipe[i]
		flavor += ingredients[i].flavor * recipe[i]
		texture += ingredients[i].texture * recipe[i]
		calories += ingredients[i].calories * recipe[i]
	}

	capacity = max(capacity, 0)
	durability = max(durability, 0)
	flavor = max(flavor, 0)
	texture = max(texture, 0)

	score := capacity * durability * flavor * texture
	return score, calories
}

func generateRecipes() [][]int {

	recipes := [][]int{}

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				recipes = append(recipes, []int{i, j, k, 100 - i - j - k})
			}
		}
	}
	return recipes
}

func solve1(lines []string) int {
	ingredients := getIngredients(lines)

	recipeCombos := generateRecipes()

	bestTotal := math.MinInt

	for _, r := range recipeCombos {
		total, _ := calcRecipeScore(ingredients, r)
		if total > bestTotal {
			bestTotal = total
		}
	}

	return bestTotal
}

func solve2(lines []string) int {
	ingredients := getIngredients(lines)

	recipeCombos := generateRecipes()

	bestTotal := math.MinInt

	for _, r := range recipeCombos {
		total, cals := calcRecipeScore(ingredients, r)

		if cals == 500 && total > bestTotal {
			bestTotal = total
		}
	}

	return bestTotal
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	fmt.Println("Part-1 — Total Score:", solve1(lines))
	fmt.Println("Part-2 — Total Score:", solve2(lines))
}
