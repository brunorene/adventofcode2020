package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/home/brsantos/go/src/github.com/brunorene/adventofcode2020/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var trees []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(trees)
	part2(trees)
}

func hasTree(trees []string, x int, y int) (bool, bool) {
	if y >= len(trees) {
		return false, true
	}
	return trees[y][x%len(trees[0])] == '#', false
}

func jump(right int, down int, x int, y int) (int, int) {
	return x + right, y + down
}

func part1(trees []string) {
	count := 0
	x := 0
	y := 0
	for true {
		x, y = jump(3, 1, x, y)
		isTree, theEnd := hasTree(trees, x, y)
		if theEnd {
			break
		}
		if isTree {
			count++
		}
	}
	log.Println(count)
}

func part2(trees []string) {
	counts := 0
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, slope := range slopes {
		count := 0
		x := 0
		y := 0
		for true {
			x, y = jump(slope[0], slope[1], x, y)
			isTree, theEnd := hasTree(trees, x, y)
			if theEnd {
				break
			}
			if isTree {
				count++
			}
		}
		if counts == 0 {
			counts = count
		} else {
			counts *= count
		}
	}
	log.Println(counts)
}
