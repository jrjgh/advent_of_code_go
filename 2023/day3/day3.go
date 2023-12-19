package day3

import (
	"github.com/jrjgh/advent_of_code/util"
	r "regexp"
	"strconv"
)

var numExp = r.MustCompile(`[0-9]+`)
var symExp = r.MustCompile(`[^\.0-9]`)
var gearRegexp = r.MustCompile(`\*`)

func PartOne() int {
	s, _ := util.ReadLinesFromFile("day3.txt")
	p := partNumbers(s)
	return util.SumInt(p)
}

func PartTwo() int {
	s, _ := util.ReadLinesFromFile("day3.txt")
	gearRatios := gearRatios(s)
	return util.SumInt(gearRatios)
}

func partNumbers(lines []string) []int {
	partNumbers := make([]int, 0)
	mtx := isSymbolMtx(lines)
	minLeft := 0
	minTop := 0
	maxBottom := len(lines) - 1
	maxRight := len(lines[0]) - 1

	for i, line := range lines {
		numIs := numExp.FindAllStringIndex(line, -1)
		for _, numI := range numIs {
			num, _ := strconv.Atoi(line[numI[0]:numI[1]])

			leftDigitIndex := numI[0]
			rightDigitIndex := numI[1] - 1
			top := max(i-1, minTop)
			bottom := min(i+1, maxBottom)
			left := max(leftDigitIndex-1, minLeft)
			right := min(rightDigitIndex+1, maxRight)
		checkForSymbol:
			for k := top; k <= bottom; k++ {
				for l := left; l <= right; l++ {
					if mtx[k][l] {
						partNumbers = append(partNumbers, num)
						break checkForSymbol
					}
				}
			}
		}
	}

	return partNumbers
}

func isSymbolMtx(lines []string) [][]bool {
	mtx := make([][]bool, len(lines))
	for i, line := range lines {
		bools := make([]bool, len(line))
		matchesI := symExp.FindAllStringIndex(line, -1)
		for _, matchI := range matchesI {
			bools[matchI[0]] = true
		}
		mtx[i] = bools
	}
	return mtx
}

func gearRatios(lines []string) []int {
	gearRatios := make([]int, 0)
	//find all *
	gearCoords := make([][2]int, 0)
	for i, line := range lines {
		gearIs := gearRegexp.FindAllStringIndex(line, -1)
		for _, gearI := range gearIs {
			gearCoords = append(gearCoords, [2]int{i, gearI[0]})
		}
	}
	//find adjacent nums
	//fmt.Printf("len(gearCoords): %d\n", len(gearCoords))
	for _, gearCoord := range gearCoords {

		adjacents := make([]int, 0)
		i, j := gearCoord[0], gearCoord[1]
		for k := max(i-1, 0); k < i+2; k++ {
			line := lines[k]
			ms := numExp.FindAllStringIndex(line, -1)
			//fmt.Println(line)
			//fmt.Printf("ms: %+v\n", ms)
			for _, m := range ms {
				//fmt.Println("hi")
				//leftmost digit index
				left := m[0]
				lb := max(left-1, 0)
				//rightmost digit index
				right := m[1] - 1
				rb := min(right+1, len(line))
				if j >= lb && j <= rb {
					num, _ := strconv.Atoi(line[left : right+1])
					adjacents = append(adjacents, num)
				}
			}
		}
		if len(adjacents) == 2 {
			gearRatios = append(gearRatios, adjacents[0]*adjacents[1])
		}

	}
	return gearRatios
}
