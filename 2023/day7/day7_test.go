package day7

import (
	"testing"
)

var testFilename = "day7_test.txt"
var filename = "day7.txt"

func TestPartOne(t *testing.T) {

	expected := 6440
	actual := PartOne(testFilename)

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}

	expected = 253638586
	actual = PartOne(filename)

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {

	expected := 5905
	actual := PartTwo(testFilename)

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}

	expected = 253253225
	actual = PartTwo(filename)

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}
