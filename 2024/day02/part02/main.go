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
		if CheckReport(report, 0, 1, 0, -1) {
			safe++
		} else if CheckReport(report, 1, 2, 0, 0) {
			// edge case: attempt to skip the first element.
			safe++
		} else if CheckReport(report, 0, 2, 0, 1) {
			// edge case: attempt to skip the second element.
			safe++
		} else {
			fmt.Printf("Unsafe: %v\n", report)
		}
	}
	return safe
}

func CheckReport(r []int, i int, j int, increasing int, skipped int) bool {
	// check if the report has finished.
	if len(r) == j {
		fmt.Printf("Safe: %v [%d]\n", r, skipped)
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

	// check if the next number should be skipped.
	if i+1 == skipped {
		i++
	}

	// check the next pair of numbers in the report.
	safe := CheckReport(r, i+1, j+1, increasing, skipped)
	if safe {
		return true
	} else if skipped != -1 {
		return false
	}

	// failed, attempt to skip next number.
	if CheckReport(r, i+1, j+2, increasing, j+1) {
		return true
	}

	// failed, attempt to skip current number.
	return CheckReport(r, i, j+1, increasing, i+1)
}
