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
	got := solve1(inputArr)
	want := 357

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
	got := solve2(inputArr)
	want := 3121910778619

	if got != want {
		t.Errorf("solve() = %d, want %d", got, want)
	}
}
