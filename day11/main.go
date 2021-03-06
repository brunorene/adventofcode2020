package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seats := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}
		for _, seat := range line {
			row = append(row, seat)
		}
		seats = append(seats, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(seats)
	part2(seats)
}

func equals(seats1 [][]rune, seats2 [][]rune) bool {
	for row := 0; row < len(seats1); row++ {
		for col := 0; col < len(seats1[row]); col++ {
			if seats1[row][col] != seats2[row][col] {
				return false
			}
		}
	}
	return true
}

func canBeEmpty(row int, col int, seats [][]rune, emptyLimit int, countViewedClear func(int, int, [][]rune) int) bool {
	if seats[row][col] != '#' {
		return false
	}
	return countViewedClear(row, col, seats) <= emptyLimit
}

func canBeOccupied(row int, col int, seats [][]rune, countViewedClear func(int, int, [][]rune) int) bool {
	if seats[row][col] != 'L' {
		return false
	}
	return countViewedClear(row, col, seats) == 8
}

func countViewedClearNear(row int, col int, seats [][]rune) int {
	isClear := 8
	for diffRow := -1; diffRow < 2; diffRow++ {
		for diffCol := -1; diffCol < 2; diffCol++ {
			if !(diffCol == 0 && diffRow == 0) &&
				row+diffRow >= 0 &&
				row+diffRow < len(seats) &&
				col+diffCol >= 0 &&
				col+diffCol < len(seats[0]) &&
				seats[row+diffRow][col+diffCol] == '#' {
				isClear--
			}
		}
	}
	return isClear
}

func countViewedClearWide(row int, col int, seats [][]rune) int {
	isClear := 8
	for offsetRow := -1; offsetRow < 2; offsetRow++ {
		for offsetCol := -1; offsetCol < 2; offsetCol++ {
			if offsetRow != 0 || offsetCol != 0 {
				currentRow := row + offsetRow
				currentCol := col + offsetCol
				for {
					if currentRow < 0 ||
						currentCol < 0 ||
						currentRow == len(seats) ||
						currentCol == len(seats[0]) ||
						seats[currentRow][currentCol] == 'L' {
						break
					}
					if seats[currentRow][currentCol] == '#' {
						isClear--
						break
					}
					currentRow += offsetRow
					currentCol += offsetCol
				}
			}
		}
	}
	return isClear
}

func changeState(seats [][]rune, emptyLimit int, countViewedClear func(int, int, [][]rune) int) [][]rune {
	newSeats := [][]rune{}
	for row := 0; row < len(seats); row++ {
		newRow := []rune{}
		for col := 0; col < len(seats[row]); col++ {
			switch {
			case canBeOccupied(row, col, seats, countViewedClear):
				newRow = append(newRow, '#')
			case canBeEmpty(row, col, seats, emptyLimit, countViewedClear):
				newRow = append(newRow, 'L')
			default:
				newRow = append(newRow, seats[row][col])
			}
		}
		newSeats = append(newSeats, newRow)
	}
	return newSeats
}

func countSeats(seats [][]rune, state rune) int {
	count := 0
	for row := 0; row < len(seats); row++ {
		for col := 0; col < len(seats[row]); col++ {
			if seats[row][col] == state {
				count++
			}
		}
	}
	return count
}

func part1(seats [][]rune) {
	newSeats := changeState(seats, 4, countViewedClearNear)
	for !equals(seats, newSeats) {
		seats = newSeats
		newSeats = changeState(seats, 4, countViewedClearNear)
	}
	log.Println(countSeats(seats, '#'))
}

func part2(seats [][]rune) {
	newSeats := changeState(seats, 3, countViewedClearWide)
	for !equals(seats, newSeats) {
		seats = newSeats
		newSeats = changeState(seats, 3, countViewedClearWide)
	}
	log.Println(countSeats(seats, '#'))
}
