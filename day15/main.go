package main

import (
	"log"
	"time"
)

func main() {

	start := []int{9, 12, 1, 4, 17, 0, 18}
	// start := []int{0, 3, 6}

	part1(start, 2020)
	startTime := time.Now()
	part1(start, 30000000)
	duration := time.Since(startTime)
	log.Println(duration)
}

func part1(start []int, end int) {

	memory := make(map[int][]int)

	var last int
	for now := 0; now < end; now++ {
		if now < len(start) {
			memory[start[now]] = []int{now}
			last = start[now]
		} else {
			// new
			if len(memory[last]) == 1 {
				last = 0
				_, exists := memory[0]
				if !exists {
					memory[0] = []int{}
				}
				memory[0] = append(memory[0], now)
			} else {
				last = memory[last][len(memory[last])-1] - memory[last][len(memory[last])-2]
				memory[last] = append(memory[last], now)
			}
		}
	}
	log.Println(last)
}
