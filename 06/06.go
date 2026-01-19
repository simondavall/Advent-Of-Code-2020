package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"aoc"
)

var title string = "# Day 6: Custom Customs #"
var url string = "https://adventofcode.com/2020/day/6"
var expectedResult1 int64 = 6585
var expectedResult2 int64 = 3276


func PartOne(input []byte) int64 {
	var blocks = aoc.RemoveEmpties(strings.Split(string(input), "\n\n"))
	var tally = 0

	for _, block := range blocks {
		var seen []rune
		lines := strings.SplitSeq(block, "\n")
		for line := range lines {
			for _, ch := range line {
				if !slices.Contains(seen, ch) {
					seen = append(seen, ch)
				}
			}
		}
		tally += len(seen)
	}

	return int64(tally)
}

var seen = make(map[rune]int)

func PartTwo(input []byte) int64 {
	var blocks = aoc.RemoveEmpties(strings.Split(string(input), "\n\n"))
	var tally = 0

	for _, block := range blocks {
		for key := range seen {
			delete(seen, key)
		}
		people := 0
		lines := strings.SplitSeq(block, "\n")
		for line := range lines {
			people++
			for _, ch := range line {
				seen[ch]++
			}
		}
		for _, item := range seen {
			if item == people {
				tally++
			}
		}
	}

	return int64(tally)
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
