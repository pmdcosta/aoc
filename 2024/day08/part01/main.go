package main

import (
	"fmt"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/pmdcosta/aoc/2024/pkg/set"
)

type Position struct {
	x, y int
}

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	antennas := GetAntennas(lines)
	antinodes := GetAllAntinodes(antennas, len(lines[0]), len(lines))
	fmt.Println(len(antinodes))
}

// GetAntennas returns all antenna positions mapped by antenna type.
func GetAntennas(lines []string) map[rune][]Position {
	antennas := make(map[rune][]Position)
	for y, l := range lines {
		for x, c := range l {
			if c != '.' && c != '#' {
				antennas[c] = append(antennas[c], Position{x, y})
			}
		}
	}
	return antennas
}

// GetAllAntinodes returns all antinodes created by all antennas within the map.
func GetAllAntinodes(antennas map[rune][]Position, lx, ly int) set.Set[Position] {
	antinodes := set.Set[Position]{}
	for _, p := range antennas {
		for _, node := range GetAntinodes(p) {
			if node.x >= 0 && node.x < lx && node.y >= 0 && node.y < ly {
				antinodes.Add(node)
			}
		}
	}
	return antinodes
}

// GetAntinodes returns all the antinodes created by all the positions of a single antenna type.
func GetAntinodes(antennas []Position) (antinodes []Position) {
	for i, a := range antennas {
		for _, b := range antennas[i+1:] {
			antinodes = append(antinodes,
				Position{a.x - (b.x - a.x), a.y - (b.y - a.y)},
				Position{b.x - (a.x - b.x), b.y - (a.y - b.y)})
		}
	}
	return antinodes
}
