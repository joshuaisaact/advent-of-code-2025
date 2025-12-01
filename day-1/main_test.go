package main

import (
	"os"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	data, err := os.ReadFile("testinput.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := strings.TrimSpace(string(data))
	inputArr := strings.Split(input, "\n")
	got := solve2(inputArr)
	want := 6

	if got != want {
		t.Errorf("solve() = %d, want %d", got, want)
	}
}
