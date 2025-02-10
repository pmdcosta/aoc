package main

import (
	"fmt"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	var count int
	for _, p := range FindTrails(ReadMap(lines)) {
		count += p
	}
	fmt.Println(count)
}

func Test_Extra(t *testing.T) {
}
