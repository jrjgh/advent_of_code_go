package day3

import (
	"github.com/jrjgh/advent_of_code/util"
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 498559
	actual := PartOne()
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 72246648
	actual := PartTwo()
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestGearRatios(t *testing.T) {
	lines := []string{
		"5..",
		".*.",
		"..3",
	}
	expected := 15
	actual := util.SumInt(gearRatios(lines))
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}
