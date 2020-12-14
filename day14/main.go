package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type memory struct {
	address int64
	value   int64
}

type masked struct {
	mask    string
	numbers []memory
}

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	masks := []masked{}

	maskRe := regexp.MustCompile(`[10X]+`)
	memRe := regexp.MustCompile(`[0-9]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch line[0:3] {
		case "mas":
			masks = append(masks, masked{maskRe.FindString(line), []memory{}})
		case "mem":
			matches := memRe.FindAllString(line, -1)
			address, _ := strconv.ParseInt(matches[0], 10, 64)
			value, _ := strconv.ParseInt(matches[1], 10, 64)
			masks[len(masks)-1].numbers = append(masks[len(masks)-1].numbers, memory{address, value})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(masks)
	part2(masks)
}

func part1(masks []masked) {
	memory := make(map[int64]int64)
	for _, mask := range masks {
		for _, value := range mask.numbers {
			binaryRep := strconv.FormatInt(value.value, 2)
			binaryRep = fmt.Sprintf("%036s", binaryRep)
			for index, bit := range mask.mask {
				if bit != 'X' {
					if index == len(mask.mask)-1 {
						binaryRep = binaryRep[:index] + string(bit)
					} else {
						binaryRep = binaryRep[:index] + string(bit) + binaryRep[index+1:]
					}
				}
			}
			memory[value.address], _ = strconv.ParseInt(binaryRep, 2, 64)
		}
	}
	sum := int64(0)
	for _, value := range memory {
		sum += value
	}
	log.Println(sum)
}

func fluctuate(binaryAddress string) []string {
	index := strings.Index(binaryAddress, "X")
	if index == -1 {
		return []string{binaryAddress}
	}
	var binaryAddress0, binaryAddress1 string
	if index == len(binaryAddress)-1 {
		binaryAddress0 = binaryAddress[:index] + "0"
		binaryAddress1 = binaryAddress[:index] + "1"
	} else {
		binaryAddress0 = binaryAddress[:index] + "0" + binaryAddress[index+1:]
		binaryAddress1 = binaryAddress[:index] + "1" + binaryAddress[index+1:]
	}
	list0 := fluctuate(binaryAddress0)
	list1 := fluctuate(binaryAddress1)
	return append(list0, list1...)
}

func allAddresses(mask string, address int64) []int64 {
	binaryRep := strconv.FormatInt(address, 2)
	binaryRep = fmt.Sprintf("%036s", binaryRep)
	for index, bit := range mask {
		if bit != '0' {
			if index == len(mask)-1 {
				binaryRep = binaryRep[:index] + string(bit)
			} else {
				binaryRep = binaryRep[:index] + string(bit) + binaryRep[index+1:]
			}
		}
	}
	result := []int64{}
	for _, memoryAddress := range fluctuate(binaryRep) {
		conv, _ := strconv.ParseInt(memoryAddress, 2, 64)
		result = append(result, conv)
	}
	return result
}

func part2(masks []masked) {
	memory := make(map[int64]int64)
	for _, mask := range masks {
		for _, value := range mask.numbers {
			for _, address := range allAddresses(mask.mask, value.address) {
				memory[address] = value.value
			}
		}
	}
	sum := int64(0)
	for _, value := range memory {
		sum += value
	}
	log.Println(sum)
}
