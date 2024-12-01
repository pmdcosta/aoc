package main

import (
	"fmt"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	a, b := BuildLocationLists(lines)
	f := CalculateListFrequency(b)
	fmt.Println(CalculateListSimilarity(a, f))
}

func BuildLocationLists(lines []string) (a, b []int) {
	for _, l := range lines {
		ids := input.SplitNumbers(l)
		a = append(a, ids[0])
		b = append(b, ids[1])
	}
	return a, b
}

func CalculateListFrequency(a []int) map[int]int {
	frequency := make(map[int]int)
	for _, v := range a {
		frequency[v]++
	}
	return frequency
}

func CalculateListSimilarity(a []int, f map[int]int) (d int) {
	for _, v := range a {
		d += v * f[v]
	}
	return d
}
