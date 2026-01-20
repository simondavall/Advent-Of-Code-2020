package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var title string = "# Day 23: Crab Cups #"
var url string = "https://adventofcode.com/2020/day/23"
var expectedResult1 int64 = 32897654
var expectedResult2 int64 = 186715244496

type Cup struct {
	value int
	prev  *Cup
	next  *Cup
}

func PartOne(input []byte) int64 {
	start, size := processInput(input)
	curPtr := start

	ptr := start
	cups := make([]*Cup, size)
	for range size {
		cups[ptr.value-1] = ptr
		ptr = ptr.next
	}

	counter := 0
	for counter < 100 {
		counter++
		val := curPtr.value
		first := curPtr.next

		selectedValues, last := getSelectedValues(first)
		destVal := getDestinationValue(val, size, selectedValues)
		destination := cups[destVal-1]

		first.prev.next = last.next
		last.next.prev = first.prev
		destination.next.prev = last
		last.next = destination.next
		destination.next = first
		first.prev = destination

		curPtr = curPtr.next
	}

	oneCup := cups[0]
	curPtr = oneCup.next

	var tally int64 = 0
	for curPtr.value != 1 {
		tally *= 10
		tally += int64(curPtr.value)
		curPtr = curPtr.next
	}

	return tally
}

func PartTwo(input []byte) int64 {
	start, _ := processInput(input)
	curPtr := start
	last := curPtr.prev
	size := 1000000

	for i := range size - 9 {
		new_Cup := Cup{i + 10, last, nil}
		last.next = &new_Cup
		last = &new_Cup
	}
	curPtr.prev = last
	last.next = curPtr

	ptr := start
	cups := make([]*Cup, size)
	for range size {
		cups[ptr.value-1] = ptr
		ptr = ptr.next
	}

	counter := 0
	for counter < 10000000 {
		counter++
		val := curPtr.value
		first := curPtr.next

		selectedValues, last := getSelectedValues(first)
		destVal := getDestinationValue(val, size, selectedValues)
		destination := cups[destVal-1]

		first.prev.next = last.next
		last.next.prev = first.prev
		destination.next.prev = last
		last.next = destination.next
		destination.next = first
		first.prev = destination

		curPtr = curPtr.next
	}

	oneCup := cups[0]

	return int64(oneCup.next.value) * int64(oneCup.next.next.value)
}

func getDestinationValue(val int, size int, selectedValues map[int]bool) int {
	destVal := nextDestination(val, size)
	for selectedValues[destVal] {
		destVal = nextDestination(destVal, size)
	}
	return destVal
}

func nextDestination(cur int, size int) int {
	next := cur - 1
	if next < 1 {
		next += size
	}
	return next
}

func getSelectedValues(first *Cup) (map[int]bool, *Cup) {
	selectedValues := make(map[int]bool)
	selectedValues[first.value] = true
	last := first.next
	selectedValues[last.value] = true
	last = last.next
	selectedValues[last.value] = true
	return selectedValues, last
}

func processInput(input []byte) (*Cup, int) { 
	var numbers []int
	for _, b := range input[:len(input)-1] {
		n, err := strconv.Atoi(string(b)) 
		check(err)
		numbers = append(numbers, n)
	}

	first := Cup{numbers[0], nil, nil}
	prev := &first
	for _, n := range numbers[1:] {
		current := Cup{n, prev, nil}
		prev.next = &current
		prev = &current
	}
	first.prev = prev
	prev.next = &first

	return &first, len(numbers)
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
