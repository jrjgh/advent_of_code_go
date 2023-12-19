package day5

import (
	"github.com/jrjgh/advent_of_code/util"
	"testing"
)

var tf, _ = util.ReadLinesFromFile("day5_test.txt")

func TestPartOne(t *testing.T) {
	//expected := 0
	//actual := PartOne(tf)
	//println(tf)
	//if expected != actual {
	//	t.Fatalf("expected %d but got %d", expected, actual)
	//}

	expected := 111627841
	actual := PartOne(lines)
	//4314294015
	//191284628
	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 69323688
	actual := PartTwo(lines)

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

//func TestParseFile(t *testing.T) {
//	parseFile()
//}
