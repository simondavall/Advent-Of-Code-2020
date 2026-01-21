package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var title string = "## Day 3: Toboggan Trajectory ##"
var url string = "https://adventofcode.com/2020/day/3"
var expectedResult1 int64 = 200
var expectedResult2 int64 = 3737923200

var (
	_mapwidth  int
	_mapheight int
)

type Slope struct{ dr, dc int }

func PartOne(input []byte) int64 {

	var lines = ProcessInput(input)
	var tally int64 = 0
	r, c := 0, 0
	dr, dc := 1, 3

	for r < _mapheight {
		if lines[r][c] == '#' {
			tally++
		}

		r = r + dr
		c = (c + dc) % _mapwidth
	}

	return tally
}

func PartTwo(input []byte) int64 {
	var lines = ProcessInput(input)
	var tally int64 = 1

	slopes := []Slope{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	for _, slope := range slopes {
		r, c := 0, 0
		var trees int64 = 0
		for r < _mapheight {
			if lines[r][c] == '#' {
				trees++
			}

			r = r + slope.dr
			c = (c + slope.dc) % _mapwidth
		}
		tally *= trees
	}

	return tally
}

func ProcessInput(input []byte) []string {
	var lines []string
	content := string(input)

	for str := range strings.SplitSeq(content, "\n") {
		if str != "" {
			lines = append(lines, str)
		}
	}

	_mapheight = len(lines)
	_mapwidth = len(lines[0])

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
		os.Exit(1)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

