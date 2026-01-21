package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"aoc"
)

var title string = "## Day 20: Jurassic Jigsaw ##"
var url string = "https://adventofcode.com/2020/day/20"
var expectedResult1 int64 = 111936085519519
var expectedResult2 int64 = 1792

type Edge int

const (
	Top Edge = iota
	Right
	Bottom
	Left
)

type Tile struct {
	id        int
	rows      []string
	edges     []string
	edge_vals []int
}

var monster = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

func PartOne(input []byte) int64 {
	var tiles = getTiles(input)
	var tally int64 = 1

	for _, tile := range tiles {
		if neighboutCount(tile, tiles) == 2 {
			tally *= int64(tile.id)
		}
	}

	return tally
}

func PartTwo(input []byte) int64 {
	var tiles = getTiles(input)
	jigsaw := constructJigsaw(tiles)
	picture := buildPicture(jigsaw, tiles)

	monster_count := getMonsterCount(picture)

	var tally = int64(countHashes(picture) - monster_count*countHashes(monster))

	return tally
}

func getMonsterCount(picture []string) int {
	monster_count := 0
	for range 4 {
		monster_count = countMonsters(picture)
		if monster_count > 0 {
			return monster_count
		}
		picture = rotate_lines(picture)
	}

	picture = flip_lines(picture)

	for range 4 {
		monster_count = countMonsters(picture)
		if monster_count > 0 {
			return monster_count
		}
		picture = rotate_lines(picture)
	}
	panic("Unable to find the monsters!!")
}

func buildPicture(puzzle [][]Tile, tiles []Tile) []string {
	picture := []string{}
	tile_size := len(tiles[0].rows)

	for _, row := range puzzle {
		for i := range tile_size - 2 {
			picture_row := []string{}
			for _, tile := range row {
				picture_row = append(picture_row, tile.rows[i+1][1:tile_size-1])
			}
			picture = append(picture, strings.Join(picture_row, ""))
		}
	}
	return picture
}

func constructJigsaw(tiles []Tile) [][]Tile {
	grid_size := int(math.Sqrt(float64(len(tiles))))
	top_left_corner := []int{0, 1, 1, 0}
	jigsaw := [][]Tile{}
	jigsaw = append(jigsaw, []Tile{})
	for _, tile := range tiles {
		if slices.Compare(sideMatches(tile, tiles), top_left_corner) == 0 {
			jigsaw[0] = append(jigsaw[0], tile)
			break
		}
	}

	for i := range grid_size - 1 {
		right := findRight(jigsaw[0][i], tiles)
		jigsaw[0] = append(jigsaw[0], right)
	}

	for i := range grid_size - 1 {
		jigsaw = append(jigsaw, []Tile{})
		bottom := findBottom(jigsaw[i][0], tiles)
		jigsaw[i+1] = append(jigsaw[i+1], bottom)

		for j := range grid_size - 1 {
			right_bottom := findRightBottom(jigsaw[i+1][j], jigsaw[i][j+1], tiles)
			jigsaw[i+1] = append(jigsaw[i+1], right_bottom)
		}
	}
	return jigsaw
}

func countHashes(lines []string) int {
	result := 0
	for _, line := range lines {
		for _, ch := range line {
			if ch == '#' {
				result += 1
			}
		}
	}
	return result
}

func containsMonster(dr int, dc int, picture []string) bool {
	for r := range 3 {
		for c := range 20 {
			if monster[r][c] == '#' {
				if picture[r+dr][c+dc] != '#' {
					return false
				}
			}
		}
	}

	return true
}

func countMonsters(picture []string) int {
	result := 0
	pic_size := len(picture)
	monster_height := len(monster)
	monster_width := len(monster[0])

	for dr := range pic_size - monster_height {
		for dc := range pic_size - monster_width {
			if containsMonster(dr, dc, picture) {
				result += 1
			}
		}
	}
	return result
}

func sidesOfRows(rows []string) []string {
	max := len(rows) - 1
	edges := make([][]byte, 4)

	for i := range rows {
		edges[Top] = append(edges[Top], rows[0][i])
		edges[Right] = append(edges[Right], rows[i][max])
		edges[Bottom] = append(edges[Bottom], rows[max][i])
		edges[Left] = append(edges[Left], rows[i][0])
	}

	var result []string
	for _, edge := range edges {
		result = append(result, string(edge))
	}
	return result
}

func rotate_lines(lines []string) []string {
	height := len(lines)
	width := len(lines[0])

	var result []string
	for c := range width {
		row := []byte{}
		for r := range height {
			row = append(row, lines[height-1-r][c])
		}
		result = append(result, string(row))
	}
	return result
}

func rotate(tile Tile) Tile {
	rows := rotate_lines(tile.rows)
	return Tile{tile.id, rows, sidesOfRows(rows), tile.edge_vals}
}

func flip_lines(lines []string) []string {
	var rows []string
	height := len(lines)
	for r := range height {
		rows = append(rows, lines[height-1-r])
	}
	return rows
}

func flip(tile Tile) Tile {
	rows := flip_lines(tile.rows)
	return Tile{tile.id, rows, sidesOfRows(rows), tile.edge_vals}
}

func transpose(strArr []string) []string {
	xl := len(strArr[0])
	yl := len(strArr)
	result := make([][]byte, xl)
	for i := range result {
		result[i] = make([]byte, yl)
	}
	var returnStr []string
	for i := range xl {
		for j := range yl {
			result[i][j] = strArr[j][i]
		}
		returnStr = append(returnStr, string(result[i]))
	}
	return returnStr
}

func getEdgeValues(lines []string) []int {
	var edges []int
	top, top_reverse := getEdgeValue(lines[0])
	edges = append(edges, top)
	edges = append(edges, top_reverse)

	bottom, bottom_reverse := getEdgeValue(lines[len(lines)-1])
	edges = append(edges, bottom)
	edges = append(edges, bottom_reverse)
	return edges
}

func getEdgeValue(edge string) (int, int) {
	forward := 0
	reverse := 0
	for _, bit := range edge {
		forward <<= 1
		if bit == '#' {
			forward += 1
		}
	}
	for i := len(edge) - 1; i >= 0; i-- {
		reverse <<= 1
		if edge[i] == '#' {
			reverse += 1
		}
	}
	return forward, reverse
}

func parseBlocks(blocks []string) ([]Tile, error) {
	var tiles []Tile
	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		tileId, err := strconv.Atoi(lines[0][5:9])
		if err != nil {
			return nil, err
		}
		lines = lines[1:]
		var edge_vals []int
		edge_vals = append(edge_vals, getEdgeValues(lines)...)
		lines = transpose(lines)
		edge_vals = append(edge_vals, getEdgeValues(lines)...)
		lines = transpose(lines)

		tiles = append(tiles, Tile{tileId, lines, sidesOfRows(lines), edge_vals})
	}
	return tiles, nil
}

func findRightBottom(left Tile, top Tile, tiles []Tile) Tile {
	var bottom_right Tile
	for _, tile := range tiles {
		if top.id == tile.id || left.id == tile.id {
			continue
		}

		bottom_right = tile
		for range 4 {
			if top.edges[Bottom] == bottom_right.edges[Top] && left.edges[Right] == bottom_right.edges[Left] {
				return bottom_right
			}
			bottom_right = rotate(bottom_right)
		}
		bottom_right = flip(bottom_right)

		for range 4 {
			if top.edges[Bottom] == bottom_right.edges[Top] && left.edges[Right] == bottom_right.edges[Left] {
				return bottom_right
			}
			bottom_right = rotate(bottom_right)
		}
	}
	panic("Could not find right bottom")
}

func findBottom(top Tile, tiles []Tile) Tile {
	for _, tile := range tiles {
		if top.id == tile.id {
			continue
		}
		bottom := tile

		for range 4 {
			if top.edges[Bottom] == bottom.edges[Top] {
				return bottom
			}
			bottom = rotate(bottom)
		}
		bottom = flip(bottom)
		for range 4 {
			if top.edges[Bottom] == bottom.edges[Top] {
				return bottom
			}
			bottom = rotate(bottom)
		}
	}
	panic("Could not find bottom")
}

func findRight(left Tile, tiles []Tile) Tile {
	for _, tile := range tiles {
		if left.id == tile.id {
			continue
		}
		right := tile

		for range 4 {
			if left.edges[Right] == right.edges[Left] {
				return right
			}
			right = rotate(right)
		}
		right = flip(right)
		for range 4 {
			if left.edges[Right] == right.edges[Left] {
				return right
			}
			right = rotate(right)
		}
	}
	panic("Could not find right tile")
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sideMatches(candidate Tile, tiles []Tile) []int {
	matches := make([]int, 4)
	for _, tile := range tiles {
		if candidate.id == tile.id {
			continue
		}
		for left := range 4 {
			for right := range 4 {
				if candidate.edges[left] == tile.edges[right] {
					matches[left] += 1
				}
				if candidate.edges[left] == reverse(tile.edges[right]) {
					matches[left] += 1
				}
			}
		}
	}

	return matches
}

func neighboutCount(candidate Tile, tiles []Tile) int {
	result := 0
	for _, tile := range tiles {
		if candidate.id == tile.id {
			continue
		}
		for _, val := range tile.edge_vals {
			if slices.Contains(candidate.edge_vals, val) {
				result += 1
			}
		}
	}

	return result / 2
}

func getTiles(input []byte) []Tile {
	blocks := aoc.RemoveEmpties(strings.Split(string(input), "\n\n"))
	tiles, err := parseBlocks(blocks)
	check(err)

	return tiles
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
