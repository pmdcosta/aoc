package main

import (
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	reports := BuildReports(lines)
	require.Equal(t, 4, CheckReports(reports))
}

func Test_Extra(t *testing.T) {
	assert.Equal(t, 1, CheckReports([][]int{{9, 1, 2, 3, 4}}))
	assert.Equal(t, 1, CheckReports([][]int{{1, 2, 2, 5, 7, 8}}))
	assert.Equal(t, 1, CheckReports([][]int{{1, 2, 5, 4, 5, 6}}))
	assert.Equal(t, 1, CheckReports([][]int{{1, 3, 4, 3, 5, 6}}))
	assert.Equal(t, 1, CheckReports([][]int{{42, 44, 47, 49, 51, 52, 54, 52}}))
	assert.Equal(t, 1, CheckReports([][]int{{2, 1, 2, 3, 4}}))
	assert.Equal(t, 1, CheckReports([][]int{{1, 2, 2, 5, 7, 8}}))
	assert.Equal(t, 1, CheckReports([][]int{{1, 3, 4, 3, 5, 6}}))
	assert.Equal(t, 1, CheckReports([][]int{{33, 36, 38, 35, 41}}))
	assert.Equal(t, 1, CheckReports([][]int{{1, 9, 3, 4, 5}}))
	assert.Equal(t, 1, CheckReports([][]int{{9, 1, 3, 4, 5}}))
}
