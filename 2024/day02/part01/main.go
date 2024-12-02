package main

import (
	"fmt"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/pmdcosta/aoc/2024/pkg/util"
)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	reports := BuildReports(lines)
	fmt.Println(CheckReports(reports))
}

func BuildReports(lines []string) (r [][]int) {
	for _, l := range lines {
		r = append(r, input.SplitNumbers(l))
	}
	return r
}

func CheckReports(reports [][]int) (safe int) {
	for _, report := range reports {
		if CheckReport(report, 0, 1, 0) {
			safe++
		}
	}
	return safe
}

func CheckReport(r []int, i int, j int, increasing int) bool {
	// check if the report has finished.
	if len(r) == j {
		return true
	}

	// set the increase direction if it hasn't been set yet.
	if increasing == 0 {
		if r[i] > r[j] {
			increasing--
		} else {
			increasing++
		}
	}

	// check if the report is unsafe due to difference in numbers being to how or too low.
	if d := util.IntDiff(r[i], r[j]); d > 3 || d == 0 {
		return false
	}

	// check if the numbers keep the expected direction of increase.
	if (increasing == 1 && r[i] > r[j]) || (increasing == -1 && r[i] < r[j]) {
		return false
	}

	// recursively check the next pair of numbers.
	return CheckReport(r, i+1, j+1, increasing)
}
