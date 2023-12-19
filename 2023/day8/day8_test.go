package day8

import "testing"

var testFilename = "day8_test.txt"
var testFilename2 = "day8_test2.txt"
var filename = "day8.txt"

func TestPartOne(t *testing.T) {
	expected := 6
	actual := PartOne(testFilename)
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}

	expected = 16343
	actual = PartOne(filename)
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 6
	actual := PartTwo(testFilename2)
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}

	expected = 15299095336639
	actual = PartTwo(filename)
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}
