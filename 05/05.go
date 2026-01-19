package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"aoc"
)

var title string = "# Day 5: Binary Boarding #"
var url string = "https://adventofcode.com/2020/day/5"
var expectedResult1 int64 = 987
var expectedResult2 int64 = 603

var maxSeatId int

func PartOne(input []byte) int64 {
	var passes = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	for _, pass := range passes {
		if pass[0] != 'B' {
			continue
		}
		seatId := calcSeatId(pass)
		maxSeatId = aoc.Max(maxSeatId, seatId)
	}

	return int64(maxSeatId)
}

func PartTwo(input []byte) int64 {
	var passes = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	var mySeatId = 0
	occupied := make([]bool, maxSeatId+1)

	for _, pass := range passes {
		current := calcSeatId(pass)
		occupied[current] = true
	}

	for i := maxSeatId; i >= 0; i-- {
		if !occupied[i] {
			mySeatId = i
			break
		}
	}

	return int64(mySeatId)
}

func calcSeatId(pass string) int {
	var rowCode = pass[:len(pass)-3]
	var seatCode  = pass[len(pass)-3:]
	var row int
	for _, x := range rowCode {
		row <<= 1
		if x == 'B' {
			row++
		}
	}
	var seat int
	for _, y := range seatCode {
		seat <<= 1
		if y == 'R' {
			seat++
		}
	}
	return row*8 + seat
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

