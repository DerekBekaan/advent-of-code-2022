package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	calories, _ := os.ReadFile("calories.txt")

	fmt.Println("Part 1: ", partOne(string(calories)))
	fmt.Println("Part 2: ", partTwo(string(calories)))
}

func sumCalories(caloriesInput string) []int {
	caloriesInput = strings.ReplaceAll(caloriesInput, "\r", "")
	caloriesSplit := strings.Split(caloriesInput, "\n")

	var summedCalories []int
	var caloriesSum = 0
	for i, calories := range caloriesSplit {
		if cals, err := strconv.Atoi(calories); err == nil {
			caloriesSum += cals
		} else if err != nil || i == len(caloriesSplit)-1 {
			summedCalories = append(summedCalories, caloriesSum)
			caloriesSum = 0
		}
	}

	return summedCalories
}

func partOne(calories string) int {
	summedCalroies := sumCalories(calories)
	sort.Ints(summedCalroies)
	return summedCalroies[len(summedCalroies)-1]
}

func partTwo(calories string) int {
	summedCalroies := sumCalories(calories)
	sort.Ints(summedCalroies)

	return summedCalroies[len(summedCalroies)-1] + summedCalroies[len(summedCalroies)-2] + summedCalroies[len(summedCalroies)-3]
}
