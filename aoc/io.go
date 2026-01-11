package aoc

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		str := scanner.Text()
		if str != "" {
			lines = append(lines, str)
		}
	}
	return lines, scanner.Err()
}

func ReadFileToIntArray(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	var lines []int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		str := scanner.Text()
		if str != "" {
			n, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			lines = append(lines, n)
		}
	}
	return lines, scanner.Err()
}

func ReadFileToInt64Array(path string) ([]int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	var lines []int64
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		str := scanner.Text()
		if str != "" {
			n, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil, err
			}
			lines = append(lines, n)
		}
	}
	return lines, scanner.Err()
}

func ReadFileSplitBy(path string, sep string) ([]string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	content := string(b)

	for str := range strings.SplitSeq(content, sep) {
		if str != "" {
			lines = append(lines, str)
		}
	}
	return lines, nil
}
