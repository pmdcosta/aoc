package main

import (
	"fmt"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	rules := GetRules(lines)
	pages, sets := GetPages(lines)

	for i := range pages {
		fmt.Println(EvaluatePage(pages[i], sets[i], rules))
	}
}
