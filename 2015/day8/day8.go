package day8

import (
	"github.com/jrjgh/advent_of_code/util"
	"strconv"
)

func partOne() int {
	lines, _ := util.ReadLinesFromFile("day8.txt")
	sum := 0
	for _, line := range lines {
		sum += len(line) - len(unquote(line))
	}
	return sum
}

func partTwo() int {
	lines, _ := util.ReadLinesFromFile("day8.txt")
	sum := 0
	for _, line := range lines {
		sum += len(quote(line)) - len(line)
	}
	return sum
}

func quote(str string) string {
	s := strconv.Quote(str)
	return s
}

func unquote(str string) string {
	s, _ := strconv.Unquote(str)
	return s
}
