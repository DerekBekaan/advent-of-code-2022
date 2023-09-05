package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1: ", partOne(string(input)))
}

func splitRanges(ranges string) (string, string) {
	split := strings.Split(ranges, ",")
	return split[0], split[1]
}

func getRangeFromString(rangeString string) Range {
	splitRange := strings.Split(rangeString, "-")
	start, _ := strconv.Atoi(splitRange[0])
	end, _ := strconv.Atoi(splitRange[1])
	return Range{start, end}
}

func rangeContainsOther(range1 Range, range2 Range) bool {
	return range1.start >= range2.start && range1.end <= range2.end ||
		range2.start >= range1.start && range2.end <= range1.end
}

func partOne(cleaningids string) int {
	scanner := bufio.NewScanner(strings.NewReader(cleaningids))
	containingRangeCount := 0
	for scanner.Scan() {
		range1, range2 := splitRanges(scanner.Text())
		if rangeContainsOther(getRangeFromString(range1), getRangeFromString(range2)) {
			containingRangeCount++
		}
	}

	return containingRangeCount
}
