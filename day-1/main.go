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

		maxCalories = addToSliceIfBigger(groupCalories, maxCalories)
	}

	return sumIntSlice(maxCalories)
}

func addToSliceIfBigger(intToAdd int, ints []int) []int {
	for i := 0; i < len(ints); i++ {
		if intToAdd > ints[i] {
			ints[i] = intToAdd
			sort.Ints(ints)
			break
		}
	}

	return ints
}

func sumIntSlice(ints []int) int {
	maxCaloriesSum := 0
	for _, i := range ints {
		maxCaloriesSum += i
	}

	return maxCaloriesSum
}
