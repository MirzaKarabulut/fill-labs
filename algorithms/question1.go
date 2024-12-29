package main

import (
	"sort"
	"strings"
)

func Question1(strArr []string) []string {
	if strArr == nil {
		return nil
	}

	if len(strArr) == 0 || len(strArr) == 1 {
		return strArr
	}

	sort.SliceStable(strArr, func(i, j int) bool {
		countA := strings.Count(strArr[i], "a")
		countB := strings.Count(strArr[j], "a")
		if countA != countB {
			return countA > countB
		}
		return len(strArr[i]) > len(strArr[j])
	})
	return strArr
}
