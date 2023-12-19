package day9

import (
	u "github.com/jrjgh/advent_of_code/util"
	"regexp"
)

var numExp = regexp.MustCompile(`(-?\d)+`)

func PartOne(filename string) int {
	lines, _ := u.ReadLinesFromFile(filename)
	ns := u.Map(parseLine, lines)
	nexts := u.Map(extrapolate, ns)
	return u.SumInt(nexts)

}

func PartTwo(filename string) int {
	lines, _ := u.ReadLinesFromFile(filename)
	ns := u.Map(parseLine, lines)
	prevs := u.Map(extrapolate2, ns)
	return u.SumInt(prevs)

}

func parseLine(line string) []int {
	return u.Map(u.ParseInt, numExp.FindAllString(line, -1))
}

func extrapolate(ns []int) int {
	if u.All[int](u.IsZero[int], ns) {
		return 0
	}
	ds := diffs(ns)
	return ns[len(ns)-1] + extrapolate(ds)
}

func extrapolate2(ns []int) int {
	//fmt.Println(ns)
	if u.All[int](u.IsZero[int], ns) {
		return 0
	}
	ds := diffs(ns)
	o := ns[0] - extrapolate2(ds)
	//fmt.Println(o)
	return o
}

func diffs(ns []int) []int {
	out := make([]int, len(ns)-1)
	for i, u2 := range ns[:len(ns)-1] {
		out[i] = ns[i+1] - u2
	}
	return out
}
