package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"aoc"
)

var title string = "# Day 21: Allergen Assessment #"
var url string = "https://adventofcode.com/2020/day/21"
var expectedResult1 int64 = 2493
var expectedResult2 string = "kqv,jxx,zzt,dklgl,pmvfzk,tsnkknk,qdlpbt,tlgrhdh"

type Food struct {
	ingredients []string
	allergens   []string
}

var allergens map[string][]string

func PartOne(input []byte) int64 {
	var foods = getFoods(input)
	var tally int64 = 0

	allergens = make(map[string][]string)
	for _, food := range foods {
		for _, alg := range food.allergens {
			val, exists := allergens[alg]
			if !exists {
				allergens[alg] = food.ingredients
				continue
			}
			allergens[alg] = aoc.Intersects(food.ingredients, val)
		}
	}

	resolved := false
	for !resolved {
		resolved = true
		for i, ing1 := range allergens {
			for j, ings := range allergens {
				if i == j || len(ing1) != 1 {
					continue
				}
				idx := slices.Index(ings, ing1[0])
				if idx == -1 {
					continue
				}
				allergens[j] = slices.Concat(allergens[j][:idx], allergens[j][idx+1:])
				resolved = false
			}
		}
	}

	ingredients := make(map[string]string)
	for k, v := range allergens {
		ingredients[v[0]] = k
	}

	for _, food := range foods {
		for _, ing := range food.ingredients {
			if len(ingredients[ing]) == 0 {
				tally++
			}
		}
	}

	return tally
}

func PartTwo() string {
	var algList []string
	for alg := range allergens {
		algList = append(algList, alg)
	}
	slices.Sort(algList)

	var dangerous []string
	for _, alg := range algList {
		dangerous = append(dangerous, allergens[alg][0])
	}

	return strings.Join(dangerous, ",")
}

func getFoods(input []byte) []Food {
	var lines = aoc.RemoveEmpties(strings.Split(string(input), "\n"))
	var foods []Food
	for _, line := range lines {
		s := strings.Split(line, " (contains ")
		f := Food{strings.Split(s[0], " "), strings.Split(strings.TrimRight(s[1], ")"), ", ")}
		foods = append(foods, f)
	}

	return foods
}

func main() {
	var resultPartOne int64 = -1
	var resultPartTwo = ""

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
		fmt.Printf("Part 2 result: %s in %s\n", resultPartTwo, time.Since(startPart2))
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
