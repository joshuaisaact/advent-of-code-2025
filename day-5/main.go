package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	lower, upper int
}

func solve1(lowers []int, uppers []int, ids []int) int {
	freshIngredients := 0
	// fmt.Println("ids", ids)
	for _, id := range ids {
		for i, lower := range lowers {
			// fmt.Println("checking id:", id, lower, uppers[i])
			if id >= lower {
				if uppers[i] >= id {
					// fmt.Println("Checking index...", i)
					// fmt.Println("This ingredient is fresh!", lower, uppers[i], id)
					freshIngredients++
					break
				}
			}
		}
	}
	return freshIngredients
}

// func solve2(lowers []int, uppers []int) int {
// 	// Bloody hell Go doesn't have Sets does it...
// 	// Would be so much easier in Javascript with a set

// 	// freshIds := make(map[int]bool)

// 	// for i, lower := range lowers {
// 	// 	for n := lower; n <= uppers[i]; n++ {
// 	// 		freshIds[n] = true
// 	// 	}
// 	// }

// 	// This took far too long, it's TOO naive a solve.

// 	fmt.Println(freshIds)
// 	return len(freshIds)
// }

func solve2(ranges []Range) int {
	// Range ranges huh? Naming things is hard

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].lower < ranges[j].lower
	})

	merged := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		last := len(merged) - 1

		if ranges[i].lower <= merged[last].upper {
			if ranges[i].upper > merged[last].upper {
				merged[last].upper = ranges[i].upper
			}
		} else {
			merged = append(merged, ranges[i])
		}
	}

	total := 0

	for _, r := range merged {
		total += r.upper - r.lower + 1
	}

	return total
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

	// fmt.Println(ids)    // Ingredient ids to check
	// fmt.Println(lowers) // Lower bounds
	// fmt.Println(uppers) // Upper bounds

	var ranges []Range
	for _, line := range strings.Split(parts[0], "\n") {
		bounds := strings.Split(line, "-")
		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, Range{lower: lower, upper: upper})
	}
	result1 := solve1(lowers, uppers, ids)
	result2 := solve2(ranges)

	fmt.Println("Total fresh ingredients:", result1)
	fmt.Println("Total ids:", result2)
}
