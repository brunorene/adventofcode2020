package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	moves := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moves = append(moves, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(moves)
	part2(moves)
}

func steer(direction rune, steerCount int) rune {
	for s := 0; s < int(math.Abs(float64(steerCount))); s++ {
		if steerCount > 0 {
			switch direction {
			case 'N':
				direction = 'E'
			case 'S':
				direction = 'W'
			case 'E':
				direction = 'S'
			case 'W':
				direction = 'N'
			}
		} else {
			switch direction {
			case 'N':
				direction = 'W'
			case 'S':
				direction = 'E'
			case 'E':
				direction = 'N'
			case 'W':
				direction = 'S'
			}
		}
	}
	return direction
}

func part1(moves []string) {
	direction := 'E'
	x := 0
	y := 0
	for _, move := range moves {
		val, _ := strconv.Atoi(move[1:])
		switch move[0] {
		case 'N':
			y -= val
		case 'S':
			y += val
		case 'E':
			x += val
		case 'W':
			x -= val
		case 'L':
			direction = steer(direction, -val/90)
		case 'R':
			direction = steer(direction, val/90)
		case 'F':
			switch direction {
			case 'N':
				y -= val
			case 'S':
				y += val
			case 'E':
				x += val
			case 'W':
				x -= val
			}
		}
	}
	log.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func steerWaypoint(wX int, wY int, steerCount int) (int, int) {
	abs := int(math.Abs(float64(steerCount)))
	for s := 0; s < abs; s++ {
		oldwWX := wX
		wX = (steerCount / abs) * wY
		wY = -(steerCount / abs) * oldwWX
	}
	return wX, wY
}

func part2(moves []string) {
	wX := 10
	wY := -1
	x := 0
	y := 0
	ninety := 90
	for _, move := range moves {
		val, _ := strconv.Atoi(move[1:])
		switch move[0] {
		case 'N':
			wY -= val
		case 'S':
			wY += val
		case 'E':
			wX += val
		case 'W':
			wX -= val
		case 'L':
			wX, wY = steerWaypoint(wX, wY, val/ninety)
		case 'R':
			wX, wY = steerWaypoint(wX, wY, -val/ninety)
		case 'F':
			x += val * wX
			y += val * wY
		}
	}
	log.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}
