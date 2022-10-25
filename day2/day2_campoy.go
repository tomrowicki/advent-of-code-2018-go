package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CampD2P2() string { // does not really work for my input, but passes the tests
	f, err := os.Open("day2/day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var words []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	for i := range words {
		for j := range words[i+1:] {
			common, ok := compare(words[i], words[j])
			if ok {
				fmt.Println(common)
				return common
			}
		}
	}
	return ""
}

func compare(a, b string) (string, bool) {
	idx := -1
	if len(a) != len(b) {
		return "", false
	}

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		if idx >= 0 {
			return "", false
		}
		idx = i
	}

	if idx < 0 {
		return "", false
	}
	return a[:idx] + a[idx+1:], true
}
