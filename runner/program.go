package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"time"
)

func main() {

	var count = 0
	var totalSolutions = 25
	var failed []int
	var executionTimes = make(map[int]time.Duration)
	timer := time.Now()

	for i := range totalSolutions {
		count++
		dir := fmt.Sprintf("../%02d", i+1)
		program := fmt.Sprintf("%02d.go", i+1)
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("go", "run", program, "input.txt")

		executionTime := time.Now()
		output, err := cmd.CombinedOutput()
		elapsed := time.Since(executionTime)
		executionTimes[i+1] = elapsed
		if err != nil {
			failed = append(failed, i+1)
			fmt.Print(string(output))
			fmt.Println("Oops, Failed!!")
		} else {
			fmt.Print(string(output))
			fmt.Println("Great Success!!")
		}
	}

	elapsed := time.Since(timer)
	formatted := fmt.Sprintf("%.4f", elapsed.Seconds())
	fmt.Printf("\nAll solutions ran in %s secs", formatted)

	if len(failed) > 0 {
		fmt.Printf("\nIncorrect results found for %d/%d solutions.", len(failed), totalSolutions)
		fmt.Println("\nFailed tests:")
		for _, f := range failed {
			fmt.Printf("%02d\n", f)
		}
	}

	keys := GetSortedKeys(executionTimes)
	fmt.Println("\n\n### Execution Times Summary ###")
	for _, k := range keys {
		fmt.Printf("Day %02d - %s\n", k, executionTimes[k])
	}
}

func GetSortedKeys(executionTimes map[int]time.Duration) []int {
	var keys []int
	for k := range executionTimes {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return executionTimes[keys[i]] > executionTimes[keys[j]]
	})

	return keys
}
