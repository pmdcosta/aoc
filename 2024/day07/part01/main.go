package main

import (
	"fmt"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	var sum int
	for _, l := range lines {
		v, values := GetEquation(l)
		if CheckEquation(v, values) {
			sum += v
		}
	}
	fmt.Println(sum)
}

func GetEquation(line string) (t int, v []int) {
	l := input.SplitNumbers(line)
	return l[0], l[1:]
}

func CheckEquation(t int, values []int) bool {
	for _, o := range []func(int, int) int{Mul, Sum} {
		if CheckIteration(t, values[0], 1, o, values) {
			return true
		}
	}
	return false
}

func CheckIteration(t int, c int, i int, f func(int, int) int, values []int) bool {
	// check if we've reached the end.
	if i == len(values) {
		// check if we hit the target value.
		if c == t {
			return true
		}
		return false
	}

	// calculate new current value.
	c = f(c, values[i])

	// stop if we're past the target value already.
	if c > t {
		return false
	}

	// check the next value.
	for _, o := range []func(int, int) int{Mul, Sum} {
		if CheckIteration(t, c, i+1, o, values) {
			return true
		}
	}
	return false
}

func Mul(a, b int) int {
	return a * b
}

func Sum(a, b int) int {
	return a + b
}
