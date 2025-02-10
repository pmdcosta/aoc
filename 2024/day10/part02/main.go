package main

import (
	"fmt"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

type Position struct {
	y, x int
}

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	var count int
	for _, p := range FindTrails(ReadMap(lines)) {
		count += p
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

// FindTrails finds all 0 height positions and follows the trails, returning the ratting of each trail.
func FindTrails(m [][]int) map[Position]int {
	trails := make(map[Position]int)
	for y, row := range m {
		for x, n := range row {
			if n == 0 {
				var p int
				FollowTrail(m, y, x, &p)
				if p > 0 {
					trails[Position{x, y}] = p
				}
			}
		}
	}
	return trails
}

// FollowTrail is a recursive function that starts at a 0 height position and follows a path until the end.
func FollowTrail(m [][]int, y, x int, rating *int) {
	// add the end position if we've reached the end of a trail.
	if m[y][x] == 9 {
		*rating = *rating + 1
		return
	}

	// check going up.
	if y > 0 && m[y-1][x] == m[y][x]+1 {
		FollowTrail(m, y-1, x, rating)
	}

	// check going down.
	if y < len(m)-1 && m[y+1][x] == m[y][x]+1 {
		FollowTrail(m, y+1, x, rating)
	}

	// check going left.
	if x > 0 && m[y][x-1] == m[y][x]+1 {
		FollowTrail(m, y, x-1, rating)
	}

	// check going right.
	if x < len(m[0])-1 && m[y][x+1] == m[y][x]+1 {
		FollowTrail(m, y, x+1, rating)
	}
}
