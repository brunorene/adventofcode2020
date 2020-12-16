package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type field struct {
	name   string
	values map[int]bool
}

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fieldRe := regexp.MustCompile(`([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)`)

	fields := []field{}
	tickets := [][]int{}

	state := "fields"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if state == "fields" && strings.TrimSpace(line) == "" {
			state = "your ticket"
			continue
		}
		if state == "your ticket" && strings.TrimSpace(line) == "" {
			state = "other tickets"
			continue
		}
		switch state {
		case "fields":
			matches := fieldRe.FindAllStringSubmatch(line, -1)
			val1, _ := strconv.Atoi(matches[0][2])
			val2, _ := strconv.Atoi(matches[0][3])
			val3, _ := strconv.Atoi(matches[0][4])
			val4, _ := strconv.Atoi(matches[0][5])
			values := make(map[int]bool)
			for i := val1; i <= val2; i++ {
				values[i] = true
			}
			for i := val3; i <= val4; i++ {
				values[i] = true
			}
			fields = append(fields, field{matches[0][1], values})
		case "your ticket":
			if line == "your ticket:" {
				continue
			}
			fallthrough
		case "other tickets":
			if line == "nearby tickets:" {
				continue
			}
			parts := strings.Split(line, ",")
			values := []int{}
			for _, part := range parts {
				val, _ := strconv.Atoi(part)
				values = append(values, val)
			}
			tickets = append(tickets, values)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	validTickets := part1(fields, tickets[1:])
	part2(fields, validTickets)
}

func part1(fields []field, tickets [][]int) [][]int {
	sum := 0
	validTickets := [][]int{}
	for _, ticket := range tickets {
		allValid := true
		for _, value := range ticket {
			valid := false
			for _, field := range fields {
				_, exists := field.values[value]
				if exists {
					valid = true
				}
			}
			if !valid {
				sum += value
				allValid = false
			}
		}
		if allValid {
			validTickets = append(validTickets, ticket)
		}
	}
	log.Println(sum)
	return validTickets
}

func singleFieldPerRow(possibilities []map[string]int) bool {
	for _, m := range possibilities {
		if len(m) != 1 {
			return false
		}
	}
	return true
}

func part2(fields []field, tickets [][]int) {
	possibilities := []map[string]int{}
	for range tickets[0] {
		possibilities = append(possibilities, make(map[string]int))
	}
	for valueRow := range possibilities {
		for _, field := range fields {
			for _, ticket := range tickets {
				_, exists := field.values[ticket[valueRow]]
				if exists {
					possibilities[valueRow][field.name]++
				}
			}
		}
	}
	// remove
	log.Println(len(tickets))

	for rowIndex := range possibilities {
		for key, value := range possibilities[rowIndex] {
			if value < len(tickets) {
				delete(possibilities[rowIndex], key)
			}
		}
	}

	for !singleFieldPerRow(possibilities) {
		for rowIndex := range possibilities {
			if len(possibilities[rowIndex]) == 1 {
				for otherRow := range possibilities {
					for key := range possibilities[rowIndex] {
						if otherRow != rowIndex {
							delete(possibilities[otherRow], key)
						}
					}
				}
			}
		}
	}

	log.Println(tickets[0])

	mul := 1
	for rowIndex := range possibilities {
		for key := range possibilities[rowIndex] {
			if strings.Contains(key, "departure") {
				log.Println(key, rowIndex, tickets[0][rowIndex])
				mul *= tickets[0][rowIndex]
			}
		}
	}

	log.Println(mul)

}
