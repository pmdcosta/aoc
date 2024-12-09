package main

import (
	"fmt"
	"testing"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func Test_ProvidedExample(t *testing.T) {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.Example)))

	fs := BuildFilesystem(lines[0])
	CompactFilesystem(fs)
	fmt.Println(fs)
	fmt.Println(Checksum(fs))

}

func Test_Extra(t *testing.T) {
}
