package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ".", "")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1()
	part2()
}

func part1() {
}

func part2() {
}
