package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isRepeatedSequence(input string) bool {
	if len(input)%2 != 0 {
		return false
	}

	mid := len(input) / 2
	return input[:mid] == input[mid:]
}

func isRepeatedSequenceSolve2(input string) bool {

	mid := len(input) / 2
	return input[:mid] == input[mid:]
}

func solve1(inputArr []string) int {
	invalidSum := 0

	for _, ids := range inputArr {
		parts := strings.Split(ids, "-")

		id1, _ := strconv.Atoi(parts[0])
		id2, _ := strconv.Atoi(parts[1])

		for i := id1; i <= id2; i++ {
			stringNum := strconv.Itoa(i)

			if isRepeatedSequence(stringNum) == true {
				invalidSum += i
			}

		}
	}

	return invalidSum
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))

	inputArr := strings.Split(input, ",")

	result1 := solve1(inputArr)
	result2 := solve2(inputArr)

	fmt.Println("Invalid IDs solve 1:", result1)
	fmt.Println("Invalid IDs solve 2:", result2)
}
