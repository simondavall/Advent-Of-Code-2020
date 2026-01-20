package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"aoc"
)

var title string = "# Day 18: Operation Order #"
var url string = "https://adventofcode.com/2020/day/18"
var expectedResult1 int64 = 25190263477788
var expectedResult2 int64 = 297139939002972

type stackItem struct {
	value int64
	op    rune
}

func PartOne(input []byte) int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	var tally int64 = 0
	var cur int64
	for _, str := range lines {
		var stack aoc.Stack[stackItem]
		cur = 0
		var op rune
		ops := []rune{'+', '*'}
		for _, ch := range str {
			if '0' <= ch && ch <= '9' {
				if op == 0 {
					cur = int64(ch - '0')
					continue
				} else {
					cur = applyOp(cur, op, int64(ch-'0'))
					op = 0
					continue
				}
			}
			if slices.Contains(ops, ch) {
				op = ch
				continue
			}
			if ch == '(' {
				stack.Push(stackItem{cur, op})
				cur = 0
				op = 0
			}
			if ch == ')' {
				p, err := stack.Pop()
				if err != nil {
					fmt.Println(err)
					return 0
				}
				if p.value > 0 && p.op > 0 {
					cur = applyOp(p.value, p.op, cur)
				}
			}
		}
		tally += cur
	}

	return tally
}

func PartTwo(input []byte) int64 {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	var tally int64 = 0
	var cur int64
	for _, str := range lines {
		var stack = make([]aoc.Stack[stackItem], 1)
		stackIndex := 0
		cur = 0
		var op rune
		ops := []rune{'+', '*'}
		for _, ch := range str {
			if '0' <= ch && ch <= '9' {
				if op == 0 {
					cur = int64(ch - '0')
					continue
				} else {
					cur = applyOp(cur, op, int64(ch-'0'))
					op = 0
					continue
				}
			}
			if slices.Contains(ops, ch) {
				if ch == '*' {
					stack[stackIndex].Push(stackItem{cur, ch})
					cur = 0
					op = 0
					continue
				}
				op = ch
				continue
			}
			if ch == '(' {

				stackIndex++
				stack = append(stack, aoc.Stack[stackItem]{})
				stack[stackIndex].Push(stackItem{cur, op})
				cur = 0
				op = 0
			}
			if ch == ')' {
				for !stack[stackIndex].IsEmpty() {
					p, _ := stack[stackIndex].Pop()
					if p.value > 0 && p.op > 0 {
						cur = applyOp(p.value, p.op, cur)
					}
				}
				stack = stack[:stackIndex]
				stackIndex--
			}
		}
		// flush stack
		for !stack[stackIndex].IsEmpty() {
			p, _ := stack[stackIndex].Pop()
			if p.value > 0 && p.op > 0 {
				cur = applyOp(p.value, p.op, cur)
			}
		}
		stackIndex--

		tally += cur
	}

	return tally
}

func applyOp(cur int64, op rune, value int64) int64 {
	switch op {
	case '+':
		return cur + value
	case '*':
		return cur * value
	default:
		panic(fmt.Sprintf("Unknown operator:%c", op))
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
