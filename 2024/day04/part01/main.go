package main

import (
	"fmt"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

const (
	LetterX rune = 88
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
			if c == LetterX {
				count += CheckLine(w, i, j)
			}
		}
	}
	return count
}

func CheckLine(w [][]rune, i, j int) (count int) {
	// check horizontal.
	if j+3 < len(w[i]) && w[i][j+1] == LetterM && w[i][j+2] == LetterA && w[i][j+3] == LetterS {
		count++
	}
	if j-3 >= 0 && w[i][j-1] == LetterM && w[i][j-2] == LetterA && w[i][j-3] == LetterS {
		count++
	}

	// check vertical.
	if i+3 < len(w) && w[i+1][j] == LetterM && w[i+2][j] == LetterA && w[i+3][j] == LetterS {
		count++
	}
	if i-3 >= 0 && w[i-1][j] == LetterM && w[i-2][j] == LetterA && w[i-3][j] == LetterS {
		count++
	}

	// check diagonals.
	if j+3 < len(w[i]) && i+3 < len(w) &&
		w[i+1][j+1] == LetterM && w[i+2][j+2] == LetterA && w[i+3][j+3] == LetterS {
		count++
	}
	if j-3 >= 0 && i-3 >= 0 &&
		w[i-1][j-1] == LetterM && w[i-2][j-2] == LetterA && w[i-3][j-3] == LetterS {
		count++
	}
	if j+3 < len(w[i]) && i-3 >= 0 &&
		w[i-1][j+1] == LetterM && w[i-2][j+2] == LetterA && w[i-3][j+3] == LetterS {
		count++
	}
	if j-3 >= 0 && i+3 < len(w) &&
		w[i+1][j-1] == LetterM && w[i+2][j-2] == LetterA && w[i+3][j-3] == LetterS {
		count++
	}
	return count
}
