package main

import (
	"fmt"
	"strconv"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	fs := BuildFilesystem(lines[0])
	CompactFilesystem(fs)
	fmt.Println(Checksum(fs))
}

func BuildFilesystem(input string) (fs []int) {
	file := true
	fileID := 0
	for _, c := range input {
		v, _ := strconv.Atoi(string(c))
		if file {
			for i := 0; i < v; i++ {
				fs = append(fs, fileID)
			}
			file = false
			fileID++
		} else {
			for i := 0; i < v; i++ {
				fs = append(fs, -1)
			}
			file = true
		}
	}
	return fs
}

func CompactFilesystem(fs []int) {
	j := len(fs) - 1
	for i, v := range fs {
		if i == j {
			return
		}
		if v == -1 {
			for ; j > i; j-- {
				if fs[j] == -1 {
					continue
				}
				fs[i] = fs[j]
				fs[j] = -1
				break
			}
		}
	}
}

func Checksum(fs []int) (sum int) {
	for i, f := range fs {
		if f != -1 {
			sum += i * f
		}
	}
	return sum
}
