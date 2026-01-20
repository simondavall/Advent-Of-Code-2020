package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"aoc"
)

var title string = "# Day 12: Rain Risk #"
var url string = "https://adventofcode.com/2020/day/12"
var expectedResult1 int64 = 759
var expectedResult2 int64 = 45763

var directions []byte = []byte{'N', 'E', 'S', 'W'}

func PartOne(input []byte) int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	hor := 0  // east/west distance
	vert := 0 // north/south distance
	dir := 1

	for _, line := range lines {
		ins := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch ins {
		case 'N':
			fallthrough
		case 'S':
			fallthrough
		case 'E':
			fallthrough
		case 'W':
			dh, dv := move(ins, value)
			hor += dh
			vert += dv
		case 'F':
			dh, dv := move(directions[dir], value)
			hor += dh
			vert += dv
		case 'L':
			fallthrough
		case 'R':
			dir = turn(ins, value, dir)
		}
	}

	return int64(aoc.Abs(hor) + aoc.Abs(vert))
}

type waypoint struct {
	hor  int
	vert int
}

func PartTwo(input []byte) int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	hor := 0  // east/west distance
	vert := 0 // north/south distance
	var w = waypoint{10, 1}

	for _, line := range lines {
		ins := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch ins {
		case 'N':
			fallthrough
		case 'S':
			fallthrough
		case 'E':
			fallthrough
		case 'W':
			dh, dv := move(ins, value)
			w.hor += dh
			w.vert += dv
		case 'F':
			hor += w.hor * value
			vert += w.vert * value
		case 'L':
			fallthrough
		case 'R':
			rotate(ins, value, &w)
		}
	}

	return int64(aoc.Abs(hor) + aoc.Abs(vert))
}

func move(ins byte, value int) (int, int) {
	hor, vert := 0, 0
	switch ins {
	case 'N':
		vert += value
	case 'S':
		vert -= value
	case 'E':
		hor += value
	case 'W':
		hor -= value
	default:
		panic("Invalid instruction passed to move()")
	}
	return hor, vert
}

func turn(ins byte, value int, dir int) int {
	dirLen := len(directions)
	var t = value / 90
	if ins == 'R' {
		dir = (dir + t) % dirLen
	} else {
		dir = (dirLen + dir - t) % dirLen
	}
	return dir
}

func rotate(ins byte, value int, w *waypoint) {
	var t = value / 90
	for range t {
		if ins == 'R' {
			w.hor, w.vert = w.vert, w.hor*-1
		} else {
			w.hor, w.vert = w.vert*-1, w.hor
		}
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
