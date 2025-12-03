package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve1(inputArr []string) int {
	voltageSum := 0

	for _, line := range inputArr {

		max := line[0]
		second := byte(0)

		for i := 1; i < len(line); i++ {

			if line[i] > second { // Edge case where line[1] is greater than line[0] - setting second to 0 as an initial value
				second = line[i]
			}

			if line[i] > max && i < len(line)-1 { // If find a new max, set it
				max = line[i]
				second = line[i+1]
			} else if line[i] > second {
				second = line[i]
			}
		}
		// fmt.Println("Current max and second:", string(max), string(second))

		result, _ := strconv.Atoi(string(max) + string(second))
		voltageSum += result

	}

	return voltageSum
}

func solve2(inputArr []string) int {
	voltageSum := 0

	for _, line := range inputArr {

		res := make([]byte, 12)
		max := byte(0)
		foundIndex := 0

		for i := 0; i < 12; i++ { // loop through the array
			for j := foundIndex; j <= len(line)-12+i; j++ { // Be cautious about going OOB

				if line[j] > max { // If find a new max, set it
					max = line[j]      // this is confirmed highest possible max
					foundIndex = j + 1 // Found the max here, let's start the next loop from next
				}
			}
			res[i] = max  // Put the max in the relevant position of the array
			max = byte(0) // reset the max

		}
		// fmt.Println("res:", string(res))
		result, _ := strconv.Atoi(string(res))
		// fmt.Println("error:", err) // Had to print error to spot null issues with earlier iteration
		// fmt.Println("result:", result)
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

	result1 := solve1(inputArr)
	result2 := solve2(inputArr)

	fmt.Println("Total Joltage solve 1:", result1)
	fmt.Println("Total Joltage solve 2:", result2)
}
