package main

import (
	"fmt"
	"os"
	"cmp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"aoc"
)

var title string = "## Day 16: Ticket Translation ##"
var url string = "https://adventofcode.com/2020/day/16"
var expectedResult1 int64 = 32835
var expectedResult2 int64 = 514662805187

type Range struct {
	lower int
	upper int
}

type Field struct {
	id     int
	name   string
	ranges []Range
}

type TicketInfo struct {
	fields    []Field
	my_ticket []int
	nearby    [][]int
}

func PartOne(input []byte) int64 {
	var ticketInfo = getTicketInfo(input)
	var tally int64 = 0

	for _, ticket := range ticketInfo.nearby {
		invalidNumber := getInvalidField(ticket, ticketInfo.fields)
		if invalidNumber > 0 {
			tally += int64(invalidNumber)
		}
	}

	return tally
}

func PartTwo(input []byte) int64 {
	var ticketInfo = getTicketInfo(input)
	whereValid := func(ticket []int) bool { return isTicketValid(ticket, ticketInfo.fields) }
	validTickets := aoc.Filter(ticketInfo.nearby, whereValid)

	var validFields [][][]string
	for _, ticket := range validTickets {
		var validIdsForTicket [][]string
		for _, n := range ticket {
			var validIdsForN []string
			fieldsForValue := validFieldsForValue(n, ticketInfo.fields)
			for _, f := range fieldsForValue {
				validIdsForN = append(validIdsForN, f.name)
			}
			sort.Strings(validIdsForN)
			validIdsForTicket = append(validIdsForTicket, validIdsForN)
		}
		validFields = append(validFields, validIdsForTicket)
	}

	transposed := aoc.Transpose(validFields)

	type Result struct {
		col_id int
		result []string
	}
	var result []Result
	for idx, vals := range transposed {
		ticket := vals[0]
		for _, items := range vals[1:] {
			ticket = intersects(ticket, items)
		}
		result = append(result, Result{idx, ticket})
	}
	lenCmp := func(a, b Result) int {
		return cmp.Compare(len(a.result), len(b.result))
	}
	slices.SortFunc(result, lenCmp)

	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			result[j].result = removeValue(result[j].result, result[i].result[0])
		}
	}

	var tally int64 = 1
	for _, res := range result {
		if strings.HasPrefix(res.result[0], "departure") {
			tally *= int64(ticketInfo.my_ticket[res.col_id])
		}
	}

	return tally
}

func intersects(first []string, second []string) []string {
	commonValues := make([]string, 0)
	for i, j := 0, 0; i < len(first) && j < len(second); {
		if first[i] == second[j] {
			commonValues = append(commonValues, second[j])
			i++
			j++
		} else if first[i] < second[j] {
			i++
		} else {
			j++
		}
	}
	return commonValues
}

func getInvalidField(ticket []int, fields []Field) int {
	for _, n := range ticket {
		isValid := false
		for _, field := range fields {
			if isValueValid(n, field) {
				isValid = true
				break
			}
		}
		if !isValid {
			return n
		}
	}
	return -1
}

func isValueValid(value int, field Field) bool {
	for _, rng := range field.ranges {
		if rng.lower <= value && value <= rng.upper {
			return true
		}
	}
	return false
}

func isTicketValid(ticket []int, fields []Field) bool {
	return getInvalidField(ticket, fields) < 0
}

func validFieldsForValue(value int, fields []Field) []Field {
	whereValid := func(field Field) bool {
		is_valid := false
		for _, rng := range field.ranges {
			if rng.lower <= value && value <= rng.upper {
				is_valid = true
				break
			}
		}
		return is_valid
	}
	return aoc.Filter(fields, whereValid)
}

func removeValue(s []string, str string) []string {
	newArray := make([]string, len(s)-1)
	idx := 0
	for _, val := range s {
		if val == str {
			continue
		}
		newArray[idx] = val
		idx++
	}
	return newArray
}

func getTicketInfo(input []byte) TicketInfo {
	var dataBlocks = aoc.RemoveEmpties(strings.Split(string(input), "\n\n"))
	ticketInfo, err := parseDataBlocks(dataBlocks)
	check(err)
	return ticketInfo
}

func parseDataBlocks(blocks []string) (TicketInfo, error) {
	var input TicketInfo

	rawFields := strings.Split(blocks[0], "\n")
	for idx, rawField := range rawFields {
		fields := strings.Split(rawField, ": ")
		var fieldRanges []Range
		for rawRanges := range strings.SplitSeq(fields[1], " or ") {
			items := strings.Split(rawRanges, "-")
			lower, err := strconv.Atoi(items[0])
			if err != nil {
				return input, err
			}
			upper, err := strconv.Atoi(items[1])
			if err != nil {
				return input, err
			}
			fieldRanges = append(fieldRanges, Range{lower, upper})
		}
		input.fields = append(input.fields, Field{idx, fields[0], fieldRanges})
	}

	rawMyTicket := strings.Split(blocks[1], "\n")
	for val := range strings.SplitSeq(rawMyTicket[1], ",") {
		ticket, err := strconv.Atoi(val)
		if err != nil {
			return input, err
		}
		input.my_ticket = append(input.my_ticket, ticket)
	}

	rawNearby := strings.Split(blocks[2], "\n")
	for _, nearby := range rawNearby[1:] {
		if nearby == "" {
			continue
		}
		numbers := strings.Split(nearby, ",")
		var ticketNumbers []int
		for _, strN := range numbers {
			n, err := strconv.Atoi(strN)
			if err != nil {
				return input, err
			}
			ticketNumbers = append(ticketNumbers, n)
		}
		input.nearby = append(input.nearby, ticketNumbers)
	}
	return input, nil
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
