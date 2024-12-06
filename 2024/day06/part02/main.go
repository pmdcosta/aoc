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
	visited := GetVisitedPositions(gx, gy, x, y)
	fmt.Println(CreateObstacles(gx, gy, x, y, visited))
}

// GetObstacles returns all the obstacles indexed in both the x and y directions along with the guard's starting position.
func GetObstacles(lines []string) (x map[int][]int, y map[int][]int, gx int, gy int) {
	x, y = make(map[int][]int), make(map[int][]int)
	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				// store the position of an obstacle.
				x[i] = append(x[i], j)
				y[j] = append(y[j], i)
				slices.Sort(x[i])
				slices.Sort(y[j])
			} else if c == '^' {
				// store the starting position of the guard.
				gx = i
				gy = j
			}
		}
	}
	// store the limits of the map.
	lx = len(lines)
	ly = len(lines[0])
	return x, y, gx, gy
}

// GetVisitedPositions simulates the guard's patrol path and returns all the positions he walked through.
func GetVisitedPositions(gx, gy int, x map[int][]int, y map[int][]int) set.Set[string] {
	// starting position and direction.
	visited, direction := set.New(S(gx, gy)), Up

	// simulate the patrol until the guard reaches the edges of the lab.
	for gx != -1 && gx != lx && gy != -1 && gy != ly {
		// simulate a walk until the guard reaches an obstacle and turns.
		nx, ny := Walk(gx, gy, x, y, direction)
		direction = Turn(direction)

		// store all the positions the guard went through.
		visited.Add(GetVisited(gx, gy, nx, ny)...)

		// update the guard's position.
		gx = nx
		gy = ny
	}
	return visited
}

// Walk simulate a patrol section until the guard reaches an obstacle.
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

// S is a helper function to store a position as a string.
func S(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

// Turn turns the guard 90 degrees.
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

// GetVisited returns all the visited positions between two points.
func GetVisited(gx, gy, nx, ny int) (v []string) {
	// check if the guard has reached the edge of the map.
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

	// calculate all the positions in between two points.
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
	return v
}

// CheckLooped simulates the guard's patrol path and returns whether he gets stuck on a loop.
func CheckLooped(gx, gy int, x map[int][]int, y map[int][]int) bool {
	// starting position and direction.
	// the guard only loops if he reaches the same position facing the same direction.
	direction := Up
	stopped := set.New(S(gx, gy) + direction)

	// simulate the patrol until the guard reaches the edges of the lab or starts looping.
	for gx != -1 && gx != lx && gy != -1 && gy != ly {
		// simulate a walk until the guard reaches an obstacle and turns.
		gx, gy = Walk(gx, gy, x, y, direction)
		direction = Turn(direction)

		// check if the guard has previously stopped at this position facing the same direction.
		if stopped.Contains(S(gx, gy) + direction) {
			return true
		}
		stopped.Add(S(gx, gy) + direction)
	}
	return false
}

// CreateObstacles iterates through a list of new obstacles and simulates the guard's patrol to find out how many of them result in a loop.
func CreateObstacles(gx, gy int, ox map[int][]int, oy map[int][]int, created set.Set[string]) int {
	var looped int
	for c := range created {
		// add the new obstacle to the list.
		nc := input.SplitNumbers(c)
		nx, ny := deepCopy(ox), deepCopy(oy)
		nx[nc[0]] = append(nx[nc[0]], nc[1])
		ny[nc[1]] = append(ny[nc[1]], nc[0])
		slices.Sort(nx[nc[0]])
		slices.Sort(ny[nc[1]])

		// simulate the patrol and check if the guard loops with the new obstacle.
		if CheckLooped(gx, gy, nx, ny) {
			looped++
		}
	}
	return looped
}

// deepCopy helper function to create a deep copy of the obstacle list.
func deepCopy(m map[int][]int) map[int][]int {
	nm := make(map[int][]int)
	for k, v := range m {
		for _, j := range v {
			nm[k] = append(nm[k], j)
		}
	}
	return nm
}
