package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func id(code string) int {
	binary := ""
	for _, c := range code {
		if strings.Contains("BR", string(c)) {
			binary += "1"
		} else {
			binary += "0"
		}
	}
	num, _ := strconv.ParseInt(binary, 2, 0)
	return int(num)
}

func main() {
	file, err := os.Open("/home/brsantos/go/src/github.com/brunorene/adventofcode2020/day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	boardings := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boardings = append(boardings, id(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(boardings)
	part2(boardings)
}

func part1(boardings []int) {
	max := 0
	for _, boarding := range boardings {
		if boarding > max {
			max = boarding
		}
	}
	log.Println(max)
}

func part2(boardings []int) {
	sort.Ints(boardings)
	for i := 1; i < len(boardings)-1; i++ {
		if boardings[i-1]+1 < boardings[i] {
			log.Println(boardings[i-1] + 1)
		}
	}
}
