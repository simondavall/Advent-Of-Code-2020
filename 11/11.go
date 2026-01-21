package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"aoc"
)

var title string = "## Day 11: Seating System ##"
var url string = "https://adventofcode.com/2020/day/11"
var expectedResult1 int64 = 2483
var expectedResult2 int64 = 2285

type direction struct {
	r int
	c int
}

var (
	_directions []direction = []direction{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	_height     int         = 0
	_width      int         = 0
)

func PartOne(input []byte) int64 {
	var grid = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	_height = len(grid)
	_width = len(grid[0])
	return int64(getOccupiedSeats(grid, getNewSeat))
}

func PartTwo(input []byte) int64 {
	var grid = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	return int64(getOccupiedSeats(grid, getNewSeat2))
}

func getOccupiedSeats(grid []string, getNewSeat func(rune, int, int, []string) rune) int {
	hasChanged := true
	for hasChanged {
		hasChanged = false
		var newGrid []string

		for r, row := range grid {
			var sb strings.Builder
			for c, oldSeat := range row {
				newSeat := getNewSeat(oldSeat, r, c, grid)
				if newSeat != oldSeat {
					hasChanged = true
				}
				sb.WriteRune(newSeat)
			}
			newGrid = append(newGrid, sb.String())
		}
		grid = newGrid
	}

	occupiedSeats := 0
	for _, row := range grid {
		for _, ch := range row {
			if ch == '#' {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}

func getNewSeat(ch rune, r int, c int, grid []string) rune {
	switch ch {

	case '.':
		return '.'

	case 'L':
		for _, dir := range _directions {
			nr := r + dir.r
			nc := c + dir.c
			if isInBounds(nr, nc) && grid[nr][nc] == '#' {
				return 'L'
			}
		}
		return '#'

	case '#':
		occupied := 0
		for _, dir := range _directions {
			nr := r + dir.r
			nc := c + dir.c
			if isInBounds(nr, nc) && grid[nr][nc] == '#' {
				occupied++
			}
		}
		if occupied >= 4 {
			return 'L'
		}
		return '#'

	default:
		panic("Oops found unknown grid item")
	}
}

func getNewSeat2(ch rune, r int, c int, grid []string) rune {
	switch ch {

	case '.':
		return '.'

	case 'L':
		for _, dir := range _directions {
			nr, nc := r, c
			for {
				nr += dir.r
				nc += dir.c
				if !isInBounds(nr, nc) {
					break
				}
				if grid[nr][nc] == '.' {
					continue
				}
				if grid[nr][nc] == '#' {
					return 'L'
				}
				break
			}
		}
		return '#'

	case '#':
		occupied := 0
		for _, dir := range _directions {
			nr, nc := r, c
			for {
				nr += dir.r
				nc += dir.c
				if !isInBounds(nr, nc) {
					break
				}
				if grid[nr][nc] == '.' {
					continue
				}
				if grid[nr][nc] == '#' {
					occupied++
				}
				break
			}
		}
		if occupied >= 5 {
			return 'L'
		}
		return '#'

	default:
		panic("Oops found unknown grid item")
	}
}

func isInBounds(r int, c int) bool {
	return 0 <= r && r < _height && 0 <= c && c < _width
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
