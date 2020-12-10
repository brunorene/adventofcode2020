package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jolts := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		jolts = append(jolts, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(jolts)
	jolts = append([]int{0}, jolts...)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	part1(jolts)
	part2(jolts)
}

func part1(jolts []int) {
	onesSum := 0
	threesSum := 0
	for i := 0; i < len(jolts)-1; i++ {
		switch jolts[i+1] - jolts[i] {
		case 1:
			onesSum++
		case 3:
			threesSum++
		}
	}
	log.Println(onesSum * threesSum)
}

func key(jolts []int) string {
	key := ""
	for _, j := range jolts {
		v := fmt.Sprint(j)
		key += v + ","
	}
	return key
}

func treeWalk(node int, rest []int, cache *map[string]int64) int64 {
	if len(rest) == 1 {
		return 1
	}
	min := int(math.Min(3, float64(len(rest)-1)))
	count := int64(0)
	for i := 0; i < min; i++ {
		if rest[i]-node < 4 {
			k := key(rest[i:])
			value, exists := (*cache)[k]
			if !exists {
				value = treeWalk(rest[i], rest[i+1:], cache)
				(*cache)[k] = value
			}
			count += value
		}
	}
	return count
}

func part2(jolts []int) {
	cache := make(map[string]int64)
	log.Println(treeWalk(0, jolts[1:], &cache))
}
