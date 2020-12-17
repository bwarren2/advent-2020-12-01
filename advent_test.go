package advent_test

import (
	"advent"
	"testing"
)

func TestSumTo(t *testing.T) {
	result := advent.SumTo("input.txt", 2020)
	if result != 964875 {
		t.Errorf("Got the wrong answer! %d", result)
	}
}

func TestTriSumTo(t *testing.T) {
	result := advent.TriSumTo("input.txt", 2020)

	if result != 158661360 {
		t.Errorf("Got the wrong answer! %d", result)
	}
}
