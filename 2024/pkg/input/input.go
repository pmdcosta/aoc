package input

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// File default location of the input and example files.
const (
	File    = "../input"
	Example = "../example"
)

// DirFile returns the absolute path of the chosen input file in the current directory.
func DirFile(name string) string {
	_, p, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(p), name)
}

// ReadFile reads a file's content to a string.
// it assumes the file is in the same directory as the go executable.
func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	return string(data)
}

// SplitFile splits a file by lines.
func SplitFile(data string) (lines []string) {
	for _, line := range strings.Split(data, "\n") {
		lines = append(lines, line)
	}
	return lines
}

// SplitTestFile splits a test string by lines.
func SplitTestFile(data string) (lines []string) {
	split := strings.Split(data, "\n")
	for i, line := range split {
		if i == 0 || i == len(split)-1 {
			continue
		}
		lines = append(lines, strings.TrimPrefix(line, "\t\t"))
	}
	return lines
}

// SplitNumbers returns all numbers from the provided string.
func SplitNumbers(s string) (n []int) {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(s, -1)
	for _, a := range numbers {
		i, err := strconv.Atoi(a)
		if err != nil {
			log.Fatalf("failed to parse number: %v", err)
		}
		n = append(n, i)
	}
	return n
}

// GetNumber converts a tune to number,
func GetNumber(r rune) int {
	n, _ := strconv.Atoi(string(r))
	return n
}
