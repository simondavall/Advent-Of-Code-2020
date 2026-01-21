package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"aoc"
)

var title string = "## Day 15: Rambunctious Recitation ##"
var url string = "https://adventofcode.com/2020/day/15"
var expectedResult1 int64 = 870
var expectedResult2 int64 = 9136

func PartOne(input []byte) int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input),  "\n"))
	line := lines[0]
	numMap := make(map[int]int)
	var diff = 0
	initial := strings.Split(line, ",")
	for idx, n := range initial {
		n, _ := strconv.Atoi(n)
		diff = addToNumMap(n, idx, numMap)
	}

	counter := len(initial)
	for counter < 2019 {
		diff = addToNumMap(diff, counter, numMap)
		counter++
	}

	return int64(diff)
}

func PartTwo(input []byte) int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input),  "\n"))
	line := lines[0]
	numMap := make(map[int]int)
	var diff = 0
	initial := strings.Split(line, ",")
	for idx, n := range initial {
		n, _ := strconv.Atoi(n)
		diff = addToNumMap(n, idx, numMap)
	}

	counter := len(initial)
	for counter < 30000000-1 {
		diff = addToNumMap(diff, counter, numMap)
		counter++
	}

	return int64(diff)
}

func addToNumMap(n int, idx int, numMap map[int]int) int {
	diff := 0
	if cur, ok := numMap[n]; ok {
		diff = idx - cur
	}
	numMap[n] = idx
	return diff
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
