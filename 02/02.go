package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var title string = "## Day 2: Password Philosophy ##"
var url string = "https://adventofcode.com/2020/day/2"
var expectedResult1 int64 = 620
var expectedResult2 int64 = 727

func PartOne(input []byte) int64 {
	var lines = processInput(input);
	var tally int64 = 0

	for _, line := range lines {
		min, max, ch, pwd := processLine(line)
		charCount := strings.Count(pwd, ch)
		if charCount >= min && charCount <= max {
			tally++
		}
	}

	return tally
}

func PartTwo(input []byte) int64 {
	var lines = processInput(input);
	var tally int64 = 0

	for _, line := range lines {
		min, max, ch, pwd := processLine(line)

		if (pwd[min-1] == ch[0] && pwd[max-1] == ch[0]) || (pwd[min-1] != ch[0] && pwd[max-1] != ch[0]) {
			continue
		}

		tally++
	}

	return tally
}

func processLine(line string) (int, int, string, string) {
	splits := strings.Split(line, " ")
	bounds := strings.Split(splits[0], "-")

	min, err := strconv.Atoi(bounds[0])
	check(err)

	max, err := strconv.Atoi(bounds[1])
	check(err)

	ch := splits[1][:len(splits[1])-1]
	pwd := splits[2]

	return min, max, ch, pwd
}

func processInput(input []byte) []string {
	var lines []string
	content := string(input)

	for str := range strings.SplitSeq(content, "\n") {
		if str != "" {
			lines = append(lines, str)
		}
	}

	return lines
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
