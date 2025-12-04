package main

import (
	"os"
	"strings"
	"testing"
)

func TestSolve1(t *testing.T) {
	data, err := os.ReadFile("testinput.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := strings.TrimSpace(string(data))
	inputArr := strings.Split(input, "\n")
	var grid [][]byte
	for _, line := range inputArr {
		grid = append(grid, []byte(line))
	}
	got := solve1(grid)
	want := 13

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
	inputArr := strings.Split(input, "\n")
	var grid [][]byte
	for _, line := range inputArr {
		grid = append(grid, []byte(line))
	}
	got := solve2(grid)
	want := 43

	if got != want {
		t.Errorf("solve() = %d, want %d", got, want)
	}
}
