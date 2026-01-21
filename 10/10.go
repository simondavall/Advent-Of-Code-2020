package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"aoc"
)

var title string = "## Day 10: Adapter Array ##"
var url string = "https://adventofcode.com/2020/day/10"
var expectedResult1 int64 = 2414
var expectedResult2 int64 = 21156911906816

var joltages []int

func PartOne(input []byte) int64 {
	joltages = getJoltages(input)
	sort.Ints(joltages[:])
	joltages = append(joltages, joltages[len(joltages)-1]+3)
	diff := make([]int, 4)
	prev := 0
	for _, n := range joltages {
		diff[n-prev]++
		prev = n
	}

	return int64(diff[1] * diff[3])
}

func PartTwo() int64 {
	dp := make([]int64, len(joltages))
	dp[0] = 1

	for i := 1; i < len(joltages); i++ {
		dp[i] = 0
		for j := i - 1; j >= 0 && joltages[i]-joltages[j] <= 3; j-- {
			dp[i] += dp[j]
		}
	}

	return dp[len(dp)-1]
}

func getJoltages(input []byte) []int {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	var joltages []int
	for _, line := range lines{
		n, err := strconv.Atoi(line)
		check(err)
		joltages = append(joltages, n)
	}
	joltages = append(joltages, 0)
	return joltages
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
		resultPartTwo = PartTwo()
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
