package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func checkSurrounds(grid [][]byte, p Point) bool {
	if grid[p.Y][p.X] != '@' {
		return false
	}

	numRolls := 0
	height := len(grid)
	width := len(grid[0])

	// fmt.Println("Check surrounds running with:", x, y)

	// Top
	if p.Y > 0 {
		if p.X > 0 && grid[p.Y-1][p.X-1] == '@' {
			numRolls++
		}
		if grid[p.Y-1][p.X] == '@' {
			numRolls++
		}
		if p.X < width-1 && grid[p.Y-1][p.X+1] == '@' {
			numRolls++
		}
	}
	// Middle
	if p.X > 0 && grid[p.Y][p.X-1] == '@' {
		numRolls++
		if numRolls >= 4 {
			return false
		}
	}
	if p.X < width-1 && grid[p.Y][p.X+1] == '@' {
		numRolls++
		if numRolls >= 4 {
			return false
		}
	}
	// Bottom
	if p.Y < height-1 {
		if p.X > 0 && grid[p.Y+1][p.X-1] == '@' {
			numRolls++
			if numRolls >= 4 {
				return false
			}
		}
		if grid[p.Y+1][p.X] == '@' {
			numRolls++
			if numRolls >= 4 {
				return false
			}
		}
		if p.X < width-1 && grid[p.Y+1][p.X+1] == '@' {
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
		for x := range line {
			if checkSurrounds(grid, Point{x, y}) {
				// fmt.Println("Roll accessible at:", x, y)

				accessibleRolls++
			}
		}
	}
	return accessibleRolls
}

func solve2(grid [][]byte) int {
	totalRemoved := 0
	pass := 0
	for {
		removedRolls := []Point{}
		for y, line := range grid {
			for x := range line {
				if checkSurrounds(grid, Point{x, y}) {
					removedRolls = append(removedRolls, Point{x, y})
				}
			}
		}

		fmt.Printf("Pass %d: found %d removable \n", pass, len(removedRolls))

		if len(removedRolls) == 0 {
			break
		}

		for _, p := range removedRolls {
			grid[p.Y][p.X] = '.'
		}

		totalRemoved += len(removedRolls)
		fmt.Printf("Pass %d: removed %d\n", pass, totalRemoved)
		pass++
	}

	return totalRemoved
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
