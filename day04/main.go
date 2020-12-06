package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("/home/brsantos/go/src/github.com/brunorene/adventofcode2020/day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`([a-z]+):([^ ]+)`)

	passports := []map[string]string{}
	scanner := bufio.NewScanner(file)
	newPassport := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.Trim(line, " ")) == 0 {
			passports = append(passports, newPassport)
			newPassport = make(map[string]string)
			continue
		}
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			newPassport[match[1]] = match[2]
		}
	}
	passports = append(passports, newPassport)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(passports)
	part2(passports)
}

func isValid(passport map[string]string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, key := range requiredFields {
		_, exists := passport[key]
		if !exists {
			return false
		}
	}
	return true
}

var (
	regexps map[string]*regexp.Regexp
	once    sync.Once
)

func regexpsInstance() map[string]*regexp.Regexp {

	once.Do(func() {
		regexps = make(map[string]*regexp.Regexp)
		regexps["byr"] = regexp.MustCompile(`^19[2-9][0-9]|200[0-2]$`)
		regexps["iyr"] = regexp.MustCompile(`^201[0-9]|2020$`)
		regexps["eyr"] = regexp.MustCompile(`^202[0-9]|2030$`)
		regexps["hgt"] = regexp.MustCompile(`^(1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in$`)
		regexps["hcl"] = regexp.MustCompile(`^#[0-9a-f]{6}$`)
		regexps["ecl"] = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		regexps["pid"] = regexp.MustCompile(`^[0-9]{9}$`)
		regexps["cid"] = regexp.MustCompile(`.*`)
	})

	return regexps
}

func hasValidData(passport map[string]string) bool {
	for key, value := range passport {
		match := regexpsInstance()[key].FindStringIndex(value)
		if match == nil {
			return false
		}
	}
	return true
}

func part1(passports []map[string]string) {
	count := 0
	for _, passport := range passports {
		if isValid(passport) {
			count++
		}
	}
	log.Println(count)
}

func part2(passports []map[string]string) {
	count := 0
	for _, passport := range passports {
		if isValid(passport) && hasValidData(passport) {
			count++
		}
	}
	log.Println(count)
}
