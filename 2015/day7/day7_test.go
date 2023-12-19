package day7

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	out := partOne()
	expected := 956
	if out != expected {
		t.Errorf("expected %d but got %d", expected, out)
	}
}

func TestPartTwo(t *testing.T) {
	out := partTwo()
	expected := 40149
	if out != expected {
		t.Errorf("expected %d but got %d", expected, out)
	}
}
