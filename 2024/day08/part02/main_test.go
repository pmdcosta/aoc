package main

import (
	"fmt"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	antennas := GetAntennas(lines)
	antinodes := GetAllAntinodes(antennas, len(lines[0]), len(lines))
	fmt.Println(len(antinodes))
}
