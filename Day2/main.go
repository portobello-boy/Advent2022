package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Hand struct {
	beats string
	loses string
	value int
}

// Set match status scores
const lose, draw, win int = 0, 3, 6

func main() {
	// Generate matchupsRoundOne (can't do memory allocation globally in Go)
	matchupsRoundOne := map[string]Hand{
		"X": {"C", "B", 1},
		"Y": {"A", "C", 2},
		"Z": {"B", "A", 3},
	}

	matchupsRoundTwo := map[string]map[string]int{
		"X": { // Lose
			"A": 3,
			"B": 1,
			"C": 2,
		},
		"Y": { // Draw
			"A": 1,
			"B": 2,
			"C": 3,
		},
		"Z": { // Win
			"A": 2,
			"B": 3,
			"C": 1,
		},
	}

	resultScoreRoundTwo := map[string]int{
		"X": lose,
		"Y": draw,
		"Z": win,
	}

	// Open input file
	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	totalScoreRoundOne := 0
	totalScoreRoundTwo := 0

	// Read from file and count calories
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := strings.Split(scanner.Text(), " ")

		result := matchupsRoundOne[v[1]]

		handScore := result.value
		matchScore := 3

		if v[0] == result.beats {
			matchScore = 6
		} else if v[0] == result.loses {
			matchScore = 0
		}

		totalScoreRoundOne += (handScore + matchScore)

		handScore = matchupsRoundTwo[v[1]][v[0]]
		matchScore = resultScoreRoundTwo[v[1]]

		totalScoreRoundTwo += (handScore + matchScore)
	}

	fmt.Println(totalScoreRoundOne)
	fmt.Println(totalScoreRoundTwo)
}
