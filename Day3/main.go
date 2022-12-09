package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	mapset "github.com/deckarep/golang-set/v2"
)

func getPriority(c rune) int {
	if unicode.IsLower(c) {
		return int(c) - int('a') + 1
	}
	return int(c) - int('A') + 27
}

func main() {
	// Open input file
	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Set variables
	totalPriority := 0
	badgePriority := 0
	groupIndex := 0
	groupSet := [3]mapset.Set[rune]{}

	// Initialize sets for finding badge priority
	for i := range groupSet {
		groupSet[i] = mapset.NewSet[rune]()
	}

	// Read from file and count calories
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		strLen := len(str)

		// Split compartment strings
		compartmentA := str[0:(strLen / 2)]
		compartmentB := str[(strLen / 2):]

		compartmentASet := mapset.NewSet[rune]()
		compartmentBSet := mapset.NewSet[rune]()

		// Add runes to sets (since sets are guaranteed to be the same size, we need just one loop)
		for i := range compartmentA {
			compartmentASet.Add(rune(compartmentA[i]))
			compartmentBSet.Add(rune(compartmentB[i]))

			// Add the compartment runes to the combined set for badge priority
			groupSet[groupIndex].Add(rune(compartmentA[i]))
			groupSet[groupIndex].Add(rune(compartmentB[i]))
		}

		// Find common items between each compartment
		intersection := compartmentASet.Intersect(compartmentBSet)

		// Calculate the priority and add it to the total
		for r, exists := intersection.Pop(); exists; r, exists = intersection.Pop() {
			totalPriority += getPriority(r)
		}

		groupIndex += 1

		// If we've seen 3 elves, find the common item and add to badge priority
		if groupIndex%3 == 0 {
			intersection = groupSet[0].Intersect(groupSet[1].Intersect(groupSet[2]))

			r, _ := intersection.Pop()
			badgePriority += getPriority(r)

			// Clear the sets for the next group
			groupSet[0].Clear()
			groupSet[1].Clear()
			groupSet[2].Clear()
		}

		groupIndex %= 3
	}

	fmt.Printf("totalPriority: %v\n", totalPriority)
	fmt.Printf("badgePriority: %v\n", badgePriority)
}
