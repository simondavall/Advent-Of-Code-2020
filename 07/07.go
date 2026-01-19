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

var title string = "# Day 7: Handy Haversacks #"
var url string = "https://adventofcode.com/2020/day/7"
var expectedResult1 int64 = 252
var expectedResult2 int64 = 35487

type Bag struct {
	name     string
	children []BagChild
}

type BagChild struct {
	name   string
	amount int
}

var (
	bags       = make(map[string]Bag)
	goldCache  = make(map[string]bool)
	countCache = make(map[string]int)
)

func PartOne(input []byte) int64 {
	clearCache()
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	loadBags(lines)

	var tally int64 = 0
	for name := range bags {
		if CanContainGoldBag(name) {
			tally++
		}
	}

	return tally - 1
}

func PartTwo() int64 {
	var tally = 0

	shinyGold := bags["shiny gold"]
	for _, child := range shinyGold.children {
		tally += GetBagCount(child)
	}

	return int64(tally)
}

func GetBagCount(bagChild BagChild) int {
	if val, ok := countCache[bagChild.name]; ok {
		return bagChild.amount + (val * bagChild.amount)
	}

	current := bags[bagChild.name]
	childrenCount := 0

	for _, next := range current.children {
		nextcount := GetBagCount(next)
		childrenCount += nextcount
	}

	countCache[bagChild.name] = childrenCount
	return bagChild.amount + (bagChild.amount * childrenCount)
}

func CanContainGoldBag(bagName string) bool {
	if bagName == "shiny gold" {
		return true
	}

	if hasGold, ok := goldCache[bagName]; ok {
		return hasGold
	}

	bag := bags[bagName]
	if bag.children == nil {
		return false
	}

	for _, next := range bag.children {
		if hasGold, ok := goldCache[next.name]; ok {
			if hasGold {
				return true
			}
			continue
		}

		hasGold := CanContainGoldBag(next.name)

		goldCache[next.name] = hasGold
		if hasGold {
			goldCache[bagName] = true
			return true
		}
	}

	goldCache[bagName] = false
	return false
}

func loadBags(lines []string) {
	pattern1 := " (\\d+) ([\\w\\s]+) bags?[,.]"
	
	for _, line := range lines {
		s := strings.Split(line, " bags contain")
		r := regexp.MustCompile(pattern1)
		matches := r.FindAllStringSubmatch(s[1], -1)
		var children []BagChild
		for _, match := range matches {
			name := match[2]
			amount, _ := strconv.Atoi(match[1])
			children = append(children, BagChild{name, amount})
		}
		var bag = Bag{s[0], children}
		bags[s[0]] = bag
	}
}

func clearCache(){
	bags       = make(map[string]Bag)
	goldCache  = make(map[string]bool)
	countCache = make(map[string]int)
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
