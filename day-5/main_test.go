package main

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSolve1(t *testing.T) {
	data, err := os.ReadFile("testinput.txt")
	if err != nil {
		t.Fatal(err)
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

	got := solve1(lowers, uppers, ids)
	want := 3

	if got != want {
		t.Errorf("solve() = %d, want %d", got, want)
	}
}

func TestSolve2(t *testing.T) {
	data, err := os.ReadFile("testinput.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := strings.TrimSpace(string(data))
	parts := strings.Split(input, "\n\n")

	var ranges []Range
	for _, line := range strings.Split(parts[0], "\n") {
		bounds := strings.Split(line, "-")
		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, Range{lower: lower, upper: upper})
	}

	got := solve2(ranges)

	want := 14

	if got != want {
		t.Errorf("solve() = %d, want %d", got, want)
	}
}
