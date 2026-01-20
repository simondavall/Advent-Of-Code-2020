package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"aoc"
)

var title string = "# Day 8: Handheld Halting #"
var url string = "https://adventofcode.com/2020/day/8"
var expectedResult1 int64 = 1749
var expectedResult2 int64 = 515

type Instruction struct {
	name  string
	value int
}

func PartOne(input []byte) int64 {
	var instructions = getInstructions(input)
	ip := 0
	var acc int64 = 0
	var seen = []int{}

	for ip >= 0 && ip < len(instructions) && !slices.Contains(seen, ip) {
		seen = append(seen, ip)
		ins := instructions[ip]
		ip, acc = ProcessInstruction(ins, ip, acc)
	}

	return acc
}

func PartTwo(input []byte) int64 {
	var instructions = getInstructions(input)
	var acc int64 = 0
	for idx, ins := range instructions {
		orig := ""
		if ins.name == "acc" {
			continue
		}
		if ins.name == "jmp" {
			orig = "jmp"
			instructions[idx].name = "nop"
		}
		if ins.name == "nop" {
			orig = "nop"
			instructions[idx].name = "jmp"
		}

		ip := 0
		acc = 0
		var seen = []int{}
		success := false

		for {
			if ip < 0 {
				panic("Cannot have negative instruction pointer")
			}
			if ip >= len(instructions) {
				success = true
				break
			}
			if slices.Contains(seen, ip) {
				break
			}
			seen = append(seen, ip)
			ins := instructions[ip]
			ip, acc = ProcessInstruction(ins, ip, acc)
		}

		instructions[idx].name = orig

		if success {
			break
		}
	}

	return acc
}

func ProcessInstruction(ins Instruction, ip int, acc int64) (int, int64) {
	switch ins.name {
	case "acc":
		acc += int64(ins.value)
		ip++
	case "jmp":
		ip += ins.value
	case "nop":
		ip++
	default:
		panic("Unknown instruction")
	}
	return ip, acc
}

func getInstructions(input []byte) []Instruction {
	var instructions []Instruction

	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	for _, line := range lines {
		s := strings.Split(line, " ")
		value, _ := strconv.Atoi(s[1])
		instructions = append(instructions, Instruction{s[0], value})
	}
	return instructions
}

func main() {
	var resultPartOne int64 = -1
	var resultPartTwo int64 = -1

	fmt.Printf("\n%s", title)
	fmt.Printf("\n%s\n", url)
	for i := 1; i < len(os.Args); i++ {
		filePath := os.Args[i]
		fmt.Printf("\nFile: %s\n", filePath)

		input, err := os.ReadFile(filePath)
		check(err)

		startPart1 := time.Now()
		resultPartOne = PartOne(input)
		fmt.Printf("Part 1 result: %d in %s\n", resultPartOne, time.Since(startPart1))

		startPart2 := time.Now()
		resultPartTwo = PartTwo(input)
		fmt.Printf("Part 2 result: %d in %s\n", resultPartTwo, time.Since(startPart2))
	}
	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
		os.Exit(1)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

