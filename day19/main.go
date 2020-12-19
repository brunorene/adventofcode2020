package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type rule struct {
	id   string
	next [][]string
	char string
}

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day19/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	isRule := true
	rules := make(map[string]rule)
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRule = false
			continue
		}
		if isRule {
			parts := strings.Split(line, ":")
			id := parts[0]
			parts = strings.Split(parts[1], "|")
			leaf := ""
			next := [][]string{}
			if strings.IndexAny(parts[0], "ab") >= 0 {
				leaf = string([]byte{parts[0][strings.IndexAny(parts[0], "ab")]})
			} else {
				for _, nextRule := range parts {
					nextRules := strings.Split(strings.Trim(nextRule, " "), " ")
					next = append(next, nextRules)
				}
			}
			rules[id] = rule{id, next, string(leaf)}
		} else {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(rules, lines)
	part2(rules, lines)
}

func generateCandidates(rule rule, rules map[string]rule) map[string]bool {
	if len(rule.char) > 0 {
		c := make(map[string]bool)
		c[rule.char] = true
		return c
	}
	candidates := make(map[string]bool)
	for _, r := range rule.next {
		parts := []map[string]bool{}
		for _, next := range r {
			parts = append(parts, generateCandidates(rules[next], rules))
		}
		if len(parts) == 1 {
			for key := range parts[0] {
				candidates[key] = true
			}
		} else {
			for left := range parts[0] {
				for right := range parts[1] {
					candidates[left+right] = true
				}
			}
		}
	}
	return candidates
}

func part1(rules map[string]rule, lines []string) {
	candidates := generateCandidates(rules["0"], rules)
	count := 0
	for _, line := range lines {
		if candidates[line] {
			count++
		}
	}
	log.Println(count)
}

func part2(rules map[string]rule, lines []string) {
	candidates0 := generateCandidates(rules["0"], rules)
	candidates42 := generateCandidates(rules["42"], rules)
	candidates31 := generateCandidates(rules["31"], rules)
	regex := regexp.MustCompile(`1+1+0+`)
	results := make(map[string]bool)
	for _, line := range lines {
		if candidates0[line] {
			results[line] = true
		}
	}
	for _, line := range lines {
		result := line
		if len(line)%8 != 0 {
			continue
		}
		index := 0
		for {
			found := false
			for key42 := range candidates42 {
				if strings.Index(result, key42) == index {
					result = strings.Replace(result, key42, "1", 1)
					index++
					found = true
					break
				}
			}
			if !found {
				break
			}
		}
		for {
			found := false
			for key31 := range candidates31 {
				if strings.Index(result, key31) == index {
					result = strings.Replace(result, key31, "0", 1)
					index++
					found = true
					break
				}
			}
			if !found {
				break
			}
		}
		if !strings.Contains(result, "a") &&
			!strings.Contains(result, "b") &&
			regex.MatchString(result) &&
			strings.Count(result, "0") < strings.Count(result, "1") {
			results[line] = true
		}
	}
	log.Println(len(results))
}
