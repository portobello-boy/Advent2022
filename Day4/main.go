package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open input file
	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	containedPairs := 0
	overlappingPairs := 0

	// Read from file and count calories
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		boundsList := [2][2]int{}

		for index, assignment := range pair {
			bounds := strings.Split(assignment, "-")
			left, _ := strconv.Atoi(bounds[0])
			right, _ := strconv.Atoi(bounds[1])
			boundsList[index] = [2]int{left, right}
		}

		fmt.Println(boundsList)

		if boundsList[0][0] >= boundsList[1][0] && boundsList[0][1] <= boundsList[1][1] {
			containedPairs += 1
		} else if boundsList[1][0] >= boundsList[0][0] && boundsList[1][1] <= boundsList[0][1] {
			containedPairs += 1
		}

		if boundsList[0][1] >= boundsList[1][0] && boundsList[0][0] <= boundsList[1][1] {
			overlappingPairs += 1
		} else if boundsList[0][0] <= boundsList[1][0] && boundsList[0][1] >= boundsList[1][0] {
			overlappingPairs += 1
		}
	}

	fmt.Printf("containedPairs: %v\n", containedPairs)
	fmt.Printf("overlappingPairs: %v\n", overlappingPairs)
}
