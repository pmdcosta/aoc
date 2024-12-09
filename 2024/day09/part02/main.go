package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pmdcosta/aoc/2024/pkg/input"
)

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))

	fs := BuildFilesystem(lines[0])
	CompactFilesystem(fs)
	fmt.Println(Checksum(fs))
}

// BuildFilesystem reads the input string and builds the filesystem as a slice of blocks.
func BuildFilesystem(input string) (fs []int) {
	file := true
	var fileID int
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

// CompactFilesystem reads through the filesystem back to front and attempts to move every file to the first fitting empty location.
func CompactFilesystem(fs []int) {
	// find the next file.
	for i := len(fs) - 1; i > 0; {
		// skip empty blocks.
		if fs[i] == -1 {
			i--
			continue
		}

		// get the current block's size.
		fSize := GetFileSize(i, fs)

		// get the first fitting empty slot.
		if eStart := GetEmptySlot(fSize, i, fs); eStart != -1 {
			// move the file to the empty slot.
			MoveFile(i-fSize+1, eStart, fSize, fs)
		}
		i -= fSize
	}
}

// GetEmptySlot returns the index of the first empty slot where the file would fit.
func GetEmptySlot(fSize int, limit int, fs []int) int {
	for e := 0; e < limit; e++ {
		// skip file blocks.
		if fs[e] != -1 {
			continue
		}

		// get the size of the current empty slot.
		size := GetEmptyBlockSize(e, fs)
		if size >= fSize {
			return e
		}
		e += size - 1
	}
	return -1
}

// MoveFile moves the file to the empty slot location and empties the previous file location.
func MoveFile(fStart, eStart, size int, fs []int) {
	for i := eStart; i < eStart+size; i++ {
		fs[i] = fs[fStart]
	}
	for i := fStart; i < fStart+size; i++ {
		fs[i] = -1
	}
}

// GetFileSize returns the size of the current file.
func GetFileSize(e int, fs []int) (size int) {
	id := fs[e]
	for i := e; i > 0; i-- {
		if fs[i] != id {
			return size
		}
		size++
	}
	return size
}

// GetEmptyBlockSize returns the size of the current empty block.
func GetEmptyBlockSize(e int, fs []int) (size int) {
	for i := e; i < len(fs); i++ {
		if fs[i] != -1 {
			return size
		}
		size++
	}
	return size
}

// Checksum calculates the checksum of the filesystem.
func Checksum(fs []int) (sum int) {
	for i, f := range fs {
		if f != -1 {
			sum += i * f
		}
	}
	return sum
}

// PrintFilesystem is a helper function to print the filesystem.
func PrintFilesystem(fs []int) {
	var s []string
	for _, f := range fs {
		if f == -1 {
			s = append(s, ".")
		} else {
			s = append(s, strconv.Itoa(f))
		}
	}
	fmt.Println(strings.Join(s, ""))
}
