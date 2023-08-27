package main

import "testing"

func TestSumPriorities(t *testing.T) {
	rucksackContents := `vJrwpWtwJgWrhcsFMMfFFhFp
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	PmmdzqPrVvPwwTWBwg
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	ttgJtRGJQctTZtZT
	CrZsJsPPZsGzwwsLwLmpwMDw`
	expected := 157

	actual := sumPriorities(rucksackContents)
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestSplitContent(t *testing.T) {
	rucksackContents := "vJrwpWtwJgWrhcsFMMfFFhFp"
	expected1 := "vJrwpWtwJgWr"
	expected2 := "hcsFMMfFFhFp"

	actual1, actual2 := splitContent(rucksackContents)

	if actual1 != expected1 {
		t.Fatalf("Expected %s, got %s for expected1", expected1, actual1)
	}
	if actual2 != expected2 {
		t.Fatalf("Expected %s, got %s for expected2", expected2, actual2)
	}
}

func TestCommonItem(t *testing.T) {
	rucksackContents1 := "jqHRNqRjqzjGDLGL"
	rucksackContents2 := "rsFMfFZSrLrFZsSL"
	expected := 'L'

	actual := getCommonItem(rucksackContents1, rucksackContents2)

	if actual != expected {
		t.Fatalf("Expected %x, got %x", expected, actual)
	}

}

func TestGetItemValue(t *testing.T) {
	expected := 16
	actual := getItemValue('p')

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
