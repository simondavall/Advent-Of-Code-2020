package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"aoc"
)

var title string = "## Day 13: Shuttle Search ##"
var url string = "https://adventofcode.com/2020/day/13"
var expectedResult1 int64 = 174
var expectedResult2 int64 = 780601154795940

type Bus struct {
	id     int64
	offset int64
}

func PartOne(input []byte) int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	time, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println(err)
		return 0
	}

	buses := getBusData(lines)
	var timestamp = int64(time)
	for {
		for _, bus := range buses {
			if timestamp%bus.id == 0 {
				return int64((timestamp - int64(time)) * bus.id)
			}
		}
		timestamp++
	}
}

func PartTwo(input []byte) int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	buses := getBusData(lines)
	var timestamp int64 = 0
	found := false
	step := buses[0].id
	busCount := 2

	for !found {
		timestamp += step
		found = true
		for _, bus := range buses[:busCount] {
			if (timestamp+bus.offset)%bus.id != 0 {
				found = false
				break
			}
		}
		if found && busCount < len(buses) {
			found = false
			step *= buses[busCount-1].id
			busCount++
		}
	}

	return int64(timestamp)
}

func getBusData(lines []string) []Bus {
	var buses []Bus
	busData := strings.Split(lines[1], ",")
	for idx, bus := range busData {
		if bus != "x" {
			busId, err := strconv.Atoi(bus)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			buses = append(buses, Bus{int64(busId), int64(idx)})
		}
	}
	return buses
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
