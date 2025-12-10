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

func isAllSpaces(rows []string, col int) bool {
	for _, row := range rows {
		if col < len(row) && row[col] != ' ' {
			return false
		}
	}
	return true
}

type problem struct{ start, end int }

func findProblems(rows []string) []problem {
	width := len(rows[0])
	var problems []problem
	start := -1

	for col := 0; col <= width; col++ {
		isSep := col == width || isAllSpaces(rows, col)
		if !isSep && start == -1 {
			start = col
		} else if isSep && start != -1 {
			problems = append(problems, problem{start, col})
			start = -1
		}
	}
	return problems
}

func solve2(lines []string) int {
	numRows := lines[:len(lines)-1]
	opRow := lines[len(lines)-1]

	problems := findProblems(numRows)

	total := 0

	for i := len(problems) - 1; i >= 0; i-- {
		p := problems[i]

		var op byte = '+'
		for col := p.start; col < p.end; col++ {
			if col < len(opRow) && (opRow[col] == '+' || opRow[col] == '*') {
				op = opRow[col]
				break
			}
		}

		var numbers []int
		for col := p.end - 1; col >= p.start; col-- {
			numStr := ""
			for _, row := range numRows {
				if col < len(row) && row[col] >= '0' && row[col] <= '9' {
					numStr += string(row[col])
				}
			}
			if numStr != "" {
				n := 0
				for _, c := range numStr {
					n = n*10 + int(c-'0')
				}
				numbers = append(numbers, n)
			}
		}

		if op == '+' {
			for _, n := range numbers {
				total += n
			}
		} else {
			result := 1
			for _, n := range numbers {
				result *= n
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

	lines := strings.Split(input, "\n")
	fmt.Println(lines)
	operations := strings.Fields(lines[len(lines)-1])
	var inputs [][]int
	for _, line := range lines[:len(lines)-1] {
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

	result2 := solve2(lines)

	fmt.Println("Total number:", result1)
	fmt.Println("Total solve 2:", result2)
}
