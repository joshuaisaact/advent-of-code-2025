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
	inputArr := strings.Split(input, ",")
	got := solve1(inputArr)
	want := 1227775554

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
	inputArr := strings.Split(input, ",")
	got := solve2(inputArr)
	want := 4174379265

	if got != want {
		t.Errorf("solve() = %d, want %d", got, want)
	}
}
