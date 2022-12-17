package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type Coord struct {
	x int
	y int
}

func dist(a, b *Coord) int {
	return int(math.Max(math.Abs(float64(a.x-b.x)), math.Abs(float64(a.y-b.y))))
}

func absDiff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func handleMove(v []string, rope []*Coord, set *mapset.Set[Coord]) {
	moveCount, _ := strconv.Atoi(v[1])

	for i := 0; i < moveCount; i++ {
		head := rope[0]

		// Move Head
		switch v[0] {
		case "U":
			(*head).y += 1
		case "R":
			(*head).x += 1
		case "D":
			(*head).y -= 1
		case "L":
			(*head).x -= 1
		}

		for knotIndex := 0; knotIndex < len(rope)-1; knotIndex++ {
			current := rope[knotIndex]
			next := rope[knotIndex+1]

			// Adjust Next
			if dist(current, next) > 1 {
				if absDiff(current.x, next.x) > 1 {
					if current.y != next.y {
						(*next).y += ((current.y - next.y) / absDiff(current.y, next.y))
					}
					(*next).x += ((current.x - next.x) / absDiff(current.x, next.x))
				} else if absDiff(current.y, next.y) > 1 {
					if current.x != next.x {
						(*next).x += ((current.x - next.x) / absDiff(current.x, next.x))
					}
					(*next).y += ((current.y - next.y) / absDiff(current.y, next.y))
				}
			}
		}

		// Print rope after move sequence
		grid := make([][]string, 10) // 60
		for i := 0; i < len(grid); i++ {
			row := []string{}
			for j := 0; j < 10; j++ {
				row = append(row, ".")
			}
			grid[i] = row
		}

		// Add to set
		(*set).Add(*rope[len(rope)-1])
	}
}

func main() {
	// Open input file
	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	visitedCoords := mapset.NewSet[Coord]()
	rope := make([]*Coord, 10)

	for i := range rope {
		rope[i] = &Coord{0, 0}
	}

	visitedCoords.Add(*rope[len(rope)-1])

	// Read from file and process moves
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		v := strings.Split(str, " ")

		handleMove(v, rope, &visitedCoords)
	}

	fmt.Printf("visitedCoords.Cardinality(): %v\n", visitedCoords.Cardinality())
}
