package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CampD1P1() int {
	f, err := os.Open("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sum := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		sum += n
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

func CampD12() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var nums []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		nums = append(nums, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	seen := map[int]bool{0: true}
	for { // this is nicer than mine because it does not require index manpilations
		for _, n := range nums {
			sum += n
			if seen[sum] {
				fmt.Println(sum)
				return sum
			}
			seen[sum] = true
		}
	}
}
