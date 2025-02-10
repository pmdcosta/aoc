package main

import (
	"fmt"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	for i := 0; i < 8; i++ {
		Iterate(lines[0], i)
	}
}

func Iterate(l string, times int) {
	stones := input.SplitNumbers(l)
	for i := 0; i < times; i++ {
		stones = Blink(stones)
	}
	fmt.Println(len(stones), stones)
}
