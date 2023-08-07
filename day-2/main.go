package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	Rock     RockPaperScissors = 1
	Paper    RockPaperScissors = 2
	Scissors RockPaperScissors = 3
)

type RockPaperScissors int8

type matchOutcomes struct {
	winner, loser RockPaperScissors
}

var matchOutcomeResults = []matchOutcomes{
	{Rock, Scissors},
	{Paper, Rock},
	{Scissors, Paper},
}

var opponentRPS = map[string]RockPaperScissors{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1: ", partOne(string(input)))
	fmt.Println("Part 2: ", partTwo(string(input)))
}

func partOne(encryptedGame string) int {
	totalPoints := 0
	scanner := bufio.NewScanner(strings.NewReader(encryptedGame))
	for scanner.Scan() {
		totalPoints += getMatchScore(scanner.Text(), getPlayedPartOne)
	}
	return totalPoints
}

func partTwo(encryptedGame string) int {
	totalPoints := 0
	scanner := bufio.NewScanner(strings.NewReader(encryptedGame))
	for scanner.Scan() {
		totalPoints += getMatchScore(scanner.Text(), getPlayedPartTwo)
	}
	return totalPoints
}

func getPlayedPartOne(encryptedGame string) (RockPaperScissors, RockPaperScissors) {
	var yourRPS = map[string]RockPaperScissors{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	playedGame := strings.Split(encryptedGame, " ")
	opponentPlay := opponentRPS[playedGame[0]]
	yourPlay := yourRPS[playedGame[1]]

	return opponentPlay, yourPlay
}

func getPlayedPartTwo(encryptedGame string) (RockPaperScissors, RockPaperScissors) {
	playedGame := strings.Split(encryptedGame, " ")
	opponentPlay := opponentRPS[playedGame[0]]

	losingPlay, _ := getLosingPlay(opponentPlay)
	winningPlay, _ := getWinningPlay(opponentPlay)
	var yourRPS = map[string]RockPaperScissors{
		"X": losingPlay,
		"Y": opponentPlay,
		"Z": winningPlay,
	}

	yourPlay := yourRPS[playedGame[1]]

	return opponentPlay, yourPlay
}

func getMatchScore(encryptedGame string, playedStrategy func(string) (RockPaperScissors, RockPaperScissors)) int {
	opponentPlay, yourPlay := playedStrategy(encryptedGame)

	matchScore := 0

	if yourPlay == opponentPlay {
		matchScore += 3
	} else if winningPlay, _ := getWinningPlay(opponentPlay); yourPlay == winningPlay {
		matchScore += 6
	}

	matchScore += int(yourPlay)

	return matchScore
}

func getLosingPlay(play RockPaperScissors) (RockPaperScissors, error) {
	for _, matchOutcomeResult := range matchOutcomeResults {
		if play == matchOutcomeResult.winner {
			return matchOutcomeResult.loser, nil
		}
	}
	return 0, errors.New("losing play not found")
}

func getWinningPlay(play RockPaperScissors) (RockPaperScissors, error) {
	for _, matchOutcomeResult := range matchOutcomeResults {
		if play == matchOutcomeResult.loser {
			return matchOutcomeResult.winner, nil
		}
	}
	return 0, errors.New("winning play not found")
}
