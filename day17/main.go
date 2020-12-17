package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type cube struct {
	x, y, z int
}

func main() {
	slice0 := []string{
		"..#..#..",
		".###..#.",
		"#..##.#.",
		"#.#.#.#.",
		".#..###.",
		".....#..",
		"#...####",
		"##....#.",
	}

	grid3d := make(map[string]bool)

	for y, line := range slice0 {
		for x, cube := range line {
			if cube == '#' {
				grid3d[fmt.Sprint(x, "|", y, "|0")] = true
			}
		}
	}

	part1(grid3d)
	part2()
}

func isActive(grid3d map[string]bool, x, y, z int) bool {
	_, exists := grid3d[fmt.Sprint(x, "|", y, "|", z)]
	if exists {
		return true
	}
	return false
}

func activeCubes(grid3d map[string]bool) ([]int, []int, []int) {
	x := []int{}
	y := []int{}
	z := []int{}
	for key := range grid3d {
		parts := strings.Split(key, "|")
		xNum, _ := strconv.Atoi(parts[0])
		yNum, _ := strconv.Atoi(parts[1])
		zNum, _ := strconv.Atoi(parts[2])
		x = append(x, xNum)
		y = append(y, yNum)
		z = append(z, zNum)
	}
	return x, y, z
}

func inactiveRelevantCubes(grid3d map[string]bool) ([]int, []int, []int) {
	inactive := make(map[string]bool)
	for key := range grid3d {
		parts := strings.Split(key, "|")
		xNum, _ := strconv.Atoi(parts[0])
		yNum, _ := strconv.Atoi(parts[1])
		zNum, _ := strconv.Atoi(parts[2])
		for cX := xNum - 1; cX < xNum+2; cX++ {
			for cY := yNum - 1; cY < yNum+2; cY++ {
				for cZ := zNum - 1; cZ < zNum+2; cZ++ {
					if !isActive(grid3d, cX, cY, cZ) {
						inactive[fmt.Sprint(cX, "|", cY, "|", cZ)] = true
					}
				}
			}
		}
	}
	x := []int{}
	y := []int{}
	z := []int{}
	for key := range inactive {
		parts := strings.Split(key, "|")
		xNum, _ := strconv.Atoi(parts[0])
		yNum, _ := strconv.Atoi(parts[1])
		zNum, _ := strconv.Atoi(parts[2])
		x = append(x, xNum)
		y = append(y, yNum)
		z = append(z, zNum)

	}
	return x, y, z
}

func activeAround(grid3d map[string]bool, x, y, z int) ([]int, []int, []int) {
	aX := []int{}
	aY := []int{}
	aZ := []int{}
	for cX := x - 1; cX < x+2; cX++ {
		for cY := y - 1; cY < y+2; cY++ {
			for cZ := z - 1; cZ < z+2; cZ++ {
				if isActive(grid3d, cX, cY, cZ) {
					aX = append(aX, cX)
					aY = append(aY, cY)
					aZ = append(aZ, cZ)
				}
			}
		}
	}
	return aX, aY, aZ
}

func activate(grid3d map[string]bool, x, y, z int) {
	grid3d[fmt.Sprint(x, "|", y, "|", z)] = true
}

func deactivate(grid3d map[string]bool, x, y, z int) {
	delete(grid3d, fmt.Sprint(x, "|", y, "|", z))
}

func part1(grid3d map[string]bool) {
	for i := 0; i < 6; i++ {
		newGrid := make(map[string]bool)
		activeX, activeY, activeZ := activeCubes(grid3d)
		for idx := 0; idx < len(activeX); idx++ {
			xList, _, _ := activeAround(grid3d, activeX[idx], activeY[idx], activeZ[idx])
			activate(newGrid, activeX[idx], activeY[idx], activeZ[idx])
			if len(xList) != 2 && len(xList) != 3 {
				deactivate(newGrid, activeX[idx], activeY[idx], activeZ[idx])
			}
		}
		inactiveX, inactiveY, inactiveZ := inactiveRelevantCubes(grid3d)
		for idx := 0; idx < len(inactiveX); idx++ {
			xList, _, _ := activeAround(grid3d, inactiveX[idx], inactiveY[idx], inactiveZ[idx])
			deactivate(newGrid, inactiveX[idx], inactiveY[idx], inactiveZ[idx])
			if len(xList) == 3 {
				activate(newGrid, inactiveX[idx], inactiveY[idx], inactiveZ[idx])
			}
		}
		grid3d = newGrid
	}
	list, _, _ := activeCubes(grid3d)
	log.Println(len(list))
}

func part2() {

}
