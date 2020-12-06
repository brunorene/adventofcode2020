package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Data password information
type Data struct {
	start    int
	end      int
	letter   string
	password string
}

func main() {
	file, err := os.Open("/home/brsantos/go/src/github.com/brunorene/adventofcode2020/day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)`)

	var passwords []Data
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		start, _ := strconv.Atoi(matches[0][1])
		end, _ := strconv.Atoi(matches[0][2])
		passwords = append(passwords, Data{start: start, end: end, letter: matches[0][3], password: matches[0][4]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(passwords)
	part2(passwords)
}

func part1(passwords []Data) {
	validCount := 0
	for _, data := range passwords {
		count := strings.Count(data.password, data.letter)
		if count >= data.start && count <= data.end {
			validCount++
		}
	}
	log.Println(validCount)
}

func part2(passwords []Data) {
	validCount := 0
	for _, data := range passwords {
		if (data.password[data.start-1] == data.letter[0]) != (data.password[data.end-1] == data.letter[0]) {
			validCount++
		}
	}
	log.Println(validCount)
}
