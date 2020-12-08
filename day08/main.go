package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	operation string
	value     int
}

type cpu struct {
	pointer      int
	acumulator   int
	instructions []instruction
}

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	regex := regexp.MustCompile(`^([a-z]{3}) ([+-][0-9]+)$`)
	cpu := cpu{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := regex.FindAllStringSubmatch(scanner.Text(), -1)
		val, _ := strconv.Atoi(parts[0][2])
		cpu.instructions = append(cpu.instructions, instruction{operation: parts[0][1], value: val})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(cpu)
	part2(cpu)
}

func execute(cpu cpu) cpu {
	switch cpu.instructions[cpu.pointer].operation {
	case "jmp":
		cpu.pointer += cpu.instructions[cpu.pointer].value
	case "acc":
		cpu.acumulator += cpu.instructions[cpu.pointer].value
		cpu.pointer++
	case "nop":
		cpu.pointer++
	}
	return cpu
}

func part1(cpu cpu) {
	visitedAddresses := make(map[int]bool)
	for !visitedAddresses[cpu.pointer] {
		visitedAddresses[cpu.pointer] = true
		cpu = execute(cpu)
	}
	log.Println(cpu.acumulator)
}

func flipInstruction(oldCPU cpu, index int) cpu {
	localCPU := cpu{oldCPU.pointer, oldCPU.acumulator, []instruction{}}
	for _, inst := range oldCPU.instructions {
		localCPU.instructions = append(localCPU.instructions, inst)
	}
	switch localCPU.instructions[index].operation {
	case "jmp":
		localCPU.instructions[index].operation = "nop"
	case "nop":
		localCPU.instructions[index].operation = "jmp"
	}
	return localCPU
}

func runProgram(cpu cpu) (cpu, bool) {
	visitedAddresses := make(map[int]bool)
	localCPU := cpu
	for !visitedAddresses[localCPU.pointer] && localCPU.pointer < len(localCPU.instructions) {
		visitedAddresses[localCPU.pointer] = true
		localCPU = execute(localCPU)
	}
	return localCPU, localCPU.pointer >= len(localCPU.instructions)
}

func part2(cpu cpu) {
	for index, instruction := range cpu.instructions {
		localCPU := flipInstruction(cpu, index)
		if instruction.operation != "acc" {
			localCPU, finished := runProgram(localCPU)
			if finished {
				log.Println(localCPU.acumulator)
				return
			}
		}
	}
}
