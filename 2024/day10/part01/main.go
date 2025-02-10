package main

import (
	"fmt"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/pmdcosta/aoc/2024/pkg/set"
)

type Position struct {
	y, x int
}

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	var count int
	for _, p := range FindTrails(ReadMap(lines)) {
		count += len(p)
	}
	fmt.Println(count)
}

// ReadMap builds the map from the input.
func ReadMap(lines []string) (m [][]int) {
	m = make([][]int, len(lines))
	for i, l := range lines {
		for _, r := range l {
			m[i] = append(m[i], input.GetNumber(r))
		}
	}
	return m
}

// FindTrails finds all 0 height positions and follows the trails, returning all the end positions for each trail.
func FindTrails(m [][]int) map[Position]set.Set[Position] {
	trails := make(map[Position]set.Set[Position])
	for y, row := range m {
		for x, n := range row {
			if n == 0 {
				ends := set.New[Position]()
				FollowTrail(m, y, x, ends)
				if len(ends) > 0 {
					trails[Position{x, y}] = ends
				}
			}
		}
	}
	return trails
}

// FollowTrail is a recursive function that starts at a 0 height position and follows a path until the end.
func FollowTrail(m [][]int, y, x int, trails set.Set[Position]) {
	// add the end position if we've reached the end of a trail.
	if m[y][x] == 9 {
		trails.Add(Position{y, x})
		return
	}

	// check going up.
	if y > 0 && m[y-1][x] == m[y][x]+1 {
		FollowTrail(m, y-1, x, trails)
	}

	// check going down.
	if y < len(m)-1 && m[y+1][x] == m[y][x]+1 {
		FollowTrail(m, y+1, x, trails)
	}

	// check going left.
	if x > 0 && m[y][x-1] == m[y][x]+1 {
		FollowTrail(m, y, x-1, trails)
	}

	// check going right.
	if x < len(m[0])-1 && m[y][x+1] == m[y][x]+1 {
		FollowTrail(m, y, x+1, trails)
	}
}
