package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func ToInt32Array(strArray []string) ([]int32, error) {
	var intArray []int32
	for _, str := range strArray {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intArray = append(intArray, int32(n))
	}
	return intArray, nil
}

func ToPair(strArr []string) (string, string, error) {
	if len(strArr) != 2 {
		return "", "", fmt.Errorf("expected array of length 2. found:%d", len(strArr))
	}
	return strArr[0], strArr[1], nil
}

func ToStrPair(str string, sep string) (string, string, error) {
	var strArr = strings.Split(str, sep)
	return ToPair(strArr)
}
