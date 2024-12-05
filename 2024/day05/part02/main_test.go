package main

import (
	"fmt"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/pmdcosta/aoc/2024/pkg/set"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	rules := GetRules(lines)
	sequences := GetSequences(lines)

	var result int
	for _, s := range sequences {
		if seq, u := IterateSequence(s, rules); u {
			result += GetMiddle(seq)
		}
	}
	fmt.Println(result)
}

func Test_Page(t *testing.T) {
	fmt.Println(GetRules([]string{"8|6", "9|6"}))
	fmt.Println(IterateSequence([]int{5, 6, 7, 8, 9}, map[int]set.Set[int]{6: set.New(8, 9)}))
}

func Test_Full(t *testing.T) {
	rules := GetRules([]string{"8|6", "9|6"})
	fmt.Println(IterateSequence([]int{5, 6, 7, 8, 9}, rules))
}
