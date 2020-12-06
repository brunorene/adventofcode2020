package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("/home/brsantos/go/src/github.com/brunorene/adventofcode2020/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(numbers)
	part2(numbers)
}

func part1(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j && numbers[i]+numbers[j] == 2020 {
				log.Println(numbers[i] * numbers[j])
				return
			}
		}
	}
}

func part2(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			for k := 0; k < len(numbers); k++ {
				if i != j && j != k && numbers[i]+numbers[j]+numbers[k] == 2020 {
					log.Println(numbers[i] * numbers[j] * numbers[k])
					return
				}
			}
		}
	}

}
