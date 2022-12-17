package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func loadTrees(file *os.File) [][]int {
	trees := make([][]int, 0, 0)

	// Read from file and count calories
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strRow := strings.Split(scanner.Text(), "")
		row := []int{}

		for _, d := range strRow {
			num, _ := strconv.Atoi(d)
			row = append(row, num)
		}

		trees = append(trees, row)
	}

	return trees
}

func countTrees(trees [][]int) int {
	totalTrees := 0
	maxScenicScore := int(0)

	// Count trees on edges (all 4 sides, minus duplicate corners)
	totalTrees = 2*len(trees) + 2*len(trees[0]) - 4

	// Iterate over all trees
	for i := 1; i < len(trees)-1; i++ {
		row := trees[i]
		for j := 1; j < len(row)-1; j++ {
			tree := trees[i][j]

			visibleFromNorth, visibleFromSouth, visibleFromEast, visibleFromWest := true, true, true, true
			countNorth, countEast, countSouth, countWest := 0, 0, 0, 0

			// Iterate in different directions to see if tree is visible
			// North
			for y := i - 1; y >= 0; y-- {
				obstacle := trees[y][j]
				countNorth += 1
				// fmt.Printf("obstacle: %v\n", obstacle)
				if obstacle >= tree {
					visibleFromNorth = false
					break
				}
			}
			// South
			for y := i + 1; y < len(trees); y++ {
				obstacle := trees[y][j]
				countSouth += 1
				if obstacle >= tree {
					visibleFromSouth = false
					break
				}
			}
			// West
			for x := j - 1; x >= 0; x-- {
				obstacle := trees[i][x]
				countWest += 1
				if obstacle >= tree {
					visibleFromWest = false
					break
				}
			}
			// East
			for x := j + 1; x < len(row); x++ {
				obstacle := trees[i][x]
				countEast += 1
				if obstacle >= tree {
					visibleFromEast = false
					break
				}
			}

			scenicScore := countNorth * countEast * countSouth * countWest
			maxScenicScore = int(math.Max(float64(maxScenicScore), float64(scenicScore)))

			if visibleFromNorth || visibleFromEast || visibleFromSouth || visibleFromWest {
				totalTrees += 1
			}
		}
	}

	fmt.Printf("maxScenicScore: %v\n", maxScenicScore)
	return totalTrees
}

func main() {
	// Open input file
	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	trees := loadTrees(file)
	visibleTrees := countTrees(trees)

	fmt.Printf("visibleTrees: %v\n", visibleTrees)
}
