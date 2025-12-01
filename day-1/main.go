package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve1(inputArr []string) int {
	zeros := 0
	sum := 50

	for _, line := range inputArr {
		letter := line[0]
		numPart := line[1:]
		num, _ := strconv.Atoi(numPart)

		if letter == 'L' {
			sum = sum - num
		} else {
			sum = sum + num
		}

		if sum == 0 || sum%100 == 0 {
			zeros += 1
			sum = 0
		}

		if sum > 100 {
			sum = sum % 100
		}

		if sum < 0 {
			res := sum % 100
			sum = res + 100
		}

	}
	return zeros
}

func solve2(inputArr []string) int {
	zeros := 0
	sum := 50
	crossed := true

	for _, line := range inputArr {
		letter := line[0]
		numPart := line[1:]
		num, _ := strconv.Atoi(numPart)

		if letter == 'L' {
			sum = sum - num
		} else {
			sum = sum + num
		}

		if sum == 0 || sum == 100 {
			zeros += 1
			sum = 0
			crossed = false
		} else if sum > 100 {

			crossed = (sum%100 != 0)

			zeros += sum / 100
			sum = sum % 100

		} else if sum < 0 {

			if crossed == true {
				zeros += 1
			}
			crossed = (sum%100 != 0)
			zeros -= (sum / 100)
			rem := sum % 100
			sum = (rem + 100) % 100

		} else {
			crossed = true
		}

	}
	return zeros
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))

	inputArr := strings.Split(input, "\n")

	result := solve1(inputArr)

	result2 := solve2(inputArr)

	fmt.Println("Number of zeros solve 1:", result)
	fmt.Println("Number of zeros solve 2:", result2)
}
