package day6

import (
	"fmt"
	"github.com/jrjgh/advent_of_code/util"
	"strings"
)

type cmd struct {
	f  func(c cmd, g [1000][1000]int) [1000][1000]int
	x1 int
	y1 int
	x2 int
	y2 int
}

func partTwo() int {
	lines, _ := util.ReadLinesFromFile("day6.txt")
	cmds := parseCommands(lines)
	grid := [1000][1000]int{}

	for _, c := range cmds {
		grid = c.f(c, grid)
	}

	sum := 0
	for _, ints := range grid {
		for _, b := range ints {
			sum += b
		}
	}
	return sum
}

func toggle(c cmd, g [1000][1000]int) [1000][1000]int {
	for i := c.y1; i <= c.y2; i++ {
		for j := c.x1; j <= c.x2; j++ {
			g[i][j] = g[i][j] + 2
		}
	}
	return g
}

func turnOff(c cmd, g [1000][1000]int) [1000][1000]int {
	for i := c.y1; i <= c.y2; i++ {
		for j := c.x1; j <= c.x2; j++ {
			if g[i][j] > 0 {
				g[i][j] = g[i][j] - 1
			}
		}
	}
	return g
}

func turnOn(c cmd, g [1000][1000]int) [1000][1000]int {
	for i := c.y1; i <= c.y2; i++ {
		for j := c.x1; j <= c.x2; j++ {
			g[i][j] = g[i][j] + 1
		}
	}
	return g
}

func parseCommand(s string) cmd {
	var x1, y1, x2, y2 int
	if strings.HasPrefix(s, "toggle") {
		fmt.Sscanf(s, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		return cmd{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
			f:  toggle,
		}
	}
	if strings.HasPrefix(s, "turn on") {
		fmt.Sscanf(s, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		return cmd{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
			f:  turnOn,
		}
	}
	if strings.HasPrefix(s, "turn off") {
		fmt.Sscanf(s, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		return cmd{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
			f:  turnOff,
		}
	}
	return cmd{}
}

func parseCommands(ss []string) []cmd {
	cmds := make([]cmd, len(ss))
	for i, s := range ss {
		cmds[i] = parseCommand(s)
	}
	return cmds
}

//type cmd struct {
//	f  func(c cmd, g [1000][1000]bool) [1000][1000]bool
//	x1 int
//	y1 int
//	x2 int
//	y2 int
//}
//
//func partOne() int {
//	lines, _ := util.ReadLinesFromFile("day6.txt")
//	cmds := parseCommands(lines)
//	grid := [1000][1000]bool{}
//
//	for _, c := range cmds {
//		grid = c.f(c, grid)
//	}
//
//	sum := 0
//	for _, bools := range grid {
//		for _, b := range bools {
//			if b {
//				sum++
//			}
//		}
//	}
//	return sum
//}
//
//func toggle(c cmd, g [1000][1000]bool) [1000][1000]bool {
//	for i := c.y1; i <= c.y2; i++ {
//		for j := c.x1; j <= c.x2; j++ {
//			g[i][j] = !g[i][j]
//		}
//	}
//	return g
//}
//
//func turnOff(c cmd, g [1000][1000]bool) [1000][1000]bool {
//	for i := c.y1; i <= c.y2; i++ {
//		for j := c.x1; j <= c.x2; j++ {
//			g[i][j] = false
//		}
//	}
//	return g
//}
//
//func turnOn(c cmd, g [1000][1000]bool) [1000][1000]bool {
//	for i := c.y1; i <= c.y2; i++ {
//		for j := c.x1; j <= c.x2; j++ {
//			g[i][j] = true
//		}
//	}
//	return g
//}
//
//func parseCommand(s string) cmd {
//	var x1, y1, x2, y2 int
//	if strings.HasPrefix(s, "toggle") {
//		fmt.Sscanf(s, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
//		return cmd{
//			x1: x1,
//			y1: y1,
//			x2: x2,
//			y2: y2,
//			f:  toggle,
//		}
//	}
//	if strings.HasPrefix(s, "turn on") {
//		fmt.Sscanf(s, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
//		return cmd{
//			x1: x1,
//			y1: y1,
//			x2: x2,
//			y2: y2,
//			f:  turnOn,
//		}
//	}
//	if strings.HasPrefix(s, "turn off") {
//		fmt.Sscanf(s, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
//		return cmd{
//			x1: x1,
//			y1: y1,
//			x2: x2,
//			y2: y2,
//			f:  turnOff,
//		}
//	}
//	return cmd{}
//}
//
//func parseCommands(ss []string) []cmd {
//	cmds := make([]cmd, len(ss))
//	for i, s := range ss {
//		cmds[i] = parseCommand(s)
//	}
//	return cmds
//}
