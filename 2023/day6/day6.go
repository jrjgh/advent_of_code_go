package day6

import (
	u "github.com/jrjgh/advent_of_code/util"
	"regexp"
	"strings"
)

var numExp = regexp.MustCompile(`\d+`)

func PartOne(filename string) int {
	times, records := parseFile(filename)
	races := u.Zip2(times, records)
	wtbrs := u.Map[[]int, int](WaysToBeatRecord, races)
	//fmt.Printf("wtbrs: %+v\n", wtbrs)
	return u.ProdInt(wtbrs)
}

func PartTwo(filename string) int {
	time, record := parseFile2(filename)
	return WaysToBeatRecord([]int{time, record})
}

func WaysToBeatRecord(race []int) int {
	time, record := race[0], race[1]
	speedTimePairs := u.PartitionInt(time)
	distances := u.Map(dist, speedTimePairs)

	beatsRecord := func(distance int) bool {
		return distance > record
	}
	winningDistances := u.Filter[int](beatsRecord, distances)
	return len(winningDistances)
}

func dist(race []int) int {
	return race[0] * race[1]
}

func parseFile(filename string) (times, distances []int) {
	lines, _ := u.ReadLinesFromFile(filename)

	times = u.Map(u.ParseInt, numExp.FindAllString(lines[0], -1))
	distances = u.Map(u.ParseInt, numExp.FindAllString(lines[1], -1))

	return times, distances
}
func parseFile2(filename string) (time, distance int) {
	lines, _ := u.ReadLinesFromFile(filename)

	time = u.ParseInt(strings.Join(numExp.FindAllString(lines[0], -1), ""))
	distance = u.ParseInt(strings.Join(numExp.FindAllString(lines[1], -1), ""))

	return time, distance
}
