package main

import (
	"fmt"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	x, y, gx, gy := GetObstacles(lines)
	visited := GetVisitedPositions(gx, gy, x, y)
	fmt.Println(CreateObstacles(gx, gy, x, y, visited))
}
