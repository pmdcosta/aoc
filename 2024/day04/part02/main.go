package main

import (
	"fmt"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

const (
	LetterM rune = 77
	LetterA rune = 65
	LetterS rune = 83
)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	stage := BuildStage(lines)
	fmt.Println(IterateStage(stage))
}

func BuildStage(in []string) (w [][]rune) {
	for _, l := range in {
		w = append(w, []rune(l))
	}
	return w
}

func IterateStage(w [][]rune) (count int) {
	for i, l := range w {
		for j, c := range l {
			if c == LetterA {
				count += CheckLine(w, i, j)
			}
		}
	}
	return count
}

func CheckLine(w [][]rune, i, j int) (count int) {
	if j-1 >= 0 && i-1 >= 0 && j+1 < len(w[i]) && i+1 < len(w) &&
		CheckWord(w[i-1][j-1], w[i][j], w[i+1][j+1]) &&
		CheckWord(w[i+1][j-1], w[i][j], w[i-1][j+1]) {
		count++
	}
	return count
}

func CheckWord(w ...rune) bool {
	return (w[0] == LetterM && w[1] == LetterA && w[2] == LetterS) || (w[0] == LetterS && w[1] == LetterA && w[2] == LetterM)
}
