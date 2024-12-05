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
	pages, sets := GetPages(lines)

	var result int
	for i := range pages {
		if EvaluatePage(pages[i], sets[i], rules) {
			result += GetMiddle(pages[i])
		}
	}
	fmt.Println(result)
}

func GetRules(lines []string) (rules map[int][]int) {
	rules = make(map[int][]int)
	for _, l := range lines {
		if strings.Contains(l, "|") {
			n := input.SplitNumbers(l)
			rules[n[1]] = append(rules[n[1]], n[0])
		}
	}
	return rules
}

func GetPages(lines []string) (pages [][]int, sets []set.Set[int]) {
	for _, l := range lines {
		if strings.Contains(l, ",") {
			n := input.SplitNumbers(l)
			pages = append(pages, n)
			sets = append(sets, set.New(n...))
		}
	}
	return pages, sets
}

func EvaluatePage(pages []int, elements set.Set[int], rules map[int][]int) bool {
	// create set holding all the elements previously read.
	var read set.Set[int]

	for _, p := range pages {
		// is current element in one of the rules.
		if req, ok := rules[p]; ok {
			// check if any of the required numbers in the rule is in the page.
			for _, r := range req {
				// required element is in the page.
				if elements.Contains(r) {
					// check if the required element has already been read.
					if len(read) == 0 || !read.Contains(r) {
						return false
					}
				}
			}
		}
		if read == nil {
			read = set.New(p)
		} else {
			read.Add(p)
		}
	}
	return true
}

func GetMiddle(pages []int) int {
	return pages[len(pages)/2]
}
