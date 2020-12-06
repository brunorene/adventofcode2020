package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/home/brsantos/go/src/github.com/brunorene/adventofcode2020/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	answersPerGroup := []map[rune]int{make(map[rune]int)}
	groupSizes := []int{0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.Trim(line, " ")) == 0 {
			answersPerGroup = append([]map[rune]int{make(map[rune]int)}, answersPerGroup...)
			groupSizes = append([]int{0}, groupSizes...)
			continue
		}
		groupSizes[0]++
		for _, answer := range line {
			_, exists := answersPerGroup[0][answer]
			if !exists {
				answersPerGroup[0][answer] = 0
			}
			answersPerGroup[0][answer]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(answersPerGroup)
	part2(answersPerGroup, groupSizes)
}

func part1(answers []map[rune]int) {
	count := 0
	for _, answer := range answers {
		count += len(answer)
	}
	log.Println(count)
}

func part2(answers []map[rune]int, groupSizes []int) {
	count := 0
	for index, answer := range answers {
		for _, personCount := range answer {
			if personCount == groupSizes[index] {
				count++
			}
		}
	}
	log.Println(count)
}
