package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve1(operations []string, inputs [][]int) int {
	total := 0
	for i, operation := range operations {
		if operation == "+" {
			for _, input := range inputs {
				total += input[i]
			}
		} else if operation == "*" {
			result := 1
			for _, input := range inputs {
				result *= input[i]
				fmt.Println("result:", result)
			}
			total += result
		}
	}
	return total
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))
	parts := strings.Split(input, "\n")
	fmt.Println(parts)
	operations := strings.Fields(parts[len(parts)-1])
	var inputs [][]int
	for _, line := range parts[:len(parts)-1] {
		fields := strings.Fields(line)
		var row []int
		for _, f := range fields {
			n, _ := strconv.Atoi(f)
			row = append(row, n)
		}
		inputs = append(inputs, row)
	}
	fmt.Println(operations)
	fmt.Println(inputs)
	result1 := solve1(operations, inputs)
	// result2 := solve2(ranges)

	fmt.Println("Total fresh ingredients:", result1)
	// fmt.Println("Total ids:", result2)
}
