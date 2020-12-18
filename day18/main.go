package main

import (
	"bufio"
	"log"
	"os"
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
	index := len(left) - 1
	level := 0
	for {
		switch left[index] {
		case ")":
			level++
		case "(":
			level--
		}
		if level == 0 {
			break
		}
		index--
	}
	rightOfLeft := make([]string, len(left[index:]))
	copy(rightOfLeft, left[index:])
	left = append(left[:index], "(")
	left = append(left, rightOfLeft...)

	index = 0
	level = 0
	for {
		switch right[index] {
		case "(":
			level++
		case ")":
			level--
		}
		if level == 0 {
			break
		}
		index++
	}
	leftOfRight := make([]string, len(right[:index+1]))
	copy(leftOfRight, right[:index+1])
	right = append([]string{")"}, right[index+1:]...)
	right = append(leftOfRight, right...)

	return left, right
}

func prioritise(oldLine []string) []string {
	line := make([]string, len(oldLine))
	copy(line, oldLine)
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
	return line
}

func part2(tokens [][]string) {
	for index, line := range tokens {
		tokens[index] = prioritise(line)
	}
	part1(tokens)
}
