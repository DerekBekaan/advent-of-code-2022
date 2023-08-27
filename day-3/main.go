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

	fmt.Println("Part 1: ", sumPriorities(string(input)))
}

func sumPriorities(rucksackContents string) int {
	scanner := bufio.NewScanner(strings.NewReader(rucksackContents))
	total := 0
	for scanner.Scan() {
		compartment1, compartment2 := splitContent(scanner.Text())
		commonItem := getCommonItem(compartment1, compartment2)

		total += getItemValue(commonItem)
	}

	return total
}

func splitContent(rucksackContents string) (string, string) {
	mid := len(rucksackContents) / 2
	return rucksackContents[:mid], rucksackContents[mid:]
}

func getCommonItem(rucksackContents1 string, rucksackContents2 string) rune {
	type rucksackCounts struct {
		sack1 int
		sack2 int
	}

	counts := map[rune]*rucksackCounts{}

	for _, char := range rucksackContents1 {
		if val, ok := counts[char]; ok {
			val.sack1 += 1
		} else {
			counts[char] = &rucksackCounts{sack1: 1}
		}
	}
	for _, char := range rucksackContents2 {
		if val, ok := counts[char]; ok {
			val.sack2 += 1
		} else {
			counts[char] = &rucksackCounts{sack2: 1}
		}
	}

	for _, item := range itemPriority {
		if counts[item] != nil && counts[item].sack1 != 0 && counts[item].sack2 != 0 {
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
