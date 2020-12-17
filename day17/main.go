package main

import (
	"log"
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

	grid3d := make(map[cube]bool)
	grid4d := make(map[hypercube]bool)

	for y, line := range slice0 {
		for x, active := range line {
			if active == '#' {
				grid3d[cube{x, y, 0}] = true
				grid4d[hypercube{x, y, 0, 0}] = true
			}
		}
	}

	part1(grid3d)
	part2(grid4d)
}

func isActive(grid3d map[cube]bool, c cube) bool {
	_, exists := grid3d[c]
	if exists {
		return true
	}
	return false
}

func activeCubes(grid3d map[cube]bool) []cube {
	cubes := []cube{}
	for key := range grid3d {
		cubes = append(cubes, key)
	}
	return cubes
}

func inactiveRelevantCubes(grid3d map[cube]bool) []cube {
	inactive := make(map[cube]bool)
	for key := range grid3d {
		for x := key.x - 1; x < key.x+2; x++ {
			for y := key.y - 1; y < key.y+2; y++ {
				for z := key.z - 1; z < key.z+2; z++ {
					c := cube{x, y, z}
					if !isActive(grid3d, c) {
						inactive[c] = true
					}
				}
			}
		}
	}
	cubes := []cube{}
	for key := range inactive {
		cubes = append(cubes, key)
	}
	return cubes
}

func activeAround(grid3d map[cube]bool, c cube) []cube {
	active := make(map[cube]bool)
	for x := c.x - 1; x < c.x+2; x++ {
		for y := c.y - 1; y < c.y+2; y++ {
			for z := c.z - 1; z < c.z+2; z++ {
				around := cube{x, y, z}
				if c != around {
					if isActive(grid3d, around) {
						active[around] = true
					}
				}
			}
		}
	}
	cubes := []cube{}
	for key := range active {
		cubes = append(cubes, key)
	}
	return cubes
}

func activate(grid3d map[cube]bool, c cube) {
	grid3d[c] = true
}

func deactivate(grid3d map[cube]bool, c cube) {
	delete(grid3d, c)
}

func part1(grid3d map[cube]bool) {
	for i := 0; i < 6; i++ {
		newGrid := make(map[cube]bool)
		activeCubes := activeCubes(grid3d)
		for _, cube := range activeCubes {
			activeAround := activeAround(grid3d, cube)
			activate(newGrid, cube)
			if len(activeAround) != 2 && len(activeAround) != 3 {
				deactivate(newGrid, cube)
			}
		}
		inactiveCubes := inactiveRelevantCubes(grid3d)
		for _, cube := range inactiveCubes {
			activeCubes := activeAround(grid3d, cube)
			deactivate(newGrid, cube)
			if len(activeCubes) == 3 {
				activate(newGrid, cube)
			}
		}
		grid3d = newGrid
	}
	list := activeCubes(grid3d)
	log.Println(len(list))
}

type hypercube struct {
	x, y, z, w int
}

func isActive4(grid4d map[hypercube]bool, c hypercube) bool {
	_, exists := grid4d[c]
	if exists {
		return true
	}
	return false
}

func activeCubes4(grid4d map[hypercube]bool) []hypercube {
	cubes := []hypercube{}
	for key := range grid4d {
		cubes = append(cubes, key)
	}
	return cubes
}

func inactiveRelevantCubes4(grid4d map[hypercube]bool) []hypercube {
	inactive := make(map[hypercube]bool)
	for key := range grid4d {
		for x := key.x - 1; x < key.x+2; x++ {
			for y := key.y - 1; y < key.y+2; y++ {
				for z := key.z - 1; z < key.z+2; z++ {
					for w := key.w - 1; w < key.w+2; w++ {
						c := hypercube{x, y, z, w}
						if !isActive4(grid4d, c) {
							inactive[c] = true
						}
					}
				}
			}
		}
	}
	cubes := []hypercube{}
	for key := range inactive {
		cubes = append(cubes, key)
	}
	return cubes
}

func activeAround4(grid4d map[hypercube]bool, c hypercube) []hypercube {
	active := make(map[hypercube]bool)
	for x := c.x - 1; x < c.x+2; x++ {
		for y := c.y - 1; y < c.y+2; y++ {
			for z := c.z - 1; z < c.z+2; z++ {
				for w := c.w - 1; w < c.w+2; w++ {
					around := hypercube{x, y, z, w}
					if c != around {
						if isActive4(grid4d, around) {
							active[around] = true
						}
					}
				}
			}
		}
	}
	cubes := []hypercube{}
	for key := range active {
		cubes = append(cubes, key)
	}
	return cubes
}

func activate4(grid4d map[hypercube]bool, c hypercube) {
	grid4d[c] = true
}

func deactivate4(grid4d map[hypercube]bool, c hypercube) {
	delete(grid4d, c)
}

func part2(grid4d map[hypercube]bool) {
	for i := 0; i < 6; i++ {
		newGrid := make(map[hypercube]bool)
		activeCubes := activeCubes4(grid4d)
		for _, cube := range activeCubes {
			activeAround := activeAround4(grid4d, cube)
			activate4(newGrid, cube)
			if len(activeAround) != 2 && len(activeAround) != 3 {
				deactivate4(newGrid, cube)
			}
		}
		inactiveCubes := inactiveRelevantCubes4(grid4d)
		for _, cube := range inactiveCubes {
			activeCubes := activeAround4(grid4d, cube)
			deactivate4(newGrid, cube)
			if len(activeCubes) == 3 {
				activate4(newGrid, cube)
			}
		}
		grid4d = newGrid
	}
	list := activeCubes4(grid4d)
	log.Println(len(list))
}
