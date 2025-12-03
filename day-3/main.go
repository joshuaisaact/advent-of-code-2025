package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve1(inputArr []string) int {
	voltageSum := 0

	for _, lines := range inputArr {

		fmt.Println("lines:", string(lines))
		fmt.Println("Lines[0]", string(lines[0]))

		max := lines[0]
		second := byte(0)

		for i := 1; i < len(lines); i++ {

			if lines[i] > second { // Edge case where lines[1] is greater than lines[0] - setting second to 0 as an initial value
				second = lines[i]
			}

			if lines[i] > max && i < len(lines)-1 { // If find a new max, set it
				max = lines[i]
				second = lines[i+1]
			} else if lines[i] > second {
				second = lines[i]
			}
		}
		fmt.Println("Current max and second:", string(max), string(second))

		result, _ := strconv.Atoi(string(max) + string(second))
		voltageSum += result

	}

	return voltageSum
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))

	inputArr := strings.Split(input, "\n")

	fmt.Println("inputarr:", inputArr)

	result1 := solve1(inputArr)
	// result2 := solve2(inputArr)

	fmt.Println("Invalid IDs solve 1:", result1)
	// fmt.Println("Invalid IDs solve 2:", result2)
}
