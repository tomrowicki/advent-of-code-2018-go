package main

//import "fmt"
import (
	"AdventOfCode2018/common"
	"AdventOfCode2018/day1"
	"AdventOfCode2018/day2"
)

func main() {
	//result := d1p1()
	//fmt.Println("Result: ", result)
	//fmt.Println("Francesk's solution: ", CampD1P1())
	d2p1()
}

func d1p1() int64 {
	return day1.SumFrequencies(common.InputLinesToSlice("day1/day1_input.txt"))
}

func d1p2() int64 {
	return day1.Rollfrequenciesbeforevaluerepeats(common.InputLinesToSlice("day1/day1_input.txt"))
}

func d2p1() {
	day2.FindDoublesTriplesAndMultiplyOccurences(common.InputLinesToSlice("day2/day2_input.txt"))
}
