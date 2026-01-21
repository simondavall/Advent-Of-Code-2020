package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var title string = "## Day 1: Report Repair ##"
var url string = "https://adventofcode.com/2020/day/1"
var expectedResult1 int64 = 605364
var expectedResult2 int64 = 128397680

func PartOne(input []byte) int64 {
	expenses := ProcessInput(input)
	for i := 0; i < len(expenses)-1; i++ {
		for j := i + 1; j < len(expenses); j++ {
			if expenses[i]+expenses[j] == 2020 {
				return int64(expenses[i] * expenses[j])
			}
		}
	}

	return 0
}

func PartTwo(input []byte) int64 {
	expenses := ProcessInput(input)
	for i := 0; i < len(expenses)-2; i++ {
		for j := i + 1; j < len(expenses)-1; j++ {
			for k := j + 1; k < len(expenses); k++ {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					return int64(expenses[i] * expenses[j] * expenses[k])
				}
			}
		}
	}

	return 0
}

func ProcessInput(input []byte) []int {
	var ints []int
	content := string(input)

	for str := range strings.SplitSeq(content, "\n") {
		if str != "" {
			var n, err = strconv.Atoi(str)
			check(err)
			ints = append(ints, n)
		}
	}

	return ints
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
		os.Exit(1);
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
