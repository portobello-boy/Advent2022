package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dariubs/uniq"
)

const length int = 14

func main() {
	// Open input file
	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	code := scanner.Text()

	for i := length; i < len(code); i++ {
		un := uniq.UniqString(strings.Split(code[i-length:i], ""))
		if len(un) == length {
			fmt.Println(un, i)
			break
		}
	}
}
