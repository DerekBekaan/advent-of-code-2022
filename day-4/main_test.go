package main

import "testing"

func TestSplitRanges(t *testing.T) {
	ranges := "2-4,6-8"
	expected1 := "2-4"
	expected2 := "6-8"

	actual1, actual2 := splitRanges(ranges)

	if actual1 != expected1 {
		t.Fatalf("Expected %v, got %v", expected1, actual1)
	}

	if actual2 != expected2 {
		t.Fatalf("Expected %v, got %v", expected2, actual2)
	}
}

func TestGetRangeFromString(t *testing.T) {
	expected := Range{2, 4}
	actual := getRangeFromString("2-4")

	if actual != expected {
		t.Fatalf("Expected %x, got %x", expected, actual)
	}
}

func TestRangeContainsOther(t *testing.T) {
	tests := []struct {
		range1, range2 Range
		expected       bool
	}{
		{Range{2, 8}, Range{3, 7}, true},
		{Range{3, 7}, Range{2, 8}, true},
		{Range{6, 6}, Range{4, 6}, true},
		{Range{2, 4}, Range{6, 8}, false},
		{Range{2, 6}, Range{4, 8}, false},
	}

	for _, test := range tests {
		actual := rangeContainsOther(test.range1, test.range2)
		if actual != test.expected {
			t.Fatalf("Expected %v, got %v", test.expected, actual)
		}
	}

}

func TestPartOne(t *testing.T) {
	cleaningIds := `2-4,6-8
	2-3,4-5
	5-7,7-9
	2-8,3-7
	6-6,4-6
	2-6,4-8`

	expected := 2

	actual := partOne(cleaningIds)

	if actual != expected {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
