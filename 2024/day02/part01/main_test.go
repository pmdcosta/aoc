package main

import (
	"fmt"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/stretchr/testify/require"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	reports := BuildReports(lines)
	require.Equal(t, 2, CheckReports(reports))
}

func Test_SkipTriesEarlier(t *testing.T) {
	fmt.Println(CheckReports([][]int{{1, 2, 5, 4, 5, 6, 7}}))
}
