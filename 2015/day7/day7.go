package day7

import (
	"github.com/jrjgh/advent_of_code/util"
	"regexp"
	"strconv"
	"strings"
)

type circuit map[string]func(circuit) int

func (c circuit) get(s string) int {
	n := c[s](c)
	c[s] = func(circuit) int {
		return n
	}
	return n
}

func parseInstr(s string) (string, func(circuit) int) {
	instrs := strings.Split(s, " -> ")
	expr := instrs[0]
	out := instrs[1]
	return out, parseExpr(expr)
}

func partOne() int {
	lines, _ := util.ReadLinesFromFile("day7.txt")
	c := circuit{}
	for _, line := range lines {
		key, instr := parseInstr(line)
		c[key] = instr
	}
	return c.get("a")
}
func partTwo() int {
	out1 := partOne()
	lines, _ := util.ReadLinesFromFile("day7.txt")
	c := circuit{}
	for _, line := range lines {
		key, instr := parseInstr(line)
		c[key] = instr
	}
	c["b"] = func(circuit) int {
		return out1
	}
	return c.get("a")
}

func parseExpr(s string) func(circuit) int {
	tokens := strings.Split(s, " ")
	//NOT
	if strings.HasPrefix(s, "NOT") {
		arg := tokens[1]
		if n, err := strconv.Atoi(arg); err == nil {
			return func(circuit) int {
				return ^n
			}
		} else {
			return func(c circuit) int {
				return ^c.get(arg)
			}
		}
	}
	//num
	if m, err := regexp.MatchString("^[0-9]+$", s); err == nil && m {
		n, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}
		return func(circuit) int {
			return n
		}
	}
	//key
	if m, err := regexp.MatchString("^[a-z]+$", s); err == nil && m {
		return func(c circuit) int {
			arg := tokens[0]
			return c.get(arg)
		}
	}
	//AND
	if m, err := regexp.MatchString("[0-9a-z]+ AND [0-9a-z]+", s); err == nil && m {
		a := parseExpr(tokens[0])
		b := parseExpr(tokens[2])
		return func(c circuit) int {
			return a(c) & b(c)
		}
	}
	//OR
	if m, err := regexp.MatchString("[0-9a-z]+ OR [0-9a-z]+", s); err == nil && m {
		a := parseExpr(tokens[0])
		b := parseExpr(tokens[2])
		return func(c circuit) int {
			return a(c) | b(c)
		}
	}
	//LSHIFT
	if m, err := regexp.MatchString("[0-9a-z]+ LSHIFT [0-9a-z]+", s); err == nil && m {
		a := parseExpr(tokens[0])
		b := parseExpr(tokens[2])
		return func(c circuit) int {
			return a(c) << b(c)
		}
	}
	//RSHIFT
	if m, err := regexp.MatchString("[0-9a-z]+ RSHIFT [0-9a-z]+", s); err == nil && m {
		a := parseExpr(tokens[0])
		b := parseExpr(tokens[2])
		return func(c circuit) int {
			return a(c) >> b(c)
		}
	}
	println("OH NO")
	return func(circuit) int {
		println(s)
		return 0
	}
}
