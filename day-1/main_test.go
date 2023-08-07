package main

import (
	"fmt"
	"testing"
)

var calories = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPartOne(t *testing.T) {
	expected := 24000

	actual := partOne(calories)

	if actual != expected {
		t.Fatalf(fmt.Sprintf("Expected %d got %d", expected, actual))
	}
}

func TestPartTwo(t *testing.T) {
	expected := 45000

	actual := partTwo(calories)

	if actual != expected {
		t.Fatalf(fmt.Sprintf("Expected %d got %d", expected, actual))
	}
}
