package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func incrementCounter(counter, register, signal *int, signalList []int, crtGrid *[]string) {
	symbol := ""
	if math.Abs(float64((*counter)%40)-float64(*register)) <= 1 {
		symbol = "#"
	} else {
		symbol = "."
	}

	(*crtGrid) = append((*crtGrid), symbol)

	fmt.Printf("symbol: %v\n", symbol)
	fmt.Printf("(*counter): %v\n", (*counter)+1)
	fmt.Printf("(*register): %v\n", (*register))

	*counter += 1
	for _, v := range signalList {
		if *counter == v {
			*signal += v * *register
		}
	}

	// *counter -= ((*counter % 40) * 40)
}

func main() {
	// Open input file
	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	counter, register, signal := 0, 1, 0
	signalList := [6]int{20, 60, 100, 140, 180, 220}
	crtGrid := []string{}

	// Read from file and process operations
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		v := strings.Split(str, " ")

		switch v[0] {
		case "noop":
			incrementCounter(&counter, &register, &signal, signalList[:], &crtGrid)
			break
		case "addx":
			incrementCounter(&counter, &register, &signal, signalList[:], &crtGrid)
			incrementCounter(&counter, &register, &signal, signalList[:], &crtGrid)
			val, _ := strconv.Atoi(v[1])
			register += val
		}
	}

	fmt.Printf("signal: %v\n", signal)

	for i := 0; i < len(crtGrid)/40; i++ {
		for j := 0; j < 40; j++ {
			fmt.Print(crtGrid[(i*40)+j])
		}
		fmt.Println("")
	}
}
