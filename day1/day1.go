package day1

import (
	"log"
	"strconv"
)

// starts with 0
var start int64 = 0

// should be 477

func SumFrequencies(freqs []string) int64 {
	var current = start
	for _, freq := range freqs {
		sign := string([]rune(freq)[0])
		freqValue := obtainFreqValue(freq)
		if sign == "+" {
			current += freqValue
		} else {
			current -= freqValue
		}
	}
	return current
}

// should be 390

func Rollfrequenciesbeforevaluerepeats(freqs []string) int64 {
	var current = start
	mapOfHits := make(map[int64]bool)
	index := 0
	for {
		if mapOfHits[current] {
			return current
		} else {
			mapOfHits[current] = true
		}
		sign := string([]rune(freqs[index])[0])
		freqValue := obtainFreqValue(freqs[index])
		if sign == "+" {
			current += freqValue
		} else {
			current -= freqValue
		}
		if index == len(freqs)-1 {
			index = 0
		} else {
			index += 1
		}
	}
}

func obtainFreqValue(freq string) int64 {
	value, err := strconv.ParseInt(freq[1:], 10, 64)
	if err != nil {
		log.Fatalln("Could not convert string to int!")
	}
	return value
}
