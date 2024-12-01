package main

import (
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/stretchr/testify/require"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	a, b := BuildLocationLists(lines)
	f := CalculateListFrequency(b)
	require.Equal(t, 31, CalculateListSimilarity(a, f))
}

func Test_Frequency(t *testing.T) {
	list := []int{4, 3, 5, 3, 9, 3}
	require.Equal(t, map[int]int{3: 3, 4: 1, 5: 1, 9: 1}, CalculateListFrequency(list))
}
