package day5

import (
	"github.com/jrjgh/advent_of_code/util"
	"regexp"
	"strconv"
	"strings"
)

var lines, _ = util.ReadLinesFromFile("day5.txt")

func PartOne(lines []string) int {
	seeds, alm := parseFile(lines)
	locations := util.Map(alm.dest, seeds)
	m, _ := util.Min[int](locations)
	return m
}

func PartTwo(lines []string) int {
	seedRanges, alm := parseFile(lines)
	expandedSeedRanges := expandSeedRanges(seedRanges)
	locations := util.Map(alm.dest, expandedSeedRanges)
	m, _ := util.Min[int](locations)
	return m
}

func expandSeedRanges(rs []int) []int {
	seeds := make([]int, 0)
	for i := 0; i < len(rs); i += 2 {
		r := expandSeedRange(rs[i], rs[i+1])
		seeds = append(seeds, r...)
	}
	return seeds
}

func expandSeedRange(a, b int) []int {
	ns := make([]int, b)
	for n := 0; n < b; n++ {
		ns[n] = a + n
	}
	return ns
}

type almanac struct {
	ranges []*almRange
	next   *almanac
}

type almRange struct {
	sourceStart int
	destStart   int
	rangeLen    int
}

func (a *almanac) dest(source int) int {
	d := source
	for _, arange := range a.ranges {
		dest := arange.dest(source)
		if dest != -1 {
			d = dest
			break
		}
	}
	if a.next != nil {
		return a.next.dest(d)
	}
	return d
}

func (r *almRange) dest(source int) int {
	xrange := r.rangeLen - 1
	if source >= r.sourceStart && source <= r.sourceStart+xrange {
		d := source - r.sourceStart
		return r.destStart + d
	}
	return -1
}

//seed -> soil -> fert -> watr -> light -> temp -> humid -> loc

func parseFile(lines []string) (seeds []int, _alm *almanac) {
	file := strings.Join(lines, "\n")
	sections := strings.Split(file, "\n\n")

	seedRegexp := regexp.MustCompile(`\d+`)
	seedMatches := seedRegexp.FindAllString(sections[0], -1)
	seeds = make([]int, len(seedMatches))
	for i, match := range seedMatches {
		s, _ := strconv.Atoi(match)
		seeds[i] = s
	}

	var head, curr *almanac

	lineRegexp := regexp.MustCompile(`(\d+) (\d+) (\d+)`)
	for _, section := range sections[1:] {
		lines := strings.Split(section, "\n")[1:]
		alm := &almanac{
			ranges: make([]*almRange, len(lines)),
		}
		for i, line := range lines {
			matches := lineRegexp.FindAllStringSubmatch(line, -1)
			dest, _ := strconv.Atoi(matches[0][1])
			source, _ := strconv.Atoi(matches[0][2])
			rlen, _ := strconv.Atoi(matches[0][3])
			alm.ranges[i] = &almRange{
				sourceStart: source,
				destStart:   dest,
				rangeLen:    rlen,
			}
		}
		if head == nil {
			head = alm
			curr = alm
		} else {
			curr.next = alm
			curr = curr.next
		}
	}
	return seeds, head
}
