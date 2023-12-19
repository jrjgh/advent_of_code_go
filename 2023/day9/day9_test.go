package day9

import "testing"

var testFilename = "day9_test.txt"
var filename = "day9.txt"

func TestPartOne(t *testing.T) {
	expected := 114
	actual := PartOne(testFilename)
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}

	expected = 1904165718
	actual = PartOne(filename)
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}
func TestPartTwo(t *testing.T) {
	expected := 2
	actual := PartTwo(testFilename)
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}

	expected = 964
	actual = PartTwo(filename)
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}
