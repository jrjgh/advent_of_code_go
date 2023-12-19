package day10

import "testing"

var testFilename = "day10_test.txt"
var filename = "day10.txt"

func TestPartOne(t *testing.T) {
	expected := 8
	actual := PartOne(testFilename)
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}

	expected = 6842
	actual = PartOne(filename)
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}
func TestPartTwo(t *testing.T) {
	//expected := 2
	//actual := PartTwo(testFilename)
	//if expected != actual {
	//	t.Fatalf("expected %d but got %d", expected, actual)
	//}
	//
	//expected = 964
	//actual = PartTwo(filename)
	//if expected != actual {
	//	t.Fatalf("expected %d but got %d", expected, actual)
	//}
}
