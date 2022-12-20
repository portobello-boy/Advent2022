package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// type Operation func(old *int, modifier int, multiplication bool)
type Operation func()

type Monkey struct {
	index            int
	items            *[]*uint64
	totalInspections *int
	operation        Operation
	testVal          int
	targetTrue       int
	targetFalse      int
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	// Open input file
	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	monkeys := make([]*Monkey, 0)
	monkeyDivisors := []int{}

	// Read from file and process monkeys
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := strings.Split(strings.Trim(scanner.Text(), ":"), " ")
		index, _ := strconv.Atoi(m[1])

		numFinder := regexp.MustCompile("[0-9]+")

		scanner.Scan()
		itemsStr := numFinder.FindAllString(scanner.Text(), -1)
		items := []*uint64{}
		for _, v := range itemsStr {
			n, _ := strconv.Atoi(v)
			u64n := uint64(n)
			items = append(items, &u64n)
		}

		initialTotal := 0

		scanner.Scan()
		op := strings.Split(scanner.Text(), " ")[5:8]
		isMult := false
		if op[1] == "*" {
			isMult = true
		}
		modifier := 0
		isSquare := false
		if op[2] == "old" {
			isSquare = true
		} else {
			m, _ := strconv.Atoi(op[2])
			modifier = m
		}

		scanner.Scan()
		test := numFinder.FindAllString(scanner.Text(), -1)
		tv, _ := strconv.Atoi(test[0])
		monkeyDivisors = append(monkeyDivisors, tv)

		scanner.Scan()
		resTrue := numFinder.FindAllString(scanner.Text(), -1)
		rt, _ := strconv.Atoi(resTrue[0])

		scanner.Scan()
		resFalse := numFinder.FindAllString(scanner.Text(), -1)
		rf, _ := strconv.Atoi(resFalse[0])

		monkeys = append(monkeys, &Monkey{
			index:            index,
			items:            &items,
			totalInspections: &initialTotal,
			testVal:          tv,
			targetTrue:       rt,
			targetFalse:      rf,
			operation: func() {
				indexToRemove := []int{}

				// Iterate over all monkeys
				for i, n := range *(monkeys[index]).items {
					*(monkeys[index].totalInspections) += 1
					// Change the number score
					if isMult {
						if isSquare {
							*n *= *n
						} else {
							*n *= uint64(modifier)
						}
					} else {
						if isSquare {
							*n += *n
						} else {
							*n += uint64(modifier)
						}
					}

					// Take module of worry
					*n %= uint64(LCM(monkeyDivisors[0], monkeyDivisors[0], monkeyDivisors...))

					// Lessen the worry
					// *n /= 3

					if (*n)%uint64(tv) == 0 {
						// Send to targetTrue
						*(*monkeys[rt]).items = append(*(*monkeys[rt]).items, n)

					} else {
						// Send to targetFalse
						*(*monkeys[rf]).items = append(*(*monkeys[rf]).items, n)
					}

					indexToRemove = append(indexToRemove, i)
				}

				for i := len(indexToRemove) - 1; i >= 0; i-- {
					// Remove from list
					(*(monkeys[index]).items)[i] = (*(monkeys[index]).items)[len(items)-1]
					*(monkeys[index]).items = (*(monkeys[index]).items)[:len(items)-1]
				}
			},
		})

		scanner.Scan()
	}

	// Perform Rounds
	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.operation()
		}
	}

	// Calculate Monkey Business
	inspections := []int{}
	monkeyBusiness := 1
	for _, m := range monkeys {
		inspections = append(inspections, *m.totalInspections)
	}

	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})

	monkeyBusiness *= inspections[0]
	monkeyBusiness *= inspections[1]

	fmt.Printf("inspections: %v\n", inspections)
	fmt.Printf("monkeyBusiness: %v\n", monkeyBusiness)
}
