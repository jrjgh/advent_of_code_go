package day6

import "testing"

var testFilename = "day6_test.txt"
var filename = "day6.txt"

func TestPartOne(t *testing.T) {

	expected := 288
	actual := PartOne(testFilename)

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}

	expected = 505494
	actual = PartOne(filename)

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {

	expected := 71503
	actual := PartTwo(testFilename)

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}

	expected = 23632299
	actual = PartTwo(filename)

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}
