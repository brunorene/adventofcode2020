package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	codes := []uint64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		codes = append(codes, uint64(number))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	index, invalid := part1(codes, 25)
	part2(codes, index, invalid)
}

func isValid(number uint64, preamble []uint64) bool {
	for i, val1 := range preamble {
		for j, val2 := range preamble {
			if i != j && val1 != val2 && val1+val2 == number {
				return true
			}
		}
	}
	return false
}

func part1(codes []uint64, preambleSize uint64) (int, uint64) {
	for i := preambleSize; i < uint64(len(codes)); i++ {
		if !isValid(codes[i], codes[i-preambleSize:i]) {
			log.Println(codes[i])
			return int(i), codes[i]
		}
	}
	return 0, 0
}

func sum(codes []uint64) uint64 {
	sum := uint64(0)
	for _, code := range codes {
		sum += code
	}
	return sum
}

func compare(codes []uint64, prop func(uint64, uint64) bool) uint64 {
	comparison := codes[0]
	for _, code := range codes {
		if prop(comparison, code) {
			comparison = code
		}
	}
	return comparison
}

func part2(codes []uint64, index int, invalid uint64) {
	max := len(codes) - index - 1
	if index > max {
		max = index
	}
	for size := 2; size < max; size++ {
		for offset := 0; offset < len(codes)-size+1; offset++ {
			if sum(codes[offset:offset+size]) == invalid {
				min := compare(codes[offset:offset+size], func(u1, u2 uint64) bool { return u1 < u2 })
				max := compare(codes[offset:offset+size], func(u1, u2 uint64) bool { return u1 > u2 })
				log.Println(min + max)
			}
		}
	}
}
