package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCrates(str string, indices []int) []string {
	crates := []string{}

	for _, i := range indices {
		crates = append(crates, string(str[i]))
	}

	return crates
}

func main() {
	// Open input file
	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Read from file and count calories
	scanner := bufio.NewScanner(file)

	// Set variables
	stacks := []*list.List{}

	// Read first line from file and determine number of stacks
	scanner.Scan()
	str := scanner.Text()
	stackNum := (len(str) + 1) / 4
	indexList := []int{}

	for i := 0; i < stackNum; i++ {
		l := list.New()
		stacks = append(stacks, l)
		indexList = append(indexList, 1+i*4)
	}

	for ; str != ""; str = scanner.Text() {
		// When we reach the final row, break
		if !strings.ContainsAny(str, "[]") {
			break
		}

		// Get the crate labels
		crates := getCrates(str, indexList)

		for i, c := range crates {
			if string(c) != " " {
				stacks[i].PushBack(string(c))
			}
		}

		// Seek to next line
		scanner.Scan()
	}

	// Jump after new line
	scanner.Scan()

	// Begin parsing instructions
	for scanner.Scan() {
		str := scanner.Text()

		v := strings.Split(str, " ")
		count, _ := strconv.Atoi(v[1])
		origin, _ := strconv.Atoi(v[3])
		dest, _ := strconv.Atoi(v[5])

		movingStack := list.New()

		for i := 0; i < count; i++ {
			crate := stacks[origin-1].Front()
			stacks[origin-1].Remove(crate)
			movingStack.PushBack(crate.Value)
		}

		stacks[dest-1].PushFrontList(movingStack)
	}

	str = ""

	for _, s := range stacks {
		if s.Front() != nil {
			str += s.Front().Value.(string)
		}
	}

	fmt.Println(str)
}
