package day10

import (
	"errors"
	"fmt"
	u "github.com/jrjgh/advent_of_code/util"
	"strings"
)

const (
	N      = 0b0001
	S      = 0b0010
	E      = 0b0100
	W      = 0b1000
	START  = 0b1111
	NS     = N | S
	NE     = N | E
	NW     = N | W
	SE     = S | E
	SW     = S | W
	EW     = E | W
	GROUND = 0b0001_0000
	MARKED = 0b0011_0000
)

func (d direction) String() string {
	switch d {
	case N:
		return "N"
	case E:
		return "E"
	case W:
		return "W"
	case S:
		return "S"
	}
	return "?"
}

type pipe int

func (p pipe) String() string {
	switch p {
	case NS:
		return "|"
	case EW:
		return "-"
	case NE:
		return "L"
	case NW:
		return "J"
	case SW:
		return "7"
	case SE:
		return "F"
	case START:
		return "S"
	case GROUND:
		return "."
	case MARKED:
		return "X"
	default:
		panic(fmt.Sprintf("unrecognized pipe: %d", p))
	}
}

type M struct {
	pipes [][]pipe
}

func (m *M) String() string {
	var b strings.Builder
	for _, pipes := range m.pipes {
		var b2 strings.Builder
		for _, pipe := range pipes {
			b2.WriteString(pipe.String())
		}
		b2.WriteString("\n")
		b.WriteString(b2.String())
	}
	return b.String()
}

func (m *M) start() *coord {
	for i := 0; i < len(m.pipes); i++ {
		for j := 0; j < len(m.pipes[i]); j++ {
			if m.pipes[i][j] == START {
				return &coord{i, j}
			}
		}
	}
	panic("start not found")
}

func (m M) get(c *coord) pipe {
	return m.pipes[c.i][c.j]
}

func (m M) set(c *coord, p pipe) {
	m.pipes[c.i][c.j] = p
}

type coord struct {
	i int
	j int
}

func (c *coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.i, c.j)
}

type direction int

func (c *coord) relativeDir(c2 *coord) (direction, error) {
	di := c2.i - c.i
	dj := c2.j - c.j

	si := u.Sign(di)
	sj := u.Sign(dj)

	if si*sj != 0 {
		return -1, errors.New("non-cardinal relative direction")
	}

	if u.Sign(di) == 1 {
		return S, nil
	} else if u.Sign(di) == -1 {
		return N, nil
	} else if u.Sign(dj) == 1 {
		return E, nil
	}
	return W, nil
}

func (p pipe) connectsTo(p2 pipe, relativeDir direction) bool {
	//fmt.Printf("%s from %s to %s\n", relativeDir, p, p2)
	var a int
	switch relativeDir {
	case N:
		return (p&N)*(p2&S) != 0
	case S:
		return (p&S)*(p2&N) != 0
	case E:
		return (p&E)*(p2&W) != 0
	case W:
		return (p&W)*(p2&E) != 0
	default:
		return a != 0
	}
}

func parseChar(c string) pipe {
	switch c {
	case "|":
		return NS
	case "-":
		return EW
	case "L":
		return NE
	case "J":
		return NW
	case "7":
		return SW
	case "F":
		return SE
	case "S":
		return START
	case ".":
		return GROUND
	default:
		panic(fmt.Sprintf("unrecognized character: %s", c))
	}
}

func parseLine(line string) []pipe {
	return u.Map(parseChar, strings.Split(line, ""))
}

func parseMap(lines []string) *M {
	return &M{
		u.Map(parseLine, lines),
	}
}

func adjacentCoords(c *coord) []*coord {
	return []*coord{
		{c.i - 1, c.j}, //left
		{c.i + 1, c.j}, //right
		{c.i, c.j - 1}, //above
		{c.i, c.j + 1}, //below
	}
}

func withinBounds(m *M) func(c *coord) bool {
	return func(c *coord) bool {
		return c.i >= 0 &&
			c.i < len(m.pipes) &&
			c.j >= 0 &&
			c.j < len(m.pipes[c.i])
	}
}

func traverse(m *M) int {
	start := m.start()
	count := -1
	for curr := []*coord{start}; len(curr) > 0; count++ {
		//fmt.Print(m.String())
		nexts := make([]*coord, 0)
		for _, c := range curr {
			adjs := u.Filter[*coord](withinBounds(m), adjacentCoords(c))
			//fmt.Printf("adjs: %+v\n", adjs)
			valids := u.Filter[*coord](validNext(m, c), adjs)
			//fmt.Printf("valids: %+v\n", valids)

			nexts = append(nexts, valids...)
			m.set(c, MARKED)
		}
		curr = nexts
	}
	return count
}

func validNext(m *M, start *coord) func(*coord) bool {
	return func(dest *coord) bool {
		dir, err := start.relativeDir(dest)
		if err != nil {
			panic(err)
		}
		p1 := m.get(start)
		p2 := m.get(dest)
		return p1.connectsTo(p2, dir)
	}
}

func PartOne(filename string) int {
	lines, _ := u.ReadLinesFromFile(filename)
	m := parseMap(lines)

	return traverse(m)
}
