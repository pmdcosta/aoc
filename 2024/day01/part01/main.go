package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	a, b := BuildLocationLists(lines)
	fmt.Println(CalculateListDistance(a, b))
}

func BuildLocationLists(lines []string) (a, b []int) {
	for _, l := range lines {
		ids := input.SplitNumbers(l)
		a = append(a, ids[0])
		b = append(b, ids[1])
	}
	return a, b
}

func CalculateListDistance(a, b []int) (d int) {
	sort.Ints(a)
	sort.Ints(b)

	for i := range a {
		d += int(math.Abs(float64(a[i] - b[i])))
	}
	return d
}
