package main

import (
	"AdventOfCode2018/common"
	"AdventOfCode2018/day1"
	"AdventOfCode2018/day2"
	"fmt"
)

func main() {
	fmt.Println("D2P2: ", d2p2())
}

func d1p1() int64 {
	return day1.SumFrequencies(common.InputLinesToSlice("day1/day1_input.txt"))
}

func d1p2() int64 {
	return day1.Rollfrequenciesbeforevaluerepeats(common.InputLinesToSlice("day1/day1_input.txt"))
}

func d2p1() int64 {
	// should be 5390
	return day2.FindDoublesTriplesAndMultiplyOccurences(common.InputLinesToSlice("day2/day2_input.txt"))
}

func d2p2() string {
	// should be nvosmkcdtdbfhyxsphzgraljq
	return day2.FindIdsThatDifferAtOnePosition(common.InputLinesToSlice("day2/day2_input.txt"))
}
