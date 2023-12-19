package day4

import (
	"github.com/jrjgh/advent_of_code/util"
	"math"
	"regexp"
	"strconv"
)

const winningNumbers, playerNumbers = 10, 25

func getCards() []*card {
	lines, _ := util.ReadLinesFromFile("day4.txt")

	cards := make([]*card, len(lines))
	for i, line := range lines {
		cards[i] = parseCard(line, winningNumbers, playerNumbers)
	}
	return cards
}

func PartOne() int {

	cards := getCards()

	totalScore := 0
	for _, c := range cards {
		totalScore += c.score()
	}

	return totalScore
}

func PartTwo() int {
	cards := getCards()
	counts := make([]int, len(cards))

	for i, c := range cards {
		counts[i] += 1
		winners := c.countWinners()
		for j := 0; j < counts[i]; j++ {
			for k := 1; k <= winners; k++ {
				counts[i+k] += 1
			}
		}
	}

	sum := 0
	for _, count := range counts {
		sum += count
	}
	return sum
}

type card struct {
	id             int
	winningNumbers map[int]struct{}
	playerNumbers  []int
}

func (c *card) score() int {
	exp := -1.0
	for _, number := range c.playerNumbers {
		if _, ok := c.winningNumbers[number]; ok {
			exp += 1.0
		}
	}
	if exp < 0.0 {
		return 0
	}
	return int(math.Pow(2, exp))
}

func (c *card) countWinners() int {
	count := 0
	for _, number := range c.playerNumbers {
		if _, ok := c.winningNumbers[number]; ok {
			count += 1
		}
	}
	return count
}

func parseCard(line string, winningNumbers, playerNumbers int) *card {
	cardPattern := `Card\s+\d+:(\d+)+|(\d+)+`
	cardRegexp := regexp.MustCompile(cardPattern)
	matches := cardRegexp.FindAllStringSubmatch(line, -1)
	id, _ := strconv.Atoi(matches[0][0])
	winningNumberMatches := matches[1 : winningNumbers+1]
	playerNumberMatches := matches[winningNumbers+1:]
	c := &card{
		id:             id,
		winningNumbers: map[int]struct{}{},
		playerNumbers:  make([]int, playerNumbers),
	}
	for _, match := range winningNumberMatches {
		n, _ := strconv.Atoi(match[0])
		c.winningNumbers[n] = struct{}{}
	}
	for i, match := range playerNumberMatches {
		n, _ := strconv.Atoi(match[0])
		c.playerNumbers[i] = n
	}
	return c
}
