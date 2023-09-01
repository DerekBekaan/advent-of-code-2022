package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var itemPriority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1: ", partOne(string(input)))
	fmt.Println("Part 2: ", partTwo(string(input)))
}

func partOne(allRucksacks string) int {
	total := 0

	scanner := bufio.NewScanner(strings.NewReader(string(allRucksacks)))
	for scanner.Scan() {
		compartment1, compartment2 := splitContent(scanner.Text())
		commonItem := getCommonItem([]string{compartment1, compartment2})

		total += getItemValue(commonItem)
	}

	return total
}

func partTwo(allRucksacks string) int {
	rucksackCount := 0
	total := 0
	rucksacks := [3]string{}

	scanner := bufio.NewScanner(strings.NewReader(string(allRucksacks)))
	for scanner.Scan() {
		rucksacks[rucksackCount] = scanner.Text()

		if rucksackCount == 2 {
			total += getItemValue(getCommonItem(rucksacks[:]))
			rucksackCount = 0
		} else {
			rucksackCount++
		}
	}

	return total
}

func splitContent(rucksackContents string) (string, string) {
	mid := len(rucksackContents) / 2
	return rucksackContents[:mid], rucksackContents[mid:]
}

func getCommonItem(rucksacks []string) rune {
	allRucksackCounts := map[rune]int{}

	for _, rucksackContents := range rucksacks {
		inCurrentRucksack := map[rune]bool{}
		for _, char := range rucksackContents {
			if !inCurrentRucksack[char] {
				inCurrentRucksack[char] = true
				allRucksackCounts[char]++
			}
		}
	}

	for _, item := range itemPriority {
		if allRucksackCounts[item] == len(rucksacks) {
			return item
		}
	}

	return rune(0)
}

func getItemValue(item rune) int {
	for index, value := range itemPriority {
		if item == value {
			return index + 1
		}
	}
	return 0
}
