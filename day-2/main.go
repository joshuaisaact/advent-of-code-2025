package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isRepeatedSequenceSolve1(input string) bool {
	if len(input)%2 != 0 || len(input) < 3 {
		return false
	}

	mid := len(input) / 2

	if input[:mid] == input[mid:] == true {
		fmt.Println("Repeated middle split!", input)
		return true
	} else {
		return false
	}
}

func solve1(inputArr []string) int {
	invalidSum := 0

	for _, ids := range inputArr {
		parts := strings.Split(ids, "-")

		id1, _ := strconv.Atoi(parts[0])
		id2, _ := strconv.Atoi(parts[1])

		for i := id1; i <= id2; i++ {
			stringNum := strconv.Itoa(i)

			if isRepeatedSequenceSolve1(stringNum) == true {
				invalidSum += i
			}

		}
	}

	return invalidSum
}

func isRepeatedSequenceSingleNumbers(input string) bool {
	if len(input) < 2 {
		return false
	}

	comparator := input[0]
	for i := 1; i < len(input); i++ {
		if comparator == input[i] {
			continue
		} else {
			return false
		}
	}
	fmt.Println("Repeated single numbers!", input)
	return true
}

func isRepeatedSequenceDoubleNumbers(input string) bool {
	if len(input) <= 3 || len(input)%2 != 0 {
		return false
	}

	comparator := input[:2]

	for i := 2; i < len(input); i += 2 {
		if comparator == input[i:i+2] {
			continue
		} else {
			return false
		}
	}
	fmt.Println("Repeated double numbers!", input)
	return true
}

func isRepeatedSequenceTripleNumbers(input string) bool {
	if len(input) <= 7 || len(input)%3 != 0 {
		return false
	}

	comparator := input[:3]

	for i := 3; i < len(input); i += 3 {
		if comparator == input[i:i+3] {
			continue
		} else {
			return false
		}
	}
	fmt.Println("Repeated triple numbers!", input)
	return true
}

func solve2(inputArr []string) int {
	invalidIDSum := 0

	for _, ids := range inputArr {
		parts := strings.Split(ids, "-")

		id1, _ := strconv.Atoi(parts[0])
		id2, _ := strconv.Atoi(parts[1])

		for i := id1; i <= id2; i++ {
			stringNum := strconv.Itoa(i)

			if isRepeatedSequenceSolve1(stringNum) ||
				isRepeatedSequenceSingleNumbers(stringNum) ||
				isRepeatedSequenceDoubleNumbers(stringNum) ||
				isRepeatedSequenceTripleNumbers(stringNum) {
				invalidIDSum += i
				continue
			}

		}
	}

	return invalidIDSum
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
