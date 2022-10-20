package day2

func FindDoublesTriplesAndMultiplyOccurences(boxIDs []string) int64 {
	doubles := 0
	triples := 0
	for _, id := range boxIDs {
		// check if the id has a double
		if hasXOccurences(id, 2) {
			doubles++
		}
		// check if the id has a tripe
		if hasXOccurences(id, 3) {
			triples++
		}
	}
	return int64(doubles * triples)
}

func FindIdsThatDifferAtOnePosition(boxIds []string) string {
	for index, id := range boxIds {
		commonString, match := lookForMatch(id, index, boxIds)
		if match {
			return commonString
		}
	}
	return ""
}

func hasXOccurences(id string, occurences int) bool {
	for _, char := range id {
		charCount := countOccurencesOfChar(char, id)
		if charCount == occurences {
			return true
		}
	}
	return false
}

func countOccurencesOfChar(givenChar rune, id string) int {
	count := 0
	for _, char := range id {
		if char == givenChar {
			count++
		}
	}
	return count
}

func lookForMatch(id string, index int, allIds []string) (string, bool) {
	for i := index + 1; i < len(allIds); i++ {
		oneDiff := hasOneDiff(id, allIds[i])
		if oneDiff {
			return getCommonChars(id, allIds[i]), true
		}
	}
	return "", false
}

func hasOneDiff(id string, secondId string) bool {
	differences := 0
	if len(id) != len(secondId) {
		return false
	}
	for i, _ := range id {
		char1 := string(id[i])
		char2 := string(secondId[i])
		if char1 != char2 {
			differences++
		}
	}
	if differences == 1 {
		return true
	}
	return false
}

func getCommonChars(id string, secondId string) string {
	commonChars := ""
	for i, _ := range id {
		if id[i] == secondId[i] {
			char := string(id[i])
			commonChars = commonChars + char
		}
	}
	return commonChars
}
