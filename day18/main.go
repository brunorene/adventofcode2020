package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	path, _ := os.Getwd()
	file, err := os.Open(path + "/day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tokens := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		line = strings.ReplaceAll(line, "  ", " ")
		lineTokens := strings.Split(line, " ")
		tokens = append(tokens, lineTokens)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(tokens)
	part2(tokens)
}

func part1(tokens [][]string) {
	sum := int64(0)
	for _, line := range tokens {
		stack := []int64{0}
		stackOp := []string{""}
		for _, token := range line {
			switch token {
			case "(":
				stack = append([]int64{0}, stack...)
				stackOp = append([]string{""}, stackOp...)
			case "+":
				stackOp = append([]string{"S"}, stackOp[1:]...)
			case "*":
				stackOp = append([]string{"M"}, stackOp[1:]...)
			default:
				var num int64
				if token == ")" {
					num = stack[0]
					stack = stack[1:]
					stackOp = stackOp[1:]
				} else {
					num, _ = strconv.ParseInt(token, 10, 64)
				}
				op := stackOp[0]
				switch op {
				case "S":
					num += stack[0]
				case "M":
					num *= stack[0]
				}
				stack = append([]int64{num}, stack[1:]...)
				stackOp = append([]string{""}, stackOp[1:]...)
			}
		}
		sum += stack[0]
	}
	log.Println(sum)
}

func wrap(left []string, right []string) ([]string, []string) {
	re := regexp.MustCompile(`[0-9]+`)
	var index int
	if re.MatchString(left[len(left)-1]) {
		index = len(left) - 2
	} else {
		level := 1
		index = len(left) - 2
		for level > 0 && index >= 0 {
			switch left[index] {
			case ")":
				level++
			case "(":
				level--
			}
			index--
		}
	}
	rightOfLeft := left[index+1:]
	left = append(left[:index+1], "(")
	left = append(left, rightOfLeft...)

	if re.MatchString(right[0]) {
		index = 0
	} else {
		level := 1
		index = 1
		for level > 0 && index < len(right)-1 {
			switch right[index] {
			case "(":
				level++
			case ")":
				level--
			}
			index++
		}
	}
	leftOfRight := right[:index+1]
	right = append([]string{")"}, right[index+1:]...)
	right = append(leftOfRight, right...)

	return left, right
}

func prioritise(line []string) []string {
	for index := 0; index < len(line); index++ {
		if line[index] == "+" {
			lenLeft := len(line[:index])
			left, right := wrap(line[:index], line[index+1:])
			line = append(left, "+")
			line = append(line, right...)
			if lenLeft < len(left) {
				index++
			}
		}
	}
	log.Println(line)
	return line
}

func part2(tokens [][]string) {
	for index, line := range tokens {
		tokens[index] = prioritise(line)
	}
	part1(tokens)
}
