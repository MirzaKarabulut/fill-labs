package main

func Question3(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}

	if len(strArr) == 1 {
		return strArr[0]
	}

	countMap := make(map[string]int)

	for _, str := range strArr {
		countMap[str]++
	}

	maxKey := ""
	maxValue := 0

	for key, value := range countMap {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}

	return maxKey
}
