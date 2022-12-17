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

func handleMove(v []string, head, tail *Coord, set *mapset.Set[Coord]) {
	moveCount, _ := strconv.Atoi(v[1])

	for i := 0; i < moveCount; i++ {
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
		// Adjust Tail
		if dist(head, tail) > 1 {
			if absDiff(head.x, tail.x) > 1 {
				if head.y != tail.y {
					(*tail).y = head.y
				}
				(*tail).x += ((head.x - tail.x) / absDiff(head.x, tail.x))
			} else if absDiff(head.y, tail.y) > 1 {
				if head.x != tail.x {
					(*tail).x = head.x
				}
				(*tail).y += ((head.y - tail.y) / absDiff(head.y, tail.y))
			}
		}

		// fmt.Printf("head: %+v\n", head)
		// fmt.Printf("tail: %+v\n", tail)

		// Add to set
		(*set).Add(*tail)
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
	headCoord := Coord{0, 0}
	tailCoord := Coord{0, 0}

	visitedCoords.Add(tailCoord)

	// Read from file and process moves
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		v := strings.Split(str, " ")
		// fmt.Printf("v: %v\n", v)

		handleMove(v, &headCoord, &tailCoord, &visitedCoords)
	}

	fmt.Printf("visitedCoords.Cardinality(): %v\n", visitedCoords.Cardinality())
}
