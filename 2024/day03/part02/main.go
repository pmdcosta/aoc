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

func CalculateMultiplications(in string) (c int) {
	do := regexp.MustCompile("do\\(\\)")
	dont := regexp.MustCompile("don't\\(\\)")

	// get all the relevant operations.
	operations := GetAllOperations(in)

	// iterate over all operations.
	enabled := true
	for _, op := range operations {
		if do.MatchString(op) {
			enabled = true
			continue
		} else if dont.MatchString(op) {
			enabled = false
			continue
		}

		if enabled {
			n := input.SplitNumbers(op)
			c += n[0] * n[1]
		}
	}
	return c
}

func GetAllOperations(in string) (s []string) {
	re := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don't\\(\\)")
	for _, e := range re.FindAllString(in, -1) {
		s = append(s, e)
	}
	return s
}
