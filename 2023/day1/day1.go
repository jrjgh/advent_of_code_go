package day1

import (
	"fmt"
	"github.com/jrjgh/advent_of_code/util"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func PartOne() int {
	lines, err := util.ReadLinesFromFile("day1.txt")

	if err != nil {
		fmt.Println(err)
	}
	ds := util.Map[string, int](decode, lines)
	sum := util.SumInt(ds)
	return sum
}

func PartTwo() int {
	lines, err := util.ReadLinesFromFile("day1.txt")

	if err != nil {
		fmt.Println(err)
	}
	ds := util.Map[string, int](decode2, lines)
	sum := util.SumInt(ds)
	return sum
}

func decode(str string) int {
	chars := strings.Split(str, "")
	digits := make([]int, 0)
	for _, char := range chars {
		n, err := strconv.ParseInt(char, 10, 32)
		if err == nil {
			digits = append(digits, int(n))
		}
	}
	fst, lst := digits[0], digits[len(digits)-1]
	return fst*10 + lst
}

func decode2(str string) int {

	replacements := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	digitPattern := "(zero|one|two|three|four|five|six|seven|eight|nine|[0-9])"
	lastDigitPattern := fmt.Sprintf("(?m)(?:.*)(%s)", digitPattern)
	r1, err := regexp.Compile(digitPattern)
	if err != nil {
		log.Fatal(err)
	}
	r2, err := regexp.Compile(lastDigitPattern)
	if err != nil {
		log.Fatal(err)
	}
	firstMatch := r1.FindString(str)
	lastMatches := r2.FindStringSubmatch(str)
	lastMatch := lastMatches[len(lastMatches)-1]
	fst, lst := replacements[firstMatch], replacements[lastMatch]

	return fst*10 + lst
}
