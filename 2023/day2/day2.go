package day2

import (
	"fmt"
	"github.com/jrjgh/advent_of_code/util"
	"regexp"
	"strconv"
	"strings"
)

type gamesubset struct {
	red   int
	blue  int
	green int
}

type game struct {
	id      int
	subsets []*gamesubset
}

func PartOne() int {
	lines, err := util.ReadLinesFromFile("day2.txt")
	if err != nil {
		fmt.Println(err)
	}
	games := make([]*game, len(lines))
	for i, line := range lines {
		games[i] = parseGame(line)
	}

	limits := &gamesubset{
		red:   12,
		green: 13,
		blue:  14,
	}
	sum := 0
g:
	for _, game := range games {
		for _, subset := range game.subsets {
			if subset.red > limits.red ||
				subset.green > limits.green ||
				subset.blue > limits.blue {
				continue g
			}
		}
		sum += game.id
	}
	return sum
}

func PartTwo() int {
	lines, err := util.ReadLinesFromFile("day2.txt")
	if err != nil {
		fmt.Println(err)
	}
	games := make([]*game, len(lines))
	for i, line := range lines {
		games[i] = parseGame(line)
	}

	sum := 0

	for _, game := range games {
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for _, subset := range game.subsets {
			if subset.red > maxRed {
				maxRed = subset.red
			}
			if subset.green > maxGreen {
				maxGreen = subset.green
			}
			if subset.blue > maxBlue {
				maxBlue = subset.blue
			}
		}
		sum += maxRed * maxBlue * maxGreen
	}
	return sum
}

func parseGame(str string) *game {
	pattern := `Game (\d+): (.*)$`
	r := regexp.MustCompile(pattern)
	matches := r.FindAllStringSubmatch(str, -1)
	g := &game{}
	g.id, _ = strconv.Atoi(matches[0][1])
	subsets := matches[0][2]
	g.subsets = parseGamesubsets(subsets)
	return g
}

func parseGamesubsets(str string) []*gamesubset {
	subsetStrs := strings.Split(str, "; ")
	subsets := make([]*gamesubset, len(subsetStrs))
	for i, subsetStr := range subsetStrs {
		subsets[i] = parseGamesubset(subsetStr)
	}
	return subsets
}

func parseGamesubset(str string) *gamesubset {
	g := &gamesubset{}
	redR := regexp.MustCompile(`(\d+)(?: red)`)
	blueR := regexp.MustCompile(`(\d+)(?: blue)`)
	greenR := regexp.MustCompile(`(\d+)(?: green)`)

	if redMatch := redR.FindStringSubmatch(str); len(redMatch) >= 2 {
		g.red, _ = strconv.Atoi(redMatch[1])
	}
	if blueMatch := blueR.FindStringSubmatch(str); len(blueMatch) >= 2 {
		g.blue, _ = strconv.Atoi(blueMatch[1])
	}
	if greenMatch := greenR.FindStringSubmatch(str); len(greenMatch) >= 2 {
		g.green, _ = strconv.Atoi(greenMatch[1])
	}

	return g
}
