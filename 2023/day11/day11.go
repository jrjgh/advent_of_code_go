package day11

import "fmt"

const (
	SPACE = region(iota)
	GALAXY
)

type region int

func (r region) String() string {
	switch r {
	case SPACE:
		return "."
	case GALAXY:
		return "#"
	default:
		panic(fmt.Sprintf("not a valid region :%d", r))
	}
}

func isSpace(r region) bool {
	if r == SPACE {
		return true
	}
	return false
}

func isGalaxy(r region) bool {
	if r == GALAXY {
		return true
	}
	return false
}

func parseRegion(str string) region {
	switch str {
	case ".":
		return SPACE
	case "#":
		return GALAXY
	default:
		panic(fmt.Sprintf("not a valid region: %s", str))
	}
}

type galaxy struct {
	regions [][]region
}

func (g *galaxy) expand() {
	var allSpace = true
	//by row
	for i, regions := range g.regions {

	}

	//by column
	allSpace = true
	for j, region := range g.regions[0] {

	}
}
