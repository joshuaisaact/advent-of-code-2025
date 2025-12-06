package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve1(lowers []int, uppers []int, ids []int) int {
	freshIngredients := 0
	fmt.Println("ids", ids)
	for _, id := range ids {
		for i, lower := range lowers {
			fmt.Println("checking id:", id, lower, uppers[i])
			if id > lower {
				if uppers[i] >= id {
					fmt.Println("Checking index...", i)
					fmt.Println("This ingredient is fresh!", lower, uppers[i], id)
					freshIngredients++
					break
				}
			}
		}
	}
	return freshIngredients
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))
	parts := strings.Split(input, "\n\n")

	var lowers []int
	var uppers []int
	for _, line := range strings.Split(parts[0], "\n") {
		bounds := strings.Split(line, "-")
		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])
		lowers = append(lowers, lower)
		uppers = append(uppers, upper)
	}

	var ids []int
	for _, line := range strings.Split(parts[1], "\n") {
		n, _ := strconv.Atoi(line)
		ids = append(ids, n)
	}

	fmt.Println(ids)    // Ingredient ids to check
	fmt.Println(lowers) // Lower bounds
	fmt.Println(uppers) // Upper bounds

	result1 := solve1(lowers, uppers, ids)
	// result2 := solve2(grid)

	fmt.Println("Total rolls solve1:", result1)
	// fmt.Println("Total rolls solve 2:", result2)
}
