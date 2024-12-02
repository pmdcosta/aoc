package util

import "math"

func IntDiff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
