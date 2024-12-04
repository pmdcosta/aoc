package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))
	fmt.Println(CalculateMultiplications(strings.Join(lines, "\n")))
}

func CalculateMultiplications(in string) (r int) {
	re := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	for _, mul := range re.FindAllString(in, -1) {
		n := input.SplitNumbers(mul)
		r += n[0] * n[1]
	}
	return r
}
