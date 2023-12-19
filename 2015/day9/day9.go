package day9

import (
	"fmt"
	"github.com/jrjgh/advent_of_code/util"
)

type Edge struct {
	Dest *Vertex
	Dist int
}
type Vertex struct {
	name  string
	Edges map[string]*Edge
}

func (v Vertex) connectsTo(v2 *Vertex) bool {
	return v.Edges[v2.name] != nil
}

func (v Vertex) getAllDestinations(seen map[string]struct{}) []string {
	if seen == nil {
		seen = map[string]struct{}{
			v.name: struct{}{},
		}
	}
	for s, edge := range v.Edges {
		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			destinations := edge.Dest.getAllDestinations(seen)
			for _, destination := range destinations {
				seen[destination] = struct{}{}
			}
		}
	}
	dests := make([]string, len(seen))
	i := 0
	for s := range seen {
		dests[i] = s
		i++
	}
	return dests
}

type Dir map[string]*Vertex

func partOne() int {
	lines, _ := util.ReadLinesFromFile("day9.txt")

	dir := Dir{}
	for _, line := range lines {
		orig, dest, dist := parseDistance(line)
		origv := dir[orig]
		destv := dir[dest]
		if origv == nil {
			origv = &Vertex{name: orig}
			dir[orig] = origv
		}
		if destv == nil {
			destv = &Vertex{name: dest}
			dir[dest] = destv
		}
		if !origv.connectsTo(destv) {
			origv.Edges[dest] = &Edge{Dest: destv, Dist: dist}
		}
		if !destv.connectsTo(origv) {
			destv.Edges[orig] = &Edge{Dest: origv, Dist: dist}
		}
	}
	//can start at any location, must visit each one exactly once
	//nah use a graph
	return -1
}

func findShortestGrandTour(v *Vertex) []string {
	//find all destinations
	//destinations := v.getAllDestinations(nil)
	return []string{}
}

func parseDistance(s string) (string, string, int) {
	var orig, dest string
	var dist int
	fmt.Sscanf(s, "%s to %s = %d", &orig, &dest, &dist)
	return orig, dest, dist
}
