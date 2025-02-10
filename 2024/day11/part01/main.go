package main

import (
	"fmt"
	"strconv"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	stones := input.SplitNumbers(lines[0])
	for i := 0; i < 25; i++ {
		stones = Blink(stones)
	}
	fmt.Println(len(stones))
}

func Blink(stones []int) (line []int) {
	for _, n := range stones {
		line = append(line, IterateStone(n)...)
	}
	return line
}

func IterateStone(s int) []int {
	// first rule.
	if s == 0 {
		return []int{1}
	}

	// second rule.
	if str := strconv.Itoa(s); len(str)%2 == 0 {
		a, _ := strconv.Atoi(str[0 : len(str)/2])
		b, _ := strconv.Atoi(str[len(str)/2:])
		return []int{a, b}
	}

	// third rule.
	return []int{s * 2024}
}
