package main

import (
	"fmt"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	var state []int
	for _, s := range input.SplitNumbers(lines[0]) {
		state = append(state, LoopStone(s, 0, 74)...)
	}
	fmt.Println(len(state))
}
