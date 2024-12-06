package main

import (
	"fmt"
	"slices"

	"github.com/pmdcosta/aoc/2024/pkg/input"
	"github.com/pmdcosta/aoc/2024/pkg/set"
)

const (
	Up    = "up"
	Down  = "down"
	Left  = "left"
	Right = "right"
)

var lx, ly int

func main() {
	lines := input.SplitFile(input.ReadFile(input.DirFile(input.File)))
	x, y, gx, gy := GetObstacles(lines)
	fmt.Println(Patrol(gx, gy, x, y))
}

func GetObstacles(lines []string) (x map[int][]int, y map[int][]int, gx int, gy int) {
	x, y = make(map[int][]int), make(map[int][]int)
	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				x[i] = append(x[i], j)
				y[j] = append(y[j], i)
				slices.Sort(x[i])
				slices.Sort(y[j])
			} else if c == '^' {
				gx = i
				gy = j
			}
		}
	}
	lx = len(lines)
	ly = len(lines[0])
	return x, y, gx, gy
}

func Patrol(gx, gy int, x map[int][]int, y map[int][]int) int {
	visited := set.New(S(gx, gy))
	direction := Up
	for gx != -1 && gx != lx && gy != -1 && gy != ly {
		nx, ny := Walk(gx, gy, x, y, direction)
		visited.Add(GetVisited(gx, gy, nx, ny)...)
		gx = nx
		gy = ny
		direction = Turn(direction)
	}
	return len(visited)
}

func Walk(gx, gy int, x map[int][]int, y map[int][]int, direction string) (int, int) {
	switch direction {
	case Down:
		if obs, ok := y[gy]; ok {
			for i := 0; i < len(obs); i++ {
				if obs[i] > gx {
					return obs[i] - 1, gy
				}
			}
		}
		return lx, gy
	case Up:
		if obs, ok := y[gy]; ok {
			for i := len(obs) - 1; i >= 0; i-- {
				if obs[i] < gx {
					return obs[i] + 1, gy
				}
			}
		}
		return -1, gy
	case Right:
		if obs, ok := x[gx]; ok {
			for i := 0; i < len(obs); i++ {
				if obs[i] > gy {
					return gx, obs[i] - 1
				}
			}
		}
		return gx, ly
	case Left:
		if obs, ok := x[gx]; ok {
			for i := len(obs) - 1; i >= 0; i-- {
				if obs[i] < gy {
					return gx, obs[i] + 1
				}
			}
		}
		return gx, -1
	}
	return -1, -1
}

func S(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func Turn(d string) string {
	switch d {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	default:
		return Up
	}
}

func GetVisited(gx, gy, nx, ny int) (v []string) {
	if nx == -1 {
		nx++
	}
	if nx == lx {
		nx--
	}
	if ny == -1 {
		ny++
	}
	if ny == ly {
		ny--
	}
	if gx > nx {
		for i := nx; i <= gx; i++ {
			v = append(v, S(i, gy))
		}
	}
	if gx < nx {
		for i := gx; i <= nx; i++ {
			v = append(v, S(i, gy))
		}
	}
	if gy > ny {
		for i := ny; i <= gy; i++ {
			v = append(v, S(gx, i))
		}
	}
	if gy < ny {
		for i := gy; i <= ny; i++ {
			v = append(v, S(gx, i))
		}
	}
	fmt.Printf("%d,%d -> %d,%d = %v\n", gx, gy, nx, ny, v)
	return v
}
