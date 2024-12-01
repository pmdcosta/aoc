package main

import (
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/stretchr/testify/require"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	a, b := BuildLocationLists(lines)
	require.Equal(t, 11, CalculateListDistance(a, b))
}
