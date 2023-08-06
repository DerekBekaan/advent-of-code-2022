package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	calories, _ := os.ReadFile("calories.txt")

	println(maxCalories(string(calories)))
}

func maxCalories(calories string) int {
	calories = strings.ReplaceAll(calories, "\r", "")
	caloriesSplit := strings.Split(calories, "\n\n")
	maxCalories := []int{0, 0, 0}

	for _, calorieGroup := range caloriesSplit {
		calorieGroupSplit := strings.Split(calorieGroup, "\n")
		groupCalories := 0

		for _, individualCalories := range calorieGroupSplit {
			individualCalories, _ := strconv.Atoi(individualCalories)

			groupCalories += individualCalories
		}

		for i := 0; i < len(maxCalories); i++ {
			if groupCalories > maxCalories[i] {
				maxCalories[i] = groupCalories
				sort.Ints(maxCalories)
				break
			}
		}
	}

	maxCaloriesSum := 0
	for _, i := range maxCalories {
		maxCaloriesSum += i
	}

	return maxCaloriesSum
}
