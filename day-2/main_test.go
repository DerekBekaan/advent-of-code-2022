package main

import (
	"fmt"
	"testing"
)

func TestGetMatchScore(t *testing.T) {
	type testParam struct {
		encryptedGame string
		points        int
	}

	testParams := []testParam{
		{encryptedGame: "A Y", points: 8},
		{encryptedGame: "B X", points: 1},
		{encryptedGame: "C Z", points: 6},
	}

	for _, testCase := range testParams {
		actualPoints := getMatchScore(testCase.encryptedGame, getPlayedPartOne)
		if actualPoints != testCase.points {
			t.Fatalf(fmt.Sprintf("Expected %d got %d", testCase.points, actualPoints))
		}
	}

}

func TestPartOne(t *testing.T) {
	games := `A Y
B X
C Z`
	expectedPoints := 15

	actualPoints := partOne(games)
	if actualPoints != expectedPoints {
		t.Fatalf(fmt.Sprintf("Expected %d got %d", expectedPoints, actualPoints))
	}
}
