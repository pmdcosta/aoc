package main

import (
	"fmt"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	var sum int
	for _, l := range lines {
		v, values := GetEquation(l)
		if CheckEquation(v, values) {
			sum += v
		}
	}
	fmt.Println(sum)
}

func Test_Extra(t *testing.T) {
	v, values := GetEquation("3267: 81 40 27")
	fmt.Println(CheckEquation(v, values))
}
