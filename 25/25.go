package main

import (
	"aoc"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var title string = "# Day 25: Combo Breaker #"
var url string = "https://adventofcode.com/2020/day/25"
var expectedResult1 int64 = 545789
var expectedResult2 int64 = 0

func PartOne(input []byte) int64 {
	cardPublicKey, doorPublicKey := processInput(input)
	subjectNumber := int64(7)

	cardLoopSize := 0
	publicKey := int64(1)
	for publicKey != cardPublicKey {
		cardLoopSize++
		publicKey = transform(publicKey, subjectNumber)
	}

	doorLoopSize := 0
	publicKey = 1
	for publicKey != doorPublicKey {
		doorLoopSize++
		publicKey = transform(publicKey, subjectNumber)
	}

	subjectNumber = cardPublicKey
	encryptionKey := int64(1)
	for range doorLoopSize {
		encryptionKey = transform(encryptionKey, subjectNumber)
	}

	return encryptionKey
}

func transform(key int64, n int64) int64 {
	key *= n
	return key % 20201227
}

func PartTwo() int64 {
	return 0
}

func processInput(input []byte) (int64, int64) {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	
	cardPublicKey, err := strconv.ParseInt(lines[0], 10, 64)
	check(err)

	doorPublicKey, err := strconv.ParseInt(lines[1], 10, 64)
	check(err)
	
	return cardPublicKey, doorPublicKey
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
