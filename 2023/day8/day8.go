package day8

import (
	u "github.com/jrjgh/advent_of_code/util"
	"regexp"
	"strings"
)

type mapFollower struct {
	key  string
	m    map[string][]string
	dirs []string
	curr int
}

func (m *mapFollower) next() {
	direction := m.dirs[m.curr]
	if direction == "L" {
		m.key = m.m[m.key][0]
	} else {
		m.key = m.m[m.key][1]
	}
	m.curr++
	if m.curr == len(m.dirs) {
		m.curr = 0
	}
}

func tripLength(m *mapFollower) int {
	count := 0
	for !isEndPoint(m.key) {
		count++
		m.next()
	}
	return count
}
func PartOne(filename string) int {
	lines, _ := u.ReadLinesFromFile(filename)

	directions := strings.Split(lines[0], "")
	m := parseMap(lines[2:])
	mf := &mapFollower{
		key:  "AAA",
		m:    m,
		dirs: directions,
		curr: 0,
	}
	count := 0

	for mf.key != "ZZZ" {
		count++
		mf.next()
	}
	return count
}

func PartTwo(filename string) int {
	lines, _ := u.ReadLinesFromFile(filename)

	ls := u.Map(parseLine, lines[2:])
	startKeys := u.Map[[]string, string](func(ss []string) string {
		return ss[0]
	}, ls)
	starts := u.Filter[string](isStartingPoint, startKeys)
	directions := strings.Split(lines[0], "")
	m := parseMap(lines[2:])

	mfs := u.Map[string, *mapFollower](func(str string) *mapFollower {
		return &mapFollower{
			key:  str,
			dirs: directions,
			m:    m,
			curr: 0,
		}
	}, starts)

	lengths := u.Map[*mapFollower](tripLength, mfs)
	return LCM(lengths[0], lengths[1], lengths[2:]...)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func parseMap(lines []string) map[string][]string {
	m := map[string][]string{}
	rows := u.Map(parseLine, lines)
	for _, row := range rows {
		key := row[0]
		branches := row[1:]
		m[key] = branches
	}
	return m
}

func isStartingPoint(str string) bool {
	b, _ := regexp.MatchString(`\w\wA`, str)
	return b
}

func isEndPoint(str string) bool {
	b, _ := regexp.MatchString(`\w\wZ`, str)
	return b
}

func parseLine(line string) []string {
	return regexp.MustCompile(`[A-Z0-9]{3}`).FindAllString(line, 3)
}
