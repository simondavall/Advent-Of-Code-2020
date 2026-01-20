package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"aoc"
)

var title string = "# Day 19: Monster Messages #"
var url string = "https://adventofcode.com/2020/day/19"
var expectedResult1 int64 = 233
var expectedResult2 int64 = 396

type Rule struct {
	kind              string
	single            string
	sub_rules         []int
	muliple_sub_rules []Rule
}

func PartOne(input []byte) int64 {
	lines, messages := processInput(input)
	var tally int64 = 0

	rules := make(map[int]Rule)
	for _, line := range lines {
		index, rule := parseLine(line)
		rules[index] = rule
	}

	for _, line := range messages {
		inputs := matchRuleIndex(line, rules, 0)
		if len(inputs) > 0 && fully_parsed(inputs) {
			tally++
		}
	}
	return tally
}

func PartTwo(input []byte) int64 {
	lines, messages := processInput(input)
	var tally int64 = 0

	rules := make(map[int]Rule)
	for _, line := range lines {
		if strings.HasPrefix(line, "8:") {
			line = "8: 42 | 42 8"
		}
		if strings.HasPrefix(line, "11:") {
			line = "11: 42 31 | 42 11 31"
		}
		index, rule := parseLine(line)
		rules[index] = rule
	}

	for _, line := range messages {
		inputs := matchRuleIndex(line, rules, 0)
		if len(inputs) > 0 && fully_parsed(inputs) {
			tally++
		}
	}
	return tally
}

func fully_parsed(inputs []string) bool {
	for _, input := range inputs {
		if len(input) == 0 {
			return true
		}
	}
	return false
}

func matchRuleIndex(line string, rules map[int]Rule, index int) []string {
	kind := rules[index].kind

	switch kind {
	case "single":
		if len(line) > 0 && line[:1] == rules[index].single {
			return []string{line[1:]}
		} else {
			return []string{}
		}
	case "sub_rules":
		return matchSubRuleIndex(line, rules, rules[index].sub_rules)
	case "multi":
		var result []string

		for _, rule := range rules[index].muliple_sub_rules {
			result = append(result, matchSubRuleIndex(line, rules, rule.sub_rules)...)
		}
		return result
	default:
		panic(fmt.Sprintf("Unknown kind of rule: %s", kind))
	}
}

func matchSubRuleIndex(line string, rules map[int]Rule, sub_rules []int) []string {
	inputs := []string{line}

	for _, subindex := range sub_rules {
		if len(inputs) == 0 {
			break
		}

		new_inputs := []string{}
		for _, input := range inputs {
			new_inputs = append(new_inputs, matchRuleIndex(input, rules, subindex)...)
		}
		inputs = new_inputs
	}
	return inputs
}

func parseLine(line string) (int, Rule) {
	s := strings.Split(line, ": ")
	index, _ := strconv.Atoi(s[0])
	rule := parseRule(s[1])

	return index, rule
}

var findStringLiteral = regexp.MustCompile("\"(.*)\"")

func parseRule(line string) Rule {
	match := findStringLiteral.FindStringSubmatch(line)
	if match != nil {
		return Rule{"single", match[1], nil, nil}
	} else {
		var result []Rule
		for subrule := range strings.SplitSeq(line, " | ") {
			var sub_rules []int
			for rule := range strings.SplitSeq(subrule, " ") {
				n, err := strconv.Atoi(rule)
				if err != nil {
					panic(fmt.Sprintf("Failed to conver rule:%s Err:%s", rule, err))
				}
				sub_rules = append(sub_rules, n)
			}
			result = append(result, Rule{"sub_rules", "", sub_rules, nil})
		}
		return Rule{"multi", "", nil, result}
	}
}

func processInput(input []byte) ([]string, []string){
	var blocks = aoc.RemoveEmpties(strings.Split(string(input), "\n\n"))
	if (len(blocks) != 2){
		panic(fmt.Sprintf("Expected 2 blocks when processing input. Found:%d", len(blocks)))
	}

	lines := strings.Split(blocks[0], "\n")
	messages := strings.Split(blocks[1], "\n")

	return lines, messages
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
