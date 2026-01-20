package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"aoc"
)

var title string = "# Day 9: Encoding Error #"
var url string = "https://adventofcode.com/2020/day/9"
var expectedResult1 int64 = 14144619
var expectedResult2 int64 = 1766397

var invalidNumber int64

func PartOne(input []byte) int64 {
	var data = getEncryptedNumbers(input)

	scope := 25

	for idx, n := range data[scope:] {
		if !isValidCode(n, data[idx:idx+scope]) {
			invalidNumber = n
			break
		}
	}

	return invalidNumber
}

func PartTwo(input []byte) int64 {
	var data = getEncryptedNumbers(input)
	contigious := findContigiousRange(invalidNumber, data)
	var min = invalidNumber
	var max int64 = 0
	for _, val := range contigious {
		min = aoc.Min64(min, val)
		max = aoc.Max64(max, val)
	}
	return min + max
}

func findContigiousRange(n int64, data []int64) []int64 {
	var lower = 0
	var upper = 0
	sum := data[upper]

	for sum != n {
		if sum > n {
			sum -= data[lower]
			lower++
		}
		if sum < n {
			upper++
			sum += data[upper]
		}
	}

	return data[lower : upper+1]
}

func isValidCode(n int64, prev []int64) bool {
	for i := 0; i < len(prev)-1; i++ {
		for j := i + 1; j < len(prev); j++ {
			if prev[i]+prev[j] == n {
				return true
			}
		}
	}
	return false
}

func getEncryptedNumbers(input []byte) []int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	var numbers []int64
	for _, line := range lines{
		n, err := strconv.ParseInt(line, 10, 64)
		check(err)
		numbers = append(numbers, n)
	}
	return numbers
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
