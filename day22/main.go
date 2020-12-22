package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day22/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p1 := []int{}
	p2 := []int{}

	state := "p1"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Player") {
			continue
		}
		if line == "" {
			state = "p2"
			continue
		}
		switch state {
		case "p1":
			val, _ := strconv.Atoi(line)
			p1 = append(p1, val)
		case "p2":
			val, _ := strconv.Atoi(line)
			p2 = append(p2, val)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(p1, p2)
	part2(p1, p2)
}

func game(p1 []int, p2 []int) ([]int, []int) {
	history := [][]int{}
	for len(p1) > 0 && len(p2) > 0 {
		for _, past := range history {
			if equals(p1, past) {
				return p1, []int{}
			}
		}
		history = append(history, p1)
		p1Hand := p1[0]
		p1 = p1[1:]
		p2Hand := p2[0]
		p2 = p2[1:]
		if p1Hand > p2Hand {
			p1 = append(p1, p1Hand, p2Hand)
		} else {
			p2 = append(p2, p2Hand, p1Hand)
		}
	}
	return p1, p2
}

func score(p1 []int, p2 []int) int {
	winner := p2
	if len(p1) > 0 {
		winner = p1
	}
	sum := 0
	for index, card := range winner {
		sum += card * (len(winner) - index)
	}
	return sum
}

func equals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for index := range a {
		if a[index] != b[index] {
			return false
		}
	}
	return true
}

func part1(p1 []int, p2 []int) {
	p1, p2 = game(p1, p2)
	fmt.Println(score(p1, p2))
}

func recursiveGame(p1 []int, p2 []int) ([]int, []int) {
	history := [][]int{}
	fmt.Println("recursive start", p1, p2)
	for len(p1) > 0 && len(p2) > 0 {
		for _, past := range history {
			if equals(p1, past) {
				return p1, []int{}
			}
		}
		history = append(history, p1)
		p1Hand := p1[0]
		p1 = p1[1:]
		p2Hand := p2[0]
		p2 = p2[1:]
		if len(p1) >= p1Hand && len(p2) >= p2Hand {
			recP1 := make([]int, p1Hand)
			copy(recP1, p1)
			recP2 := make([]int, p2Hand)
			copy(recP2, p2)
			recP1, _ = recursiveGame(recP1, recP2)
			if len(recP1) > 0 {
				p1 = append(p1, p1Hand, p2Hand)
			} else {
				p2 = append(p2, p2Hand, p1Hand)
			}
		} else if p1Hand > p2Hand {
			p1 = append(p1, p1Hand, p2Hand)
		} else {
			p2 = append(p2, p2Hand, p1Hand)
		}
	}
	return p1, p2
}

func part2(p1 []int, p2 []int) {
	p1, p2 = recursiveGame(p1, p2)
	fmt.Println(score(p1, p2))
}
