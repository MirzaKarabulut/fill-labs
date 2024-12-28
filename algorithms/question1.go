package main

import (
	"sort"
	"strings"
)

func Question1(charArr []string) []string {
	sort.SliceStable(charArr, func(i, j int) bool {
		countA := strings.Count(charArr[i], "a")
		countB := strings.Count(charArr[j], "a")
		if countA != countB {
			return countA > countB
		}
		return len(charArr[i]) > len(charArr[j])
	})
	return charArr
}
