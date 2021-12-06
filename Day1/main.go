package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// log.Print(string(input))

	countIncreases(string(input))
}

func countIncreases(input string) int {
	var inputAsInt []int
	for _, line := range strings.Split(input, "\n") {
		number, err := strconv.Atoi(line)			
		if err != nil {
			log.Fatalf("Failed to parse string to integer")
		}

		inputAsInt = append(inputAsInt, number)
	}

	var increaseCount int
	for i, j := 0, 1; j < len(inputAsInt); i, j = i + 1, j + 1 {
		if inputAsInt[i] < inputAsInt[j] {
			increaseCount++
		}
	}
	// var last int
	// for i, next := range inputAsInt {
	// 	last = next
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	if last < next {
	// 		increaseCount++
	// 	}
	// }

	log.Printf("Count: %d", increaseCount)
	return increaseCount
}