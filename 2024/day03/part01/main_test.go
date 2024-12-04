package main

import (
	"strings"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/stretchr/testify/require"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	require.Equal(t, 161, CalculateMultiplications(strings.Join(lines, "\n")))
}
