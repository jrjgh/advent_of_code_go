package day8

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	out := partOne()
	expected := 1350
	println(out)
	if out != expected {
		t.Errorf("expected %d but got %d", expected, out)
	}
}

func TestPartTwo(t *testing.T) {
	out := partTwo()
	wrong := 2685
	println(out)
	if out == wrong {
		t.Errorf("got wrong answer: %d", wrong)
	}
}
