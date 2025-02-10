package main

import (
	"fmt"
	"strconv"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

var cache = make(map[int]map[int][]int)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	stones := input.SplitNumbers(lines[0])
	for i := 0; i < 75; i++ {
		fmt.Println(i, len(stones))
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

func LoopStone(s int, times int, limit int) []int {
	// check if we've reached the end.
	if times == limit {
		return IterateStone(s)
	}

	// check cache for value.
	if iter, ok := cache[s]; ok {
		if stones, ok := iter[limit-times]; ok {
			// fmt.Println("cached", s, limit-times)
			return stones
		}
	}

	// iterate each new stone.
	times++
	var line []int
	for _, v := range IterateStone(s) {
		val := LoopStone(v, times, limit)
		if cache[v] == nil {
			cache[v] = map[int][]int{limit - times: val}
		} else {
			cache[v][limit-times] = val
		}
		line = append(line, val...)
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
