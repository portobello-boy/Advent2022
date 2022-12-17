/*
--- Day 1: Calorie Counting ---

Santa's reindeer typically eat regular reindeer food, but they need a lot of magical energy to deliver presents on Christmas. For that, their favorite snack is a special type of star fruit that only grows deep in the jungle. The Elves have brought you on their annual expedition to the grove where the fruit grows.

To supply enough magical energy, the expedition needs to retrieve a minimum of fifty stars by December 25th. Although the Elves assure you that the grove has plenty of fruit, you decide to grab any fruit you see along the way, just in case.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies. One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).

The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items' Calories and end up with the following list:

1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

This list represents the Calories of the food carried by five Elves:

    The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
    The second Elf is carrying one food item with 4000 Calories.
    The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
    The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
    The fifth Elf is carrying one food item with 10000 Calories.

In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the most Calories. In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

*/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Elf struct {
	index    int
	calories int
}

// Priority Queue code pulled from https://golang.google.cn/pkg/container/heap/
type PriorityQueueItem struct {
	elf   Elf
	index int
}

type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].elf.calories > pq[j].elf.calories

}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j

}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item

}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *PriorityQueueItem, calories int) {
	item.elf.calories = calories
	heap.Fix(pq, item.index)
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
	elfIndex := 1
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Create elf container
	elf := Elf{
		index:    elfIndex,
		calories: 0,
	}

	// Read from file and count calories
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		if str != "" {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err)
			}

			elf.calories += num
		} else {
			elfIndex += 1

			pqItem := &PriorityQueueItem{
				elf:   elf,
				index: elfIndex,
			}

			heap.Push(&pq, pqItem)

			elf = Elf{
				index:    elfIndex,
				calories: 0,
			}
		}
	}

	pqItem := &PriorityQueueItem{
		elf:   elf,
		index: elfIndex + 1,
	}

	heap.Push(&pq, pqItem)

	topThreeSum := 0

	for i := 0; i < 3; i++ {
		item := heap.Pop(&pq).(*PriorityQueueItem)
		elf := item.elf

		fmt.Println(elf)
		topThreeSum += elf.calories
	}

	fmt.Println(topThreeSum)

	// for pq.Len() > 0 {
	// 	item := heap.Pop(&pq).(*PriorityQueueItem)
	// 	fmt.Println(item)
	// }
}
