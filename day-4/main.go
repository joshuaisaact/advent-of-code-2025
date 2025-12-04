package main

import (
	"fmt"
	"os"
	"strings"
)

func checkSurrounds(grid [][]byte, y int, x int) bool {
	if grid[y][x] != '@' {
		return false
	}

	numRolls := 0
	height := len(grid)
	width := len(grid[0])

	fmt.Println("Check surrounds running with:", x, y)

	// Top
	if y > 0 {
		if x > 0 && grid[y-1][x-1] == '@' {
			numRolls++
		}
		if grid[y-1][x] == '@' {
			numRolls++
		}
		if x < width-1 && grid[y-1][x+1] == '@' {
			numRolls++
		}
	}
	// Middle
	if x > 0 && grid[y][x-1] == '@' {
		numRolls++
		if numRolls >= 4 {
			return false
		}
	}
	if x < width-1 && grid[y][x+1] == '@' {
		numRolls++
		if numRolls >= 4 {
			return false
		}
	}
	// Bottom
	if y < height-1 {
		if x > 0 && grid[y+1][x-1] == '@' {
			numRolls++
			if numRolls >= 4 {
				return false
			}
		}
		if grid[y+1][x] == '@' {
			numRolls++
			if numRolls >= 4 {
				return false
			}
		}
		if x < width-1 && grid[y+1][x+1] == '@' {
			numRolls++
			if numRolls >= 4 {
				return false
			}
		}
	}

	return true
}

func solve1(grid [][]byte) int {
	accessibleRolls := 0
	for y, line := range grid {
		for x, _ := range line {
			if checkSurrounds(grid, x, y) {
				fmt.Println("Roll accessible at:", x, y)
				accessibleRolls += 1
			}
		}
	}
	return accessibleRolls
}

func solve2(grid [][]byte) int {
	accessibleRolls := 0
	for y, line := range grid {
		for x, _ := range line {
			if checkSurrounds(grid, x, y) == true {
				fmt.Println("Roll accessible at:", x, y)
				accessibleRolls += 1
			}
		}
	}
	return accessibleRolls
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))

	inputArr := strings.Split(input, "\n")

	var grid [][]byte
	for _, line := range inputArr {
		grid = append(grid, []byte(line))
	}

	result1 := solve1(grid)
	result2 := solve2(grid)

	fmt.Println("Total rolls solve1:", result1)
	fmt.Println("Total rolls solve 2:", result2)
}
