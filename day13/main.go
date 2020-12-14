package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	departure, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, ",")
	buses := []int64{}
	for _, part := range parts {
		if part == "x" {
			buses = append(buses, 1)
		} else {
			bus, _ := strconv.ParseInt(part, 10, 64)
			buses = append(buses, bus)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(int64(departure), buses)
	// part2(buses, diffs)
	log.Println(part2V2(line))

}

func part1(departure int64, buses []int64) {
	min := int64(math.MaxInt64)
	var minBus int64
	for _, bus := range buses {
		waitingTime := departure - (departure % bus) + bus - departure
		if waitingTime < min && waitingTime > 1 {
			min = waitingTime
			minBus = bus
		}
	}
	log.Println(minBus * min)
}

func allMatches(buses []int64) []int64 {
	result := []int64{}
	result = append(result, buses[0])
	for i := int64(1); i < int64(len(buses)); i++ {
		if buses[0] == buses[i]-i {
			result = append(result, buses[i])
		}
	}
	return result
}

func part2(buses []int64) {
	matches := allMatches(buses)
	for len(matches) < len(buses) {
		mult := int64(1)
		for _, match := range matches {
			mult *= match
		}
	}

	log.Println(buses[0])
}

func part2V2(schedule string) int {
	var buses []int
	for _, id := range strings.Split(schedule, ",") {
		if id == "x" {
			id = "1"
		}
		number, _ := strconv.Atoi(id)
		buses = append(buses, number)
	}

	timestamp := 1

	for {
		timeToSkipIfNoMatch := 1
		valid := true

		for offset := 0; offset < len(buses); offset++ {
			if (timestamp+offset)%buses[offset] != 0 {
				valid = false
				break
			}
			timeToSkipIfNoMatch *= buses[offset]
		}

		if valid {
			return timestamp
		}

		timestamp += timeToSkipIfNoMatch
	}
}
