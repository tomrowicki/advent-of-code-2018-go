package day2

func FindDoublesTriplesAndMultiplyOccurences(boxIDs []string) int64 {

	doubles := 0
	triples := 0

	for _, id := range boxIDs {
		// check if the id has a double
		if hasDoubleLetter(id) {
			doubles++
		}
		// check if the id has a tripe
		if hasTripleLetter(id) {
			triples++
		}
	}
	return int64(doubles * triples)
}

func hasDoubleLetter(id string) bool {
	for _, char := range id {
		charCount := countOccurencesOfChar(char, id)
		if charCount == 2 {
			return true
		}
	}
	return false
}

func hasTripleLetter(id string) bool {
	for _, char := range id {
		charCount := countOccurencesOfChar(char, id)
		if charCount == 3 {
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

//func establishUniqueRunes(id string) []rune {
//	mapOfRunes := make(map[rune]int)
//
//	for pos, char := range "日本語" {
//		fmt.Printf("character %c starts at byte position %d\n", char, pos)
//	}
//}
