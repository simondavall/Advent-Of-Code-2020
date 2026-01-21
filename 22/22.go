package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"aoc"
)

var title string = "## Day 22: Crab Combat ##"
var url string = "https://adventofcode.com/2020/day/22"
var expectedResult1 int64 = 32179
var expectedResult2 int64 = 30498

func PartOne(input []byte) int64 {
	deck1, deck2 := getDecks(input)
	players := []aoc.Queue[int]{deck1, deck2}

	winner, _ := playNewGame(players)
	return getScore(winner)
}

func PartTwo(input []byte) int64 {
	deck1, deck2 := getDecks(input)
	players := []aoc.Queue[int]{deck1, deck2}

	winner, _ := playNewRecursiveGame(players, 1)
	return getScore(winner)
}

func playNewGame(players []aoc.Queue[int]) (aoc.Queue[int], int) {
	for !players[0].IsEmpty() && !players[1].IsEmpty() {
		card1, _ := players[0].Dequeue()
		card2, _ := players[1].Dequeue()

		if card1 > card2 {
			players[0].Enqueue(card1)
			players[0].Enqueue(card2)
		} else {
			players[1].Enqueue(card2)
			players[1].Enqueue(card1)
		}
	}
	if players[0].Count() > 0 {
		return players[0], 0
	} else {
		return players[1], 1
	}
}

func roundSeenBefore(players []aoc.Queue[int], cache []map[string]bool) bool {
	key0 := getKey(players[0])
	key1 := getKey(players[1])

	if cache[0][key0] || cache[1][key1] {
		return true
	}

	cache[0][key0] = true
	cache[1][key1] = true

	return false
}

func playNewRecursiveGame(players []aoc.Queue[int], game int) (aoc.Queue[int], int) {
	cache := make([]map[string]bool, 2)
	cache[0] = make(map[string]bool)
	cache[1] = make(map[string]bool)
	round := 0
	for !players[0].IsEmpty() && !players[1].IsEmpty() {
		round++

		if roundSeenBefore(players, cache) {
			return nil, 0
		}

		card1, _ := players[0].Dequeue()
		card2, _ := players[1].Dequeue()

		var winner int
		if players[0].Count() >= card1 && players[1].Count() >= card2 {
			newPlayer := make([]aoc.Queue[int], 2)
			newPlayer[0] = slices.Clone(players[0][:card1])
			newPlayer[1] = slices.Clone(players[1][:card2])

			_, winner = playNewRecursiveGame(newPlayer, game+1)
		} else {
			if card1 > card2 {
				winner = 0
			} else {
				winner = 1
			}
		}
		if winner == 0 {
			players[0].Enqueue(card1)
			players[0].Enqueue(card2)
		} else {
			players[1].Enqueue(card2)
			players[1].Enqueue(card1)
		}
	}

	if players[0].Count() > 0 {
		return players[0], 0
	} else {
		return players[1], 1
	}
}

func getDeck(cards string) aoc.Queue[int] {
	var q aoc.Queue[int]
	for _, card := range strings.Split(cards, "\n")[1:] {
		c, err := strconv.Atoi(card)
		if err != nil {
			panic(fmt.Sprintf("Error convering card to int: card:%s, err:%s", card, err))
		}
		q.Enqueue(c)
	}
	return q
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.ReplaceAll(fmt.Sprint(a), " ", delim), "[]")
}

func getKey(q aoc.Queue[int]) string {
	return arrayToString(q, ",")
}

func getScore(q aoc.Queue[int]) int64 {
	var score int64 = 0
	multiplier := q.Count()
	for !q.IsEmpty() {
		card, _ := q.Dequeue()
		score += int64(card * multiplier)
		multiplier--
	}
	return score
}

func getDecks(input []byte) (aoc.Queue[int], aoc.Queue[int]){
	var blocks = aoc.RemoveEmpties(strings.Split(string(input), "\n\n"))
	if (len(blocks) != 2){
		panic(fmt.Sprintf("Expected 2 blocks when processing input. Found:%d", len(blocks)))
	}
	var deck1 = getDeck(blocks[0])
	var deck2 = getDeck(blocks[1])
	return deck1, deck2
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
