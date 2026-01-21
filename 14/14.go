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

var title string = "## Day 14: Docking Data ##"
var url string = "https://adventofcode.com/2020/day/14"
var expectedResult1 int64 = 13105044880745
var expectedResult2 int64 = 3505392154485

func PartOne(input []byte) int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	var tally int64 = 0
	mask := ""
	memory := make(map[int]int64)
	binaryNumber := make([]byte, 36)

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			msk := strings.Split(line, " = ")
			mask = msk[1]
			continue
		}

		memAddr, value, err := getMemory(line)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		setBinaryNumber(value, binaryNumber)

		var memStr strings.Builder
		for i := 0; i < len(mask); i++ {
			if mask[i] == 'X' {
				memStr.WriteByte(binaryNumber[i])
			} else {
				memStr.WriteByte(mask[i])
			}
		}

		decimalNumber, err := strconv.ParseInt(memStr.String(), 2, 64)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		memory[memAddr] = decimalNumber
	}

	for _, memoryValue := range memory {
		tally += memoryValue
	}

	return tally
}

func PartTwo(input []byte) int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	var tally int64 = 0
	mask := ""
	memory := make(map[int64]int)
	binaryNumber := make([]byte, 36)

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			msk := strings.Split(line, " = ")
			mask = msk[1]
			continue
		}

		memAddr, value, err := getMemory(line)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		setBinaryNumber(memAddr, binaryNumber)

		first := make([]byte, 36)
		memAddresses := [][]byte{first}
		for i := 0; i < len(mask); i++ {
			switch mask[i] {
			case 'X':
				var newEntries [][]byte
				for _, current := range memAddresses {
					newEntry := slices.Clone(current)
					newEntry[i] = '1'
					current[i] = '0'
					newEntries = append(newEntries, newEntry)
				}
				memAddresses = append(memAddresses, newEntries...)
			case '1':
				for _, address := range memAddresses {
					address[i] = mask[i]
				}
			default:
				for _, address := range memAddresses {
					address[i] = binaryNumber[i]
				}
			}
		}

		for _, address := range memAddresses {
			memAddr, err := strconv.ParseInt(string(address[:]), 2, 64)
			if err != nil {
				fmt.Println(err)
				return 0
			}
			memory[memAddr] = value
		}
	}

	for _, memValue := range memory {
		tally += int64(memValue)
	}

	return tally
}

func getMemory(line string) (int, int, error) {
	memory := strings.Split(line, "] = ")
	value, err := strconv.Atoi(memory[1])
	if err != nil {
		return 0, 0, err
	}
	maddr := strings.Split(memory[0], "[")
	memAddr, err := strconv.Atoi(maddr[1])
	if err != nil {
		return 0, 0, err
	}

	return memAddr, value, nil
}

func setBinaryNumber(value int, binaryNumber []byte) {
	for i := range binaryNumber {
		binaryNumber[i] = '0'
	}

	valueAsBinary := strconv.FormatInt(int64(value), 2)
	offset := 36 - len(valueAsBinary)
	for i := 0; i < len(valueAsBinary); i++ {
		binaryNumber[i+offset] = valueAsBinary[i]
	}
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
