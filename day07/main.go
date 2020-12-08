package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type bagContent struct {
	ascendants map[string]bool
	name       string
	bagsInside map[string]*child
}

type child struct {
	bag   *bagContent
	count int
}

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bags := make(map[string]*bagContent)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ".", "")
		line = strings.ReplaceAll(line, " bags", "")
		line = strings.ReplaceAll(line, " bag", "")
		parts := strings.Split(line, " contain ")
		bag, exists := bags[parts[0]]
		if !exists {
			bag = &bagContent{ascendants: make(map[string]bool), name: parts[0], bagsInside: make(map[string]*child)}
			bags[bag.name] = bag
		}
		parts = strings.Split(parts[1], ", ")
		for _, part := range parts {
			countAndName := strings.SplitN(part, " ", 2)
			count, _ := strconv.Atoi(countAndName[0])
			if count == 0 {
				continue
			}
			innerBag, exists := bags[countAndName[1]]
			if !exists {
				innerBag = &bagContent{ascendants: make(map[string]bool), name: countAndName[1], bagsInside: make(map[string]*child)}
				bags[innerBag.name] = innerBag
			}
			bag.bagsInside[innerBag.name] = &child{bag: innerBag, count: count}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(bags)
	part2(bags)
}

func countAscendants(bag *bagContent) {
	for _, descendant := range *&bag.bagsInside {
		descendant.bag.ascendants[bag.name] = true
		for key := range bag.ascendants {
			descendant.bag.ascendants[key] = true
		}
		countAscendants(descendant.bag)
	}
}

func part1(bags map[string]*bagContent) {
	for _, bag := range bags {
		countAscendants(bag)
	}
	log.Println(len(bags["shiny gold"].ascendants))
}

func totalInnerBags(bag bagContent) int {
	sum := 0
	for _, child := range bag.bagsInside {
		sum += child.count
		sum += child.count * totalInnerBags(*child.bag)
	}
	return sum
}

func part2(bags map[string]*bagContent) {
	log.Println(totalInnerBags(*bags["shiny gold"]))
}
