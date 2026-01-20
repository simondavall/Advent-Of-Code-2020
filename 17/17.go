package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

var title string = "# Day 17: Conway Cubes #"
var url string = "https://adventofcode.com/2020/day/17"
var expectedResult1 int64 = 240
var expectedResult2 int64 = 1180


type WXYZ struct {
	w int
	x int
	y int
	z int
}

func PartOne(input []byte) int64 {
	var grid = getGrid(input)
	check := make(map[WXYZ]bool)
	active := make(map[WXYZ]bool)
	for z, plain := range grid {
		for y, row := range plain {
			for x, ch := range row {
				if ch == '#' {
					cur := WXYZ{0, x, y, z}
					active[cur] = true
					check = addNeighboursToCheck(cur, check)
				}
			}
		}
	}

	counter := 0
	for counter < 6 {
		counter++

		new_active := make(map[WXYZ]bool)
		new_check := make(map[WXYZ]bool)
		for cur := range check {
			if cur.w != 0 {
				continue
			}
			c := countActiveNeighbours(cur, active)
			if active[cur] && (c == 2 || c == 3) {
				new_active[cur] = true
				new_check = addNeighboursToCheck(cur, new_check)
			} else if !active[cur] && c == 3 {
				new_active[cur] = true
				new_check = addNeighboursToCheck(cur, new_check)
			}
		}
		active = new_active
		check = new_check
	}
	return int64(len(active))
}

func PartTwo(input []byte) int64 {
	var grid = getGrid(input)
	check := make(map[WXYZ]bool)
	active := make(map[WXYZ]bool)
	for z, plain := range grid {
		for y, row := range plain {
			for x, ch := range row {
				if ch == '#' {
					cur := WXYZ{0, x, y, z}
					active[cur] = true
					check = addNeighboursToCheck(cur, check)
				}
			}
		}
	}

	counter := 0
	for counter < 6 {
		counter++

		new_active := make(map[WXYZ]bool)
		new_check := make(map[WXYZ]bool)
		for cur := range check {
			c := countActiveNeighbours(cur, active)
			if active[cur] && (c == 2 || c == 3) {
				new_active[cur] = true
				new_check = addNeighboursToCheck(cur, new_check)
			} else if !active[cur] && c == 3 {
				new_active[cur] = true
				new_check = addNeighboursToCheck(cur, new_check)
			}
		}
		active = new_active
		check = new_check
	}
	return int64(len(active))
}

func countActiveNeighbours(cur WXYZ, active map[WXYZ]bool) int {
	count := 0
	for dw := -1; dw <= 1; dw++ {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				for dz := -1; dz <= 1; dz++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					neighbour := WXYZ{cur.w + dw, cur.x + dx, cur.y + dy, cur.z + dz}
					if active[neighbour] {
						count++
					}
				}
			}
		}
	}
	return count
}

func addNeighboursToCheck(cur WXYZ, check map[WXYZ]bool) map[WXYZ]bool {
	for dw := -1; dw <= 1; dw++ {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				for dz := -1; dz <= 1; dz++ {
					check[WXYZ{cur.w + dw, cur.x + dx, cur.y + dy, cur.z + dz}] = true
				}
			}
		}
	}
	return check
}

func getGrid(input []byte) [][][]byte {
 	var plainGrid [][]byte
	var width = bytes.IndexByte(input, '\n') + 1

	idx := 0
	for idx < len(input) - (width - 1){
		plainGrid = append(plainGrid, input[idx:idx + width - 1])
		idx += width
	}

	var grid [][][]byte
	grid = append(grid, plainGrid)
	
	return grid
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
