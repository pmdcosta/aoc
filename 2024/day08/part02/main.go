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
		for _, node := range GetAntinodes(p, lx, ly) {
			antinodes.Add(node)
		}
	}
	return antinodes
}

// GetAntinodes returns all the antinodes created by all the positions of a single antenna type.
func GetAntinodes(antennas []Position, lx, ly int) (antinodes []Position) {
	for i, a := range antennas {
		for _, b := range antennas[i+1:] {
			antinodes = append(antinodes, GetResonantAntinodes(a, b, lx, ly)...)
			antinodes = append(antinodes, GetResonantAntinodes(b, a, lx, ly)...)
		}
	}
	return antinodes
}

// GetResonantAntinodes returns all antinodes created by the resonance of two antennas in a line.
func GetResonantAntinodes(a, b Position, lx, ly int) (nodes []Position) {
	// calculate the difference between positions.
	dx, dy := b.x-a.x, b.y-a.y

	// start with the first antenna and create antinodes until we reach outside the map.
	for x, y := a.x, a.y; x >= 0 && x < lx && y >= 0 && y < ly; x, y = x-dx, y-dy {
		nodes = append(nodes, Position{x, y})
	}
	return nodes
}
