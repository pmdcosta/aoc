package main

import (
	"fmt"
	"strings"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/pmdcosta/aoc/2024/pkg/set"
)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

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

func GetMiddle(pages []int) int {
	return pages[len(pages)/2]
}

func GetRules(lines []string) (rules map[int]set.Set[int]) {
	rules = make(map[int]set.Set[int])
	for _, l := range lines {
		if strings.Contains(l, "|") {
			n := input.SplitNumbers(l)
			if len(rules[n[1]]) == 0 {
				rules[n[1]] = set.New(n[0])
			} else {
				rules[n[1]].Add(n[0])
			}
		}
	}
	return rules
}

func GetSequences(lines []string) (pages [][]int) {
	for _, l := range lines {
		if strings.Contains(l, ",") {
			n := input.SplitNumbers(l)
			pages = append(pages, n)
		}
	}
	return pages
}

func IterateSequence(sequence []int, rules map[int]set.Set[int]) (s []int, u bool) {
	// iterate through the sequence.
	for i := 0; i < len(sequence)-1; i++ {
		// check if the current number is on the rules
		if _, ok := rules[sequence[i]]; !ok {
			continue
		}

		// swap the current number if it should be.
		swapped := true
		for swapped {
			sequence, swapped = SwapNumbers(i, sequence, rules[sequence[i]])
			if swapped {
				u = true
			}
		}
	}
	return sequence, u
}

func SwapNumbers(c int, sequence []int, rule set.Set[int]) ([]int, bool) {
	// read the rest of the sequence.
	for i := c; i < len(sequence); i++ {
		// if any of the next numbers in the sequence are in the current rule.
		if rule.Contains(sequence[i]) {
			// swap numbers in the sequence.
			t := sequence[i]
			sequence[i] = sequence[c]
			sequence[c] = t
			return sequence, true
		}
	}
	return sequence, false
}
