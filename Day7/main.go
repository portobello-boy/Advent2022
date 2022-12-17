package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Name     string
	Size     int
	Parent   *Node
	Children []*Node
}

type DirInfo struct {
	Name string
	Size int
}

// Priority Queue code pulled from https://golang.google.cn/pkg/container/heap/
type PriorityQueueItem struct {
	directory DirInfo
	index     int
}

type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].directory.Size < pq[j].directory.Size

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
func (pq *PriorityQueue) update(item *PriorityQueueItem, size int) {
	item.directory.Size = size
	heap.Fix(pq, item.index)
}

const filesystemSize = 70000000
const sizeGoal = 30000000
const sizeLimit = 100000

var dirIndex int = 0

func calculateTreeSize(root Node, sizeUnderLimit *int, pq *PriorityQueue) int {
	directorySize := 0

	if len(root.Children) == 0 {
		return root.Size
	}

	for _, child := range root.Children {
		directorySize += calculateTreeSize(*child, sizeUnderLimit, pq)
	}

	heap.Push(pq, &PriorityQueueItem{
		directory: DirInfo{
			Name: root.Name,
			Size: directorySize,
		},
		index: dirIndex,
	})

	dirIndex += 1

	if directorySize < sizeLimit {
		*sizeUnderLimit += directorySize
	}
	return directorySize
}

func seekPath(current **Node, root *Node, path string) {
	// If there is nowhere left to seek, return
	if path == "" {
		return
	}

	v := strings.Split(path, "/")

	if v[0] == ".." {
		if (*current).Parent == nil {
			fmt.Println("Cannot go up from root!")
			os.Exit(-1)
		}
		*current = (*current).Parent
		seekPath(current, root, strings.Join(v[1:], "/"))
	}

	// If going to child directory, change current to there and finish
	if path[0] != '/' {
		for _, child := range (*current).Children {
			if child.Name == v[0] {
				*current = child
				seekPath(current, root, strings.Join(v[1:], "/"))
			}
		}
	} else {
		*current = root
		seekPath(current, root, strings.Join(v, "/")[1:])
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

	// Create File Tree
	root := Node{
		Name:     "/",
		Size:     0,
		Parent:   nil,
		Children: make([]*Node, 0),
	}

	// Create a Pointer to Root Directory
	current := &root

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	str := scanner.Text()

	for {
		getNextLine := true

		v := strings.Split(str, " ")

		switch v[0] {
		case "$":
			switch v[1] {
			case "cd":
				// Change current position in file tree...
				seekPath(&current, &root, v[2])
				break
			case "ls":
				// Look ahead at lines and add to tree...
			listLoop:
				for {

					// Check if we're at the end of the document
					if !scanner.Scan() {
						break
					}

					// Get next line
					entry := scanner.Text()
					e := strings.Split(entry, " ")

					// If next line is a command, stop here and mark flag
					if e[0] == "$" {
						getNextLine = false
						break listLoop
					}

					// Check if item is directory
					if e[0] == "dir" {
						current.Children = append(current.Children, &Node{
							Name:     e[1],
							Size:     0,
							Parent:   current,
							Children: make([]*Node, 0),
						})
					} else {
						size, _ := strconv.Atoi(e[0])
						current.Children = append(current.Children, &Node{
							Name:     e[1],
							Size:     size,
							Parent:   current,
							Children: nil,
						})
					}
				}
				break
			}
			break
		}

		if getNextLine {
			// Check if we're at the end of the document
			if !scanner.Scan() {
				break
			}
		}
		str = scanner.Text()
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	sizeUnderLimit := 0
	filesystemRootSize := calculateTreeSize(root, &sizeUnderLimit, &pq)
	minimumDeleteSize := filesystemRootSize - (filesystemSize - sizeGoal)

	fmt.Printf("sizeUnderLimit: %v\n", sizeUnderLimit)

	smallestSuitableSize := 0

	for {
		pqItem := heap.Pop(&pq)
		if pqItem.(*PriorityQueueItem).directory.Size > minimumDeleteSize {
			smallestSuitableSize = pqItem.(*PriorityQueueItem).directory.Size
			break
		}
	}

	fmt.Printf("smallestSuitableSize: %v\n", smallestSuitableSize)
}
